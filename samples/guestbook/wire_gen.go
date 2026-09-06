//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"gocloud.dev/blob"
	"gocloud.dev/blob/azureblob"
	"gocloud.dev/gcp"
	"gocloud.dev/mysql/awsmysql"
	"gocloud.dev/mysql/gcpmysql"
	"gocloud.dev/runtimevar"
	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
	runtimeconfig "google.golang.org/genproto/googleapis/cloud/runtimeconfig/v1beta1"
)

func setupAWS(ctx context.Context, flags *cliFlags) (*server.Server, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var (
	_wireClientValue        = http.DefaultClient
	_wireDefaultDriverValue = &server.DefaultDriver{}
)

func setupAzure(ctx context.Context, flags *cliFlags) (*server.Server, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var (
	_wireLoggerValue            = requestlog.Logger(nil)
	_wireTextMapPropagatorValue = propagation.TextMapPropagator(nil)
	_wireTracerProviderValue    = trace.TracerProvider(nil)
	_wireMeterProviderValue     = metric.MeterProvider(nil)
)

func setupGCP(ctx context.Context, flags *cliFlags) (*server.Server, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func setupLocal(ctx context.Context, flags *cliFlags) (*server.Server, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

var (
	_wireRequestlogLoggerValue = requestlog.Logger(nil)
)

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

func bucketName(flags *cliFlags) azureblob.ContainerName {
	_ = "STUB: not implemented"
	return *new(azureblob.ContainerName)
}

func azureBucket(ctx context.Context, client *container.Client, flags *cliFlags) (*blob.Bucket, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func azureMOTDVar(ctx context.Context, b *blob.Bucket, flags *cliFlags) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gcpBucket(ctx context.Context, flags *cliFlags, client *gcp.HTTPClient) (*blob.Bucket, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func openGCPDatabase(ctx context.Context, opener *gcpmysql.URLOpener, id gcp.ProjectID, flags *cliFlags) (*sql.DB, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func gcpMOTDVar(ctx context.Context, client runtimeconfig.RuntimeConfigManagerClient, project gcp.ProjectID, flags *cliFlags) (*runtimevar.Variable, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func localBucket(flags *cliFlags) (*blob.Bucket, error) { _ = "STUB: not implemented"; return nil, nil }

func dialLocalSQL(flags *cliFlags) (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func localRuntimeVar(flags *cliFlags) (*runtimevar.Variable, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
