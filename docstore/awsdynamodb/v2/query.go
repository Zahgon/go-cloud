package awsdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	dyn "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dyn2Types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"gocloud.dev/docstore/driver"
)

type avmap = map[string]dyn2Types.AttributeValue

func (c *collection) RunGetQuery(ctx context.Context, q *driver.Query) (driver.DocumentIterator, error) {
	_ = "STUB: not implemented"
	return *new(driver.DocumentIterator), nil
}

func (c *collection) checkPlan(qr *queryRunner) error { _ = "STUB: not implemented"; return nil }

func (c *collection) planQuery(q *driver.Query) (*queryRunner, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) bestQueryable(q *driver.Query) (indexName *string, pkey, skey string) {
	_ = "STUB: not implemented"
	return nil, "", ""
}

func localFieldsIncluded(q *driver.Query, li dyn2Types.LocalSecondaryIndexDescription) bool {
	_ = "STUB: not implemented"
	return false
}

func orderingConsistent(q *driver.Query, sortField string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *collection) globalFieldsIncluded(q *driver.Query, gi dyn2Types.GlobalSecondaryIndexDescription) bool {
	_ = "STUB: not implemented"
	return false
}

func keyAttributes(ks []dyn2Types.KeySchemaElement) (pkey, skey string) {
	_ = "STUB: not implemented"
	return "", ""
}

func hasFilter(q *driver.Query, field string) bool { _ = "STUB: not implemented"; return false }

func hasEqualityFilter(q *driver.Query, field string) bool { _ = "STUB: not implemented"; return false }

type queryRunner struct {
	c         *collection
	scanIn    *dyn.ScanInput
	queryIn   *dyn.QueryInput
	beforeRun func(asFunc func(i any) bool) error
}

func (qr *queryRunner) run(ctx context.Context, startAfter avmap) (items []avmap, last avmap, asFunc func(i any) bool, err error) {
	_ = "STUB: not implemented"
	return nil, *new(avmap), nil, nil
}

func processFilters(cb expression.Builder, fs []driver.Filter, pkey, skey string) expression.Builder {
	_ = "STUB: not implemented"
	return *new(expression.Builder)
}

func filtersToConditionBuilder(fs []driver.Filter) expression.ConditionBuilder {
	_ = "STUB: not implemented"
	return *new(expression.ConditionBuilder)
}

func toKeyCondition(f driver.Filter, pkey, skey string) (expression.KeyConditionBuilder, bool) {
	_ = "STUB: not implemented"
	return *new(expression.KeyConditionBuilder), false
}

func toFilter(f driver.Filter) expression.ConditionBuilder {
	_ = "STUB: not implemented"
	return *new(expression.ConditionBuilder)
}

func toInCondition(f driver.Filter) expression.ConditionBuilder {
	_ = "STUB: not implemented"
	return *new(expression.ConditionBuilder)
}

type documentIterator struct {
	qr     *queryRunner
	items  []map[string]dyn2Types.AttributeValue
	curr   int
	offset int
	limit  int
	count  int
	last   map[string]dyn2Types.AttributeValue
	asFunc func(i any) bool
}

func (it *documentIterator) Next(ctx context.Context, doc driver.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (it *documentIterator) next(ctx context.Context, doc driver.Document, decode bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (it *documentIterator) Stop() { _ = "STUB: not implemented"; return }

func (it *documentIterator) As(i any) bool { _ = "STUB: not implemented"; return false }

func (c *collection) QueryPlan(q *driver.Query) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (qr *queryRunner) queryPlan() string { _ = "STUB: not implemented"; return "" }

func InMemorySortFallback(createDocument func() any) FallbackFunc {
	_ = "STUB: not implemented"
	return *new(FallbackFunc)
}

type docsForSorting struct {
	docs      []driver.Document
	vals      []any
	ascending bool
}

func (d docsForSorting) Len() int { _ = "STUB: not implemented"; return 0 }

func (d docsForSorting) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (d docsForSorting) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func compare(v1, v2 any) int { _ = "STUB: not implemented"; return 0 }

type sliceIterator struct {
	docs []driver.Document
	next int
}

func (it *sliceIterator) Next(ctx context.Context, doc driver.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func copyTopLevel(dest, src driver.Document) error { _ = "STUB: not implemented"; return nil }

func (*sliceIterator) Stop()       { _ = "STUB: not implemented"; return }
func (*sliceIterator) As(any) bool { _ = "STUB: not implemented"; return false }
