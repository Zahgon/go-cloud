package awsdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	dyn "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dyn2Types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/wire"
	"gocloud.dev/docstore"
	"gocloud.dev/docstore/driver"
	"gocloud.dev/gcerrors"
)

var Set = wire.NewSet(
	wire.Struct(new(URLOpener)),
)

type collection struct {
	db           *dyn.Client
	table        string
	partitionKey string
	sortKey      string
	description  *dyn2Types.TableDescription
	opts         *Options
}

type FallbackFunc func(context.Context, *driver.Query, RunQueryFunc) (driver.DocumentIterator, error)

type Options struct {
	AllowScans bool

	RevisionField string

	RunQueryFallback FallbackFunc

	MaxOutstandingActionRPCs int

	ConsistentRead bool
}

type RunQueryFunc func(context.Context, *driver.Query) (driver.DocumentIterator, error)

func OpenCollection(db *dyn.Client, tableName, partitionKey, sortKey string, opts *Options) (*docstore.Collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCollection(db *dyn.Client, tableName, partitionKey, sortKey string, opts *Options) (*collection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) Key(doc driver.Document) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (c *collection) RevisionField() string { _ = "STUB: not implemented"; return "" }

func (c *collection) RunActions(ctx context.Context, actions []*driver.Action, opts *driver.RunActionsOptions) driver.ActionListError {
	_ = "STUB: not implemented"
	return *new(driver.ActionListError)
}

func (c *collection) runGets(ctx context.Context, actions []*driver.Action, errs []error, opts *driver.RunActionsOptions) {
	_ = "STUB: not implemented"
	return
}

func (c *collection) batchGet(ctx context.Context, gets []*driver.Action, errs []error, opts *driver.RunActionsOptions, start, end int) {
	_ = "STUB: not implemented"
	return
}

func mapActionIndices(actions []*driver.Action, start, end int) map[any]int {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) runWrites(ctx context.Context, writes []*driver.Action, errs []error, opts *driver.RunActionsOptions) {
	_ = "STUB: not implemented"
	return
}

type writeOp struct {
	action          *driver.Action
	writeItem       dyn2Types.TransactWriteItem
	newPartitionKey string
	newRevision     string
	run             func(context.Context) error
}

func (c *collection) newWriteOp(a *driver.Action, opts *driver.RunActionsOptions) (*writeOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) newPut(a *driver.Action, opts *driver.RunActionsOptions) (*writeOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) runPut(ctx context.Context, dput *dyn2Types.Put, a *driver.Action, opts *driver.RunActionsOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *collection) newDelete(a *driver.Action, opts *driver.RunActionsOptions) (*writeOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) newUpdate(a *driver.Action, opts *driver.RunActionsOptions) (*writeOp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) onSuccess(op *writeOp) error { _ = "STUB: not implemented"; return nil }

func (c *collection) missingKeyField(m map[string]dyn2Types.AttributeValue) string {
	_ = "STUB: not implemented"
	return ""
}

func (c *collection) precondition(a *driver.Action) (*expression.ConditionBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func revisionPrecondition(doc driver.Document, revField string) (*expression.ConditionBuilder, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *collection) transactWrite(ctx context.Context, actions []*driver.Action, errs []error, opts *driver.RunActionsOptions) {
	_ = "STUB: not implemented"
	return
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

func (c *collection) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

var errorCodeMap = map[string]gcerrors.ErrorCode{

	"ValidationException": gcerrors.InvalidArgument,
}

func (c *collection) Close() error { _ = "STUB: not implemented"; return nil }
