package gcppubsub

import (
	"context"
	"net/url"
	"regexp"
	"sync"

	raw "cloud.google.com/go/pubsub/apiv1"
	pb "cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"github.com/google/wire"
	"gocloud.dev/gcerrors"
	"gocloud.dev/gcp"
	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/batcher"
	"gocloud.dev/pubsub/driver"
	"google.golang.org/grpc"
)

var endPoint = "pubsub.googleapis.com:443"

var sendBatcherOpts = &batcher.Options{
	MaxBatchSize: 1000,
	MaxHandlers:  2,

	MaxBatchByteSize: 8 * 1024 * 1024,
}

var defaultRecvBatcherOpts = &batcher.Options{

	MaxBatchSize: 1000,
	MaxHandlers:  10,
}

var ackBatcherOpts = &batcher.Options{

	MaxBatchSize: 1000,
	MaxHandlers:  2,
}

func init() {
	o := new(lazyCredsOpener)
	pubsub.DefaultURLMux().RegisterTopic(Scheme, o)
	pubsub.DefaultURLMux().RegisterSubscription(Scheme, o)
}

var Set = wire.NewSet(
	Dial,
	PublisherClient,
	SubscriberClient,
	wire.Struct(new(SubscriptionOptions)),
	wire.Struct(new(TopicOptions)),
	wire.Struct(new(URLOpener), "Conn", "TopicOptions", "SubscriptionOptions"),
)

type lazyCredsOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *lazyCredsOpener) defaultConn(ctx context.Context) (*URLOpener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *lazyCredsOpener) OpenTopicURL(ctx context.Context, u *url.URL) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *lazyCredsOpener) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "gcppubsub"

type URLOpener struct {
	Conn *grpc.ClientConn

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
	path   string
	client *raw.PublisherClient
}

func Dial(ctx context.Context, ts gcp.TokenSource) (*grpc.ClientConn, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func dialEmulator(ctx context.Context, e string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PublisherClient(ctx context.Context, conn *grpc.ClientConn) (*raw.PublisherClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SubscriberClient(ctx context.Context, conn *grpc.ClientConn) (*raw.SubscriberClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type TopicOptions struct {
	BatcherOptions batcher.Options
}

func OpenTopic(client *raw.PublisherClient, projectID gcp.ProjectID, topicName string, opts *TopicOptions) *pubsub.Topic {
	_ = "STUB: not implemented"
	return nil
}

var topicPathRE = regexp.MustCompile("^projects/.+/topics/.+$")

func OpenTopicByPath(client *raw.PublisherClient, topicPath string, opts *TopicOptions) (*pubsub.Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openTopic(client *raw.PublisherClient, topicPath string) driver.Topic {
	_ = "STUB: not implemented"
	return *new(driver.Topic)
}

func (t *topic) SendBatch(ctx context.Context, dms []*driver.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *topic) IsRetryable(error) bool { _ = "STUB: not implemented"; return false }

func (t *topic) As(i any) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func errorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*topic) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (*topic) Close() error { _ = "STUB: not implemented"; return nil }

type subscription struct {
	client  *raw.SubscriberClient
	path    string
	options *SubscriptionOptions
}

type SubscriptionOptions struct {
	MaxBatchSize int

	NackLazy bool

	ReceiveBatcherOptions batcher.Options

	AckBatcherOptions batcher.Options
}

func OpenSubscription(client *raw.SubscriberClient, projectID gcp.ProjectID, subscriptionName string, opts *SubscriptionOptions) *pubsub.Subscription {
	_ = "STUB: not implemented"
	return nil
}

var subscriptionPathRE = regexp.MustCompile("^projects/.+/subscriptions/.+$")

func OpenSubscriptionByPath(client *raw.SubscriberClient, subscriptionPath string, opts *SubscriptionOptions) (*pubsub.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openSubscription(client *raw.SubscriberClient, subscriptionPath string, opts *SubscriptionOptions) *subscription {
	_ = "STUB: not implemented"
	return nil
}

func (s *subscription) ReceiveBatch(ctx context.Context, maxMessages int) ([]*driver.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func messageAsFunc(pm *pb.PubsubMessage, rm *pb.ReceivedMessage) func(any) bool {
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

func (s *subscription) IsRetryable(err error) bool { _ = "STUB: not implemented"; return false }

func (s *subscription) As(i any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (*subscription) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (*subscription) Close() error { _ = "STUB: not implemented"; return nil }

func queryParameterInt(value []string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func queryParameterBool(value []string) (bool, error) { _ = "STUB: not implemented"; return false, nil }
