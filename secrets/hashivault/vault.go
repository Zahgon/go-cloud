package hashivault

import (
	"context"
	"net/url"
	"sync"

	"github.com/hashicorp/vault/api"
	"gocloud.dev/gcerrors"
	"gocloud.dev/secrets"
)

type Config struct {
	Token string

	APIConfig api.Config
}

func Dial(ctx context.Context, cfg *Config) (*api.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func init() {
	secrets.DefaultURLMux().RegisterKeeper(Scheme, new(defaultDialer))
}

func getVaultURL() (string, error) { _ = "STUB: not implemented"; return "", nil }

func getVaultToken() string { _ = "STUB: not implemented"; return "" }

type defaultDialer struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *defaultDialer) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "hashivault"

type URLOpener struct {
	Client *api.Client

	Options KeeperOptions
}

func (o *URLOpener) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newKeeper(client *api.Client, keyID string, opts *KeeperOptions) *keeper {
	_ = "STUB: not implemented"
	return nil
}

func OpenKeeper(client *api.Client, keyID string, opts *KeeperOptions) *secrets.Keeper {
	_ = "STUB: not implemented"
	return nil
}

type keeper struct {
	keyID  string
	client *api.Client
	opts   KeeperOptions
}

func (k *keeper) Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *keeper) Encrypt(ctx context.Context, plaintext []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *keeper) Close() error { _ = "STUB: not implemented"; return nil }

func (k *keeper) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (k *keeper) ErrorCode(error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

type KeeperOptions struct {
	Engine string
}
