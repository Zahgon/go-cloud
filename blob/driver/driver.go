package driver

import (
	"context"
	"io"
	"time"

	"gocloud.dev/gcerrors"
)

type ReaderOptions struct {
	BeforeRead func(asFunc func(any) bool) error
}

type Reader interface {
	io.ReadCloser

	Attributes() *ReaderAttributes

	As(any) bool
}

type Downloader interface {
	Download(w io.Writer) error
}

type Writer interface {
	io.WriteCloser
}

type Uploader interface {
	Upload(r io.Reader) error
}

type WriterOptions struct {
	BufferSize int

	MaxConcurrency int

	CacheControl string

	ContentDisposition string

	ContentEncoding string

	ContentLanguage string

	ContentMD5 []byte

	Metadata map[string]string

	DisableContentTypeDetection bool

	BeforeWrite func(asFunc func(any) bool) error

	IfNotExist bool
}

type CopyOptions struct {
	BeforeCopy func(asFunc func(any) bool) error
}

type DeleteOptions struct {
	BeforeDelete func(asFunc func(any) bool) error
}

type ReaderAttributes struct {
	ContentType string

	ModTime time.Time

	Size int64
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

	AsFunc func(any) bool
}

type ListOptions struct {
	Prefix string

	Delimiter string

	PageSize int

	PageToken []byte

	BeforeList func(asFunc func(any) bool) error
}

type ListObject struct {
	Key string

	ModTime time.Time

	Size int64

	MD5 []byte

	IsDir bool

	AsFunc func(any) bool
}

type ListPage struct {
	Objects []*ListObject

	NextPageToken []byte
}

type Bucket interface {
	ErrorCode(error) gcerrors.ErrorCode

	As(i any) bool

	ErrorAs(error, any) bool

	Attributes(ctx context.Context, key string) (*Attributes, error)

	ListPaged(ctx context.Context, opts *ListOptions) (*ListPage, error)

	NewRangeReader(ctx context.Context, key string, offset, length int64, opts *ReaderOptions) (Reader, error)

	NewTypedWriter(ctx context.Context, key, contentType string, opts *WriterOptions) (Writer, error)

	Copy(ctx context.Context, dstKey, srcKey string, opts *CopyOptions) error

	Delete(ctx context.Context, key string, opts *DeleteOptions) error

	SignedURL(ctx context.Context, key string, opts *SignedURLOptions) (string, error)

	Close() error
}

type SignedURLOptions struct {
	Expiry time.Duration

	Method string

	ContentType string

	EnforceAbsentContentType bool

	BeforeSign func(asFunc func(any) bool) error
}

type prefixedBucket struct {
	base   Bucket
	prefix string
}

func NewPrefixedBucket(b Bucket, prefix string) Bucket {
	_ = "STUB: not implemented"
	return *new(Bucket)
}

func (b *prefixedBucket) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
func (b *prefixedBucket) As(i any) bool                 { _ = "STUB: not implemented"; return false }
func (b *prefixedBucket) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }
func (b *prefixedBucket) Attributes(ctx context.Context, key string) (*Attributes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *prefixedBucket) ListPaged(ctx context.Context, opts *ListOptions) (*ListPage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *prefixedBucket) NewRangeReader(ctx context.Context, key string, offset, length int64, opts *ReaderOptions) (Reader, error) {
	_ = "STUB: not implemented"
	return *new(Reader), nil
}

func (b *prefixedBucket) NewTypedWriter(ctx context.Context, key, contentType string, opts *WriterOptions) (Writer, error) {
	_ = "STUB: not implemented"
	return *new(Writer), nil
}

func (b *prefixedBucket) Copy(ctx context.Context, dstKey, srcKey string, opts *CopyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *prefixedBucket) Delete(ctx context.Context, key string, opts *DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *prefixedBucket) SignedURL(ctx context.Context, key string, opts *SignedURLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *prefixedBucket) Close() error { _ = "STUB: not implemented"; return nil }

type singleKeyBucket struct {
	base Bucket
	key  string
}

func NewSingleKeyBucket(b Bucket, key string) Bucket {
	_ = "STUB: not implemented"
	return *new(Bucket)
}

func (b *singleKeyBucket) ErrorCode(err error) gcerrors.ErrorCode {
	_ = "STUB: not implemented"
	return *new(gcerrors.ErrorCode)
}
func (b *singleKeyBucket) As(i any) bool                 { _ = "STUB: not implemented"; return false }
func (b *singleKeyBucket) ErrorAs(err error, i any) bool { _ = "STUB: not implemented"; return false }
func (b *singleKeyBucket) Attributes(ctx context.Context, _ string) (*Attributes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *singleKeyBucket) ListPaged(ctx context.Context, opts *ListOptions) (*ListPage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *singleKeyBucket) NewRangeReader(ctx context.Context, _ string, offset, length int64, opts *ReaderOptions) (Reader, error) {
	_ = "STUB: not implemented"
	return *new(Reader), nil
}

func (b *singleKeyBucket) NewTypedWriter(ctx context.Context, _, contentType string, opts *WriterOptions) (Writer, error) {
	_ = "STUB: not implemented"
	return *new(Writer), nil
}

func (b *singleKeyBucket) Copy(ctx context.Context, dstKey, _ string, opts *CopyOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *singleKeyBucket) Delete(ctx context.Context, _ string, opts *DeleteOptions) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *singleKeyBucket) SignedURL(ctx context.Context, _ string, opts *SignedURLOptions) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (b *singleKeyBucket) Close() error { _ = "STUB: not implemented"; return nil }
