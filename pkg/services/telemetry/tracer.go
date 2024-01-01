package telemetry

import (
	"context"

	"github.com/samber/lo"
	"go.opentelemetry.io/otel/trace"
)

type Tracer struct {
	trace.Tracer
	prefix string
	opts   []trace.SpanStartOption
}

func NewTracerWithOptions(
	tracer trace.Tracer,
	namePrefix string,
	opts []trace.SpanStartOption,
) trace.Tracer {
	if namePrefix != "" {
		namePrefix += "."
	}

	return &Tracer{
		Tracer: tracer,
		prefix: namePrefix,
		opts:   opts,
	}
}

func (t *Tracer) Start(
	ctx context.Context,
	name string,
	opts ...trace.SpanStartOption,
) (context.Context, trace.Span) {
	return t.Tracer.Start(
		ctx,
		t.prefix+name,
		lo.Flatten([][]trace.SpanStartOption{t.opts, opts})...,
	)
}
