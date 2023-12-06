package operation

import (
	"go.opentelemetry.io/otel/sdk/trace"
)

var events = []string{
	"start",
	"end",
	"start processing",
	"end processing",
	"api called",
	"message enqueued",
}

func NewEvent() trace.Event {
	return trace.Event{
		Name:       random(events),
		Attributes: GenerateRandomAttributes(3),
	}
}
