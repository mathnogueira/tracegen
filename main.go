package main

import (
	"context"
	"log"

	"github.com/mathnogueira/tracegen/generator"
)

func main() {
	executionGraph, err := generator.CreateExecutionGraph(generator.Config{
		NumberServices: 5,
		NumberSpans:    20,
		Collector: generator.CollectorConfig{
			Endpoint: "localhost:4317",
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	executionGraph.Execute(context.Background())
}
