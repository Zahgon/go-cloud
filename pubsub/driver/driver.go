package driver

import (
	"context"

	"gocloud.dev/gcerrors"
)

type AckID any

type AckInfo struct {
	AckID AckID

	IsAck bool
}

type Message struct {
	LoggableID string

	Body []byte

	Metadata map[string]string

	AckID AckID

	AsFunc func(any) bool

	BeforeSend func(asFunc func(any) bool) error

	AfterSend func(asFunc func(any) bool) error
}

func (m *Message) ByteSize() int { _ = "STUB: not implemented"; return 0 }

type Topic interface {
	SendBatch(ctx context.Context, ms []*Message) error

	IsRetryable(err error) bool

	As(i any) bool

	ErrorAs(error, any) bool

	ErrorCode(error) gcerrors.ErrorCode

	Close() error
}

type Subscription interface {
	ReceiveBatch(ctx context.Context, maxMessages int) ([]*Message, error)

	SendAcks(ctx context.Context, ackIDs []AckID) error

	CanNack() bool

	SendNacks(ctx context.Context, ackIDs []AckID) error

	IsRetryable(err error) bool

	As(i any) bool

	ErrorAs(error, any) bool

	ErrorCode(error) gcerrors.ErrorCode

	Close() error
}
