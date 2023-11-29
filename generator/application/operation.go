package application

import (
	"context"

	"github.com/kubeshop/tracegen/generator/application/operation"
	"go.opentelemetry.io/otel/trace"
)

type Operation struct {
	Type   operation.Type
	Name   string
	Entity string
	tracer trace.Tracer
}

func NewOperation(entity string, tracer trace.Tracer) *Operation {
	opType := operation.GetRandomType()
	opName := operation.GetRandomName(opType, entity)
	return &Operation{
		Type:   opType,
		Name:   opName,
		Entity: entity,
		tracer: tracer,
	}
}

func (o *Operation) CreateSpan(ctx context.Context) (context.Context, trace.Span) {
	ctx, span := o.tracer.Start(ctx, o.Name)
	return ctx, span
}
