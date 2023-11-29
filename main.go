package main

import (
	"context"
	"log"

	"github.com/kubeshop/tracegen/generator"
)

func main() {
	executionGraph, err := generator.CreateExecutionGraph(generator.Config{
		NumberServices: 10,
		NumberSpans:    150,
		Collector: generator.CollectorConfig{
			Endpoint: "localhost:4317",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	executionGraph.Execute(context.Background())
}
