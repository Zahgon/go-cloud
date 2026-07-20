package azureblob

import (
	"context"
	"io"
	"net/url"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	azblobblob "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/google/wire"
	"gocloud.dev/blob"
	"gocloud.dev/blob/driver"
	"gocloud.dev/gcerrors"
)

const (
	defaultPageSize        = 1000
	defaultUploadBuffers   = 5
	defaultUploadBlockSize = 8 * 1024 * 1024
)

func ptrVal[T any](p *T) (v T) { _ = "STUB: not implemented"; return *new(T) }

func init() {
	blob.DefaultURLMux().RegisterBucket(Scheme, new(lazyOpener))
}

var Set = wire.NewSet(
	NewDefaultServiceURLOptions,
	NewServiceURL,
	NewDefaultClient,
)

type Options struct{}

type ServiceURL string

type ContainerName string

type ServiceURLOptions struct {
	AccountName string

	SASToken string

	StorageDomain string

	Protocol string

	IsCDN bool

	IsLocalEmulator bool
}

func NewDefaultServiceURLOptions() *ServiceURLOptions { _ = "STUB: not implemented"; return nil }

func (o *ServiceURLOptions) withOverrides(urlValues url.Values) (*ServiceURLOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewServiceURL(opts *ServiceURLOptions) (ServiceURL, error) {
	_ = "STUB: not implemented"
	return *new(ServiceURL), nil
}

type lazyOpener struct {
	init   sync.Once
	opener *URLOpener
}

func (o *lazyOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type credTypeEnumT int

const (
	credTypeDefault credTypeEnumT = iota
	credTypeSharedKey
	credTypeSASViaNone
	credTypeConnectionString
)

type credInfoT struct {
	CredType credTypeEnumT

	AccountName string
	AccountKey  string

	ConnectionString string
}

func newCredInfoFromEnv() *credInfoT { _ = "STUB: not implemented"; return nil }

func (i *credInfoT) NewClient(svcURL ServiceURL, containerName ContainerName) (*container.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "azblob"

type URLOpener struct {
	MakeClient func(svcURL ServiceURL, containerName ContainerName) (*container.Client, error)

	ServiceURLOptions ServiceURLOptions

	Options Options
}

func (o *URLOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type bucket struct {
	client *container.Client
	opts   *Options
}

func NewDefaultClient(svcURL ServiceURL, containerName ContainerName) (*container.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenBucket(ctx context.Context, client *container.Client, opts *Options) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func openBucket(ctx context.Context, client *container.Client, opts *Options) (*bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bucket) Close() error { _ = "STUB: not implemented"; return nil }

func (b *bucket) Copy(ctx context.Context, dstKey, srcKey string, opts *driver.CopyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

type reader struct {
	body  io.ReadCloser
	attrs driver.ReaderAttributes
	raw   *azblobblob.DownloadStreamResponse
}

func (r *reader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *reader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *reader) Attributes() *driver.ReaderAttributes { _ = "STUB: not implemented"; return nil }

func (r *reader) As(i any) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) NewRangeReader(ctx context.Context, key string, offset, length int64, opts *driver.ReaderOptions) (driver.Reader, error) {
	_ = "STUB: not implemented"
	return *new(driver.Reader), nil
}

func getSize(contentLength *int64, contentRange string) int64 { _ = "STUB: not implemented"; return 0 }

func (b *bucket) As(i any) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (b *bucket) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (b *bucket) Attributes(ctx context.Context, key string) (*driver.Attributes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bucket) ListPaged(ctx context.Context, opts *driver.ListOptions) (*driver.ListPage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bucket) SignedURL(ctx context.Context, key string, opts *driver.SignedURLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type writer struct {
	ctx        context.Context
	client     *blockblob.Client
	uploadOpts *azblob.UploadStreamOptions

	pw *io.PipeWriter
	pr *io.PipeReader

	upload bool

	donec chan struct{}

	err error
}

func escapeKey(key string, isPrefix bool) string { _ = "STUB: not implemented"; return "" }

func unescapeKey(key string) string { _ = "STUB: not implemented"; return "" }

func (b *bucket) NewTypedWriter(ctx context.Context, key, contentType string, opts *driver.WriterOptions) (driver.Writer, error) {
	_ = "STUB: not implemented"
	return *new(driver.Writer), nil
}

func (w *writer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *writer) Upload(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (w *writer) open(r io.Reader, closePipeOnError bool) { _ = "STUB: not implemented"; return }

func (w *writer) Close() (err error) { _ = "STUB: not implemented"; return nil }
