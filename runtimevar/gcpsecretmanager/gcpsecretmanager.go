package gcpsecretmanager

import (
	"context"
	"net/url"
	"regexp"
	"sync"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/google/wire"
	"gocloud.dev/gcerrors"
	"gocloud.dev/gcp"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

func Dial(ctx context.Context, ts gcp.TokenSource) (*secretmanager.Client, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
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

const Scheme = "gcpsecretmanager"

type URLOpener struct {
	Client *secretmanager.Client

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

func OpenVariable(client *secretmanager.Client, secretKey string, decoder *runtimevar.Decoder, opts *Options) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var secretKeyRE = regexp.MustCompile(`^projects/[a-z][a-z0-9_\-]{4,28}[a-z0-9_]/secrets/[a-zA-Z0-9_\-]{1,255}$`)

const latestVersion = "/versions/latest"

func newWatcher(client *secretmanager.Client, secretKey string, decoder *runtimevar.Decoder, opts *Options) (driver.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(driver.Watcher), nil
}

func SecretKey(projectID gcp.ProjectID, secretID string) string {
	_ = "STUB: not implemented"
	return ""
}

type state struct {
	val        any
	raw        *secretmanagerpb.AccessSecretVersionResponse
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
	client  *secretmanager.Client
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
