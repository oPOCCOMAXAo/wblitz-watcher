package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
)

func RecordError(ctx context.Context, err error) {
	if err == nil {
		return
	}

	SpanFromContext(ctx).
		RecordError(
			err,
			trace.WithStackTrace(true),
			trace.WithTimestamp(time.Now()),
		)
}

func RecordErrorBackground(err error) {
	ctx, span := errTracer.Start(context.Background(), "error")
	defer span.End()

	RecordError(ctx, err)
}
