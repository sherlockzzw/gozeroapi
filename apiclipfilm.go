package main

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/config"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/global/errorMsg"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/handler"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func main() {
	flag.Parse()
	c := config.Init()
	var lc logc.LogConf
	logc.MustSetup(lc)
	errorMsg.New()
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		result.SystemError(r.Context(), w, err)
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(*c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
