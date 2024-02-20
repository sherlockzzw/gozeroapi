package bgm

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"net/http"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/logic/bgm"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBgmListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBgmListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := bgm.NewGetBgmListLogic(r.Context(), svcCtx)
		resp, err := l.GetBgmList(&req)
		if err != nil {
			result.ResultFail(r.Context(), w, result.ServerError, err)
		} else {
			result.ResultSuccess(r.Context(), w, resp)
		}
	}
}
