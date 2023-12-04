package cmd

import (
	"context"
	"log"

	"github.com/mathnogueira/tracegen/generator"
)

func run(ctx context.Context) {
	executionGraph, err := generator.CreateExecutionGraph(generator.Config{
		NumberServices: uint(services),
		NumberSpans:    uint(minSpans),
		Collector: generator.CollectorConfig{
			Endpoint: collectorEndpoint,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	executionGraph.Execute(
		ctx,
		generator.WithTraceStates(traceState),
	)
}
