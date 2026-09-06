package drivertest

import (
	"context"
	"io"
	"net/http"
	"testing"

	"gocloud.dev/blob"
	"gocloud.dev/blob/driver"
)

type Harness interface {
	MakeDriver(ctx context.Context) (driver.Bucket, error)

	MakeDriverForNonexistentBucket(ctx context.Context) (driver.Bucket, error)

	HTTPClient() *http.Client

	Close()
}

type HarnessMaker func(ctx context.Context, t *testing.T) (Harness, error)

type AsTest interface {
	Name() string

	BucketCheck(b *blob.Bucket) error

	ErrorCheck(b *blob.Bucket, err error) error

	BeforeRead(as func(any) bool) error

	BeforeWrite(as func(any) bool) error

	BeforeCopy(as func(any) bool) error

	BeforeDelete(as func(any) bool) error

	BeforeList(as func(any) bool) error

	BeforeSign(as func(any) bool) error

	AttributesCheck(attrs *blob.Attributes) error

	ReaderCheck(r *blob.Reader) error

	ListObjectCheck(o *blob.ListObject) error
}

func closeWithErrorCheck(t testing.TB, c io.Closer) { _ = "STUB: not implemented"; return }

func deleteWithErrorCheck(ctx context.Context, t testing.TB, b *blob.Bucket, key string) {
	_ = "STUB: not implemented"
	return
}

func writeWithErrorCheck(ctx context.Context, t testing.TB, b *blob.Bucket, key string, content []byte) {
	_ = "STUB: not implemented"
	return
}

type verifyAsFailsOnNil struct{}

func (verifyAsFailsOnNil) Name() string { _ = "STUB: not implemented"; return "" }

func (verifyAsFailsOnNil) BucketCheck(b *blob.Bucket) error { _ = "STUB: not implemented"; return nil }

func (verifyAsFailsOnNil) ErrorCheck(b *blob.Bucket, err error) (ret error) {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) BeforeRead(as func(any) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) BeforeWrite(as func(any) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) BeforeCopy(as func(any) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) BeforeDelete(as func(any) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) BeforeList(as func(any) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) BeforeSign(as func(any) bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) AttributesCheck(attrs *blob.Attributes) error {
	_ = "STUB: not implemented"
	return nil
}

func (verifyAsFailsOnNil) ReaderCheck(r *blob.Reader) error { _ = "STUB: not implemented"; return nil }

func (verifyAsFailsOnNil) ListObjectCheck(o *blob.ListObject) error {
	_ = "STUB: not implemented"
	return nil
}

func RunConformanceTests(t *testing.T, newHarness HarnessMaker, asTests []AsTest) {
	_ = "STUB: not implemented"
	return
}

func RunBenchmarks(b *testing.B, bkt *blob.Bucket) { _ = "STUB: not implemented"; return }

func testNonexistentBucket(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testList(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testListWeirdKeys(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

type listResult struct {
	Key   string
	IsDir bool

	Sub []listResult
}

func doList(ctx context.Context, b *blob.Bucket, prefix, delim string, recurse bool) ([]listResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func testListDelimiters(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testDirsWithCharactersBeforeDelimiter(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func iterToSetOfKeys(ctx context.Context, t *testing.T, iter *blob.ListIterator) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func testRead(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testAttributes(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func loadTestData(tb testing.TB, name string) []byte { _ = "STUB: not implemented"; return nil }

func testWrite(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testCanceledWrite(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testMetadata(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testMD5(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testCopy(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testDelete(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testConcurrentWriteAndRead(t *testing.T, newHarness HarnessMaker) {
	_ = "STUB: not implemented"
	return
}

func testUploadDownload(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testKeys(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testSignedURL(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func testAs(t *testing.T, newHarness HarnessMaker, st AsTest) { _ = "STUB: not implemented"; return }

func testIfNotExist(t *testing.T, newHarness HarnessMaker) { _ = "STUB: not implemented"; return }

func benchmarkRead(b *testing.B, bkt *blob.Bucket) { _ = "STUB: not implemented"; return }

func benchmarkWriteReadDelete(b *testing.B, bkt *blob.Bucket) { _ = "STUB: not implemented"; return }
