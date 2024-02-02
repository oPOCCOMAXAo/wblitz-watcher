package models

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type SpanType string

const (
	SpanTypeCommand SpanType = "command"
	SpanTypeEvent   SpanType = "event"
	SpanTypeHTTP    SpanType = "http"
	SpanTypeTask    SpanType = "task"
)

func (s SpanType) Attribute() attribute.KeyValue {
	return attribute.String("span.type", string(s))
}

func (s SpanType) Option() trace.SpanStartEventOption {
	attr := s.Attribute()

	return trace.WithAttributes(attr)
}
