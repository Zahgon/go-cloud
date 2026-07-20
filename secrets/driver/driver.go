package driver

import (
	"context"

	"gocloud.dev/gcerrors"
)

type Keeper interface {
	Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error)

	Encrypt(ctx context.Context, plaintext []byte) ([]byte, error)

	Close() error

	ErrorAs(err error, i any) bool

	ErrorCode(error) gcerrors.ErrorCode
}
