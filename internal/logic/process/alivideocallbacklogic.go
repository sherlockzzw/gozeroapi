package process

import (
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AliVideoCallBackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAliVideoCallBackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AliVideoCallBackLogic {
	return &AliVideoCallBackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AliVideoCallBackLogic) AliVideoCallBack(req *types.AliVideoCallBackReq) (resp *types.AliVideoCallBackRes, err error) {
	// todo: add your logic here and delete this line

	return
}
