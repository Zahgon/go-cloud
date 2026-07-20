package azuredb

import (
	"context"
	"crypto/x509"
	"net/http"
)

const caBundleURL = "https://www.digicert.com/CACerts/BaltimoreCyberTrustRoot.crt.pem"

type CertPoolProvider interface {
	AzureCertPool(context.Context) (*x509.CertPool, error)
}

type CertFetcher struct {
	Client *http.Client
}

func (cf *CertFetcher) AzureCertPool(ctx context.Context) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cf *CertFetcher) Fetch(ctx context.Context) (certs []*x509.Certificate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
