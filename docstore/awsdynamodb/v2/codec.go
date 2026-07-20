package awsdynamodb

import (
	"reflect"
	"time"

	dyn2Types "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"gocloud.dev/docstore/driver"
)

var nullValue = &dyn2Types.AttributeValueMemberNULL{Value: true}

type encoder struct {
	av dyn2Types.AttributeValue
}

func (e *encoder) EncodeNil()        { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeBool(x bool) { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeInt(x int64) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeUint(x uint64) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeBytes(x []byte)  { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeFloat(x float64) { _ = "STUB: not implemented"; return }

func (e *encoder) ListIndex(int) { _ = "STUB: not implemented"; return }
func (e *encoder) MapKey(string) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeString(x string) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeComplex(x complex128) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeList(n int) driver.Encoder {
	_ = "STUB: not implemented"
	return *new(driver.Encoder)
}

func (e *encoder) EncodeMap(n int) driver.Encoder {
	_ = "STUB: not implemented"
	return *new(driver.Encoder)
}

var typeOfGoTime = reflect.TypeFor[time.Time]()

func (e *encoder) EncodeSpecial(v reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type listEncoder struct {
	s []dyn2Types.AttributeValue
	encoder
}

func (e *listEncoder) ListIndex(i int) { _ = "STUB: not implemented"; return }

type mapEncoder struct {
	m map[string]dyn2Types.AttributeValue
	encoder
}

func (e *mapEncoder) MapKey(k string) { _ = "STUB: not implemented"; return }

func encodeDoc(doc driver.Document) (dyn2Types.AttributeValue, error) {
	_ = "STUB: not implemented"
	return *new(dyn2Types.AttributeValue), nil
}

func encodeDocKeyFields(doc driver.Document, pkey, skey string) (*dyn2Types.AttributeValueMemberM, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeValue(v any) (dyn2Types.AttributeValue, error) {
	_ = "STUB: not implemented"
	return *new(dyn2Types.AttributeValue), nil
}

func encodeFloat(f float64) dyn2Types.AttributeValue {
	_ = "STUB: not implemented"
	return *new(dyn2Types.AttributeValue)
}

func decodeDoc(item dyn2Types.AttributeValue, doc driver.Document) error {
	_ = "STUB: not implemented"
	return nil
}

type decoder struct {
	av dyn2Types.AttributeValue
}

func (d decoder) String() string { _ = "STUB: not implemented"; return "" }

func (d decoder) AsBool() (bool, bool) { _ = "STUB: not implemented"; return false, false }

func (d decoder) AsNull() bool { _ = "STUB: not implemented"; return false }

func (d decoder) AsString() (string, bool) { _ = "STUB: not implemented"; return "", false }

func (d decoder) AsInt() (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsUint() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsFloat() (float64, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsComplex() (complex128, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsBytes() ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

func (d decoder) ListLen() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) DecodeList(f func(i int, vd driver.Decoder) bool) {
	_ = "STUB: not implemented"
	return
}

func (d decoder) MapLen() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) DecodeMap(f func(key string, vd driver.Decoder, exactMatch bool) bool) {
	_ = "STUB: not implemented"
	return
}

func (d decoder) AsInterface() (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func toGoValue(av dyn2Types.AttributeValue) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (d decoder) AsSpecial(v reflect.Value) (bool, any, error) {
	_ = "STUB: not implemented"
	return false, *new(any), nil
}
