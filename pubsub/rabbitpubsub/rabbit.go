package rabbitpubsub

import (
	"context"
	"net/url"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
	"gocloud.dev/gcerrors"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/driver"
)

func init() {
	o := new(defaultDialer)
	pubsub.DefaultURLMux().RegisterTopic(Scheme, o)
	pubsub.DefaultURLMux().RegisterSubscription(Scheme, o)
}

type defaultDialer struct {
	mu     sync.Mutex
	conn   *amqp.Connection
	opener *URLOpener
}

func (o *defaultDialer) defaultConn(ctx context.Context) (*URLOpener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *defaultDialer) OpenTopicURL(ctx context.Context, u *url.URL) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *defaultDialer) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "rabbit"

type URLOpener struct {
	Connection *amqp.Connection

	TopicOptions TopicOptions

	SubscriptionOptions SubscriptionOptions
}

func (o *URLOpener) OpenTopicURL(ctx context.Context, u *url.URL) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *URLOpener) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type topic struct {
	exchange string
	conn     amqpConnection
	opts     *TopicOptions

	mu     sync.Mutex
	ch     amqpChannel
	pubc   <-chan amqp.Confirmation
	retc   <-chan amqp.Return
	closec <-chan *amqp.Error
}

type TopicOptions struct {
	KeyName string
}

type SubscriptionOptions struct {
	KeyName string

	PrefetchCount *int
}

func OpenTopic(conn *amqp.Connection, name string, opts *TopicOptions) *pubsub.Topic {
	_ = "STUB: not implemented"
	return nil
}

func newTopic(conn amqpConnection, name string, opts *TopicOptions) *topic {
	_ = "STUB: not implemented"
	return nil
}

func (t *topic) establishChannel(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func runWithContext(ctx context.Context, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *topic) SendBatch(ctx context.Context, ms []*driver.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *topic) receiveFromPublishChannels(ctx context.Context, nMessages int) error {
	_ = "STUB: not implemented"
	return nil
}

type MultiError []error

func (m MultiError) Error() string { _ = "STUB: not implemented"; return "" }

func closeErr(closec <-chan *amqp.Error) error { _ = "STUB: not implemented"; return nil }

func toRoutingKeyAndAMQPPublishing(m *driver.Message, opts *TopicOptions) (routingKey string, msg amqp.Publishing) {
	_ = "STUB: not implemented"
	return "", *new(amqp.Publishing)
}

func (*topic) IsRetryable(err error) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

var errorCodes = map[int]gcerrors.ErrorCode{
	amqp.NotFound:           gcerrors.NotFound,
	amqp.PreconditionFailed: gcerrors.FailedPrecondition,

	amqp.SyntaxError:    gcerrors.Internal,
	amqp.CommandInvalid: gcerrors.Internal,
	amqp.InternalError:  gcerrors.Internal,
	amqp.NotImplemented: gcerrors.Unimplemented,
	amqp.ChannelError:   gcerrors.FailedPrecondition,
}

func errorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func isRetryable(err error) bool { _ = "STUB: not implemented"; return false }

func (t *topic) As(i any) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func errorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*topic) Close() error { _ = "STUB: not implemented"; return nil }

func OpenSubscription(conn *amqp.Connection, name string, opts *SubscriptionOptions) *pubsub.Subscription {
	_ = "STUB: not implemented"
	return nil
}

type subscription struct {
	conn     amqpConnection
	queue    string
	consumer string

	opts *SubscriptionOptions

	mu     sync.Mutex
	ch     amqpChannel
	delc   <-chan amqp.Delivery
	closec <-chan *amqp.Error

	receiveBatchHook func()
}

var nextConsumer int64

func newSubscription(conn amqpConnection, name string, opts *SubscriptionOptions) *subscription {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) establishChannel(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func applyOptionsToChannel(opts *SubscriptionOptions, ch amqpChannel) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) ReceiveBatch(ctx context.Context, maxMessages int) ([]*driver.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toDriverMessage(d amqp.Delivery, opts *SubscriptionOptions) *driver.Message {
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

func (s *subscription) sendAcksOrNacks(ctx context.Context, ackIDs []driver.AckID, ack bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (*subscription) IsRetryable(err error) bool { _ = "STUB: not implemented"; return false }

func (*subscription) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (s *subscription) As(i any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) Close() error { _ = "STUB: not implemented"; return nil }
