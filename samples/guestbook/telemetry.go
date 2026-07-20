package main

import (
	"context"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var otelTracesProviderSet = wire.NewSet(
	NewTraceProvider,
	wire.Bind(new(trace.TracerProvider), new(*sdktrace.TracerProvider)),
)

var otelMetricsProviderSet = wire.NewSet(
	NewMeterProvider,
	wire.Bind(new(metric.MeterProvider), new(*sdkmetric.MeterProvider)),
)

var OtelLogsSet = wire.NewSet(
	NewLogsExporter,
	NewLoggerProvider,
	wire.Bind(new(log.LoggerProvider), new(*sdklog.LoggerProvider)),
)

func newResource() *resource.Resource { _ = "STUB: not implemented"; return nil }

func newPropagationTextMap() propagation.TextMapPropagator {
	_ = "STUB: not implemented"
	return *new(propagation.TextMapPropagator)
}

func newTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanExporter), nil
}

func newTraceSampler(ctx context.Context) sdktrace.Sampler {
	_ = "STUB: not implemented"
	return *new(sdktrace.Sampler)
}

func NewTraceProvider(ctx context.Context, exporter sdktrace.SpanExporter, sampler sdktrace.Sampler) (*sdktrace.TracerProvider, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newMetricsReader(ctx context.Context) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

func NewMeterProvider(ctx context.Context, reader sdkmetric.Reader) (*sdkmetric.MeterProvider, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLogsExporter(ctx context.Context) (sdklog.Exporter, error) {
	_ = "STUB: not implemented"
	return *new(sdklog.Exporter), nil
}

func NewLoggerProvider(ctx context.Context, res *resource.Resource, exporter sdklog.Exporter) (*sdklog.LoggerProvider, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
