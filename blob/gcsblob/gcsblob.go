package gcsblob

import (
	"context"
	"io"
	"net/url"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/google/wire"
	"google.golang.org/api/option"

	"gocloud.dev/blob"
	"gocloud.dev/blob/driver"
	"gocloud.dev/gcerrors"
	"gocloud.dev/gcp"
)

const defaultPageSize = 1000

func init() {
	blob.DefaultURLMux().RegisterBucket(Scheme, new(lazyCredsOpener))
}

var Set = wire.NewSet(
	wire.Struct(new(URLOpener), "Client"),
)

func readDefaultCredentials(credFileAsJSON []byte) (AccessID string, PrivateKey []byte) {
	_ = "STUB: not implemented"
	return "", nil
}

type lazyCredsOpener struct {
	init   sync.Once
	opener *URLOpener
	err    error
}

func (o *lazyCredsOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "gs"

type URLOpener struct {
	Client *gcp.HTTPClient

	Options Options
}

func (o *URLOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *URLOpener) forParams(ctx context.Context, q url.Values) (*Options, *gcp.HTTPClient, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

type Options struct {
	GoogleAccessID string

	PrivateKey []byte

	SignBytes func([]byte) ([]byte, error)

	MakeSignBytes func(requestCtx context.Context) SignBytesFunc

	Client *storage.Client

	ClientOptions []option.ClientOption
}

func (o *Options) clear() { _ = "STUB: not implemented"; return }

type SignBytesFunc func([]byte) ([]byte, error)

func openBucket(ctx context.Context, client *gcp.HTTPClient, bucketName string, opts *Options) (*bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenBucket(ctx context.Context, client *gcp.HTTPClient, bucketName string, opts *Options) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type bucket struct {
	name   string
	client *storage.Client
	opts   *Options
}

type reader struct {
	body  io.ReadCloser
	attrs driver.ReaderAttributes
	raw   *storage.Reader
}

func (r *reader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *reader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *reader) Attributes() *driver.ReaderAttributes { _ = "STUB: not implemented"; return nil }

func (r *reader) As(i any) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (b *bucket) Close() error { _ = "STUB: not implemented"; return nil }

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

func escapeKey(key string) string { _ = "STUB: not implemented"; return "" }

func unescapeKey(key string) string { _ = "STUB: not implemented"; return "" }

func (b *bucket) NewTypedWriter(ctx context.Context, key, contentType string, opts *driver.WriterOptions) (driver.Writer, error) {
	_ = "STUB: not implemented"
	return *new(driver.Writer), nil
}

type CopyObjectHandles struct {
	Dst, Src *storage.ObjectHandle
}

func (b *bucket) Copy(ctx context.Context, dstKey, srcKey string, opts *driver.CopyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) Delete(ctx context.Context, key string, opts *driver.DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) SignedURL(ctx context.Context, key string, dopts *driver.SignedURLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func bufferSize(size int) int { _ = "STUB: not implemented"; return 0 }
