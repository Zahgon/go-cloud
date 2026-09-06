//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"database/sql"

	"gocloud.dev/blob"
	"gocloud.dev/runtimevar"
	"gocloud.dev/server"
)

func setupLocal(ctx context.Context, flags *cliFlags) (*server.Server, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func localBucket(flags *cliFlags) (*blob.Bucket, error) { _ = "STUB: not implemented"; return nil, nil }

func dialLocalSQL(flags *cliFlags) (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func localRuntimeVar(flags *cliFlags) (*runtimevar.Variable, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
