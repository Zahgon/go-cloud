package constantvar

import (
	"context"
	"errors"
	"net/url"
	"time"

	"gocloud.dev/gcerrors"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

func init() {
	runtimevar.DefaultURLMux().RegisterVariable(Scheme, &URLOpener{})
}

const Scheme = "constant"

type URLOpener struct {
	Decoder *runtimevar.Decoder
}

func (o *URLOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var errNotExist = errors.New("variable does not exist")

func New(value any) *runtimevar.Variable { _ = "STUB: not implemented"; return nil }

func NewBytes(b []byte, decoder *runtimevar.Decoder) *runtimevar.Variable {
	_ = "STUB: not implemented"
	return nil
}

func NewFromEnv(envVarName string, decoder *runtimevar.Decoder) *runtimevar.Variable {
	_ = "STUB: not implemented"
	return nil
}

func NewError(err error) *runtimevar.Variable { _ = "STUB: not implemented"; return nil }

type watcher struct {
	value any
	err   error
	t     time.Time
}

func (w *watcher) Value() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (w *watcher) UpdateTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (w *watcher) As(i any) bool { _ = "STUB: not implemented"; return false }

func (w *watcher) WatchVariable(ctx context.Context, prev driver.State) (driver.State, time.Duration) {
	_ = "STUB: not implemented"
	return *new(driver.State), *new(time.Duration)
}

func (*watcher) Close() error { _ = "STUB: not implemented"; return nil }

func (*watcher) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*watcher) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
