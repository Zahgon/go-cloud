package pubsub

import (
	"context"
	"net/url"
	"sync"
	"time"

	"gocloud.dev/gcerrors"
	"gocloud.dev/internal/gcerr"
	"gocloud.dev/internal/openurl"
	gcdkotel "gocloud.dev/internal/otel"
	"gocloud.dev/pubsub/batcher"
	"gocloud.dev/pubsub/driver"
)

type Message struct {
	LoggableID string

	Body []byte

	Metadata map[string]string

	BeforeSend func(asFunc func(any) bool) error

	AfterSend func(asFunc func(any) bool) error

	asFunc func(any) bool

	ack func(isAck bool)

	nackable bool

	mu sync.Mutex

	isAcked bool
}

func (m *Message) Ack() { _ = "STUB: not implemented"; return }

func (m *Message) Nackable() bool { _ = "STUB: not implemented"; return false }

func (m *Message) Nack() { _ = "STUB: not implemented"; return }

func (m *Message) As(i any) bool { _ = "STUB: not implemented"; return false }

type Topic struct {
	driver  driver.Topic
	batcher *batcher.Batcher
	tracer  *gcdkotel.Tracer
	mu      sync.Mutex
	err     error

	cancel func()
}

func (t *Topic) Send(ctx context.Context, m *Message) (err error) {
	_ = "STUB: not implemented"
	return nil
}

var errTopicShutdown = gcerr.Newf(gcerr.FailedPrecondition, nil, "pubsub: Topic has been Shutdown")

func (t *Topic) Shutdown(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (t *Topic) As(i any) bool { _ = "STUB: not implemented"; return false }

func (t *Topic) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

var NewTopic = newTopic

func newSendBatcher(ctx context.Context, t *Topic, dt driver.Topic, opts *batcher.Options) *batcher.Batcher {
	_ = "STUB: not implemented"
	return nil
}

func newTopic(d driver.Topic, opts *batcher.Options) *Topic { _ = "STUB: not implemented"; return nil }

const pkgName = "gocloud.dev/pubsub"

var (
	OpenTelemetryViews = gcdkotel.Views(pkgName)
)

type Subscription struct {
	driver driver.Subscription
	tracer *gcdkotel.Tracer

	ackBatcher    *batcher.Batcher
	canNack       bool
	backgroundCtx context.Context
	cancel        func()

	recvBatchOpts *batcher.Options

	mu               sync.Mutex
	q                []*driver.Message
	err              error
	unreportedAckErr error
	waitc            chan struct{}
	runningBatchSize float64
	throughputStart  time.Time
	throughputCount  int

	preReceiveBatchHook func(maxMessages int)
}

const (
	desiredQueueDuration = 2 * time.Second

	expectedReceiveBatchDuration = 1 * time.Second

	prefetchRatio = float64(expectedReceiveBatchDuration) / float64(desiredQueueDuration)

	initialBatchSize = 1

	decay = 0.5

	maxGrowthFactor = 2.0

	maxShrinkFactor = 0.75

	maxBatchSize = 3000
)

func (s *Subscription) updateBatchSize() int { _ = "STUB: not implemented"; return 0 }

func (s *Subscription) Receive(ctx context.Context) (_ *Message, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type msgsOrError struct {
	msgs []*driver.Message
	err  error
}

func (s *Subscription) getNextBatch(nMessages int) chan msgsOrError {
	_ = "STUB: not implemented"
	return nil
}

var errSubscriptionShutdown = gcerr.Newf(gcerr.FailedPrecondition, nil, "pubsub: Subscription has been Shutdown")

func (s *Subscription) Shutdown(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Subscription) As(i any) bool { _ = "STUB: not implemented"; return false }

func (s *Subscription) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

var NewSubscription = newSubscription

func newSubscription(ds driver.Subscription, recvBatchOpts, ackBatcherOpts *batcher.Options) *Subscription {
	_ = "STUB: not implemented"
	return nil
}

func newAckBatcher(ctx context.Context, s *Subscription, ds driver.Subscription, opts *batcher.Options) *batcher.Batcher {
	_ = "STUB: not implemented"
	return nil
}

type errorCoder interface {
	ErrorCode(error) gcerrors.ErrorCode
}

func wrapError(ec errorCoder, err error) error { _ = "STUB: not implemented"; return nil }

type TopicURLOpener interface {
	OpenTopicURL(ctx context.Context, u *url.URL) (*Topic, error)
}

type SubscriptionURLOpener interface {
	OpenSubscriptionURL(ctx context.Context, u *url.URL) (*Subscription, error)
}

type URLMux struct {
	subscriptionSchemes openurl.SchemeMap
	topicSchemes        openurl.SchemeMap
}

func (mux *URLMux) TopicSchemes() []string { _ = "STUB: not implemented"; return nil }

func (mux *URLMux) ValidTopicScheme(scheme string) bool { _ = "STUB: not implemented"; return false }

func (mux *URLMux) SubscriptionSchemes() []string { _ = "STUB: not implemented"; return nil }

func (mux *URLMux) ValidSubscriptionScheme(scheme string) bool {
	_ = "STUB: not implemented"
	return false
}

func (mux *URLMux) RegisterTopic(scheme string, opener TopicURLOpener) {
	_ = "STUB: not implemented"
	return
}

func (mux *URLMux) RegisterSubscription(scheme string, opener SubscriptionURLOpener) {
	_ = "STUB: not implemented"
	return
}

func (mux *URLMux) OpenTopic(ctx context.Context, urlstr string) (*Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mux *URLMux) OpenSubscription(ctx context.Context, urlstr string) (*Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mux *URLMux) OpenTopicURL(ctx context.Context, u *url.URL) (*Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mux *URLMux) OpenSubscriptionURL(ctx context.Context, u *url.URL) (*Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var defaultURLMux = &URLMux{}

func DefaultURLMux() *URLMux { _ = "STUB: not implemented"; return nil }

func OpenTopic(ctx context.Context, urlstr string) (*Topic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenSubscription(ctx context.Context, urlstr string) (*Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
