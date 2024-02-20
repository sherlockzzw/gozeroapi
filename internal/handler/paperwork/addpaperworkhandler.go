package paperwork

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/global/errorMsg"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/result"
	"net/http"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/logic/paperwork"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddPaperworkHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddPaperworkReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := paperwork.NewAddPaperworkLogic(r.Context(), svcCtx)
		resp, errMsg, err := l.AddPaperwork(&req)
		if err != nil {
			result.ResultFail(r.Context(), w, result.ServerError, err)
		} else if errMsg != "" {
			result.ResultFailCustom(r.Context(), w, errorMsg.ForbiddenWordError, errMsg)
		} else {
			result.ResultSuccess(r.Context(), w, resp)
		}
	}
}
