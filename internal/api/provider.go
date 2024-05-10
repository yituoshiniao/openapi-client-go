package api

import (
	"github.com/google/wire"
	"github.com/yituoshiniao/gin-api-http/internal/api/cron"
	"github.com/yituoshiniao/gin-api-http/internal/api/http"
	"github.com/yituoshiniao/gin-api-http/internal/api/http/servicev1"
	"github.com/yituoshiniao/gin-api-http/internal/task"
)

var WireSet = wire.NewSet(
	cron.NewEnter,
	http.NewResponse,
	task.NewTaskCron,
	servicev1.NewGenerateIDSrv,
	servicev1.NewTestSrv,
	servicev1.NewTestV2Srv,
)
