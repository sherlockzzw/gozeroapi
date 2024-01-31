package process

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoBatchLogic {
	return &VideoBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoBatchLogic) VideoBatch(req *types.VideoBatchReq) (resp *types.VideoBatchRes, err error) {
	// todo: add your logic here and delete this line

	return
}
