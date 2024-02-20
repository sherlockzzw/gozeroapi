package paperwork

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/client/paperworkservice"
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaperworkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPaperworkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaperworkLogic {
	return &GetPaperworkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPaperworkLogic) GetPaperwork(req *types.GetPaperworkReq) (resp *types.GetPaperworkRes, err error) {
	resp = &types.GetPaperworkRes{}
	paperwork, err := l.svcCtx.RpcClient.PaperworkRpcClient.GetAddPaperworkList(l.ctx, &paperworkservice.GetAddPaperworkListRequest{
		Type:     req.Type,
		Page:     req.Page,
		PageSize: req.PageSize,
		UUID:     l.ctx.Value("uuid").(string),
	})
	if err != nil {
		return resp, err
	}
	for _, info := range paperwork.List {
		resp.List = append(resp.List, types.PaperworkList{
			Id:      info.GetId(),
			Content: info.GetContent(),
		})
	}
	return
}
