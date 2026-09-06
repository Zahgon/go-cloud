package docstore

import (
	"context"
	"net/url"

	"gocloud.dev/internal/openurl"
)

type CollectionURLOpener interface {
	OpenCollectionURL(ctx context.Context, u *url.URL) (*Collection, error)
}

type URLMux struct {
	schemes openurl.SchemeMap
}

func (mux *URLMux) CollectionSchemes() []string { _ = "STUB: not implemented"; return nil }

func (mux *URLMux) ValidCollectionScheme(scheme string) bool {
	_ = "STUB: not implemented"
	return false
}

func (mux *URLMux) RegisterCollection(scheme string, opener CollectionURLOpener) {
	_ = "STUB: not implemented"
	return
}

func (mux *URLMux) OpenCollection(ctx context.Context, urlstr string) (*Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mux *URLMux) OpenCollectionURL(ctx context.Context, u *url.URL) (*Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var defaultURLMux = new(URLMux)

func DefaultURLMux() *URLMux { _ = "STUB: not implemented"; return nil }

func OpenCollection(ctx context.Context, urlstr string) (*Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
