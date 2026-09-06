package awscloud

import (
	"net/http"

	"github.com/google/wire"
	"gocloud.dev/aws"
	"gocloud.dev/aws/rds"
	"gocloud.dev/blob/s3blob"
	"gocloud.dev/docstore/awsdynamodb/v2"
	"gocloud.dev/pubsub/awssnssqs"
	"gocloud.dev/runtimevar/awsparamstore"
	"gocloud.dev/secrets/awskms"
	"gocloud.dev/server/xrayserver"
)

var AWS = wire.NewSet(
	Services,
	aws.NewDefaultV2Config,
	wire.Value(http.DefaultClient),
)

var Services = wire.NewSet(
	s3blob.Set,
	awssnssqs.Set,
	awsparamstore.Set,
	awskms.Set,
	rds.CertFetcherSet,
	awsdynamodb.Set,
	xrayserver.Set,
)
