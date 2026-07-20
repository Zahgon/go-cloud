package natspubsub

import (
	"context"
	"errors"
	"net/url"
	"regexp"
	"sync"

	"github.com/nats-io/nats.go"

	"gocloud.dev/gcerrors"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/batcher"
	"gocloud.dev/pubsub/driver"
)

var errNotInitialized = errors.New("natspubsub: topic not initialized")

var recvBatcherOpts = &batcher.Options{

	MaxBatchSize: 1,
	MaxHandlers:  1,
}

func init() {
	o := new(defaultDialer)
	pubsub.DefaultURLMux().RegisterTopic(Scheme, o)
	pubsub.DefaultURLMux().RegisterSubscription(Scheme, o)
}

type defaultDialer struct {
	init sync.Once
	err  error

	opener   URLOpener
	openerV2 URLOpener
}

func (o *defaultDialer) defaultConn(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type serverVersion struct {
	major, minor, patch int
}

func (o *defaultDialer) OpenTopicURL(ctx context.Context, u *url.URL) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *defaultDialer) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var semVerRegexp = regexp.MustCompile(`\Av?([0-9]+)\.?([0-9]+)?\.?([0-9]+)?`)

func parseServerVersion(version string) (serverVersion, error) {
	_ = "STUB: not implemented"
	return *new(serverVersion), nil
}

const Scheme = "nats"

type URLOpener struct {
	Connection *nats.Conn

	TopicOptions TopicOptions

	SubscriptionOptions SubscriptionOptions

	UseV2 bool
}

const natsV2QueryParameter = "natsv2"

func (o *URLOpener) OpenTopicURL(ctx context.Context, u *url.URL) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *URLOpener) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TopicOptions struct{}

type SubscriptionOptions struct {
	Queue string
}

type topic struct {
	useV2 bool
	nc    *nats.Conn
	subj  string
}

func OpenTopic(nc *nats.Conn, subject string, _ *TopicOptions) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenTopicV2(nc *nats.Conn, subject string, _ *TopicOptions) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openTopic(nc *nats.Conn, subject string, useV2 bool) (driver.Topic, error) {
	_ = "STUB: not implemented"
	return *new(driver.Topic), nil
}

func (t *topic) SendBatch(ctx context.Context, msgs []*driver.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *topic) sendMessage(m *driver.Message) error { _ = "STUB: not implemented"; return nil }

func (t *topic) sendMessageV2(m *driver.Message) error { _ = "STUB: not implemented"; return nil }

func (*topic) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (t *topic) As(i any) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorAs(error, any) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (*topic) Close() error { _ = "STUB: not implemented"; return nil }

type subscription struct {
	useV2  bool
	nc     *nats.Conn
	nsub   *nats.Subscription
	nextID int
}

func OpenSubscription(nc *nats.Conn, subject string, opts *SubscriptionOptions) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenSubscriptionV2(nc *nats.Conn, subject string, opts *SubscriptionOptions) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openSubscription(nc *nats.Conn, subject string, opts *SubscriptionOptions, useV2 bool) (driver.Subscription, error) {
	_ = "STUB: not implemented"
	return *new(driver.Subscription), nil
}

func (s *subscription) ReceiveBatch(ctx context.Context, _ int) ([]*driver.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decode(msg *nats.Msg) (*driver.Message, error) { _ = "STUB: not implemented"; return nil, nil }

func messageAsFunc(msg *nats.Msg) func(any) bool { _ = "STUB: not implemented"; return nil }

func (s *subscription) SendAcks(ctx context.Context, ids []driver.AckID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) CanNack() bool { _ = "STUB: not implemented"; return false }

func (s *subscription) SendNacks(ctx context.Context, ids []driver.AckID) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) As(i any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) ErrorAs(error, any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (*subscription) Close() error { _ = "STUB: not implemented"; return nil }

func encodeMessage(dm *driver.Message) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func decodeMessage(data []byte, dm *driver.Message) error { _ = "STUB: not implemented"; return nil }

func queryUseV2(q url.Values) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func encodeMessageV2(dm *driver.Message, sub string) *nats.Msg {
	_ = "STUB: not implemented"
	return nil
}

func decodeMessageV2(msg *nats.Msg) (*driver.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
