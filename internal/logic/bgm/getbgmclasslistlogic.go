package bgm

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/tool"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/client/bgmservice"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBgmClassListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBgmClassListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBgmClassListLogic {
	return &GetBgmClassListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBgmClassListLogic) GetBgmClassList() (resp *types.GetBgmClassListRes, err error) {
	list, err := l.svcCtx.RpcClient.BgmRpcClient.GetBgmClassList(l.ctx, &bgmservice.GetBgmClassListRequest{State: tool.QuoteInt32(1)})
	resp = &types.GetBgmClassListRes{}
	for _, item := range list.List {
		resp.List = append(resp.List, types.BgmClassList{
			Id:   item.GetId(),
			Name: item.GetName(),
			Sort: item.GetSort(),
		})
	}
	return resp, err
}
