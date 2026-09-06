package awsmysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net/http"
	"net/url"

	"github.com/XSAM/otelsql"
	"github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gocloud.dev/aws/rds"
	gcmysql "gocloud.dev/mysql"
)

var Set = wire.NewSet(
	wire.Struct(new(URLOpener), "CertSource", "HTTPClient"),
	rds.CertFetcherSet,
)

type URLOpener struct {
	HTTPClient *http.Client

	CertSource rds.CertPoolProvider

	TraceOpts []otelsql.Option
}

const Scheme = "awsmysql"

func init() {
	gcmysql.DefaultURLMux().RegisterMySQL(Scheme, &URLOpener{})
}

func (uo *URLOpener) OpenMySQLURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type connector struct {
	sem      chan struct{}
	provider CertPoolProvider

	ready chan struct{}
	cfg   *mysql.Config
	iam   func(context.Context) (string, error)
}

func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (c *connector) Driver() driver.Driver { _ = "STUB: not implemented"; return *new(driver.Driver) }

type CertPoolProvider = rds.CertPoolProvider

type CertFetcher = rds.CertFetcher
