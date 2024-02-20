package svc

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/config"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/middleware"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/client/bgmservice"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/client/paperworkservice"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/client/templateservice"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	RedisCli  *redis.Client
	AuthToken rest.Middleware
	RpcClient RpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	config.InItWatchEtcd(c)
	config.RegisterEtcd(context.Background(), &c)
	return &ServiceContext{
		Config:    c,
		RedisCli:  config.RegisterRedisCli(c),
		AuthToken: middleware.NewAuthTokenMiddleware(c).Handle,
		RpcClient: NewRpcClient(c),
	}
}

type RpcClient struct {
	BgmRpcClient       bgmservice.BgmService
	TemplateRpcClient  templateservice.TemplateService
	PaperworkRpcClient paperworkservice.PaperworkService
}

func NewRpcClient(c config.Config) RpcClient {
	return RpcClient{
		BgmRpcClient:       bgmservice.NewBgmService(zrpc.MustNewClient(c.ClipFilmRpcConf)),
		TemplateRpcClient:  templateservice.NewTemplateService(zrpc.MustNewClient(c.ClipFilmRpcConf)),
		PaperworkRpcClient: paperworkservice.NewPaperworkService(zrpc.MustNewClient(c.ClipFilmRpcConf)),
	}
}
