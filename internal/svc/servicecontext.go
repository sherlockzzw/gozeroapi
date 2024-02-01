package svc

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/config"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/middleware"
	"context"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config    config.Config
	AuthToken rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	config.InItWatchEtcd(c)
	config.RegisterEtcd(context.Background(), &c)
	return &ServiceContext{
		Config:    c,
		AuthToken: middleware.NewAuthTokenMiddleware(c).Handle,
	}
}
