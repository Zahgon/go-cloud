package mempubsub

import (
	"context"
	"errors"
	"net/url"
	"sync"
	"time"

	"gocloud.dev/gcerrors"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/batcher"
	"gocloud.dev/pubsub/driver"
)

func init() {
	o := new(URLOpener)
	pubsub.DefaultURLMux().RegisterTopic(Scheme, o)
	pubsub.DefaultURLMux().RegisterSubscription(Scheme, o)
}

const Scheme = "mem"

type URLOpener struct {
	mu     sync.Mutex
	topics map[string]*pubsub.Topic
}

func (o *URLOpener) OpenTopicURL(ctx context.Context, u *url.URL) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *URLOpener) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var errNotExist = errors.New("mempubsub: topic does not exist")

type topic struct {
	mu        sync.Mutex
	subs      []*subscription
	nextAckID int
}

type TopicOptions struct {
	BatcherOptions batcher.Options
}

func NewTopic() *pubsub.Topic { _ = "STUB: not implemented"; return nil }

func NewTopicWithOptions(opts *TopicOptions) *pubsub.Topic { _ = "STUB: not implemented"; return nil }

func (t *topic) SendBatch(ctx context.Context, ms []*driver.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (*topic) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (t *topic) As(i any) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorAs(error, any) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (*topic) Close() error { _ = "STUB: not implemented"; return nil }

type SubscriptionOptions struct {
	ReceiveBatcherOptions batcher.Options

	AckBatcherOptions batcher.Options
}

type subscription struct {
	mu          sync.Mutex
	topic       *topic
	ackDeadline time.Duration
	msgs        map[driver.AckID]*message
}

func NewSubscription(pstopic *pubsub.Topic, ackDeadline time.Duration) *pubsub.Subscription {
	_ = "STUB: not implemented"
	return nil
}

func NewSubscriptionWithOptions(pstopic *pubsub.Topic, ackDeadline time.Duration, opts *SubscriptionOptions) *pubsub.Subscription {
	_ = "STUB: not implemented"
	return nil
}

func newSubscription(topic *topic, ackDeadline time.Duration) *subscription {
	_ = "STUB: not implemented"
	return nil
}

type message struct {
	msg        *driver.Message
	expiration time.Time
}

func (s *subscription) add(ms []*driver.Message) { _ = "STUB: not implemented"; return }

func (s *subscription) receiveNoWait(now time.Time, max int) []*driver.Message {
	_ = "STUB: not implemented"
	return nil
}

const pollDuration = 250 * time.Millisecond

func (s *subscription) ReceiveBatch(ctx context.Context, maxMessages int) ([]*driver.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *subscription) wait(ctx context.Context, dur time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) SendAcks(ctx context.Context, ackIDs []driver.AckID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) CanNack() bool { _ = "STUB: not implemented"; return false }

func (s *subscription) SendNacks(ctx context.Context, ackIDs []driver.AckID) error {
	_ = "STUB: not implemented"
	return nil
}

func (*subscription) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) As(i any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) ErrorAs(error, any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (*subscription) Close() error { _ = "STUB: not implemented"; return nil }
