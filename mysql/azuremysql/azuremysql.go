package azuremysql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net/url"

	"github.com/XSAM/otelsql"
	"github.com/go-sql-driver/mysql"
	"gocloud.dev/azure/azuredb"
	cdkmysql "gocloud.dev/mysql"
)

type URLOpener struct {
	CertSource azuredb.CertPoolProvider

	TraceOpts []otelsql.Option
}

const Scheme = "azuremysql"

func init() {
	cdkmysql.DefaultURLMux().RegisterMySQL(Scheme, &URLOpener{})
}

func (uo *URLOpener) OpenMySQLURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type connector struct {
	sem      chan struct{}
	provider CertPoolProvider
	ready    chan struct{}
	cfg      *mysql.Config
}

func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (c *connector) Driver() driver.Driver { _ = "STUB: not implemented"; return *new(driver.Driver) }

type CertPoolProvider = azuredb.CertPoolProvider

type CertFetcher = azuredb.CertFetcher
