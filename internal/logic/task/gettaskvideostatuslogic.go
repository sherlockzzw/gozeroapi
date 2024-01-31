package task

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskVideoStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaskVideoStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskVideoStatusLogic {
	return &GetTaskVideoStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaskVideoStatusLogic) GetTaskVideoStatus(req *types.GetTaskVideoStatusReq) (resp *types.GetTaskVideoStatusRes, err error) {
	// todo: add your logic here and delete this line

	return
}
