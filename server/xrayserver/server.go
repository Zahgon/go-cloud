package xrayserver

import (
	"context"

	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"

	"github.com/google/wire"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var Set = wire.NewSet(
	server.Set,
	TracesSet,
	MetricsSet,

	NewRequestLogger,
	wire.Bind(new(requestlog.Logger), new(*requestlog.NCSALogger)),
)

var TracesSet = wire.NewSet(
	NewTextMapPropagator,
	wire.Bind(new(propagation.TextMapPropagator), new(*xray.Propagator)),
	NewTraceSampler,
	NewTraceExporter,
	NewTraceProvider,
	wire.Bind(new(trace.TracerProvider), new(*sdktrace.TracerProvider)),
)

var MetricsSet = wire.NewSet(
	NewMetricsReader,
	NewMeterProvider,
	wire.Bind(new(metric.MeterProvider), new(*sdkmetric.MeterProvider)),
)

func NewResource(ctx context.Context) (*resource.Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewTextMapPropagator() *xray.Propagator { _ = "STUB: not implemented"; return nil }

func NewTraceSampler() sdktrace.Sampler { _ = "STUB: not implemented"; return *new(sdktrace.Sampler) }

func NewTraceExporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanExporter), nil
}

func NewTraceProvider(ctx context.Context, exp sdktrace.SpanExporter, sampler sdktrace.Sampler) (*sdktrace.TracerProvider, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func NewMetricsReader(ctx context.Context) (sdkmetric.Reader, error) {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Reader), nil
}

func NewMeterProvider(ctx context.Context, reader sdkmetric.Reader) (*sdkmetric.MeterProvider, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func NewRequestLogger() *requestlog.NCSALogger { _ = "STUB: not implemented"; return nil }
