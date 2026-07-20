package server

import (
	"context"
	"net/http"
	"sync"

	"github.com/google/wire"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"gocloud.dev/server/driver"
	"gocloud.dev/server/health"
	"gocloud.dev/server/requestlog"
)

var Set = wire.NewSet(
	New,
	wire.Struct(new(Options), "RequestLogger", "HealthChecks",
		"TraceTextMapPropagator", "TraceProvider", "MetricsProvider", "Driver"),
	wire.Value(&DefaultDriver{}),
	wire.Bind(new(driver.Server), new(*DefaultDriver)),
)

type Server struct {
	reqlog            requestlog.Logger
	handler           http.Handler
	wrappedHandler    http.Handler
	healthHandler     health.Handler
	textMapPropagator propagation.TextMapPropagator
	traceProvider     trace.TracerProvider
	meterProvider     metric.MeterProvider
	once              sync.Once
	driver            driver.Server
}

type Options struct {
	RequestLogger requestlog.Logger

	HealthChecks []health.Checker

	TraceTextMapPropagator propagation.TextMapPropagator

	TraceProvider trace.TracerProvider

	MetricsProvider metric.MeterProvider

	Driver driver.Server
}

func New(h http.Handler, opts *Options) *Server { _ = "STUB: not implemented"; return nil }

func (srv *Server) init() {
	srv.once.Do(func() {

		if srv.textMapPropagator != nil {
			otel.SetTextMapPropagator(srv.textMapPropagator)
		}

		if srv.traceProvider != nil {
			otel.SetTracerProvider(srv.traceProvider)
		}

		if srv.meterProvider != nil {
			otel.SetMeterProvider(srv.meterProvider)
		}

		if srv.driver == nil {
			srv.driver = NewDefaultDriver()
		}
		if srv.handler == nil {
			srv.handler = http.DefaultServeMux
		}

		const healthPrefix = "/healthz/"

		mux := http.NewServeMux()
		mux.HandleFunc(healthPrefix+"liveness", health.HandleLive)
		mux.Handle(healthPrefix+"readiness", &srv.healthHandler)
		h := srv.handler
		if srv.reqlog != nil {
			h = requestlog.NewHandler(srv.reqlog, h)
		}

		h = otelhttp.NewHandler(h, "", otelhttp.WithPublicEndpointFn(func(*http.Request) bool { return true }))
		mux.Handle("/", h)
		srv.wrappedHandler = mux
	})
}

func (srv *Server) ListenAndServe(addr string) error { _ = "STUB: not implemented"; return nil }

func (srv *Server) ListenAndServeTLS(addr, certFile, keyFile string) error {
	_ = "STUB: not implemented"
	return nil
}

func (srv *Server) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type DefaultDriver struct {
	Server http.Server
}

func NewDefaultDriver() *DefaultDriver { _ = "STUB: not implemented"; return nil }

func (dd *DefaultDriver) ListenAndServe(addr string, h http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

func (dd *DefaultDriver) ListenAndServeTLS(addr, certFile, keyFile string, h http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

func (dd *DefaultDriver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
