package template

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTemplateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTemplateListLogic {
	return &GetTemplateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTemplateListLogic) GetTemplateList() (resp *types.GetTemplateListRes, err error) {
	// todo: add your logic here and delete this line
	resp = &types.GetTemplateListRes{
		Data: types.ApiTemplateListData{
			Id:   1,
			Name: "1",
		},
		IsSellingPoint: 1,
	}
	return
}
