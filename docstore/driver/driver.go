package driver

import (
	"context"

	"gocloud.dev/gcerrors"
)

type Collection interface {
	Key(Document) (any, error)

	RevisionField() string

	RunActions(ctx context.Context, actions []*Action, opts *RunActionsOptions) ActionListError

	RunGetQuery(context.Context, *Query) (DocumentIterator, error)

	QueryPlan(*Query) (string, error)

	RevisionToBytes(any) ([]byte, error)

	BytesToRevision([]byte) (any, error)

	As(i any) bool

	ErrorAs(err error, i any) bool

	ErrorCode(error) gcerrors.ErrorCode

	Close() error
}

type DeleteQueryer interface {
	RunDeleteQuery(context.Context, *Query) error
}

type UpdateQueryer interface {
	RunUpdateQuery(context.Context, *Query, []Mod) error
}

type ActionKind int

const (
	Create ActionKind = iota
	Replace
	Put
	Get
	Delete
	Update
)

//go:generate stringer -type=ActionKind

type Action struct {
	Kind          ActionKind
	Doc           Document
	Key           any
	FieldPaths    [][]string
	Mods          []Mod
	Index         int
	InAtomicWrite bool
}

type Mod struct {
	FieldPath []string
	Value     any
}

type IncOp struct {
	Amount any
}

type ActionListError []struct {
	Index int
	Err   error
}

func NewActionListError(errs []error) ActionListError {
	_ = "STUB: not implemented"
	return *new(ActionListError)
}

type RunActionsOptions struct {
	BeforeDo func(asFunc func(any) bool) error
}

type Query struct {
	FieldPaths [][]string

	Filters []Filter

	Offset int

	Limit int

	OrderByField string

	OrderAscending bool

	BeforeQuery func(asFunc func(any) bool) error
}

type Filter struct {
	FieldPath []string
	Op        string
	Value     any
}

func NewFilter(fieldPath []string, op string, value any) Filter {
	_ = "STUB: not implemented"
	return *new(Filter)
}

type DocumentIterator interface {
	Next(context.Context, Document) error

	Stop()

	As(i any) bool
}

const EqualOp = "="
