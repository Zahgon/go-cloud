package gcpruntimeconfig

import (
	"context"
	"net/url"
	"regexp"
	"sync"
	"time"

	"github.com/google/wire"
	"gocloud.dev/gcerrors"
	"gocloud.dev/gcp"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
	pb "google.golang.org/genproto/googleapis/cloud/runtimeconfig/v1beta1"
)

const (
	endPoint = "runtimeconfig.googleapis.com:443"
)

func Dial(ctx context.Context, ts gcp.TokenSource) (pb.RuntimeConfigManagerClient, func(), error) {
	_ = "STUB: not implemented"
	return *new(pb.RuntimeConfigManagerClient), nil, nil
}

func init() {
	runtimevar.DefaultURLMux().RegisterVariable(Scheme, new(lazyCredsOpener))
}

var Set = wire.NewSet(
	Dial,
	wire.Struct(new(URLOpener), "Client"),
)

type lazyCredsOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *lazyCredsOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "gcpruntimeconfig"

type URLOpener struct {
	Client pb.RuntimeConfigManagerClient

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

func OpenVariable(client pb.RuntimeConfigManagerClient, variableKey string, decoder *runtimevar.Decoder, opts *Options) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var variableKeyRE = regexp.MustCompile("^projects/.+/configs/.+/variables/.+$")

func newWatcher(client pb.RuntimeConfigManagerClient, variableKey string, decoder *runtimevar.Decoder, opts *Options) (driver.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(driver.Watcher), nil
}

func VariableKey(projectID gcp.ProjectID, configID, variableName string) string {
	_ = "STUB: not implemented"
	return ""
}

type state struct {
	val        any
	raw        *pb.Variable
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

func equivalentError(err1, err2 error) bool { _ = "STUB: not implemented"; return false }

type watcher struct {
	client  pb.RuntimeConfigManagerClient
	wait    time.Duration
	name    string
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

func bytesFromProto(vpb *pb.Variable) []byte { _ = "STUB: not implemented"; return nil }

func parseUpdateTime(vpb *pb.Variable) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}
