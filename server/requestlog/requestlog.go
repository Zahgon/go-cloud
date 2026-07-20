package requestlog

import (
	"bufio"
	"io"
	"net"
	"net/http"
	"time"

	"go.opentelemetry.io/otel/trace"
)

type Logger interface {
	Log(*Entry)
}

type Handler struct {
	log Logger
	h   http.Handler
}

func NewHandler(log Logger, h http.Handler) *Handler { _ = "STUB: not implemented"; return nil }

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func cloneRequestWithoutBody(r *http.Request) *http.Request { _ = "STUB: not implemented"; return nil }

type Entry struct {
	Request *http.Request

	ReceivedTime    time.Time
	RequestBodySize int64

	Status             int
	ResponseHeaderSize int64
	ResponseBodySize   int64
	Latency            time.Duration
	TraceID            trace.TraceID
	SpanID             trace.SpanID

	Referer string

	Proto string

	RequestMethod string

	RequestURL string

	RequestHeaderSize int64

	UserAgent string

	RemoteIP string

	ServerIP string
}

func ipFromHostPort(hp string) string { _ = "STUB: not implemented"; return "" }

type readCounterCloser struct {
	r   io.ReadCloser
	n   int64
	err error
}

func (rcc *readCounterCloser) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rcc *readCounterCloser) Close() error { _ = "STUB: not implemented"; return nil }

type writeCounter int64

func (wc *writeCounter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func headerSize(h http.Header) int64 { _ = "STUB: not implemented"; return 0 }

type responseStats struct {
	w        http.ResponseWriter
	hsize    int64
	wc       writeCounter
	code     int
	hijacked bool
}

func (r *responseStats) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (r *responseStats) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

func (r *responseStats) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *responseStats) size() (hdr, body int64) { _ = "STUB: not implemented"; return 0, 0 }

func (r *responseStats) Hijack() (_ net.Conn, _ *bufio.ReadWriter, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

func (r *responseStats) Flush() { _ = "STUB: not implemented"; return }
