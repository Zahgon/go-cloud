package azurecloud

import (
	"github.com/google/wire"
	"gocloud.dev/blob/azureblob"
	"gocloud.dev/secrets/azurekeyvault"
)

var Azure = wire.NewSet(
	azurekeyvault.Set,
	azureblob.Set,
)
