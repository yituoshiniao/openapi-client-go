package conn

import (
	"github.com/go-redis/redis"
	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/kit/xrds"
)

func NewRedis(conf config.Config) *redis.Client {
	return xrds.Open(conf.XRedis)
}
