package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"go.opentelemetry.io/otel"

	"gocloud.dev/gcp"
	"gocloud.dev/server"
	"gocloud.dev/server/health"
	"gocloud.dev/server/sdserver"
)

type GlobalMonitoredResource struct {
	projectID string
}

func (g GlobalMonitoredResource) MonitoredResource() (string, map[string]string) {
	_ = "STUB: not implemented"
	return "", nil
}

func helloHandler(w http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func mainHandler(w http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

type customHealthCheck struct {
	mu      sync.RWMutex
	healthy bool
}

func (h *customHealthCheck) CheckHealth() error { _ = "STUB: not implemented"; return nil }

func main() {
	addr := flag.String("listen", ":8080", "HTTP port to listen on")
	doTrace := flag.Bool("trace", true, "Export traces to Stackdriver")
	flag.Parse()

	ctx := context.Background()
	credentials, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		log.Fatal(err)
	}

	projectID, err := gcp.DefaultProjectID(credentials)
	if err != nil {
		log.Fatal(err)
	}

	if *doTrace {
		fmt.Println("Exporting traces to Stackdriver")

		traceSampler := sdserver.NewTraceSampler(ctx)

		spanExporter, err0 := sdserver.NewTraceExporter(projectID)
		if err0 != nil {
			log.Fatal(err0)
		}
		tp, cleanup, err0 := sdserver.NewTraceProvider(ctx, spanExporter, traceSampler)
		if err0 != nil {
			log.Fatal(err0)
		}
		defer cleanup()
		otel.SetTracerProvider(tp)

		metricsReader, err0 := sdserver.NewMetricsReader(projectID)
		if err0 != nil {
			log.Fatal(err0)
		}
		mp, cleanup2, err0 := sdserver.NewMeterProvider(ctx, metricsReader)
		if err0 != nil {
			log.Fatal(err0)
		}
		defer cleanup2()
		otel.SetMeterProvider(mp)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", mainHandler)

	healthCheck := new(customHealthCheck)
	time.AfterFunc(10*time.Second, func() {
		healthCheck.mu.Lock()
		defer healthCheck.mu.Unlock()
		healthCheck.healthy = true
	})

	options := &server.Options{
		RequestLogger: sdserver.NewRequestLogger(),
		HealthChecks:  []health.Checker{healthCheck},

		Driver: &server.DefaultDriver{},
	}

	s := server.New(mux, options)
	fmt.Printf("Listening on %s\n", *addr)

	err = s.ListenAndServe(*addr)
	if err != nil {
		log.Fatal(err)
	}
}
