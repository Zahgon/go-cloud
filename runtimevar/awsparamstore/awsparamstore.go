package awsparamstore

import (
	"context"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/google/wire"
	"gocloud.dev/gcerrors"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

func init() {
	runtimevar.DefaultURLMux().RegisterVariable(Scheme, new(lazySessionOpener))
}

var Set = wire.NewSet(
	Dial,
)

func Dial(cfg aws.Config) *ssm.Client { _ = "STUB: not implemented"; return nil }

type URLOpener struct {
	Decoder *runtimevar.Decoder

	Options Options
}

type lazySessionOpener struct{}

func (o *lazySessionOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "awsparamstore"

func (o *URLOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Options struct {
	WaitDuration time.Duration
}

func OpenVariable(client *ssm.Client, name string, decoder *runtimevar.Decoder, opts *Options) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var OpenVariableV2 = OpenVariable

func newWatcher(client *ssm.Client, name string, decoder *runtimevar.Decoder, opts *Options) *watcher {
	_ = "STUB: not implemented"
	return nil
}

type state struct {
	val        any
	rawGet     *ssm.GetParameterOutput
	updateTime time.Time
	version    int64
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
	client *ssm.Client

	name string

	wait time.Duration

	decoder *runtimevar.Decoder
}

func getParameter(ctx context.Context, client *ssm.Client, name string) (int64, []byte, time.Time, *ssm.GetParameterOutput, error) {
	_ = "STUB: not implemented"
	return 0, nil, *new(time.Time), nil, nil
}

func (w *watcher) WatchVariable(ctx context.Context, prev driver.State) (driver.State, time.Duration) {
	_ = "STUB: not implemented"
	return *new(driver.State), *new(time.Duration)
}

func (w *watcher) Close() error { _ = "STUB: not implemented"; return nil }

func (w *watcher) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func getErrorCode(err error) string { _ = "STUB: not implemented"; return "" }

func (w *watcher) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
