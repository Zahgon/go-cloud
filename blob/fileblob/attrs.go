package fileblob

import (
	"fmt"
)

const attrsExt = ".attrs"

var errAttrsExt = fmt.Errorf("file extension %q is reserved", attrsExt)

type xattrs struct {
	CacheControl       string            `json:"user.cache_control"`
	ContentDisposition string            `json:"user.content_disposition"`
	ContentEncoding    string            `json:"user.content_encoding"`
	ContentLanguage    string            `json:"user.content_language"`
	ContentType        string            `json:"user.content_type"`
	Metadata           map[string]string `json:"user.metadata"`
	MD5                []byte            `json:"md5"`
}

func setAttrs(path string, xa xattrs) error { _ = "STUB: not implemented"; return nil }

func getAttrs(path string) (xattrs, error) { _ = "STUB: not implemented"; return *new(xattrs), nil }
