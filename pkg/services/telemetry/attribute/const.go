package attribute

import (
	"go.opentelemetry.io/otel/attribute"
)

const (
	Action          attribute.Key = "action"           // should be string.
	EventStack      attribute.Key = "event.stack"      // should be string slice.
	ErrorsExtended  attribute.Key = "errors.extended"  // should be string slice.
	RequestBody     attribute.Key = "request.body"     // should be string.
	ResponseBody    attribute.Key = "response.body"    // should be string.
	ResponseHeaders attribute.Key = "response.headers" // should be string map.
	ResponseStatus  attribute.Key = "response.status"  // should be int.
	TaskName        attribute.Key = "task.name"        // should be string.
	URLFull         attribute.Key = "url.full"         // should be string.
	URLQuery        attribute.Key = "url.query"        // should be string.
	URLPath         attribute.Key = "url.path"         // should be string.
)
