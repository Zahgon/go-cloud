package memblob

import (
	"bytes"
	"context"
	"errors"
	"hash"
	"io"
	"net/url"
	"sync"

	"gocloud.dev/blob"
	"gocloud.dev/blob/driver"
	"gocloud.dev/gcerrors"
)

const defaultPageSize = 1000

var (
	errNotFound       = errors.New("blob not found")
	errNotImplemented = errors.New("not implemented")
)

func init() {
	blob.DefaultURLMux().RegisterBucket(Scheme, &URLOpener{})
}

const Scheme = "mem"

type URLOpener struct{}

func (*URLOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Options struct {
	NoMD5 bool
}

type blobEntry struct {
	Content    []byte
	Attributes *driver.Attributes
}

type bucket struct {
	options Options

	mu    sync.Mutex
	blobs map[string]*blobEntry
}

func openBucket(opts *Options) driver.Bucket { _ = "STUB: not implemented"; return *new(driver.Bucket) }

func OpenBucket(opts *Options) *blob.Bucket { _ = "STUB: not implemented"; return nil }

func (b *bucket) Close() error { _ = "STUB: not implemented"; return nil }

func (b *bucket) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (b *bucket) ListPaged(ctx context.Context, opts *driver.ListOptions) (*driver.ListPage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bucket) As(i any) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) Attributes(ctx context.Context, key string) (*driver.Attributes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bucket) NewRangeReader(ctx context.Context, key string, offset, length int64, opts *driver.ReaderOptions) (driver.Reader, error) {
	_ = "STUB: not implemented"
	return *new(driver.Reader), nil
}

type reader struct {
	r     io.Reader
	attrs driver.ReaderAttributes
}

func (r *reader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *reader) Download(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func (r *reader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *reader) Attributes() *driver.ReaderAttributes { _ = "STUB: not implemented"; return nil }

func (r *reader) As(i any) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) NewTypedWriter(ctx context.Context, key, contentType string, opts *driver.WriterOptions) (driver.Writer, error) {
	_ = "STUB: not implemented"
	return *new(driver.Writer), nil
}

type writer struct {
	ctx         context.Context
	b           *bucket
	key         string
	contentType string
	metadata    map[string]string
	opts        *driver.WriterOptions
	buf         bytes.Buffer

	md5hash    hash.Hash
	ifNotExist bool
}

func (w *writer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (w *writer) Upload(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (w *writer) Close() error { _ = "STUB: not implemented"; return nil }

func (b *bucket) Copy(ctx context.Context, dstKey, srcKey string, opts *driver.CopyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) Delete(ctx context.Context, key string, opts *driver.DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) SignedURL(ctx context.Context, key string, opts *driver.SignedURLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
