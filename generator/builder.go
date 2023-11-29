package generator

import (
	"github.com/kubeshop/tracegen/generator/application"
)

func CreateExecutionGraph(config Config) (*ExecutionGraph, error) {
	numberSpansPerService := int(config.NumberSpans / config.NumberServices)
	services := make([]*application.Service, 0, config.NumberServices)
	for i := 0; i < int(config.NumberServices); i++ {
		services = append(services, application.NewService(
			application.WithCollectorEndpoint(config.Collector.Endpoint),
		))
	}

	flatten := make([]*application.Operation, 0)
	for _, service := range services {
		for i := 0; i < numberSpansPerService; i++ {
			operation := application.NewOperation(service.Domain, service.Tracer)
			service.Operations = append(service.Operations, operation)
			flatten = append(flatten, operation)
		}
	}

	executionGraph := newExecutionGraph(flatten)
	return executionGraph, nil
}
