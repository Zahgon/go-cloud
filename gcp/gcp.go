package gcp

import (
	"context"
	"net/http"

	"github.com/google/wire"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var DefaultIdentity = wire.NewSet(
	CredentialsTokenSource,
	DefaultCredentials,
	DefaultProjectID)

type ProjectID string

type TokenSource oauth2.TokenSource

type HTTPClient struct {
	http.Client
}

func NewAnonymousHTTPClient(transport http.RoundTripper) *HTTPClient {
	_ = "STUB: not implemented"
	return nil
}

func NewHTTPClient(transport http.RoundTripper, ts TokenSource) (*HTTPClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DefaultTransport() http.RoundTripper {
	_ = "STUB: not implemented"
	return *new(http.RoundTripper)
}

func DefaultCredentials(ctx context.Context) (*google.Credentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DefaultCredentialsWithParams(ctx context.Context, params google.CredentialsParams) (*google.Credentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CredentialsTokenSource(creds *google.Credentials) TokenSource {
	_ = "STUB: not implemented"
	return *new(TokenSource)
}

func DefaultProjectID(creds *google.Credentials) (ProjectID, error) {
	_ = "STUB: not implemented"
	return *new(ProjectID), nil
}
