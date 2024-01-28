package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/codes"
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

//nolint:nolintlint
//nolint:contextcheck
func RecordErrorBackground(err error) {
	ctx, span := errTracer.Start(context.Background(), "error")
	defer span.End()

	RecordError(ctx, err)
}

func RecordErrorFail(ctx context.Context, err error) {
	if err == nil {
		return
	}

	span := SpanFromContext(ctx)

	span.RecordError(
		err,
		trace.WithStackTrace(true),
		trace.WithTimestamp(time.Now()),
	)

	span.SetStatus(codes.Error, err.Error())
}
