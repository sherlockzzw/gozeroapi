package bgm

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/commonx/tool"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/client/bgmservice"
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBgmListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBgmListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBgmListLogic {
	return &GetBgmListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBgmListLogic) GetBgmList(req *types.GetBgmListReq) (resp *types.GetBgmListRes, err error) {
	res, err := l.svcCtx.RpcClient.BgmRpcClient.GetBgmClassByBgmList(l.ctx, &bgmservice.GetBgmClassByBgmListRequest{
		ClassId:  req.ClassId,
		Page:     req.Page,
		PageSize: req.PageSize,
		State:    tool.QuoteInt32(1),
	})
	resp = &types.GetBgmListRes{
		Page:  res.GetPage(),
		Total: res.GetTotal(),
	}
	for _, item := range res.List {
		resp.List = append(resp.List, types.BgmList{
			Id:       item.GetId(),
			Name:     item.Name,
			BgmUrl:   item.BgmUrl,
			ClassId:  item.ClassId,
			Sort:     item.Sort,
			CoverUrl: item.CoverUrl,
			Duration: item.Duration,
		})
	}
	return
}
