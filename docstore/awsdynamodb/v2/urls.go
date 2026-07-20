package awsdynamodb

import (
	"context"
	"net/url"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	dyn "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"gocloud.dev/docstore"
)

func init() {
	docstore.DefaultURLMux().RegisterCollection(Scheme, new(lazySessionOpener))
}

type lazySessionOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *lazySessionOpener) OpenCollectionURL(ctx context.Context, u *url.URL) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "dynamodb"

type URLOpener struct {
}

func (o *URLOpener) OpenCollectionURL(_ context.Context, u *url.URL) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *URLOpener) processURL(u *url.URL) (db *dyn.Client, tableName, partitionKey, sortKey string, opts *Options, err error) {
	_ = "STUB: not implemented"
	return nil, "", "", "", nil, nil
}

func Dial(p aws.Config) (*dyn.Client, error) { _ = "STUB: not implemented"; return nil, nil }
