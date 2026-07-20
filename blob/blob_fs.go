package blob

import (
	"context"
	"io/fs"
	"time"
)

var (
	_ = fs.FS(&Bucket{})
	_ = fs.SubFS(&Bucket{})
)

type iofsFileInfo struct {
	lo   *ListObject
	name string
}

func (f *iofsFileInfo) Name() string       { _ = "STUB: not implemented"; return "" }
func (f *iofsFileInfo) Size() int64        { _ = "STUB: not implemented"; return 0 }
func (f *iofsFileInfo) Mode() fs.FileMode  { _ = "STUB: not implemented"; return *new(fs.FileMode) }
func (f *iofsFileInfo) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (f *iofsFileInfo) IsDir() bool        { _ = "STUB: not implemented"; return false }
func (f *iofsFileInfo) Sys() any           { _ = "STUB: not implemented"; return *new(any) }
func (f *iofsFileInfo) Info() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}
func (f *iofsFileInfo) Type() fs.FileMode { _ = "STUB: not implemented"; return *new(fs.FileMode) }

type iofsOpenFile struct {
	*Reader
	name string
}

func (f *iofsOpenFile) Name() string      { _ = "STUB: not implemented"; return "" }
func (f *iofsOpenFile) Mode() fs.FileMode { _ = "STUB: not implemented"; return *new(fs.FileMode) }
func (f *iofsOpenFile) IsDir() bool       { _ = "STUB: not implemented"; return false }
func (f *iofsOpenFile) Sys() any          { _ = "STUB: not implemented"; return *new(any) }
func (f *iofsOpenFile) Stat() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}

type iofsDir struct {
	b    *Bucket
	key  string
	name string

	opened  bool
	entries []fs.DirEntry
	offset  int
}

func newDir(b *Bucket, key, name string) *iofsDir { _ = "STUB: not implemented"; return nil }

func (d *iofsDir) Name() string       { _ = "STUB: not implemented"; return "" }
func (d *iofsDir) Size() int64        { _ = "STUB: not implemented"; return 0 }
func (d *iofsDir) Mode() fs.FileMode  { _ = "STUB: not implemented"; return *new(fs.FileMode) }
func (d *iofsDir) Type() fs.FileMode  { _ = "STUB: not implemented"; return *new(fs.FileMode) }
func (d *iofsDir) ModTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
func (d *iofsDir) IsDir() bool        { _ = "STUB: not implemented"; return false }
func (d *iofsDir) Sys() any           { _ = "STUB: not implemented"; return *new(any) }
func (d *iofsDir) Info() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}
func (d *iofsDir) Stat() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}
func (d *iofsDir) Read([]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (d *iofsDir) Close() error { _ = "STUB: not implemented"; return nil }
func (d *iofsDir) ReadDir(count int) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *iofsDir) openOnce() error { _ = "STUB: not implemented"; return nil }

func (b *Bucket) SetIOFSCallback(fn func() (context.Context, *ReaderOptions)) {
	_ = "STUB: not implemented"
	return
}

func (b *Bucket) Open(path string) (fs.File, error) {
	_ = "STUB: not implemented"
	return *new(fs.File), nil
}

func (b *Bucket) Sub(dir string) (fs.FS, error) { _ = "STUB: not implemented"; return *new(fs.FS), nil }
