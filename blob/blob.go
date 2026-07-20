package blob

import (
	"bytes"
	"context"
	"hash"
	"io"
	"iter"
	"net/url"
	"sync"
	"time"

	"go.opentelemetry.io/otel/metric"
	"gocloud.dev/blob/driver"
	"gocloud.dev/internal/gcerr"
	"gocloud.dev/internal/openurl"
	gcdkotel "gocloud.dev/internal/otel"
)

var _ = io.ReadSeekCloser(&Reader{})

type Reader struct {
	b              driver.Bucket
	r              driver.Reader
	key            string
	ctx            context.Context
	dopts          *driver.ReaderOptions
	baseOffset     int64
	baseLength     int64
	relativeOffset int64
	savedOffset    int64
	end            func(error)

	bytesReadCounter metric.Int64Counter
	bytesRead        int
	closed           bool
}

func (r *Reader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (r *Reader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *Reader) ContentType() string { _ = "STUB: not implemented"; return "" }

func (r *Reader) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (r *Reader) Size() int64 { _ = "STUB: not implemented"; return 0 }

func (r *Reader) As(i any) bool { _ = "STUB: not implemented"; return false }

func (r *Reader) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *Reader) downloadAndClose(w io.Writer) (err error) { _ = "STUB: not implemented"; return nil }

func readFromWriteTo(r io.Reader, w io.Writer) (int64, int64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

type Attributes struct {
	CacheControl string

	ContentDisposition string

	ContentEncoding string

	ContentLanguage string

	ContentType string

	Metadata map[string]string

	CreateTime time.Time

	ModTime time.Time

	Size int64

	MD5 []byte

	ETag string

	asFunc func(any) bool
}

func (a *Attributes) As(i any) bool { _ = "STUB: not implemented"; return false }

type Writer struct {
	b          driver.Bucket
	w          driver.Writer
	key        string
	end        func(err error)
	cancel     func()
	contentMD5 []byte
	md5hash    hash.Hash

	bytesWrittenCounter metric.Int64Counter
	bytesWritten        int
	closed              bool

	ctx  context.Context
	opts *driver.WriterOptions
	buf  *bytes.Buffer
}

const sniffLen = 512

func (w *Writer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *Writer) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (w *Writer) open(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *Writer) write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *Writer) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *Writer) uploadAndClose(r io.Reader) (err error) { _ = "STUB: not implemented"; return nil }

type ListOptions struct {
	Prefix string

	Delimiter string

	BeforeList func(asFunc func(any) bool) error
}

type ListIterator struct {
	b       *Bucket
	opts    *driver.ListOptions
	page    *driver.ListPage
	nextIdx int
}

