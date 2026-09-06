package batcher

import (
	"context"
	"errors"
	"reflect"
	"sync"
)

func Split(n int, opts *Options) []int { _ = "STUB: not implemented"; return nil }

type Batcher struct {
	opts          Options
	handler       func(any) error
	itemSliceZero reflect.Value
	wg            sync.WaitGroup

	mu        sync.Mutex
	pending   []waiter
	nHandlers int
	shutdown  bool
}

var ErrMessageTooLarge = errors.New("batcher: message too large")

type sizableItem interface {
	ByteSize() int
}

type waiter struct {
	item any
	errc chan error
}

type Options struct {
	MaxHandlers int

	MinBatchSize int

	MaxBatchSize int

	MaxBatchByteSize int
}

func newOptionsWithDefaults(opts *Options) Options { _ = "STUB: not implemented"; return *new(Options) }

func (o *Options) NewMergedOptions(opts *Options) *Options { _ = "STUB: not implemented"; return nil }

func New(itemType reflect.Type, opts *Options, handler func(any) error) *Batcher {
	_ = "STUB: not implemented"
	return nil
}

func (b *Batcher) Add(ctx context.Context, item any) error { _ = "STUB: not implemented"; return nil }

func (b *Batcher) AddNoWait(item any) <-chan error { _ = "STUB: not implemented"; return nil }

func (b *Batcher) handleBatch(batch []waiter) { _ = "STUB: not implemented"; return }

func (b *Batcher) nextBatch() []waiter { _ = "STUB: not implemented"; return nil }

func (b *Batcher) callHandler(batch []waiter) { _ = "STUB: not implemented"; return }

func (b *Batcher) Shutdown() { _ = "STUB: not implemented"; return }
