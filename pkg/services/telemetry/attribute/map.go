package attribute

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func SetMapStringSlice(
	span trace.Span,
	key attribute.Key,
	value map[string][]string,
) {
	for mKey, mVal := range value {
		span.SetAttributes(attribute.KeyValue{
			Key:   key + "." + attribute.Key(mKey),
			Value: attribute.StringSliceValue(mVal),
		})
	}
}
