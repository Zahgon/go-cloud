package main

import (
	"context"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/memdocstore"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/mempubsub"
)

type frontend struct {
	requestTopic *pubsub.Topic
	bucket       *blob.Bucket
	coll         *docstore.Collection
}

var (
	listTemplate      *template.Template
	orderFormTemplate *template.Template
)

func init() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if filepath.Base(dir) != "order" {

		log.Printf("working around #33016, which put us in %s", dir)
		dir = filepath.Join(filepath.Dir(dir), "order")
	}
	listTemplate = template.Must(template.ParseFiles(filepath.Join(dir, "list.htmlt")))
	orderFormTemplate = template.Must(template.ParseFiles(filepath.Join(dir, "order-form.htmlt")))
}

func (f *frontend) run(ctx context.Context, port int) error { _ = "STUB: not implemented"; return nil }

func wrapHTTPError(f func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

func (*frontend) orderForm(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frontend) createOrder(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frontend) doCreateOrder(ctx context.Context, email string, file io.Reader, now time.Time) (id string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (f *frontend) listOrders(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frontend) showImage(w http.ResponseWriter, r *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *frontend) newID() string { _ = "STUB: not implemented"; return "" }

func executeTemplate(t *template.Template, data any, w http.ResponseWriter) error {
	_ = "STUB: not implemented"
	return nil
}
