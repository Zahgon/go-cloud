package etcdvar

import (
	"context"
	"errors"
	"net/url"
	"sync"
	"time"

	"gocloud.dev/gcerrors"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

func init() {
	runtimevar.DefaultURLMux().RegisterVariable(Scheme, &defaultDialer{})
}

const Scheme = "etcd"

type defaultDialer struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *defaultDialer) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type URLOpener struct {
	Client *clientv3.Client

	Decoder *runtimevar.Decoder

	Options Options
}

func (o *URLOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Options struct {
	Timeout time.Duration
}

func OpenVariable(cli *clientv3.Client, name string, decoder *runtimevar.Decoder, opts *Options) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWatcher(cli *clientv3.Client, name string, decoder *runtimevar.Decoder, opts *Options) *watcher {
	_ = "STUB: not implemented"
	return nil
}

var errNotExist = errors.New("variable does not exist")

type state struct {
	val        any
	raw        *clientv3.GetResponse
	updateTime time.Time
	version    int64
	err        error
}

func (s *state) Value() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (s *state) UpdateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *state) As(i any) bool { _ = "STUB: not implemented"; return false }

type watcher struct {
	ch chan *state

	shutdown func()
}

func (w *watcher) WatchVariable(ctx context.Context, _ driver.State) (driver.State, time.Duration) {
	_ = "STUB: not implemented"
	return *new(driver.State), *new(time.Duration)
}

func (w *watcher) updateState(s, prev *state) *state { _ = "STUB: not implemented"; return nil }

func equivalentError(err1, err2 error) bool { _ = "STUB: not implemented"; return false }

func (w *watcher) watch(ctx context.Context, cli *clientv3.Client, name string, decoder *runtimevar.Decoder, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (w *watcher) Close() error { _ = "STUB: not implemented"; return nil }

func (w *watcher) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*watcher) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
