package httpvar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"gocloud.dev/gcerrors"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

func init() {
	o := &URLOpener{Client: http.DefaultClient}
	for _, scheme := range Schemes {
		runtimevar.DefaultURLMux().RegisterVariable(scheme, o)
	}
}

var Schemes = []string{"http", "https"}

type URLOpener struct {
	Client *http.Client

	Decoder *runtimevar.Decoder

	Options Options
}

func (o *URLOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Options struct {
	WaitDuration time.Duration
}

type RequestError struct {
	Response *http.Response
}

func (e *RequestError) Error() string { _ = "STUB: not implemented"; return "" }

func newRequestError(response *http.Response) *RequestError { _ = "STUB: not implemented"; return nil }

func OpenVariable(client *http.Client, urlStr string, decoder *runtimevar.Decoder, opts *Options) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type state struct {
	val        any
	raw        *http.Response
	rawBytes   []byte
	updateTime time.Time
	err        error
}

func (s *state) Value() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (s *state) UpdateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *state) As(i any) bool { _ = "STUB: not implemented"; return false }

func errorState(err error, prevS driver.State) driver.State {
	_ = "STUB: not implemented"
	return *new(driver.State)
}

func equivalentError(err1, err2 error) bool { _ = "STUB: not implemented"; return false }

type watcher struct {
	client   *http.Client
	endpoint *url.URL
	decoder  *runtimevar.Decoder
	wait     time.Duration
}

func (w *watcher) WatchVariable(ctx context.Context, prev driver.State) (driver.State, time.Duration) {
	_ = "STUB: not implemented"
	return *new(driver.State), *new(time.Duration)
}

func (w *watcher) Close() error { _ = "STUB: not implemented"; return nil }

func (w *watcher) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*watcher) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func newWatcher(client *http.Client, endpoint *url.URL, decoder *runtimevar.Decoder, opts *Options) driver.Watcher {
	_ = "STUB: not implemented"
	return *new(driver.Watcher)
}
