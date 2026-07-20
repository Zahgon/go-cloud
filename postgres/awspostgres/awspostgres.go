package awspostgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/XSAM/otelsql"
	"gocloud.dev/aws/rds"
	"gocloud.dev/postgres"
)

type URLOpener struct {
	HTTPClient *http.Client

	CertSource rds.CertPoolProvider

	TraceOpts []otelsql.Option
}

const Scheme = "awspostgres"

func init() {
	postgres.DefaultURLMux().RegisterPostgres(Scheme, &URLOpener{})
}

func (uo *URLOpener) OpenPostgresURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type pqDriver struct {
	provider  rds.CertPoolProvider
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
	provider  rds.CertPoolProvider
	pqConn    string
	traceOpts []otelsql.Option
	iam       func(context.Context) (string, error)
}

func (c connector) Connect(ctx context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (c connector) Driver() driver.Driver { _ = "STUB: not implemented"; return *new(driver.Driver) }

type dialer struct {
	provider rds.CertPoolProvider
}

func (d dialer) dial(ctx context.Context, network, address string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d dialer) Dial(network, address string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d dialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
