package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/gin-api-http/internal/api/http/servicev1"
	"github.com/yituoshiniao/gin-api-http/internal/conn"
)

type GoodsCenterRouter struct {
	idClient      *conn.IDClient
	generateIDSrv *servicev1.GenerateIDSrv
}

func NewGoodsCenterRouter(
	idClient *conn.IDClient,
	generateIDSrv *servicev1.GenerateIDSrv,
) *GoodsCenterRouter {
	return &GoodsCenterRouter{
		idClient:      idClient,
		generateIDSrv: generateIDSrv,
	}
}

func (h *GoodsCenterRouter) Register(v1 *gin.RouterGroup) {
	health := v1.Group("")
	health.Use()
	{
		health.GET("health", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
	}

	commonGroup := v1.Group("v1/common") //模块
	{
		commonGroup.GET("/generateId", h.generateIDSrv.Handle) //雪花生成id
	}

}
