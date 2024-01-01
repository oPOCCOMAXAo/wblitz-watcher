package telemetry

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

func RecordEvent(
	ctx context.Context,
	name string,
	opts ...trace.EventOption,
) {
	span := SpanFromContext(ctx)
	span.AddEvent(name, append(opts, trace.WithStackTrace(true))...)
}
