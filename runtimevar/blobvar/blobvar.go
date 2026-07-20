package blobvar

import (
	"context"
	"net/url"
	"sync"
	"time"

	"gocloud.dev/blob"
	"gocloud.dev/gcerrors"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

func init() {
	runtimevar.DefaultURLMux().RegisterVariable(Scheme, &defaultOpener{})
}

type defaultOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *defaultOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "blob"

type URLOpener struct {
	Bucket *blob.Bucket

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

func OpenVariable(bucket *blob.Bucket, key string, decoder *runtimevar.Decoder, opts *Options) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWatcher(bucket *blob.Bucket, key string, decoder *runtimevar.Decoder, opener *URLOpener, opts *Options) driver.Watcher {
	_ = "STUB: not implemented"
	return *new(driver.Watcher)
}

type state struct {
	val        any
	updateTime time.Time
	rawBytes   []byte
	err        error
}

func (s *state) Value() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (s *state) UpdateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (s *state) As(i any) bool { _ = "STUB: not implemented"; return false }

func errorState(err error, prevS driver.State) driver.State {
	_ = "STUB: not implemented"
	return *new(driver.State)
}

type watcher struct {
	bucket  *blob.Bucket
	opener  *URLOpener
	key     string
	wait    time.Duration
	decoder *runtimevar.Decoder
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
