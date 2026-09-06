package driver

import (
	"context"
	"time"

	"gocloud.dev/gcerrors"
)

const DefaultWaitDuration = 30 * time.Second

func WaitDuration(d time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type State interface {
	Value() (any, error)

	UpdateTime() time.Time

	As(any) bool
}

type Watcher interface {
	WatchVariable(ctx context.Context, prev State) (state State, wait time.Duration)

	Close() error

	ErrorAs(error, any) bool

	ErrorCode(error) gcerrors.ErrorCode
}
