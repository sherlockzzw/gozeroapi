package bgm

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"net/http"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/logic/bgm"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
)

func GetBgmClassListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := bgm.NewGetBgmClassListLogic(r.Context(), svcCtx)
		resp, err := l.GetBgmClassList()
		if err != nil {
			result.ResultFail(r.Context(), w, result.ServerError, err)
		} else {
			result.ResultSuccess(r.Context(), w, resp)
		}
	}
}
