package fileblob

import (
	"context"
	"hash"
	"io"
	"net/url"
	"os"
	"sync"

	"gocloud.dev/blob"
	"gocloud.dev/blob/driver"
	"gocloud.dev/gcerrors"
)

const defaultPageSize = 1000

func init() {
	blob.DefaultURLMux().RegisterBucket(Scheme, &URLOpener{})
}

const Scheme = "file"

type URLOpener struct {
	Options Options
}

func (o *URLOpener) OpenBucketURL(ctx context.Context, u *url.URL) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var recognizedParams = map[string]bool{
	"create_dir":      true,
	"base_url":        true,
	"secret_key_path": true,
	"metadata":        true,
	"no_tmp_dir":      true,
	"dir_file_mode":   true,
}

type metadataOption string

const (
	MetadataInSidecar metadataOption = ""

	MetadataDontWrite metadataOption = "skip"
)

func (o *URLOpener) forParams(ctx context.Context, q url.Values) (*Options, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Options struct {
	URLSigner URLSigner

	CreateDir bool

	DirFileMode os.FileMode

	NoTempDir bool

	Metadata metadataOption
}

type bucket struct {
	dir  string
	opts *Options
}

func openBucket(dir string, opts *Options) (driver.Bucket, error) {
	_ = "STUB: not implemented"
	return *new(driver.Bucket), nil
}

func OpenBucket(dir string, opts *Options) (*blob.Bucket, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bucket) Close() error { _ = "STUB: not implemented"; return nil }

func escapeKey(s string) string { _ = "STUB: not implemented"; return "" }

func unescapeKey(s string) string { _ = "STUB: not implemented"; return "" }

func (b *bucket) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}

func (b *bucket) path(key string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (b *bucket) forKey(key string) (string, os.FileInfo, *xattrs, error) {
	_ = "STUB: not implemented"
	return "", *new(os.FileInfo), nil, nil
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

func (b *bucket) isEscapingPath(path string) bool { _ = "STUB: not implemented"; return false }

type reader struct {
	r     io.Reader
	c     io.Closer
	attrs driver.ReaderAttributes
}

func (r *reader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *reader) Close() error { _ = "STUB: not implemented"; return nil }

func (r *reader) Attributes() *driver.ReaderAttributes { _ = "STUB: not implemented"; return nil }

func (r *reader) As(i any) bool { _ = "STUB: not implemented"; return false }

func createTemp(path string, noTempDir bool) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *bucket) NewTypedWriter(ctx context.Context, key, contentType string, opts *driver.WriterOptions) (driver.Writer, error) {
	_ = "STUB: not implemented"
	return *new(driver.Writer), nil
}

type writerWithSidecar struct {
	ctx        context.Context
	f          *os.File
	path       string
	attrs      xattrs
	contentMD5 []byte

	md5hash    hash.Hash
	ifNotExist bool
	mu         *sync.Mutex
}

func (w *writerWithSidecar) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *writerWithSidecar) Close() error { _ = "STUB: not implemented"; return nil }

type writer struct {
	*os.File
	ctx        context.Context
	path       string
	ifNotExist bool
	mu         *sync.Mutex
}

func (w *writer) Upload(r io.Reader) error { _ = "STUB: not implemented"; return nil }

func (w *writer) Close() error { _ = "STUB: not implemented"; return nil }

func (b *bucket) Copy(ctx context.Context, dstKey, srcKey string, opts *driver.CopyOptions) (err error) {
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

type URLSigner interface {
	URLFromKey(ctx context.Context, key string, opts *driver.SignedURLOptions) (*url.URL, error)

	KeyFromURL(ctx context.Context, surl *url.URL) (string, error)
}

type URLSignerHMAC struct {
	baseURL   *url.URL
	secretKey []byte
}

func NewURLSignerHMAC(baseURL *url.URL, secretKey []byte) *URLSignerHMAC {
	_ = "STUB: not implemented"
	return nil
}

func (h *URLSignerHMAC) URLFromKey(ctx context.Context, key string, opts *driver.SignedURLOptions) (*url.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *URLSignerHMAC) getMAC(q url.Values) string { _ = "STUB: not implemented"; return "" }

func (h *URLSignerHMAC) KeyFromURL(ctx context.Context, sURL *url.URL) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (h *URLSignerHMAC) checkMAC(q url.Values) bool { _ = "STUB: not implemented"; return false }
