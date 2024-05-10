module github.com/yituoshiniao/gin-api-http

go 1.13

//replace github.com/yituoshiniao/kit => /Users/gq/work/project/hh_go/kit-hh
//replace github.com/yituoshiniao/openapi-client-go => ./gen/go

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/GUAIK-ORG/go-snowflake v0.0.0-20200116064823-220c4260e85f
	github.com/alibaba/sentinel-golang v1.0.4
	github.com/allegro/bigcache v1.2.1
	github.com/avast/retry-go v2.7.0+incompatible
	github.com/awa/go-iap v1.21.1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dghubble/sling v1.4.0
	github.com/felixge/fgprof v0.9.3
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-acme/lego/v3 v3.4.0
	github.com/go-kit/kit v0.10.0
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.14.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/glog v1.1.2 // indirect
	github.com/google/uuid v1.3.1
	github.com/google/wire v0.5.0
	github.com/onsi/gomega v1.18.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7
	github.com/prometheus/client_golang v1.12.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/shirou/gopsutil/v3 v3.21.6
	github.com/spf13/viper v1.7.1
	github.com/swaggo/files v1.0.1
	github.com/swaggo/gin-swagger v1.6.0
	github.com/swaggo/swag v1.16.3
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	github.com/yituoshiniao/kit v0.3.6
	go.uber.org/zap v1.25.0
	golang.org/x/tools v0.21.0 // indirect
	google.golang.org/api v0.146.0
	google.golang.org/grpc v1.58.2
	gorm.io/datatypes v1.2.0 // indirect
	gorm.io/driver/mysql v1.5.1
	gorm.io/gen v0.3.23
	gorm.io/gorm v1.25.4
	gorm.io/hints v1.1.2 // indirect
	gorm.io/plugin/dbresolver v1.4.7
)
