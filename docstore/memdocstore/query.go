package memdocstore

import (
	"context"

	"gocloud.dev/docstore/driver"
)

func (c *collection) RunGetQuery(_ context.Context, q *driver.Query) (driver.DocumentIterator, error) {
	_ = "STUB: not implemented"
	return *new(driver.DocumentIterator), nil
}

func filtersMatch(fs []driver.Filter, doc storedDoc, nested bool) bool {
	_ = "STUB: not implemented"
	return false
}

func filterMatches(f driver.Filter, doc storedDoc, nested bool) bool {
	_ = "STUB: not implemented"
	return false
}

func applyComparison(op string, c int) bool { _ = "STUB: not implemented"; return false }

func compare(x1, x2 any, op string) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func sortDocs(docs []storedDoc, field string, asc bool) { _ = "STUB: not implemented"; return }

type docIterator struct {
	docs       []storedDoc
	fieldPaths [][]string
	revField   string
	err        error
}

func (it *docIterator) Next(ctx context.Context, doc driver.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (it *docIterator) Stop() { _ = "STUB: not implemented"; return }

func (it *docIterator) As(i any) bool { _ = "STUB: not implemented"; return false }

func (c *collection) QueryPlan(q *driver.Query) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
