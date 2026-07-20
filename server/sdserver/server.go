package sdserver

import (
	"context"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
	"gocloud.dev/server"

	"github.com/google/wire"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"gocloud.dev/gcp"
	"gocloud.dev/server/requestlog"
)

var Set = wire.NewSet(
	server.Set,
	NewTextMapPropagator,
	NewTraceSampler,
	NewTraceExporter,
	NewTraceProvider,
	wire.Bind(new(trace.TracerProvider), new(*sdktrace.TracerProvider)),
	NewMetricsReader,
	NewMeterProvider,
	wire.Bind(new(metric.MeterProvider), new(*sdkmetric.MeterProvider)),

	NewRequestLogger,
	wire.Bind(new(requestlog.Logger), new(*requestlog.StackdriverLogger)),
)

func NewResource(ctx context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTextMapPropagator() propagation.TextMapPropagator {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator)
}

func NewTraceSampler(ctx context.Context) sdktrace.Sampler {
	_ = "STUB: not implemented"
	return *new(sdktrace.Sampler)
}

func NewTraceExporter(projectID gcp.ProjectID) (sdktrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanExporter), nil
}

func NewTraceProvider(ctx context.Context, exporter sdktrace.SpanExporter, sampler sdktrace.Sampler) (*sdktrace.TracerProvider, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func NewMetricsReader(projectID gcp.ProjectID) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

func NewMeterProvider(ctx context.Context, reader sdkmetric.Reader) (*sdkmetric.MeterProvider, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func NewRequestLogger() *requestlog.StackdriverLogger { _ = "STUB: not implemented"; return nil }