func (i *ListIterator) Next(ctx context.Context) (*ListObject, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type errorState struct {
	mu   sync.Mutex
	done bool
	err  error
}

func (es *errorState) Done() { _ = "STUB: not implemented"; return }

func (es *errorState) Set(err error) { _ = "STUB: not implemented"; return }

func (es *errorState) Err() error { _ = "STUB: not implemented"; return nil }

func (es *errorState) Func() func() error { _ = "STUB: not implemented"; return nil }

func (i *ListIterator) All(ctx context.Context) (iter.Seq2[*ListObject, func(io.Writer, *ReaderOptions) error], func() error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ListObject struct {
	Key string

	ModTime time.Time

	Size int64

	MD5 []byte

	IsDir bool

	asFunc func(any) bool
}

func (o *ListObject) As(i any) bool { _ = "STUB: not implemented"; return false }

type Bucket struct {
	b      driver.Bucket
	tracer *gcdkotel.Tracer

	bytesReadCounter    metric.Int64Counter
	bytesWrittenCounter metric.Int64Counter

	ioFSCallback func() (context.Context, *ReaderOptions)

	mu     sync.RWMutex
	closed bool
}

const pkgName = "gocloud.dev/blob"

var (
	OpenTelemetryViews = append(
		append(
			gcdkotel.Views(pkgName),
			gcdkotel.CounterView(pkgName, "/bytes_read", "Sum of bytes read from the service.")...),
		gcdkotel.CounterView(pkgName, "/bytes_written", "Sum of bytes written to the service.")...)
)

var NewBucket = newBucket

func newBucket(b driver.Bucket) *Bucket { _ = "STUB: not implemented"; return nil }

func (b *Bucket) As(i any) bool { _ = "STUB: not implemented"; return false }

func (b *Bucket) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }

func (b *Bucket) ReadAll(ctx context.Context, key string) (_ []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bucket) Download(ctx context.Context, key string, w io.Writer, opts *ReaderOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bucket) List(opts *ListOptions) *ListIterator { _ = "STUB: not implemented"; return nil }

var FirstPageToken = []byte("first page")

func (b *Bucket) ListPage(ctx context.Context, pageToken []byte, pageSize int, opts *ListOptions) (retval []*ListObject, nextPageToken []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (b *Bucket) IsAccessible(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *Bucket) Exists(ctx context.Context, key string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (b *Bucket) Attributes(ctx context.Context, key string) (_ *Attributes, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bucket) NewReader(ctx context.Context, key string, opts *ReaderOptions) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bucket) NewRangeReader(ctx context.Context, key string, offset, length int64, opts *ReaderOptions) (_ *Reader, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bucket) newRangeReader(ctx context.Context, key string, offset, length int64, opts *ReaderOptions) (_ *Reader, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bucket) WriteAll(ctx context.Context, key string, p []byte, opts *WriterOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bucket) Upload(ctx context.Context, key string, r io.Reader, opts *WriterOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bucket) NewWriter(ctx context.Context, key string, opts *WriterOptions) (_ *Writer, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *Bucket) Copy(ctx context.Context, dstKey, srcKey string, opts *CopyOptions) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bucket) Delete(ctx context.Context, key string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *Bucket) SignedURL(ctx context.Context, key string, opts *SignedURLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *Bucket) Close() error { _ = "STUB: not implemented"; return nil }

const DefaultSignedURLExpiry = 1 * time.Hour

type SignedURLOptions struct {
	Expiry time.Duration

	Method string

	ContentType string

	EnforceAbsentContentType bool

	BeforeSign func(asFunc func(any) bool) error
}

type ReaderOptions struct {
	BeforeRead func(asFunc func(any) bool) error
}

type WriterOptions struct {
	BufferSize int

	MaxConcurrency int

	CacheControl string

	ContentDisposition string

	ContentEncoding string

	ContentLanguage string

	ContentType string

	DisableContentTypeDetection bool

	ContentMD5 []byte

	Metadata map[string]string

	BeforeWrite func(asFunc func(any) bool) error

	IfNotExist bool
}

type CopyOptions struct {
	BeforeCopy func(asFunc func(any) bool) error
}

type BucketURLOpener interface {
	OpenBucketURL(ctx context.Context, u *url.URL) (*Bucket, error)
}

type URLMux struct {
	schemes openurl.SchemeMap
}

func (mux *URLMux) BucketSchemes() []string { _ = "STUB: not implemented"; return nil }

func (mux *URLMux) ValidBucketScheme(scheme string) bool { _ = "STUB: not implemented"; return false }

func (mux *URLMux) RegisterBucket(scheme string, opener BucketURLOpener) {
	_ = "STUB: not implemented"
	return
}

func (mux *URLMux) OpenBucket(ctx context.Context, urlstr string) (*Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mux *URLMux) OpenBucketURL(ctx context.Context, u *url.URL) (*Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func applyPrefixParam(ctx context.Context, opener BucketURLOpener, u *url.URL) (*Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var defaultURLMux = new(URLMux)

func DefaultURLMux() *URLMux { _ = "STUB: not implemented"; return nil }

func OpenBucket(ctx context.Context, urlstr string) (*Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func wrapError(b driver.Bucket, err error, key string) error { _ = "STUB: not implemented"; return nil }

var errClosed = gcerr.Newf(gcerr.FailedPrecondition, nil, "blob: Bucket has been closed")

func PrefixedBucket(bucket *Bucket, prefix string) *Bucket { _ = "STUB: not implemented"; return nil }

func SingleKeyBucket(bucket *Bucket, singleKey string) *Bucket {
	_ = "STUB: not implemented"
	return nil
}
