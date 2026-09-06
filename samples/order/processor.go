package main

import (
	"context"
	_ "image/jpeg"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/memdocstore"
	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/mempubsub"
)

type processor struct {
	requestSub *pubsub.Subscription
	bucket     *blob.Bucket
	coll       *docstore.Collection
}

func (p *processor) run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *processor) handleRequest(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func createOrFindOrder(ctx context.Context, coll *docstore.Collection, req *OrderRequest) (*Order, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *processor) processOrder(ctx context.Context, order *Order) error {
	_ = "STUB: not implemented"
	return nil
}
