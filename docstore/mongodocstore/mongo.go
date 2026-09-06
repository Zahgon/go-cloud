package mongodocstore

import (
	"context"

	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gocloud.dev/docstore"
	"gocloud.dev/docstore/driver"
	"gocloud.dev/gcerrors"
)

func Dial(ctx context.Context, uri string) (*mongo.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var Set = wire.NewSet(
	Dial,
	wire.Struct(new(URLOpener), "Client"),
)

type collection struct {
	coll          *mongo.Collection
	idField       string
	idFunc        func(docstore.Document) interface{}
	revisionField string
	opts          *Options
}

type Options struct {
	LowercaseFields bool

	RevisionField string

	NoWriteQueryUpdateRevisions bool
}

func OpenCollection(mcoll *mongo.Collection, idField string, opts *Options) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenCollectionWithIDFunc(mcoll *mongo.Collection, idFunc func(docstore.Document) interface{}, opts *Options) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCollection(mcoll *mongo.Collection, idField string, idFunc func(docstore.Document) interface{}, opts *Options) (*collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) Key(doc driver.Document) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) RevisionField() string { _ = "STUB: not implemented"; return "" }

const mongoIDField = "_id"

func (c *collection) RunActions(ctx context.Context, actions []*driver.Action, opts *driver.RunActionsOptions) driver.ActionListError {
	_ = "STUB: not implemented"
	return *new(driver.ActionListError)
}

type indexedError = struct {
	Index int
	Err   error
}

func (c *collection) runGets(ctx context.Context, gets []*driver.Action, errs []error, opts *driver.RunActionsOptions) {
	_ = "STUB: not implemented"
	return
}

func (c *collection) bulkFind(ctx context.Context, gets []*driver.Action, errs []error, dopts *driver.RunActionsOptions) {
	_ = "STUB: not implemented"
	return
}

func (c *collection) projectionDoc(fps [][]string) bson.D {
	_ = "STUB: not implemented"
	return *new(bson.D)
}

func (c *collection) toMongoFieldPath(fp []string) string { _ = "STUB: not implemented"; return "" }

func sliceToLower(s []string) { _ = "STUB: not implemented"; return }

func (c *collection) prepareCreate(a *driver.Action) (mdoc, createdID interface{}, rev string, err error) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}

func (c *collection) prepareReplace(a *driver.Action) (filter bson.D, mdoc map[string]interface{}, rev string, err error) {
	_ = "STUB: not implemented"
	return *new(bson.D), nil, "", nil
}

func (c *collection) encodeDoc(doc driver.Document, id interface{}) (map[string]interface{}, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (c *collection) prepareUpdate(a *driver.Action) (filter bson.D, updateDoc map[string]bson.D, rev string, err error) {
	_ = "STUB: not implemented"
	return *new(bson.D), nil, "", nil
}

func (c *collection) newUpdateDoc(mods []driver.Mod, writeRevision bool) (map[string]bson.D, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (c *collection) makeFilter(id interface{}, doc driver.Document) (filter bson.D, rev interface{}, err error) {
	_ = "STUB: not implemented"
	return *new(bson.D), nil, nil
}

func (c *collection) bulkWrite(ctx context.Context, actions []*driver.Action, errs []error, dopts *driver.RunActionsOptions) []error {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) txWrite(ctx context.Context, actions []*driver.Action, errs []error, dopts *driver.RunActionsOptions) []error {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) determineDeleteErrors(ctx context.Context, models []mongo.WriteModel, actions []*driver.Action, errs []error) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *collection) newCreateModel(a *driver.Action) (*mongo.InsertOneModel, interface{}, string, error) {
	_ = "STUB: not implemented"
	return nil, nil, "", nil
}

func (c *collection) newDeleteModel(a *driver.Action) (*mongo.DeleteOneModel, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) newReplaceModel(a *driver.Action, upsert bool) (*mongo.ReplaceOneModel, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (c *collection) newUpdateModel(a *driver.Action) (*mongo.UpdateOneModel, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func (c *collection) RevisionToBytes(rev interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) hasField(doc driver.Document, field string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *collection) BytesToRevision(b []byte) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) As(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (c *collection) ErrorAs(err error, i interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *collection) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (c *collection) Close() error { _ = "STUB: not implemented"; return nil }

const mongoDupKeyCode = 11000

func translateMongoCode(code int) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
