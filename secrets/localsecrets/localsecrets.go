package localsecrets

import (
	"context"
	"encoding/base64"
	"net/url"

	"gocloud.dev/gcerrors"
	"gocloud.dev/secrets"
)

func init() {
	secrets.DefaultURLMux().RegisterKeeper(Scheme, &URLOpener{})
}

const (
	Scheme = "base64key"
)

type URLOpener struct{}

func (o *URLOpener) OpenKeeperURL(ctx context.Context, u *url.URL) (*secrets.Keeper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type keeper struct {
	secretKey [32]byte
}

func NewKeeper(sk [32]byte) *secrets.Keeper { _ = "STUB: not implemented"; return nil }

func Base64KeyStd(base64str string) ([32]byte, error) {
	_ = "STUB: not implemented"
	return [32]byte{}, nil
}

func Base64Key(base64str string) ([32]byte, error) {
	_ = "STUB: not implemented"
	return [32]byte{}, nil
}

func base64Key(base64str string, encoding *base64.Encoding) ([32]byte, error) {
	_ = "STUB: not implemented"
	return [32]byte{}, nil
}

func NewRandomKey() ([32]byte, error) { _ = "STUB: not implemented"; return [32]byte{}, nil }

const nonceSize = 24

func (k *keeper) Encrypt(ctx context.Context, message []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *keeper) Decrypt(ctx context.Context, message []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *keeper) Close() error { _ = "STUB: not implemented"; return nil }

func (k *keeper) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (k *keeper) ErrorCode(error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
