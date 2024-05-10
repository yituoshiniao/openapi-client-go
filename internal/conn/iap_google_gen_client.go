package conn

import (
	"context"
	"fmt"
	"github.com/yituoshiniao/gin-api-http/config"
	"github.com/yituoshiniao/kit/xlog"
	"google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
)

type IapGoogleGenClient struct {
	httpClient *HttpClient
	conf       config.Config
}

func NewIapGoogleGenClient(
	httpClient *HttpClient,
	conf config.Config,
) *IapGoogleGenClient {
	return &IapGoogleGenClient{
		httpClient: httpClient,
		conf:       conf,
	}
}

func (i *IapGoogleGenClient) IapGoogleGenClient(ctx context.Context) (svc *androidpublisher.Service) {
	url := fmt.Sprintf("%s/%s/", i.conf.GpConnectApi.Host, i.conf.GpConnectApi.Prefix)
	options := []option.ClientOption{
		option.WithEndpoint(url),
		option.WithHTTPClient(i.httpClient.Client),
	}
	svc, err := androidpublisher.NewService(ctx, options...)
	if err != nil {
		xlog.S(ctx).Fatalw("androidpublisher-NewService-错误", err, err)
	}
	return
}
