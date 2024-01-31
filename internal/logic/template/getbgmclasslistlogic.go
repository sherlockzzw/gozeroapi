package template

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
