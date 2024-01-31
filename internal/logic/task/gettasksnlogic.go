package task

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskSnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskSnLogic {
	return &GetTaskSnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskSnLogic) GetTaskSn() (resp *types.GetTaskSnRes, err error) {
	// todo: add your logic here and delete this line

	return
}
