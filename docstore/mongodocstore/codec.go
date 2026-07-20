package mongodocstore

import (
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gocloud.dev/docstore/driver"
)

func encodeDoc(doc driver.Document, lowercaseFields bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeValue(x interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

type encoder struct {
	val             interface{}
	lowercaseFields bool
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

var (
	typeOfGoTime   = reflect.TypeOf(time.Time{})
	typeOfObjectID = reflect.TypeOf(primitive.ObjectID{})
)

func (e *encoder) EncodeSpecial(v reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *encoder) EncodeList(n int) driver.Encoder {
	_ = "STUB: not implemented"
	return *new(driver.Encoder)
}

type listEncoder struct {
	s []interface{}
	encoder
}

func (e *listEncoder) ListIndex(i int) { _ = "STUB: not implemented"; return }

type mapEncoder struct {
	m        map[string]interface{}
	isStruct bool
	encoder
}

func (e *encoder) EncodeMap(n int) driver.Encoder {
	_ = "STUB: not implemented"
	return *new(driver.Encoder)
}

func (e *mapEncoder) MapKey(k string) { _ = "STUB: not implemented"; return }

func decodeDoc(m map[string]interface{}, ddoc driver.Document, idField string, lowercaseFields bool) error {
	_ = "STUB: not implemented"
	return nil
}

type decoder struct {
	val             interface{}
	lowercaseFields bool
}

func (d decoder) String() string { _ = "STUB: not implemented"; return "" }

func (d decoder) AsNull() bool { _ = "STUB: not implemented"; return false }

func (d decoder) AsBool() (bool, bool) { _ = "STUB: not implemented"; return false, false }

func (d decoder) AsString() (string, bool) { _ = "STUB: not implemented"; return "", false }

func (d decoder) AsInt() (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsUint() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsFloat() (float64, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) AsBytes() ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

func (d decoder) AsInterface() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func toGoValue(v interface{}) (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (d decoder) ListLen() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) DecodeList(f func(i int, d2 driver.Decoder) bool) {
	_ = "STUB: not implemented"
	return
}

func (d decoder) MapLen() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) DecodeMap(f func(key string, d2 driver.Decoder, exactMatch bool) bool) {
	_ = "STUB: not implemented"
	return
}

func (d decoder) AsSpecial(v reflect.Value) (bool, interface{}, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func bsonDateTimeToTime(dt primitive.DateTime) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
