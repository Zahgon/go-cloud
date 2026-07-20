package memdocstore

import (
	"reflect"
	"time"

	"gocloud.dev/docstore/driver"
)

func encodeDoc(doc driver.Document) (storedDoc, error) {
	_ = "STUB: not implemented"
	return *new(storedDoc), nil
}

func encodeValue(v any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

type encoder struct {
	val any
}

func (e *encoder) EncodeNil()            { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeBool(x bool)     { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeInt(x int64)     { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeUint(x uint64)   { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeBytes(x []byte)  { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeFloat(x float64) { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeString(x string) { _ = "STUB: not implemented"; return }
func (e *encoder) ListIndex(int)         { _ = "STUB: not implemented"; return }
func (e *encoder) MapKey(string)         { _ = "STUB: not implemented"; return }

var typeOfGoTime = reflect.TypeFor[time.Time]()

func (e *encoder) EncodeSpecial(v reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *encoder) EncodeList(n int) driver.Encoder {
	_ = "STUB: not implemented"
	return *new(driver.Encoder)
}

type listEncoder struct {
	s []any
	encoder
}

func (e *listEncoder) ListIndex(i int) { _ = "STUB: not implemented"; return }

type mapEncoder struct {
	m map[string]any
	encoder
}

func (e *encoder) EncodeMap(n int) driver.Encoder {
	_ = "STUB: not implemented"
	return *new(driver.Encoder)
}

func (e *mapEncoder) MapKey(k string) { _ = "STUB: not implemented"; return }

func decodeDoc(m storedDoc, ddoc driver.Document, fps [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

type decoder struct {
	val any
}

func (d decoder) String() string { _ = "STUB: not implemented"; return "" }

func (d decoder) AsNull() bool { _ = "STUB: not implemented"; return false }

func (d decoder) AsBool() (bool, bool) { _ = "STUB: not implemented"; return false, false }

func (d decoder) AsString() (string, bool) { _ = "STUB: not implemented"; return "", false }

func (d decoder) AsInt() (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsUint() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsFloat() (float64, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsBytes() ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

func (d decoder) AsInterface() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (d decoder) ListLen() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) DecodeList(f func(i int, d2 driver.Decoder) bool) {
	_ = "STUB: not implemented"
	return
}

func (d decoder) MapLen() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) DecodeMap(f func(key string, d2 driver.Decoder, _ bool) bool) {
	_ = "STUB: not implemented"
	return
}

func (d decoder) AsSpecial(v reflect.Value) (bool, any, error) {
	_ = "STUB: not implemented"
	return false, *new(any), nil
}
