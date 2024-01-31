package task

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskVideoListLogic {
	return &GetTaskVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskVideoListLogic) GetTaskVideoList(req *types.GetTaskVideoListReq) (resp *types.GetTaskVideoListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
