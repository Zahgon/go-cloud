package s3blob

import (
	"context"
	"io"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/wire"
	"gocloud.dev/blob"
	"gocloud.dev/blob/driver"
	"gocloud.dev/gcerrors"
)

const defaultPageSize = 1000

func init() {
	blob.DefaultURLMux().RegisterBucket(Scheme, new(urlSessionOpener))
}

var Set = wire.NewSet(
	Dial,
)

func Dial(cfg aws.Config) *s3.Client { _ = "STUB: not implemented"; return nil }

type urlSessionOpener struct{}

func (o *urlSessionOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const Scheme = "s3"

type URLOpener struct {
	Options Options
}

const (
	sseTypeParamKey            = "ssetype"
	kmsKeyIdParamKey           = "kmskeyid"
	accelerateParamKey         = "accelerate"
	usePathStyleParamKey       = "use_path_style"
	legacyUsePathStyleParamKey = "s3ForcePathStyle"
	disableHTTPSParamKey       = "disable_https"
)

func toServerSideEncryptionType(value string) (types.ServerSideEncryption, error) {
	_ = "STUB: not implemented"
	return *new(types.ServerSideEncryption), nil
}

func (o *URLOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Options struct {
	UseLegacyList bool

	EncryptionType types.ServerSideEncryption

	KMSEncryptionID string

	RequestChecksumCalculation aws.RequestChecksumCalculation
}

func openBucket(ctx context.Context, client *s3.Client, bucketName string, opts *Options) (*bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func OpenBucket(ctx context.Context, client *s3.Client, bucketName string, opts *Options) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var OpenBucketV2 = OpenBucket

type reader struct {
	body  io.ReadCloser
	attrs driver.ReaderAttributes
	raw   *s3.GetObjectOutput
}

func (r *reader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *reader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *reader) As(i any) bool { _ = "STUB: not implemented"; return false }

func (r *reader) Attributes() *driver.ReaderAttributes { _ = "STUB: not implemented"; return nil }

type writer struct {
	pw *io.PipeWriter
	pr *io.PipeReader

	upload bool

	ctx context.Context
	tm  *transfermanager.Client
	req *transfermanager.UploadObjectInput

	donec chan struct{}

	err error
}

func (w *writer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *writer) Upload(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (w *writer) open(r io.Reader, closePipeOnError bool) { _ = "STUB: not implemented"; return }

func (w *writer) Close() (err error) { _ = "STUB: not implemented"; return nil }

type bucket struct {
	name          string
	client        *s3.Client
	useLegacyList bool

	encryptionType             types.ServerSideEncryption
	kmsKeyId                   string
	requestChecksumCalculation aws.RequestChecksumCalculation
}

func (b *bucket) Close() error { _ = "STUB: not implemented"; return nil }

func (b *bucket) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (b *bucket) ListPaged(ctx context.Context, opts *driver.ListOptions) (*driver.ListPage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bucket) listObjects(ctx context.Context, in *s3.ListObjectsV2Input, opts *driver.ListOptions) (*s3.ListObjectsV2Output, error) {
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

func eTagToMD5(etag *string) []byte { _ = "STUB: not implemented"; return nil }

func getSize(contentLength int64, contentRange string) int64 { _ = "STUB: not implemented"; return 0 }

func escapeKey(key string) string { _ = "STUB: not implemented"; return "" }

func unescapeKey(key string) string { _ = "STUB: not implemented"; return "" }

func (b *bucket) NewTypedWriter(ctx context.Context, key, contentType string, opts *driver.WriterOptions) (driver.Writer, error) {
	_ = "STUB: not implemented"
	return *new(driver.Writer), nil
}

func (b *bucket) Copy(ctx context.Context, dstKey, srcKey string, opts *driver.CopyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *bucket) SignedURL(ctx context.Context, key string, opts *driver.SignedURLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
