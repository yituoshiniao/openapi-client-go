package conn

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewIapGoogleGenClient,
	NewIDClient,
	NewRedis,
	NewGoodsCenterDB,
	NewCamexamDB,
	NewTokenDB,
	NewAppStoreServerClient,
	NewWxClient,
	NewHttpClient,
	NewRetry,
	NewIapStoreClient,
	NewDbProxyClient,
	NewCsUserClient,
)
