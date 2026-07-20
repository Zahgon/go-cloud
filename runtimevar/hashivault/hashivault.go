package hashivault

import (
	"context"
	"net/url"
	"sync"
	"time"

	"github.com/hashicorp/vault/api"
	"gocloud.dev/gcerrors"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

func init() {
	runtimevar.DefaultURLMux().RegisterVariable(Scheme, new(defaultDialer))
}

const Scheme = "hashivault"

type SecretError struct {
	Code int

	Message string
}

func (e *SecretError) Error() string { _ = "STUB: not implemented"; return "" }

func newNotFoundError(path string) *SecretError { _ = "STUB: not implemented"; return nil }

func newInvalidDataError(path, reason string) *SecretError { _ = "STUB: not implemented"; return nil }

type Config struct {
	Token string

	APIConfig api.Config
}

func Dial(ctx context.Context, cfg *Config) (*api.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getVaultURL() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getVaultToken() string { _ = "STUB: not implemented"; return "" }

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
	Client *api.Client

	Decoder *runtimevar.Decoder

	Options Options
}

func (o *URLOpener) OpenVariableURL(ctx context.Context, u *url.URL) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Options struct {
	WaitDuration time.Duration

	EngineVersion int

	Mount string
}

func OpenVariable(client *api.Client, secretPath string, decoder *runtimevar.Decoder, opts *Options) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newWatcher(client *api.Client, secretPath string, decoder *runtimevar.Decoder, opts *Options) (driver.Watcher, error) {
	_ = "STUB: not implemented"
	return *new(driver.Watcher), nil
}

type state struct {
	val        any
	raw        *api.Secret
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
	client        *api.Client
	path          string
	decoder       *runtimevar.Decoder
	wait          time.Duration
	engineVersion int
	mount         string
}

func (w *watcher) WatchVariable(ctx context.Context, prev driver.State) (driver.State, time.Duration) {
	_ = "STUB: not implemented"
	return *new(driver.State), *new(time.Duration)
}

func (w *watcher) Close() error { _ = "STUB: not implemented"; return nil }

func (w *watcher) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (w *watcher) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
