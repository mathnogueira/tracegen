package generator

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
)

type Option func(context.Context) context.Context

func WithTraceStates(traceStates []string) Option {
	return func(ctx context.Context) context.Context {
		if len(traceStates) == 0 {
			return ctx
		}

		spanContext := trace.SpanContextFromContext(ctx)
		for _, tracestate := range traceStates {

			traceState, err := trace.ParseTraceState(tracestate)
			if err != nil {
				panic(fmt.Errorf(`invalid tracestate "%s": %w`, tracestate, err))
			}
			spanContext = spanContext.WithTraceState(traceState)
		}

		return trace.ContextWithSpanContext(ctx, spanContext)
	}
}
