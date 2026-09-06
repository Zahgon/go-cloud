package driver

import (
	"sync"
)

func UniqueString() string { _ = "STUB: not implemented"; return "" }

func SplitActions(actions []*Action, split func(a, b *Action) bool) [][]*Action {
	_ = "STUB: not implemented"
	return nil
}

func GroupActions(actions []*Action) (beforeGets, getList, writeList, writesTxList, afterGets []*Action) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil, nil
}

func AsFunc(val any) func(any) bool { _ = "STUB: not implemented"; return nil }

func GroupByFieldPath(gets []*Action) [][]*Action { _ = "STUB: not implemented"; return nil }

func fpsEqual(fps1, fps2 [][]string) bool { _ = "STUB: not implemented"; return false }

func FieldPathsEqual(fp1, fp2 []string) bool { _ = "STUB: not implemented"; return false }

func FieldPathEqualsField(fp []string, s string) bool { _ = "STUB: not implemented"; return false }

type Throttle struct {
	c  chan struct{}
	wg sync.WaitGroup
}

func NewThrottle(max int) *Throttle { _ = "STUB: not implemented"; return nil }

func (t *Throttle) Acquire() { _ = "STUB: not implemented"; return }

func (t *Throttle) Release() { _ = "STUB: not implemented"; return }

func (t *Throttle) Wait() { _ = "STUB: not implemented"; return }
