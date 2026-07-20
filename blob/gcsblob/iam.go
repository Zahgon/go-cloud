package gcsblob

import (
	"context"
	"sync"

	"cloud.google.com/go/iam/credentials/apiv1/credentialspb"
	gax "github.com/googleapis/gax-go/v2"
)

type credentialsClient struct {
	init sync.Once
	err  error

	client interface {
		SignBlob(context.Context, *credentialspb.SignBlobRequest, ...gax.CallOption) (*credentialspb.SignBlobResponse, error)
	}
}

func (c *credentialsClient) CreateMakeSignBytesWith(lifetimeCtx context.Context, googleAccessID string) func(context.Context) SignBytesFunc {
	_ = "STUB: not implemented"
	return nil
}
