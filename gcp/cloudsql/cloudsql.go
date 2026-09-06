package cloudsql

import (
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/certs"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/proxy"
	"github.com/google/wire"
	"gocloud.dev/gcp"
	"golang.org/x/oauth2"
)

var CertSourceSet = wire.NewSet(
	NewCertSource,
	wire.Bind(new(proxy.CertSource), new(*certs.RemoteCertSource)))

func NewCertSource(c *gcp.HTTPClient) *certs.RemoteCertSource {
	_ = "STUB: not implemented"
	return nil
}

func NewCertSourceWithIAM(c *gcp.HTTPClient, t oauth2.TokenSource) *certs.RemoteCertSource {
	_ = "STUB: not implemented"
	return nil
}
