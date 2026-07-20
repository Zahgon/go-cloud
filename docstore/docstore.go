package docstore

import (
	"context"
	"sync"

	"gocloud.dev/docstore/driver"
	"gocloud.dev/internal/gcerr"
	gcdkotel "gocloud.dev/internal/otel"
)

type Document = any

type Collection struct {
	driver driver.Collection
	tracer *gcdkotel.Tracer
	mu     sync.Mutex
	closed bool
}

const pkgName = "gocloud.dev/docstore"

var (
	OpenTelemetryViews = gcdkotel.Views(pkgName)
)

var NewCollection = newCollection

func newCollection(d driver.Collection) *Collection { _ = "STUB: not implemented"; return nil }

const DefaultRevisionField = "DocstoreRevision"

func (c *Collection) revisionField() string { _ = "STUB: not implemented"; return "" }

type FieldPath string

func (c *Collection) Actions() *ActionList { _ = "STUB: not implemented"; return nil }

type ActionList struct {
	coll               *Collection
	actions            []*Action
	enableAtomicWrites bool
	beforeDo           func(asFunc func(any) bool) error
}

type Action struct {
	kind          driver.ActionKind
	doc           Document
	fieldpaths    []FieldPath
	mods          Mods
	inAtomicWrite bool
}

func (l *ActionList) add(a *Action) *ActionList { _ = "STUB: not implemented"; return nil }

func (l *ActionList) Create(doc Document) *ActionList { _ = "STUB: not implemented"; return nil }

func (l *ActionList) Replace(doc Document) *ActionList { _ = "STUB: not implemented"; return nil }

func (l *ActionList) Put(doc Document) *ActionList { _ = "STUB: not implemented"; return nil }

func (l *ActionList) Delete(doc Document) *ActionList { _ = "STUB: not implemented"; return nil }

func (l *ActionList) Get(doc Document, fps ...FieldPath) *ActionList {
	_ = "STUB: not implemented"
	return nil
}

func (l *ActionList) Update(doc Document, mods Mods) *ActionList {
	_ = "STUB: not implemented"
	return nil
}

type Mods map[FieldPath]any

func Increment(amount any) any { _ = "STUB: not implemented"; return *new(any) }

type ActionListError []struct {
	Index int
	Err   error
}

func (e ActionListError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ActionListError) Unwrap() []error { _ = "STUB: not implemented"; return nil }

func (l *ActionList) BeforeDo(f func(asFunc func(any) bool) error) *ActionList {
	_ = "STUB: not implemented"
	return nil
}

func (l *ActionList) Do(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (l *ActionList) do(ctx context.Context, withTracing bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (l *ActionList) toDriverActions() ([]*driver.Action, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Collection) toDriverAction(a *Action) (*driver.Action, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFieldPaths(fps []FieldPath) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toDriverMods(mods Mods) ([]driver.Mod, error) { _ = "STUB: not implemented"; return nil, nil }

func fpHasPrefix(fp, prefix []string) bool { _ = "STUB: not implemented"; return false }

func isIncNumber(x any) bool { _ = "STUB: not implemented"; return false }

func (l *ActionList) String() string { _ = "STUB: not implemented"; return "" }

func (l *ActionList) AtomicWrites() *ActionList { _ = "STUB: not implemented"; return nil }

func (a *Action) String() string { _ = "STUB: not implemented"; return "" }

func (c *Collection) Create(ctx context.Context, doc Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collection) Replace(ctx context.Context, doc Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collection) Put(ctx context.Context, doc Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collection) Delete(ctx context.Context, doc Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collection) Get(ctx context.Context, doc Document, fps ...FieldPath) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Collection) Update(ctx context.Context, doc Document, mods Mods) error {
	_ = "STUB: not implemented"
	return nil
}

func parseFieldPath(fp FieldPath) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Collection) RevisionToString(rev any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Collection) StringToRevision(s string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (c *Collection) As(i any) bool { _ = "STUB: not implemented"; return false }

var errClosed = gcerr.Newf(gcerr.FailedPrecondition, nil, "docstore: Collection has been closed")

func (c *Collection) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Collection) checkClosed() error { _ = "STUB: not implemented"; return nil }

func wrapError(c driver.Collection, err error) error { _ = "STUB: not implemented"; return nil }

func (c *Collection) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }
