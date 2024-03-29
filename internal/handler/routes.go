// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/handler/user"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/list",
				Handler: user.GetUserListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)
}
