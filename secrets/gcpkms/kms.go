package gcpkms

import (
	"context"
	"net/url"
	"sync"

	cloudkms "cloud.google.com/go/kms/apiv1"
	"github.com/google/wire"
	"gocloud.dev/gcerrors"
	"gocloud.dev/gcp"
	"gocloud.dev/secrets"
)

const endPoint = "cloudkms.googleapis.com:443"

func Dial(ctx context.Context, ts gcp.TokenSource) (*cloudkms.KeyManagementClient, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func init() {
	secrets.DefaultURLMux().RegisterKeeper(Scheme, new(lazyCredsOpener))
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

func (o *lazyCredsOpener) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "gcpkms"

type URLOpener struct {
	Client *cloudkms.KeyManagementClient

	Options KeeperOptions
}

func (o *URLOpener) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenKeeper(client *cloudkms.KeyManagementClient, keyResourceID string, opts *KeeperOptions) *secrets.Keeper {
	_ = "STUB: not implemented"
	return nil
}

func KeyResourceID(projectID, location, keyRing, key string) string {
	_ = "STUB: not implemented"
	return ""
}

type keeper struct {
	keyResourceID string
	client        *cloudkms.KeyManagementClient
	opts          KeeperOptions
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

func (k *keeper) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

type KeeperOptions struct {
	AdditionalAuthenticatedData []byte
}
