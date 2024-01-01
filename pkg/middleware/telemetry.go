package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"

	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/telemetry/attribute"
)

func TelemetryInit() gin.HandlerFunc {
	host, _ := os.Hostname()

	return otelgin.Middleware(host)
}

func TelemetryExtra() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		telemetry.ApplyAttributes(ctx,
			attribute.URLPath.String(ctx.Request.URL.Path),
			attribute.URLQuery.String(ctx.Request.URL.RawQuery),
		)

		span := telemetry.SpanFromContext(ctx)
		span.SetName(ctx.Request.Method + " " + ctx.FullPath())
	}
}
