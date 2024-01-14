package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/codes"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry/attribute"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/htmx"
	"github.com/opoccomaxao/wblitz-watcher/pkg/utils/ioutils"
)

type errorsHandler struct{}

func HandleErrors() gin.HandlerFunc {
	return (&errorsHandler{}).Handle
}

func (h *errorsHandler) Handle(ctx *gin.Context) {
	body, err := ioutils.SniffReadCloser(&ctx.Request.Body)
	if err != nil {
		ctx.Error(err).SetType(gin.ErrorTypePrivate)
	} else {
		ctx.Next()
	}

	if len(ctx.Errors) > 0 {
		telemetry.ApplyAttributes(ctx, attribute.ErrorsExtended.StringSlice(h.extendErrors(ctx.Errors)))

		if body != nil {
			telemetry.ApplyAttributes(ctx, attribute.RequestBody.String(string(body)))
		}

		telemetry.ApplyStatus(ctx, codes.Error)
	} else {
		telemetry.ApplyStatus(ctx, codes.Ok)
	}

	for _, err := range ctx.Errors {
		telemetry.RecordError(ctx, err)
	}

	if len(ctx.Errors) == 0 {
		return
	}

	log.Printf("%s %s:\n%s", ctx.Request.Method, ctx.Request.RequestURI, h.extendErrors(ctx.Errors))

	if h.processErrorsByType(ctx, gin.ErrorTypeBind, http.StatusBadRequest, false) {
		return
	}

	if h.processErrorsByType(ctx, gin.ErrorTypePublic, http.StatusBadRequest, false) {
		return
	}

	if h.processErrorsByType(ctx, gin.ErrorTypeAny, http.StatusInternalServerError, true) {
		return
	}

	ctx.Status(http.StatusInternalServerError)
}

func (h *errorsHandler) extendErrors(errs []*gin.Error) []string {
	return lo.Map(errs, h.stringifyError)
}

func (h *errorsHandler) stringifyError(err *gin.Error, _ int) string {
	return fmt.Sprintf("%+v", err)
}

func (h *errorsHandler) tryParseBindError(err *gin.Error, _ int) []string {
	//nolint:errorlint // validation errors will be handled properly.
	switch typed := err.Err.(type) {
	case validator.ValidationErrors:
		return lo.Map(typed, h.tryParseFieldError)
	default:
		return []string{err.Error()}
	}
}

func (h *errorsHandler) tryParseFieldError(err validator.FieldError, _ int) string {
	return fmt.Sprintf("%s should be %s %s", err.Field(), err.ActualTag(), err.Param())
}

func (h *errorsHandler) processErrorsByType(
	ctx *gin.Context,
	typ gin.ErrorType,
	status int,
	hideText bool,
) bool {
	errs := ctx.Errors.ByType(typ)
	if len(errs) == 0 {
		return false
	}

	var texts []string

	if hideText {
		texts = []string{http.StatusText(status)}
	} else {
		texts = lo.Flatten(lo.Map(errs, h.tryParseBindError))
	}

	if htmx.IsHxRequest(ctx) {
		ctx.String(status, strings.Join(texts, "\n"))
	} else {
		ctx.JSON(status, &models.ErrorResponse{
			Errors: texts,
		})
	}

	return true
}
