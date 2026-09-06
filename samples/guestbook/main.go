package main

import (
	"context"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	"gocloud.dev/blob"
	"gocloud.dev/runtimevar"

	"gocloud.dev/server"
	"gocloud.dev/server/health"
)

type cliFlags struct {
	bucket          string
	dbHost          string
	dbName          string
	dbUser          string
	dbPassword      string
	motdVar         string
	motdVarWaitTime time.Duration

	cloudSQLRegion    string
	runtimeConfigName string
}

var envFlag string

func main() {

	cf := new(cliFlags)
	flag.StringVar(&envFlag, "env", "local", "environment to run under (gcp, aws, azure, or local)")
	addr := flag.String("listen", ":8080", "port to listen for HTTP on")
	flag.StringVar(&cf.bucket, "bucket", "", "bucket name")
	flag.StringVar(&cf.dbHost, "db_host", "", "database host or Cloud SQL instance name")
	flag.StringVar(&cf.dbName, "db_name", "guestbook", "database name")
	flag.StringVar(&cf.dbUser, "db_user", "guestbook", "database user")
	flag.StringVar(&cf.dbPassword, "db_password", "", "database user password")
	flag.StringVar(&cf.motdVar, "motd_var", "", "message of the day variable location")
	flag.DurationVar(&cf.motdVarWaitTime, "motd_var_wait_time", 5*time.Second, "polling frequency of message of the day")
	flag.StringVar(&cf.cloudSQLRegion, "cloud_sql_region", "", "region of the Cloud SQL instance (GCP only)")
	flag.StringVar(&cf.runtimeConfigName, "runtime_config", "", "Runtime Configurator config resource (GCP only)")
	flag.Parse()

	ctx := context.Background()
	var srv *server.Server
	var cleanup func()
	var err error
	switch envFlag {
	case "gcp":
		srv, cleanup, err = setupGCP(ctx, cf)
	case "aws":
		srv, cleanup, err = setupAWS(ctx, cf)
	case "azure":
		if cf.dbHost == "" {
			cf.dbHost = "localhost"
		}
		if cf.dbPassword == "" {
			cf.dbPassword = "xyzzy"
		}
		srv, cleanup, err = setupAzure(ctx, cf)
	case "local":

		if cf.dbHost == "" {
			cf.dbHost = "localhost"
		}
		if cf.dbPassword == "" {
			cf.dbPassword = "xyzzy"
		}
		srv, cleanup, err = setupLocal(ctx, cf)
	default:
		log.Fatalf("unknown -env=%s", envFlag)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	log.Printf("Running, connected to %q cloud", envFlag)
	log.Fatal(srv.ListenAndServe(*addr))
}

var applicationSet = wire.NewSet(
	newApplication,
	appHealthChecks,
	newRouter,
	wire.Bind(new(http.Handler), new(*mux.Router)),
)

func newRouter(app *application) *mux.Router { _ = "STUB: not implemented"; return nil }

type application struct {
	db      *sql.DB
	bucket  *blob.Bucket
	motdVar *runtimevar.Variable
}

func newApplication(db *sql.DB, bucket *blob.Bucket, motdVar *runtimevar.Variable) *application {
	_ = "STUB: not implemented"
	return nil
}

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

type greeting struct {
	Content string
}

var tmpl = template.Must(template.New("index.html").Parse(`<!DOCTYPE html>
<title>Guestbook - {{.Env}}</title>
<style type="text/css">
html, body {
	font-family: Helvetica, sans-serif;
}
blockquote {
	font-family: cursive, Helvetica, sans-serif;
}
.banner {
	height: 125px;
	width: 250px;
}
.greeting {
	font-size: 85%;
}
.motd {
	font-weight: bold;
}
</style>
<h1>Guestbook</h1>
<div><img class="banner" src="{{.BannerSrc}}"></div>
{{with .MOTD}}<p class="motd">Admin says: {{.}}</p>{{end}}
{{range .Greetings}}
<div class="greeting">
	Someone wrote:
	<blockquote>{{.Content}}</blockquote>
</div>
{{end}}
<form action="/sign" method="POST">
	<div><textarea name="content" rows="3"></textarea></div>
	<div><input type="submit" value="Sign"></div>
</form>
`))

func (app *application) sign(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (app *application) serveBlob(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func appHealthChecks(db *sql.DB) ([]health.Checker, func()) {
	_ = "STUB: not implemented"
	return nil, nil
}
