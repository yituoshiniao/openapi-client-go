//go:generate msgp -tests=false
package conn

import (
	"context"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/gin-api-http/internal/api/http/dto"
	"github.com/yituoshiniao/kit/xhttp/hclient"
	"github.com/yituoshiniao/kit/xlog"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type CsUserClient struct {
	sling *sling.Sling
}

// NewCsUserClient
func NewCsUserClient(conf config.Config, logger *zap.Logger, tracer opentracing.Tracer) *CsUserClient {
	return &CsUserClient{
		sling: hclient.New(
			hclient.WithLogger(logger),
			hclient.WithTracer(tracer),
			//hclient.WithTarget(conf.CsUserApi.Host),
			hclient.WithTimeout(time.Duration(5*int64(time.Second))), // 5秒超时
			hclient.WithMetrics(true),
			hclient.WithInsecure(), hclient.WithServiceName("CsUserClient"),
		),
	}
}

//type CsUserPurchaseConListRequest struct {
//	ExtendByDays      string `json:"extendByDays"`
//	ExtendReasonCode  string `json:"extendReasonCode"`
//	RequestIdentifier string `json:"requestIdentifier"`
//}

type CsUserPurchaseConListResponse struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Ret  string      `json:"ret"`
	Msg  string      `json:"msg"`
}

// CsUserPurchaseConList 通过订单号延长订阅续订日期
func (a *CsUserClient) CsUserPurchaseConList(ctx context.Context, host string, req []*dto.PurchaseConList, platform string) (CsUserPurchaseConListResponse, error) {
	resp := CsUserPurchaseConListResponse{}
	authorization := "authorization"
	//pathURL := fmt.Sprintf("/inApps/v1/subscriptions/extend?platform=%s", platform)
	pathURL := fmt.Sprintf("/goodsCenterLogic/auth/v1/token/chinaMobileRsaDecode?platform=%s", platform)
	request, err := a.sling.New().
		Base(host).
		Set("Authorization", authorization).
		Post(pathURL).
		BodyJSON(req).
		Request()
	if err != nil {
		xlog.S(ctx).Errorw("CsUserPurchaseConList-Request--错误", "err", err)
		return resp, errors.WithMessage(err, host+pathURL)
	}
	res, err := a.sling.Do(request.WithContext(ctx), &resp, &resp)
	if err != nil {
		xlog.S(ctx).Errorw("CsUserPurchaseConList--错误", "err", err, "pathURL", pathURL)
		//err =
		return resp, errors.WithMessage(err, host+pathURL)
	}
	//xlog.S(ctx).Infow("响应", "resp", resp, "res.Header", res.Header)
	if res.StatusCode != http.StatusOK {
		return resp, errors.WithMessage(errors.New(fmt.Sprintf("响应错误code: %s", res.Status)), host+pathURL)
	}
	return resp, nil
}
