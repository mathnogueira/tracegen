package cmd

import (
	"context"
	"log"
	"time"

	"github.com/mathnogueira/tracegen/generator"
)

type runOption func(*generator.Config)

func WithDelay(delay time.Duration) runOption {
	return func(c *generator.Config) {
		c.Delay = delay
	}
}

func run(ctx context.Context, opts ...runOption) {
	config := generator.Config{
		NumberServices: uint(services),
		NumberSpans:    uint(minSpans),
		Collector: generator.CollectorConfig{
			Endpoint: collectorEndpoint,
			Insecure: insecure,
		},
	}

	for _, opt := range opts {
		opt(&config)
	}

	executionGraph, err := generator.CreateExecutionGraph(config)
	if err != nil {
		log.Fatal(err)
	}

	executionGraph.Execute(
		ctx,
		generator.WithTraceStates(traceState),
	)
}
