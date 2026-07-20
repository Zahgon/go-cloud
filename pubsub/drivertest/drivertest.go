package drivertest

import (
	"context"
	"testing"

	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/batcher"
	"gocloud.dev/pubsub/driver"
)

type Harness interface {
	CreateTopic(ctx context.Context, testName string) (dt driver.Topic, cleanup func(), err error)

	MakeNonexistentTopic(ctx context.Context) (driver.Topic, error)

	CreateSubscription(ctx context.Context, t driver.Topic, testName string) (ds driver.Subscription, cleanup func(), err error)

	MakeNonexistentSubscription(ctx context.Context) (ds driver.Subscription, cleanup func(), err error)

	Close()

	MaxBatchSizes() (int, int)

	SupportsMultipleSubscriptions() bool
}

type HarnessMaker func(ctx context.Context, t *testing.T) (Harness, error)

type AsTest interface {
	Name() string

	TopicCheck(t *pubsub.Topic) error

	SubscriptionCheck(s *pubsub.Subscription) error

	TopicErrorCheck(t *pubsub.Topic, err error) error

	SubscriptionErrorCheck(s *pubsub.Subscription, err error) error

	MessageCheck(m *pubsub.Message) error

	BeforeSend(as func(any) bool) error

	AfterSend(as func(any) bool) error
}

var batchSizeOne = &batcher.Options{MaxBatchSize: 1, MaxHandlers: 1}

type verifyAsFailsOnNil struct{}

func (verifyAsFailsOnNil) Name() string { _ = "STUB: not implemented"; return "" }

func (verifyAsFailsOnNil) TopicCheck(t *pubsub.Topic) error { _ = "STUB: not implemented"; return nil }

func (verifyAsFailsOnNil) SubscriptionCheck(s *pubsub.Subscription) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) TopicErrorCheck(t *pubsub.Topic, err error) (ret error) {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) SubscriptionErrorCheck(s *pubsub.Subscription, err error) (ret error) {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) MessageCheck(m *pubsub.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) BeforeSend(as func(any) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) AfterSend(as func(any) bool) error { _ = "STUB: not implemented"; return nil }

func RunConformanceTests(t *testing.T, newHarness HarnessMaker, asTests []AsTest) {
	_ = "STUB: not implemented"
	return
}

func RunBenchmarks(b *testing.B, topic *pubsub.Topic, sub *pubsub.Subscription) {
	_ = "STUB: not implemented"
	return
}

func testNonExistentTopicSucceedsOnOpenButFailsOnSend(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testNonExistentSubscriptionSucceedsOnOpenButFailsOnReceive(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testSendReceive(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testSendReceiveTwo(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testSendReceiveJSON(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testNack(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testBatching(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testDoubleAck(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func publishN(ctx context.Context, t *testing.T, topic *pubsub.Topic, n int) []*pubsub.Message {
	_ = "STUB: not implemented"
	return nil
}

func receiveN(ctx context.Context, t *testing.T, sub *pubsub.Subscription, n int) []*pubsub.Message {
	_ = "STUB: not implemented"
	return nil
}

func diffMessageSets(got, want []*pubsub.Message) string { _ = "STUB: not implemented"; return "" }

func testErrorOnSendToClosedTopic(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testErrorOnReceiveFromClosedSubscription(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testCancelSendReceive(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testMetadata(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testNonUTF8MessageBody(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func isCanceled(err error) bool { _ = "STUB: not implemented"; return false }

func makePair(ctx context.Context, t *testing.T, h Harness) (*pubsub.Topic, *pubsub.Subscription, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func testAs(t *testing.T, newHarness HarnessMaker, st AsTest) { _ = "STUB: not implemented"; return }

func benchmark(b *testing.B, topic *pubsub.Topic, sub *pubsub.Subscription, timeSend bool) {
	_ = "STUB: not implemented"
	return
}

func publishNConcurrently(topic *pubsub.Topic, nMessages, nGoroutines int, attrs map[string]string, body []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func receiveNConcurrently(sub *pubsub.Subscription, nMessages, nGoroutines int) error {
	_ = "STUB: not implemented"
	return nil
}

func runConcurrently(n, g int, f func(context.Context) error) error {
	_ = "STUB: not implemented"
	return nil
}
