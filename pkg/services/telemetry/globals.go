//nolint:gochecknoglobals
package telemetry

import (
	"go.opentelemetry.io/otel/trace"
)

var errTracer trace.Tracer = trace.NewNoopTracerProvider().Tracer("")
