// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package inject

import (
	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/gin-api-http/internal/api/cron"
	"github.com/yituoshiniao/gin-api-http/internal/api/http"
	"github.com/yituoshiniao/gin-api-http/internal/api/http/servicev1"
	"github.com/yituoshiniao/gin-api-http/internal/app"
	"github.com/yituoshiniao/gin-api-http/internal/conn"
	"github.com/yituoshiniao/gin-api-http/internal/metrics"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/repository"
	"github.com/yituoshiniao/gin-api-http/internal/module/mockv2/application/service"
	"github.com/yituoshiniao/gin-api-http/internal/router"
	"github.com/yituoshiniao/gin-api-http/internal/router/v1"
	"github.com/yituoshiniao/gin-api-http/internal/task"
	"github.com/yituoshiniao/gin-api-http/internal/util"
)

// Injectors from wire.go:

func InitApp() (*app.App, func(), error) {
	string2 := app.InitEnv()
	configConfig := config.ParseConfig(string2)
	engine := app.InitGin(configConfig)
	logger, cleanup, err := ProvideLogger(configConfig)
	if err != nil {
		return nil, nil, err
	}
	tracer, cleanup2 := ProvideTracer(configConfig)
	idClient := conn.NewIDClient(configConfig, logger, tracer)
	counterMetrics := metrics.NewCounterMetrics()
	response := http.NewResponse(counterMetrics, tracer, logger)
	context := app.InitCtx()
	snowflakeID, err := util.NewSnowflakeIDClient(configConfig, context)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	generateIDSrv := servicev1.NewGenerateIDSrv(response, snowflakeID)
	goodsCenterRouter := v1.NewGoodsCenterRouter(idClient, generateIDSrv)
	goodsCenterDB, cleanup3 := conn.NewGoodsCenterDB(configConfig, logger)
	userScoreRepo := repository.NewUserScoreRepo(goodsCenterDB)
	client := conn.NewRedis(configConfig)
	userScoreSrv := service.NewUserScoreSrv(userScoreRepo, client)
	testSrv := servicev1.NewTestSrv(response, userScoreSrv)
	testV2Srv := servicev1.NewTestV2Srv(response)
	goodsCenterTestV2 := v1.NewGoodsCenterTestV2(testSrv, testV2Srv)
	routerRouter := router.NewRouter(engine, goodsCenterRouter, goodsCenterTestV2, configConfig)
	enter := app.NewEnter(configConfig)
	cronTestSrv := cron.NewTestSrv(context)
	factoryTemplateSrv := cron.NewFactoryTemplateSrv(cronTestSrv)
	cronEnter := cron.NewEnter(counterMetrics, factoryTemplateSrv)
	cronSingleAppleProductPriceTask := task.NewCronSingleAppleProductPriceTask(context)
	taskEnter, cleanup4 := task.NewEnter(context, counterMetrics, logger, cronSingleAppleProductPriceTask)
	summaryMetrics := metrics.NewSummaryMetrics()
	appApp := app.NewApp(configConfig, routerRouter, context, engine, enter, cronEnter, taskEnter, counterMetrics, summaryMetrics)
	return appApp, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
