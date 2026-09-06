package gcpfirestore

import (
	"reflect"
	"time"

	pb "cloud.google.com/go/firestore/apiv1/firestorepb"
	"gocloud.dev/docstore/driver"
	"google.golang.org/genproto/googleapis/type/latlng"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

func encodeDoc(doc driver.Document, nameField string) (*pb.Document, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func encodeValue(x any) (*pb.Value, error) { _ = "STUB: not implemented"; return nil, nil }

type encoder struct {
	pv *pb.Value
}

var nullValue = &pb.Value{ValueType: &pb.Value_NullValue{}}

func (e *encoder) EncodeNil()        { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeBool(x bool) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeInt(x int64) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeUint(x uint64) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeBytes(x []byte) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeFloat(x float64) { _ = "STUB: not implemented"; return }
func (e *encoder) EncodeString(x string) { _ = "STUB: not implemented"; return }

func (e *encoder) ListIndex(int) { _ = "STUB: not implemented"; return }
func (e *encoder) MapKey(string) { _ = "STUB: not implemented"; return }

func (e *encoder) EncodeList(n int) driver.Encoder {
	_ = "STUB: not implemented"
	return *new(driver.Encoder)
}

func (e *encoder) EncodeMap(n int) driver.Encoder {
	_ = "STUB: not implemented"
	return *new(driver.Encoder)
}

var (
	typeOfGoTime         = reflect.TypeFor[time.Time]()
	typeOfProtoTimestamp = reflect.TypeFor[*tspb.Timestamp]()
	typeOfLatLng         = reflect.TypeFor[*latlng.LatLng]()
)

func (e *encoder) EncodeSpecial(v reflect.Value) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type listEncoder struct {
	s []*pb.Value
	encoder
}

func (e *listEncoder) ListIndex(i int) { _ = "STUB: not implemented"; return }

type mapEncoder struct {
	m map[string]*pb.Value
	encoder
}

func (e *mapEncoder) MapKey(k string) { _ = "STUB: not implemented"; return }

func floatval(x float64) *pb.Value { _ = "STUB: not implemented"; return nil }

func decodeDoc(pdoc *pb.Document, ddoc driver.Document, nameField, revField string) error {
	_ = "STUB: not implemented"
	return nil
}

type decoder struct {
	pv *pb.Value
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

func decodeValue(v *pb.Value) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (d decoder) ListLen() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) DecodeList(f func(int, driver.Decoder) bool) { _ = "STUB: not implemented"; return }

func (d decoder) MapLen() (int, bool) { _ = "STUB: not implemented"; return 0, false }

func (d decoder) DecodeMap(f func(string, driver.Decoder, bool) bool) {
	_ = "STUB: not implemented"
	return
}

func (d decoder) AsSpecial(v reflect.Value) (bool, any, error) {
	_ = "STUB: not implemented"
	return false, *new(any), nil
}
