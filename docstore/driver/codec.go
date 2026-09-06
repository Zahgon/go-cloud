package driver

import (
	"encoding"
	"reflect"

	"gocloud.dev/docstore/internal/fields"
	"gocloud.dev/internal/gcerr"
	"google.golang.org/protobuf/proto"
)

var (
	binaryMarshalerType   = reflect.TypeFor[encoding.BinaryMarshaler]()
	binaryUnmarshalerType = reflect.TypeFor[encoding.BinaryUnmarshaler]()
	textMarshalerType     = reflect.TypeFor[encoding.TextMarshaler]()
	textUnmarshalerType   = reflect.TypeFor[encoding.TextUnmarshaler]()
	protoMessageType      = reflect.TypeFor[proto.Message]()
)

type Encoder interface {
	EncodeNil()
	EncodeBool(bool)
	EncodeString(string)
	EncodeInt(int64)
	EncodeUint(uint64)
	EncodeFloat(float64)
	EncodeBytes([]byte)

	EncodeList(n int) Encoder
	ListIndex(i int)

	EncodeMap(n int) Encoder
	MapKey(string)

	EncodeSpecial(v reflect.Value) (bool, error)
}

func Encode(v reflect.Value, e Encoder) error { _ = "STUB: not implemented"; return nil }

func encode(v reflect.Value, enc Encoder) error { _ = "STUB: not implemented"; return nil }

func encodeList(v reflect.Value, enc Encoder) error { _ = "STUB: not implemented"; return nil }

func encodeMap(v reflect.Value, enc Encoder) error { _ = "STUB: not implemented"; return nil }

func stringifyMapKey(k reflect.Value) (string, error) { _ = "STUB: not implemented"; return "", nil }

func encodeStructWithFields(v reflect.Value, fields fields.List, e Encoder) error {
	_ = "STUB: not implemented"
	return nil
}

func fieldByIndex(v reflect.Value, index []int) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

type Decoder interface {
	AsString() (string, bool)
	AsInt() (int64, bool)
	AsUint() (uint64, bool)
	AsFloat() (float64, bool)
	AsBytes() ([]byte, bool)
	AsBool() (bool, bool)
	AsNull() bool

	ListLen() (int, bool)

	DecodeList(func(int, Decoder) bool)

	MapLen() (int, bool)

	DecodeMap(func(string, Decoder, bool) bool)

	AsInterface() (any, error)

	AsSpecial(reflect.Value) (bool, any, error)

	String() string
}

func Decode(v reflect.Value, d Decoder) error { _ = "STUB: not implemented"; return nil }

func decode(v reflect.Value, d Decoder) error { _ = "STUB: not implemented"; return nil }

func decodeList(v reflect.Value, d Decoder) error { _ = "STUB: not implemented"; return nil }

func prepareLength(v reflect.Value, wantLen int) error { _ = "STUB: not implemented"; return nil }

func decodeMap(v reflect.Value, d Decoder) error { _ = "STUB: not implemented"; return nil }

func unstringifyMapKey(key string, keyType reflect.Type) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func decodeStruct(v reflect.Value, d Decoder) error { _ = "STUB: not implemented"; return nil }

func fieldByIndexCreate(v reflect.Value, index []int) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func decodingError(v reflect.Value, d Decoder) error { _ = "STUB: not implemented"; return nil }

func overflowError(x any, t reflect.Type) error { _ = "STUB: not implemented"; return nil }

func wrap(err error, code gcerr.ErrorCode) error { _ = "STUB: not implemented"; return nil }

var fieldCache = fields.NewCache(parseTag, nil, nil)

func IsEmptyValue(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

type tagOptions struct {
	omitEmpty bool
}

func parseTag(t reflect.StructTag) (name string, keep bool, other any, err error) {
	_ = "STUB: not implemented"
	return "", false, *new(any), nil
}
