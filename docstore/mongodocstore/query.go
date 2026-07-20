package mongodocstore

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gocloud.dev/docstore/driver"
)

func (c *collection) RunGetQuery(ctx context.Context, q *driver.Query) (driver.DocumentIterator, error) {
	_ = "STUB: not implemented"
	return *new(driver.DocumentIterator), nil
}

var mongoQueryOps = map[string]string{
	driver.EqualOp: "$eq",
	">":            "$gt",
	">=":           "$gte",
	"<":            "$lt",
	"<=":           "$lte",
	"in":           "$in",
	"not-in":       "$nin",
}

func (c *collection) filtersToBSON(fs []driver.Filter) (bson.D, error) {
	_ = "STUB: not implemented"
	return *new(bson.D), nil
}

func (c *collection) filterToBSON(f driver.Filter) (bson.E, error) {
	_ = "STUB: not implemented"
	return *new(bson.E), nil
}

type docIterator struct {
	cursor          *mongo.Cursor
	idField         string
	ctx             context.Context
	lowercaseFields bool
}

func (it *docIterator) Next(ctx context.Context, doc driver.Document) error {
	_ = "STUB: not implemented"
	return nil
}

func (it *docIterator) nextMap(ctx context.Context) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (it *docIterator) Stop() { _ = "STUB: not implemented"; return }

func (it *docIterator) As(i interface{}) bool { _ = "STUB: not implemented"; return false }

func (c *collection) QueryPlan(q *driver.Query) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *collection) RunDeleteQuery(ctx context.Context, q *driver.Query) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) RunUpdateQuery(ctx context.Context, q *driver.Query, mods []driver.Mod) error {
	_ = "STUB: not implemented"
	return nil
}
