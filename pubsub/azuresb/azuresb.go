package azuresb

import (
	"context"
	"net/url"
	"sync"
	"time"

	servicebus "github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"gocloud.dev/gcerrors"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/batcher"
	"gocloud.dev/pubsub/driver"
)

const (
	defaultListenerTimeout = 2 * time.Second
)

var sendBatcherOpts = &batcher.Options{
	MaxBatchSize: 1,
	MaxHandlers:  100,
}

var recvBatcherOpts = &batcher.Options{
	MaxBatchSize: 50,
	MaxHandlers:  100,
}

var ackBatcherOpts = &batcher.Options{
	MaxBatchSize: 1,
	MaxHandlers:  100,
}

func init() {
	o := new(defaultOpener)
	pubsub.DefaultURLMux().RegisterTopic(Scheme, o)
	pubsub.DefaultURLMux().RegisterSubscription(Scheme, o)
}

type defaultOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *defaultOpener) defaultOpener() (*URLOpener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *defaultOpener) OpenTopicURL(ctx context.Context, u *url.URL) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *defaultOpener) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "azuresb"

type URLOpener struct {
	ConnectionString string

	ServiceBusHostname string

	ServiceBusClientOptions *servicebus.ClientOptions

	ServiceBusSenderOptions   *servicebus.NewSenderOptions
	ServiceBusReceiverOptions *servicebus.ReceiverOptions

	TopicOptions TopicOptions

	SubscriptionOptions SubscriptionOptions
}

func redactSharedAccessKey(connString string) string { _ = "STUB: not implemented"; return "" }

func (o *URLOpener) sbClient(kind string, u *url.URL) (*servicebus.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
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
	sbSender *servicebus.Sender
}

type TopicOptions struct {
	BatcherOptions batcher.Options
}

func NewClientFromConnectionString(connectionString string, opts *servicebus.ClientOptions) (*servicebus.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewClientFromServiceBusHostname(serviceBusHostname string, opts *servicebus.ClientOptions) (*servicebus.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSender(sbClient *servicebus.Client, topicName string, opts *servicebus.NewSenderOptions) (*servicebus.Sender, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewReceiver(sbClient *servicebus.Client, topicName, subscriptionName string, opts *servicebus.ReceiverOptions) (*servicebus.Receiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenTopic(ctx context.Context, sbSender *servicebus.Sender, opts *TopicOptions) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openTopic(ctx context.Context, sbSender *servicebus.Sender, _ *TopicOptions) (driver.Topic, error) {
	_ = "STUB: not implemented"
	return *new(driver.Topic), nil
}

func (t *topic) SendBatch(ctx context.Context, dms []*driver.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *topic) IsRetryable(err error) bool { _ = "STUB: not implemented"; return false }

func (t *topic) As(i any) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func errorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (*topic) Close() error { _ = "STUB: not implemented"; return nil }

type subscription struct {
	sbReceiver *servicebus.Receiver
	opts       *SubscriptionOptions
}

type SubscriptionOptions struct {
	ReceiveAndDelete bool

	ReceiveBatcherOptions batcher.Options

	AckBatcherOptions batcher.Options

	ListenerTimeout time.Duration
}

func OpenSubscription(ctx context.Context, sbClient *servicebus.Client, sbReceiver *servicebus.Receiver, opts *SubscriptionOptions) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openSubscription(ctx context.Context, sbClient *servicebus.Client, sbReceiver *servicebus.Receiver, opts *SubscriptionOptions) (driver.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(driver.Subscription), nil
}

func (s *subscription) IsRetryable(err error) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) As(i any) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (s *subscription) ReceiveBatch(ctx context.Context, maxMessages int) ([]*driver.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func messageAsFunc(sbmsg *servicebus.ReceivedMessage) func(any) bool {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) SendAcks(ctx context.Context, ids []driver.AckID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) CanNack() bool { _ = "STUB: not implemented"; return false }

func (s *subscription) SendNacks(ctx context.Context, ids []driver.AckID) error {
	_ = "STUB: not implemented"
	return nil
}

func errorCode(err error) (gcerrors.ErrorCode, bool) {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode), false
}

func (*subscription) Close() error { _ = "STUB: not implemented"; return nil }
