package kafkapubsub

import (
	"context"
	"net/url"
	"sync"
	"time"

	"github.com/IBM/sarama"
	"gocloud.dev/gcerrors"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/batcher"
	"gocloud.dev/pubsub/driver"
)

var sendBatcherOpts = &batcher.Options{
	MaxBatchSize: 100,
	MaxHandlers:  100,
}

var recvBatcherOpts = &batcher.Options{

	MaxBatchSize: 1,
	MaxHandlers:  1,
}

func init() {
	opener := new(defaultOpener)
	pubsub.DefaultURLMux().RegisterTopic(Scheme, opener)
	pubsub.DefaultURLMux().RegisterSubscription(Scheme, opener)
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

const Scheme = "kafka"

type URLOpener struct {
	Brokers []string

	Config *sarama.Config

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

func MinimalConfig() *sarama.Config { _ = "STUB: not implemented"; return nil }

type topic struct {
	producer  sarama.SyncProducer
	topicName string
	opts      TopicOptions
}

type TopicOptions struct {
	KeyName string

	BatcherOptions batcher.Options
}

func OpenTopic(brokers []string, config *sarama.Config, topicName string, opts *TopicOptions) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openTopic(brokers []string, config *sarama.Config, topicName string, opts *TopicOptions) (*topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *topic) SendBatch(ctx context.Context, dms []*driver.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *topic) Close() error { _ = "STUB: not implemented"; return nil }

func (t *topic) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (t *topic) As(i any) bool { _ = "STUB: not implemented"; return false }

func (t *topic) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (t *topic) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func errorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

type subscription struct {
	opts          SubscriptionOptions
	closeCh       chan struct{}
	joinCh        chan struct{}
	cancel        func()
	closeErr      error
	consumerGroup sarama.ConsumerGroup

	mu             sync.Mutex
	unacked        []*ackInfo
	sess           sarama.ConsumerGroupSession
	expectedClaims int
	claims         []sarama.ConsumerGroupClaim
}

type ackInfo struct {
	msg   *sarama.ConsumerMessage
	acked bool
}

type SubscriptionOptions struct {
	KeyName string

	WaitForJoin time.Duration
}

func OpenSubscription(brokers []string, config *sarama.Config, group string, topics []string, opts *SubscriptionOptions) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openSubscription(brokers []string, config *sarama.Config, group string, topics []string, opts *SubscriptionOptions) (driver.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(driver.Subscription), nil
}

func (s *subscription) Setup(sess sarama.ConsumerGroupSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) Cleanup(sarama.ConsumerGroupSession) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) ReceiveBatch(ctx context.Context, maxMessages int) ([]*driver.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
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

func (s *subscription) Close() error { _ = "STUB: not implemented"; return nil }

func (*subscription) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) As(i any) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func errorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }
