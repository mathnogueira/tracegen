package telemetry

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetTracer(ctx context.Context, config Config) (trace.Tracer, error) {
	err := config.Validate()
	if err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}
	propagator := propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{})
	otel.SetTextMapPropagator(propagator)

	provider, err := getTraceProvider(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("could not get trace provider: %w", err)
	}

	tracer := provider.Tracer(config.TracerName)
	return tracer, nil
}

func getTraceProvider(ctx context.Context, config Config) (*sdktrace.TracerProvider, error) {
	exporter, err := getCollectorExporter(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("could not get collector exporter: %w", err)
	}

	r, err := getResources(ctx, config)
	if err != nil {
		exporter.Shutdown(ctx)
		return nil, fmt.Errorf("could not get tracer resources")
	}

	processor := sdktrace.NewBatchSpanProcessor(exporter, sdktrace.WithBatchTimeout(100*time.Millisecond))

	return sdktrace.NewTracerProvider(
		sdktrace.WithResource(r),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(1.0)),
		sdktrace.WithSpanProcessor(processor),
	), nil
}

func getResources(ctx context.Context, config Config) (*resource.Resource, error) {
	resource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(config.ServiceName),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("could not merge resources: %w", err)
	}

	return resource, nil
}

func getCollectorExporter(ctx context.Context, config Config) (sdktrace.SpanExporter, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, config.CollectorEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("could not connect to collector: %w", err)
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("could not create trace exporter: %w", err)
	}

	return exporter, nil
}
