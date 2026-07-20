package drivertest

import (
	"math/rand"
	"sync"

	"gocloud.dev/docstore/driver"
)

func MakeUniqueStringDeterministicForTesting(seed int64) { _ = "STUB: not implemented"; return }

type randReader struct {
	mu sync.Mutex
	r  *rand.Rand
}

func (r *randReader) Read(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func MustDocument(doc any) driver.Document { _ = "STUB: not implemented"; return *new(driver.Document) }
