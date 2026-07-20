package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"net/url"

	"github.com/XSAM/otelsql"
	"gocloud.dev/internal/openurl"
)

const Scheme = "postgres"

func init() {
	DefaultURLMux().RegisterPostgres(Scheme, &URLOpener{})
}

type URLOpener struct {
	TraceOpts []otelsql.Option
}

func (uo *URLOpener) OpenPostgresURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type connector struct {
	dsn       string
	traceOpts []otelsql.Option
}

func (c connector) Connect(ctx context.Context) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func (c connector) Driver() driver.Driver { _ = "STUB: not implemented"; return *new(driver.Driver) }

type PostgresURLOpener interface {
	OpenPostgresURL(ctx context.Context, u *url.URL) (*sql.DB, error)
}

type URLMux struct {
	schemes openurl.SchemeMap
}

func (mux *URLMux) RegisterPostgres(scheme string, opener PostgresURLOpener) {
	_ = "STUB: not implemented"
	return
}

func (mux *URLMux) OpenPostgres(ctx context.Context, urlstr string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mux *URLMux) OpenPostgresURL(ctx context.Context, u *url.URL) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var defaultURLMux = new(URLMux)

func DefaultURLMux() *URLMux { _ = "STUB: not implemented"; return nil }

func Open(ctx context.Context, urlstr string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
