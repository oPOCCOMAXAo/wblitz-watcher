package telemetry

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

//nolint:contextcheck
func SpanFromContext(ctx context.Context) trace.Span {
	ginCtx, _ := ctx.(*gin.Context)
	if ginCtx != nil {
		ctx = ginCtx.Request.Context()
	}

	return trace.SpanFromContext(ctx)
}

func DetouchContext(src context.Context) context.Context {
	span := SpanFromContext(src)

	return trace.ContextWithSpan(context.Background(), span)
}

func KVFromCarrier(carrier propagation.TextMapCarrier) []attribute.KeyValue {
	return lo.Map(carrier.Keys(), func(key string, _ int) attribute.KeyValue {
		return attribute.String(key, carrier.Get(key))
	})
}

func ApplyAttributes(
	ctx context.Context,
	attrs ...attribute.KeyValue,
) {
	span := SpanFromContext(ctx)
	if span == nil {
		return
	}

	span.SetAttributes(attrs...)
}

func ApplyCarrier(
	ctx context.Context,
	carrier propagation.TextMapCarrier,
) {
	ApplyAttributes(ctx, KVFromCarrier(carrier)...)
}

func ApplyStatus(
	ctx context.Context,
	code codes.Code,
) {
	span := SpanFromContext(ctx)
	if span == nil {
		return
	}

	// description doesn't work.
	span.SetStatus(code, "")
}
