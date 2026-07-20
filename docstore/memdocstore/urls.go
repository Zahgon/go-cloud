package memdocstore

import (
	"context"
	"net/url"
	"sync"

	"gocloud.dev/docstore"
)

func init() {
	docstore.DefaultURLMux().RegisterCollection(Scheme, &URLOpener{})
}

const Scheme = "mem"

type URLOpener struct {
	mu          sync.Mutex
	collections map[string]urlColl
}

type urlColl struct {
	keyName string
	coll    *docstore.Collection
}

func (o *URLOpener) OpenCollectionURL(ctx context.Context, u *url.URL) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
