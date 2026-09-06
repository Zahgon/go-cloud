package gcpfirestore

import (
	"context"
	"net/url"
	"sync"

	vkit "cloud.google.com/go/firestore/apiv1"
	"gocloud.dev/docstore"
)

func init() {
	docstore.DefaultURLMux().RegisterCollection(Scheme, &lazyCredsOpener{})
}

type lazyCredsOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *lazyCredsOpener) OpenCollectionURL(ctx context.Context, u *url.URL) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "firestore"

type URLOpener struct {
	Client *vkit.Client
}

func (o *URLOpener) OpenCollectionURL(ctx context.Context, u *url.URL) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
