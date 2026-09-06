package azurekeyvault

import (
	"context"
	"net/url"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/azkeys"
	"github.com/google/wire"
	"gocloud.dev/gcerrors"
	"gocloud.dev/secrets"
)

var errorCodeMap = map[int]gcerrors.ErrorCode{
	200: gcerrors.OK,
	400: gcerrors.InvalidArgument,
	401: gcerrors.PermissionDenied,
	403: gcerrors.PermissionDenied,
	404: gcerrors.NotFound,
	408: gcerrors.DeadlineExceeded,
	429: gcerrors.ResourceExhausted,
	500: gcerrors.Internal,
	501: gcerrors.Unimplemented,
}

func init() {
	secrets.DefaultURLMux().RegisterKeeper(Scheme, new(defaultDialer))
}

var Set = wire.NewSet(
	DefaultClientMaker,
	wire.Struct(new(URLOpener), "Client"),
)

type ClientMakerT func(keyVaultURI string) (*azkeys.Client, error)

type defaultDialer struct{}

func (o *defaultDialer) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "azurekeyvault"

type URLOpener struct {
	ClientMaker ClientMakerT

	Options KeeperOptions
}

func (o *URLOpener) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type keeper struct {
	client      *azkeys.Client
	keyVaultURI string
	keyName     string
	keyVersion  string
	options     *KeeperOptions
}

type KeeperOptions struct {
	Algorithm azkeys.EncryptionAlgorithm

	EncryptOptions *azkeys.EncryptOptions

	DecryptOptions *azkeys.DecryptOptions
}

func DefaultClientMaker(keyVaultURI string) (*azkeys.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var keyIDRE = regexp.MustCompile(`^(https://.+\.vault\.(?:[a-z\d-.]+)/)keys/(.+)$`)

func OpenKeeper(clientMaker ClientMakerT, keyID string, opts *KeeperOptions) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openKeeper(clientMaker ClientMakerT, keyID string, opts *KeeperOptions) (*keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *keeper) Encrypt(ctx context.Context, plaintext []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *keeper) Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *keeper) Close() error { _ = "STUB: not implemented"; return nil }

func (k *keeper) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (k *keeper) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
