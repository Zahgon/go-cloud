//go:build wireinject
// +build wireinject

package main

import (
	"context"

	azcontainer "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"gocloud.dev/blob"
	"gocloud.dev/blob/azureblob"
	"gocloud.dev/runtimevar"
	"gocloud.dev/server"
)

func setupAzure(ctx context.Context, flags *cliFlags) (*server.Server, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func bucketName(flags *cliFlags) azureblob.ContainerName {
	_ = "STUB: not implemented"
	return *new(azureblob.ContainerName)
}

func azureBucket(ctx context.Context, client *azcontainer.Client, flags *cliFlags) (*blob.Bucket, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func azureMOTDVar(ctx context.Context, b *blob.Bucket, flags *cliFlags) (*runtimevar.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
