package gcpmysql

import (
	"context"
	"database/sql"
	"net/url"
	"sync"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/proxy"
	"github.com/XSAM/otelsql"
	"github.com/go-sql-driver/mysql"
	cdkmysql "gocloud.dev/mysql"
)

const Scheme = "gcpmysql"

func init() {
	cdkmysql.DefaultURLMux().RegisterMySQL(Scheme, new(lazyCredsOpener))
}

type lazyCredsOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *lazyCredsOpener) OpenMySQLURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type URLOpener struct {
	CertSource proxy.CertSource

	TraceOpts []otelsql.Option
}

func (uo *URLOpener) OpenMySQLURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func configFromURL(u *url.URL) (*mysql.Config, error) { _ = "STUB: not implemented"; return nil, nil }

func instanceFromURL(u *url.URL) (instance, db string, _ error) {
	_ = "STUB: not implemented"
	return "", "", nil
}
