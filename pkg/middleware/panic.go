package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/opoccomaxao/wblitz-watcher/pkg/models"
)

// HandlePanic should be after HandleErrors as it passes panic as error.
func HandlePanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}

			ctx.Error(errors.Wrapf(models.ErrPanic, "%+v", err))
			ctx.Abort()
		}()

		ctx.Next()
	}
}
