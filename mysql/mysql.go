package mysql

import (
	"context"
	"database/sql"
	"net/url"
	"regexp"

	"github.com/XSAM/otelsql"
	"github.com/go-sql-driver/mysql"
	"gocloud.dev/internal/openurl"
)

const Scheme = "mysql"

func init() {
	DefaultURLMux().RegisterMySQL(Scheme, &URLOpener{})
}

type URLOpener struct {
	TraceOpts []otelsql.Option
}

func (uo *URLOpener) OpenMySQLURL(_ context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var netAddrRE = regexp.MustCompile(`^(.+)\((.+)\)$`)

func ConfigFromURL(u *url.URL) (cfg *mysql.Config, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type MySQLURLOpener interface {
	OpenMySQLURL(ctx context.Context, u *url.URL) (*sql.DB, error)
}

type URLMux struct {
	schemes openurl.SchemeMap
}

func (mux *URLMux) RegisterMySQL(scheme string, opener MySQLURLOpener) {
	_ = "STUB: not implemented"
	return
}

func (mux *URLMux) OpenMySQL(ctx context.Context, urlstr string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mux *URLMux) OpenMySQLURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var defaultURLMux = new(URLMux)

func DefaultURLMux() *URLMux { _ = "STUB: not implemented"; return nil }

func Open(ctx context.Context, urlstr string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
