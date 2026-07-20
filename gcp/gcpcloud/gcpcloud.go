package gcpcloud

import (
	"github.com/google/wire"
	"gocloud.dev/blob/gcsblob"
	"gocloud.dev/docstore/gcpfirestore"
	"gocloud.dev/gcp"
	"gocloud.dev/gcp/cloudsql"
	"gocloud.dev/pubsub/gcppubsub"
	"gocloud.dev/runtimevar/gcpruntimeconfig"
	"gocloud.dev/secrets/gcpkms"
	"gocloud.dev/server/sdserver"
)

var GCP = wire.NewSet(Services, gcp.DefaultIdentity)

var Services = wire.NewSet(
	gcp.DefaultTransport,
	gcp.NewHTTPClient,

	gcpruntimeconfig.Set,
	gcpkms.Set,
	gcppubsub.Set,
	gcsblob.Set,
	cloudsql.CertSourceSet,
	gcpfirestore.Set,
	sdserver.Set,
)
