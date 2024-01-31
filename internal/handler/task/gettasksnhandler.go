package task

import (
	"net/http"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/logic/task"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTaskSnHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := task.NewGetTaskSnLogic(r.Context(), svcCtx)
		resp, err := l.GetTaskSn()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
