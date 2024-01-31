package template

import (
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
	// todo: add your logic here and delete this line

	return
}
