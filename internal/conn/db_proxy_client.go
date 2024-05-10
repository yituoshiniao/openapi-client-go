//go:generate msgp -tests=false
package conn

import (
	"context"
	"fmt"
	"github.com/dghubble/sling"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/kit/xhttp/hclient"
	"github.com/yituoshiniao/kit/xlog"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"time"
)

type DbProxyClient struct {
	sling *sling.Sling
}

// EmptyResp 空数据,查询条件结果为空时的返回
const EmptyResp string = "{}"

// NewDbProxyClient // db proxy
func NewDbProxyClient(conf config.Config, logger *zap.Logger, tracer opentracing.Tracer) *DbProxyClient {
	return &DbProxyClient{
		sling: hclient.New(
			hclient.WithLogger(logger),
			hclient.WithTracer(tracer),
			hclient.WithTarget(conf.DbProxyApi.Host),
			hclient.WithTimeout(time.Duration(conf.DbProxyApi.Timeout*int64(time.Second))),
			hclient.WithMetrics(true),
			hclient.WithInsecure(), hclient.WithServiceName("DbProxyClient"),
		),
	}
}

// HTMLEscapeString 转义sql 防注入
func (d *DbProxyClient) HTMLEscapeString(ctx context.Context, sql string) (safeSql string) {
	return template.HTMLEscapeString(sql)
}

type DbProxyRequest struct {
	Sql     string `json:"sql"`
	DbName  string `json:"db_name"`
	LockKey string `json:"lock_key"`
}

// Do 执行sql到dbproxy服务
func (d *DbProxyClient) Do(ctx context.Context, dbName, sql string) (resp interface{}, err error) {
	req := DbProxyRequest{
		//Sql:    d.HTMLEscapeString(ctx, sql),
		Sql:    sql,
		DbName: dbName,
	}
	request, err := d.sling.New().
		Post("/dbproxy/user_query_common").
		BodyJSON(req).
		Request()

	if err != nil {
		xlog.S(ctx).Errorw("DbProxyClient Do Request--错误", "err", err)
		return resp, err
	}

	res, err := d.sling.Do(request.WithContext(ctx), &resp, &resp)
	if err != nil {
		xlog.S(ctx).Errorw("DbProxyClient 执行错误--错误", "err", err)
		return resp, err
	}

	xlog.S(ctx).Infow("DbProxyClient响应", "resp", resp)
	if res.StatusCode != http.StatusOK {
		xlog.S(ctx).Errorw("DbProxyClient StatusCode错误", "err", err)
		return resp, errors.New(fmt.Sprintf("响应错误code: %s", res.Status))
	}
	return resp, nil
}

// DoRespToStruct 数据转换 dbproxy 结果转 对应的struct
func (l *DbProxyClient) DoRespToStruct(ctx context.Context, dbName, sql string, model interface{}) (result interface{}, err error) {
	resp, err := l.Do(ctx, dbName, sql)
	if err != nil {
		xlog.S(ctx).Errorw("Do-错误", "err", err)
		return result, err
	}

	byteMarshal, err := ffjson.Marshal(resp)
	if err != nil {
		xlog.S(ctx).Errorw("Marshal-错误", "err", err)
		return result, err
	}
	//items := make([]model.TsLoginToken, 0)
	err = ffjson.Unmarshal(byteMarshal, &model)
	if err != nil {
		xlog.S(ctx).Errorw("Unmarshal-错误", "err", err)
		return result, err
	}
	return model, nil
}
