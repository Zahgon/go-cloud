package mongodocstore

import (
	"context"
	"net/url"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"gocloud.dev/docstore"
)

func init() {
	docstore.DefaultURLMux().RegisterCollection(Scheme, new(defaultDialer))
}

type defaultDialer struct {
	mongoServerURL string
	mu             sync.Mutex
	opener         *URLOpener
	err            error
}

func (o *defaultDialer) OpenCollectionURL(ctx context.Context, u *url.URL) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "mongo"

type URLOpener struct {
	Client *mongo.Client

	Options Options
}

func (o *URLOpener) OpenCollectionURL(ctx context.Context, u *url.URL) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
