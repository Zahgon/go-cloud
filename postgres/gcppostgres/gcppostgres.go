package gcppostgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net"
	"net/url"
	"sync"
	"time"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/proxy"
	"github.com/XSAM/otelsql"
	"gocloud.dev/postgres"
)

const Scheme = "gcppostgres"

func init() {
	postgres.DefaultURLMux().RegisterPostgres(Scheme, new(lazyCredsOpener))
}

type lazyCredsOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *lazyCredsOpener) OpenPostgresURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type URLOpener struct {
	CertSource proxy.CertSource

	TraceOpts []otelsql.Option
}

func (uo *URLOpener) OpenPostgresURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func instanceFromURL(u *url.URL) (instance, db string, _ error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

type pqDriver struct {
	client    *proxy.Client
	instance  string
	traceOpts []otelsql.Option
}

func (d pqDriver) Open(name string) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (d pqDriver) OpenConnector(name string) (driver.Connector, error) {
	_ = "STUB: not implemented"
	return *new(driver.Connector), nil
}

type connector struct {
	client    *proxy.Client
	instance  string
	pqConn    string
	traceOpts []otelsql.Option
}

func (c connector) Connect(context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (c connector) Driver() driver.Driver { _ = "STUB: not implemented"; return *new(driver.Driver) }

type dialer struct {
	client   *proxy.Client
	instance string
}

func (d dialer) Dial(network, address string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d dialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
