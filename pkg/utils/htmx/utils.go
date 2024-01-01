package htmx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsHxRequest(ctx *gin.Context) bool {
	return ctx.GetHeader("HX-Request") == "true"
}

func Redirect(ctx *gin.Context, url string) {
	if IsHxRequest(ctx) {
		ctx.Header("HX-Redirect", url)
	} else {
		ctx.Redirect(http.StatusFound, url)
	}
}
