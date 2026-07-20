package secrets

import (
	"context"
	"net/url"
	"sync"

	"gocloud.dev/internal/gcerr"
	"gocloud.dev/internal/openurl"
	gcdkotel "gocloud.dev/internal/otel"
	"gocloud.dev/secrets/driver"
)

type Keeper struct {
	k      driver.Keeper
	tracer *gcdkotel.Tracer

	mu     sync.RWMutex
	closed bool
}

var NewKeeper = newKeeper

func newKeeper(k driver.Keeper) *Keeper { _ = "STUB: not implemented"; return nil }

const pkgName = "gocloud.dev/secrets"

var (
	OpenTelemetryViews = gcdkotel.Views(pkgName)
)

func (k *Keeper) Encrypt(ctx context.Context, plaintext []byte) (ciphertext []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Keeper) Decrypt(ctx context.Context, ciphertext []byte) (plaintext []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var errClosed = gcerr.Newf(gcerr.FailedPrecondition, nil, "secrets: Keeper has been closed")

func (k *Keeper) Close() error { _ = "STUB: not implemented"; return nil }

func (k *Keeper) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func wrapError(k *Keeper, err error) error { _ = "STUB: not implemented"; return nil }

type KeeperURLOpener interface {
	OpenKeeperURL(ctx context.Context, u *url.URL) (*Keeper, error)
}

type URLMux struct {
	schemes openurl.SchemeMap
}

func (mux *URLMux) KeeperSchemes() []string { _ = "STUB: not implemented"; return nil }

func (mux *URLMux) ValidKeeperScheme(scheme string) bool { _ = "STUB: not implemented"; return false }

func (mux *URLMux) RegisterKeeper(scheme string, opener KeeperURLOpener) {
	_ = "STUB: not implemented"
	return
}

func (mux *URLMux) OpenKeeper(ctx context.Context, urlstr string) (*Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mux *URLMux) OpenKeeperURL(ctx context.Context, u *url.URL) (*Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var defaultURLMux = new(URLMux)

func DefaultURLMux() *URLMux { _ = "STUB: not implemented"; return nil }

func OpenKeeper(ctx context.Context, urlstr string) (*Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
