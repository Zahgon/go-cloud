package docstore

import (
	"context"

	"gocloud.dev/docstore/driver"
)

type Query struct {
	coll *Collection
	dq   *driver.Query
	err  error
}

func (c *Collection) Query() *Query { _ = "STUB: not implemented"; return nil }

func (q *Query) Where(fp FieldPath, op string, value any) *Query {
	_ = "STUB: not implemented"
	return nil
}

type valueValidator func(any) bool

var validOp = map[string]valueValidator{
	"=":      validEqualValue,
	">":      validFilterValue,
	"<":      validFilterValue,
	">=":     validFilterValue,
	"<=":     validFilterValue,
	"in":     validFilterSlice,
	"not-in": validFilterSlice,
}

func validEqualValue(v any) bool { _ = "STUB: not implemented"; return false }

func validFilterValue(v any) bool { _ = "STUB: not implemented"; return false }

func validFilterSlice(v any) bool { _ = "STUB: not implemented"; return false }

func (q *Query) Offset(n int) *Query { _ = "STUB: not implemented"; return nil }

func (q *Query) Limit(n int) *Query { _ = "STUB: not implemented"; return nil }

const (
	Ascending  = "asc"
	Descending = "desc"
)

func (q *Query) OrderBy(field, direction string) *Query { _ = "STUB: not implemented"; return nil }

func (q *Query) BeforeQuery(f func(asFunc func(any) bool) error) *Query {
	_ = "STUB: not implemented"
	return nil
}

func (q *Query) Get(ctx context.Context, fps ...FieldPath) *DocumentIterator {
	_ = "STUB: not implemented"
	return nil
}

func (q *Query) get(ctx context.Context, withTracing bool, fps ...FieldPath) *DocumentIterator {
	_ = "STUB: not implemented"
	return nil
}

func (q *Query) initGet(fps []FieldPath) error { _ = "STUB: not implemented"; return nil }

func (q *Query) invalidf(format string, args ...any) *Query { _ = "STUB: not implemented"; return nil }

type DocumentIterator struct {
	iter driver.DocumentIterator
	coll *Collection
	err  error
}

func (it *DocumentIterator) Next(ctx context.Context, dst Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (it *DocumentIterator) Stop() { _ = "STUB: not implemented"; return }

func (it *DocumentIterator) As(i any) bool { _ = "STUB: not implemented"; return false }

func (q *Query) Plan(fps ...FieldPath) (string, error) { _ = "STUB: not implemented"; return "", nil }
