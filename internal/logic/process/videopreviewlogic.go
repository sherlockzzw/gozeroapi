package process

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoPreviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoPreviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoPreviewLogic {
	return &VideoPreviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoPreviewLogic) VideoPreview(req *types.VideoPreviewReq) (resp *types.VideoPreviewRes, err error) {
	// todo: add your logic here and delete this line

	return
}
