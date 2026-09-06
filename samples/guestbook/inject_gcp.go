//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"database/sql"

	"gocloud.dev/blob"
	"gocloud.dev/gcp"
	"gocloud.dev/mysql/gcpmysql"
	"gocloud.dev/runtimevar"
	"gocloud.dev/server"
	pb "google.golang.org/genproto/googleapis/cloud/runtimeconfig/v1beta1"
)

func setupGCP(ctx context.Context, flags *cliFlags) (*server.Server, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func gcpBucket(ctx context.Context, flags *cliFlags, client *gcp.HTTPClient) (*blob.Bucket, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func openGCPDatabase(ctx context.Context, opener *gcpmysql.URLOpener, id gcp.ProjectID, flags *cliFlags) (*sql.DB, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func gcpMOTDVar(ctx context.Context, client pb.RuntimeConfigManagerClient, project gcp.ProjectID, flags *cliFlags) (*runtimevar.Variable, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
