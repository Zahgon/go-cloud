package rds

import (
	"context"
	"crypto/x509"
	"net/http"

	"github.com/google/wire"
)

var CertFetcherSet = wire.NewSet(
	wire.Struct(new(CertFetcher), "Client"),
	wire.Bind(new(CertPoolProvider), new(*CertFetcher)),
)

type CertPoolProvider interface {
	RDSCertPool(context.Context) (*x509.CertPool, error)
}

const caBundleURL = "https://truststore.pki.rds.amazonaws.com/global/global-bundle.pem"

type CertFetcher struct {
	Client *http.Client
}

func (cf *CertFetcher) RDSCertPool(ctx context.Context) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cf *CertFetcher) Fetch(ctx context.Context) (certs []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
