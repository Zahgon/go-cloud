package drivertest

import (
	"context"
	"testing"
	"time"

	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/driver"
)

type Harness interface {
	MakeWatcher(ctx context.Context, name string, decoder *runtimevar.Decoder) (driver.Watcher, error)

	CreateVariable(ctx context.Context, name string, val []byte) error

	UpdateVariable(ctx context.Context, name string, val []byte) error

	DeleteVariable(ctx context.Context, name string) error

	Close()

	Mutable() bool
}

type HarnessMaker func(t *testing.T) (Harness, error)

type AsTest interface {
	Name() string

	SnapshotCheck(s *runtimevar.Snapshot) error

	ErrorCheck(v *runtimevar.Variable, err error) error
}

type verifyAsFailsOnNil struct{}

func (verifyAsFailsOnNil) Name() string { _ = "STUB: not implemented"; return "" }

func (verifyAsFailsOnNil) SnapshotCheck(v *runtimevar.Snapshot) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) ErrorCheck(v *runtimevar.Variable, err error) (ret error) {
	_ = "STUB: not implemented"
	return nil
}

func RunConformanceTests(t *testing.T, newHarness HarnessMaker, asTests []AsTest) {
	_ = "STUB: not implemented"
	return
}

func waitTimeForBlockingCheck() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func testNonExistentVariable(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testString(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

type Message struct {
	Name, Text string
}

func testJSON(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testInvalidJSON(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testUpdate(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testDelete(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testUpdateWithErrors(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testAs(t *testing.T, newHarness HarnessMaker, st AsTest) { _ = "STUB: not implemented"; return }
