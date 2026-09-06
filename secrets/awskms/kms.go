package awskms

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/google/wire"
	"gocloud.dev/gcerrors"
	"gocloud.dev/secrets"
)

func init() {
	secrets.DefaultURLMux().RegisterKeeper(Scheme, new(lazySessionOpener))
}

var Set = wire.NewSet(
	Dial,
)

func Dial(cfg aws.Config) (*kms.Client, error) { _ = "STUB: not implemented"; return nil, nil }

var DialV2 = Dial

type lazySessionOpener struct{}

func (o *lazySessionOpener) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "awskms"

type URLOpener struct {
	Options KeeperOptions
}

func addEncryptionContextFromURLParams(opts *KeeperOptions, u url.Values) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *URLOpener) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenKeeper(client *kms.Client, keyID string, opts *KeeperOptions) *secrets.Keeper {
	_ = "STUB: not implemented"
	return nil
}

var OpenKeeperV2 = OpenKeeper

type keeper struct {
	keyID  string
	opts   KeeperOptions
	client *kms.Client
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

var errorCodeMap = map[string]gcerrors.ErrorCode{
	(&types.NotFoundException{}).ErrorCode():          gcerrors.NotFound,
	(&types.InvalidCiphertextException{}).ErrorCode(): gcerrors.InvalidArgument,
	(&types.InvalidKeyUsageException{}).ErrorCode():   gcerrors.InvalidArgument,
	(&types.KMSInternalException{}).ErrorCode():       gcerrors.Internal,
	(&types.KMSInvalidStateException{}).ErrorCode():   gcerrors.FailedPrecondition,
	(&types.DisabledException{}).ErrorCode():          gcerrors.PermissionDenied,
	(&types.InvalidGrantTokenException{}).ErrorCode(): gcerrors.PermissionDenied,
	(&types.KeyUnavailableException{}).ErrorCode():    gcerrors.ResourceExhausted,
	(&types.DependencyTimeoutException{}).ErrorCode(): gcerrors.DeadlineExceeded,
}

type KeeperOptions struct {
	EncryptionContext map[string]string
}
