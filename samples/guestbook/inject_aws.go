//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"database/sql"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"gocloud.dev/blob"
	"gocloud.dev/mysql/awsmysql"
	"gocloud.dev/runtimevar"
	"gocloud.dev/server"
)

func setupAWS(ctx context.Context, flags *cliFlags) (*server.Server, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func awsBucket(ctx context.Context, client *s3.Client, flags *cliFlags) (*blob.Bucket, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func openAWSDatabase(ctx context.Context, opener *awsmysql.URLOpener, flags *cliFlags) (*sql.DB, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func awsMOTDVar(ctx context.Context, client *ssm.Client, flags *cliFlags) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
