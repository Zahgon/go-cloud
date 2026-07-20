package drivertest

import (
	"context"
	"io"
	"testing"

	"gocloud.dev/secrets"
	"gocloud.dev/secrets/driver"
)

type Harness interface {
	MakeDriver(ctx context.Context) (driver.Keeper, driver.Keeper, error)

	Close()
}

type HarnessMaker func(ctx context.Context, t *testing.T) (Harness, error)

type AsTest interface {
	Name() string

	ErrorCheck(k *secrets.Keeper, err error) error
}

type verifyAsFailsOnNil struct{}

func (v verifyAsFailsOnNil) Name() string { _ = "STUB: not implemented"; return "" }

func (v verifyAsFailsOnNil) ErrorCheck(k *secrets.Keeper, err error) (ret error) {
	_ = "STUB: not implemented"
	return nil
}

func closeWithErrorCheck(t testing.TB, c io.Closer) { _ = "STUB: not implemented"; return }

func RunConformanceTests(t *testing.T, newHarness HarnessMaker, asTests []AsTest) {
	_ = "STUB: not implemented"
	return
}

func testEncryptDecrypt(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testMultipleEncryptionsNotEqual(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testMultipleKeys(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testDecryptMalformedError(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testAs(t *testing.T, newHarness HarnessMaker, tc AsTest) { _ = "STUB: not implemented"; return }
