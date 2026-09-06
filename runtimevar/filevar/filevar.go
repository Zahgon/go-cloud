package filevar

import (
	"context"
	"net/url"
	"time"

	"github.com/fsnotify/fsnotify"
	"gocloud.dev/gcerrors"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

func init() {
	runtimevar.DefaultURLMux().RegisterVariable(Scheme, &URLOpener{})
}

const Scheme = "file"

type URLOpener struct {
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

func OpenVariable(path string, decoder *runtimevar.Decoder, opts *Options) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWatcher(path string, decoder *runtimevar.Decoder, opts *Options) (*watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type errNotExist struct {
	err error
}

func (e *errNotExist) Error() string { _ = "STUB: not implemented"; return "" }

type state struct {
	val        any
	updateTime time.Time
	raw        []byte
	err        error
}

func (s *state) Value() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (s *state) UpdateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *state) As(i any) bool { _ = "STUB: not implemented"; return false }

type watcher struct {
	path string

	ch chan *state

	closeCh chan error

	shutdown func()
}

func (w *watcher) WatchVariable(ctx context.Context, _ driver.State) (driver.State, time.Duration) {
	_ = "STUB: not implemented"
	return *new(driver.State), *new(time.Duration)
}

func (w *watcher) updateState(s, prev *state) *state { _ = "STUB: not implemented"; return nil }

func (w *watcher) watch(ctx context.Context, notifier *fsnotify.Watcher, file string, decoder *runtimevar.Decoder, wait time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (w *watcher) Close() error { _ = "STUB: not implemented"; return nil }

func (w *watcher) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*watcher) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
