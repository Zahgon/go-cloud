package driver

import (
	"reflect"

	"gocloud.dev/docstore/internal/fields"
)

type Document struct {
	Origin any
	m      map[string]any
	s      reflect.Value
	fields fields.List
}

func NewDocument(doc any) (Document, error) { _ = "STUB: not implemented"; return *new(Document), nil }

func (d Document) GetField(field string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (d Document) getDocument(fp []string, create bool) (Document, error) {
	_ = "STUB: not implemented"
	return *new(Document), nil
}

func (d Document) Get(fp []string) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (d Document) structField(name string) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (d Document) Set(fp []string, val any) error { _ = "STUB: not implemented"; return nil }

func (d Document) SetField(field string, value any) error { _ = "STUB: not implemented"; return nil }

func (d Document) FieldNames() []string { _ = "STUB: not implemented"; return nil }

func (d Document) Encode(e Encoder) error { _ = "STUB: not implemented"; return nil }

func (d Document) Decode(dec Decoder) error { _ = "STUB: not implemented"; return nil }

func (d Document) HasField(field string) bool { _ = "STUB: not implemented"; return false }

func (d Document) HasFieldFold(field string) bool { _ = "STUB: not implemented"; return false }

func (d Document) hasField(field string, exactMatch bool) bool {
	_ = "STUB: not implemented"
	return false
}
