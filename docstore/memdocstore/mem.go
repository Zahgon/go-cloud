package memdocstore

import (
	"context"
	"sync"

	"gocloud.dev/docstore"
	"gocloud.dev/docstore/driver"
	"gocloud.dev/gcerrors"
)

type Options struct {
	RevisionField string

	MaxOutstandingActions int

	Filename string

	AllowNestedSliceQueries bool

	onClose func()
}

func OpenCollection(keyField string, opts *Options) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenCollectionWithKeyFunc(keyFunc func(docstore.Document) any, opts *Options) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCollection(keyField string, keyFunc func(docstore.Document) any, opts *Options) (driver.Collection, error) {
	_ = "STUB: not implemented"
	return *new(driver.Collection), nil
}

type storedDoc map[string]any

type collection struct {
	keyField    string
	keyFunc     func(docstore.Document) any
	opts        *Options
	mu          sync.Mutex
	docs        map[any]storedDoc
	curRevision int64
}

func (c *collection) Key(doc driver.Document) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (c *collection) RevisionField() string { _ = "STUB: not implemented"; return "" }

func (c *collection) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (c *collection) RunActions(ctx context.Context, actions []*driver.Action, opts *driver.RunActionsOptions) driver.ActionListError {
	_ = "STUB: not implemented"
	return *new(driver.ActionListError)
}

func (c *collection) runAtomicWrites(ctx context.Context, actions []*driver.Action, errs []error) {
	_ = "STUB: not implemented"
	return
}

func (c *collection) runAction(ctx context.Context, a *driver.Action) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) executeAction(a *driver.Action, current storedDoc, exists bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) update(doc storedDoc, mods []driver.Mod) error {
	_ = "STUB: not implemented"
	return nil
}

func add(x, y any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (c *collection) changeRevision(doc storedDoc) { _ = "STUB: not implemented"; return }

func (c *collection) checkRevision(arg driver.Document, current storedDoc) error {
	_ = "STUB: not implemented"
	return nil
}

func getAtFieldPath(m map[string]any, fp []string, nested bool) (result any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func setAtFieldPath(m map[string]any, fp []string, val any) error {
	_ = "STUB: not implemented"
	return nil
}

func getParentMap(m map[string]any, fp []string, create bool) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) RevisionToBytes(rev any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) BytesToRevision(b []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (c *collection) As(i any) bool { _ = "STUB: not implemented"; return false }

func (c *collection) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (c *collection) Close() error { _ = "STUB: not implemented"; return nil }

type mapOfDocs = map[any]storedDoc

func loadDocs(filename string) (m mapOfDocs, err error) {
	_ = "STUB: not implemented"
	return *new(mapOfDocs), nil
}

func saveDocs(filename string, m mapOfDocs) error { _ = "STUB: not implemented"; return nil }
