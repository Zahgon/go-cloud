package driver

import (
	"math/big"
	"reflect"
	"time"
)

func CompareTimes(t1, t2 time.Time) int { _ = "STUB: not implemented"; return 0 }

func CompareNumbers(n1, n2 any) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func toBigFloat(x reflect.Value) (*big.Float, error) { _ = "STUB: not implemented"; return nil, nil }
