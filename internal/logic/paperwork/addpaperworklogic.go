package paperwork

import (
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/rpc-clipFilm/client/paperworkservice"
	"context"

	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/svc"
	"codeup.aliyun.com/64df1ec7dba61e96ebf612bf/jiandaoshou/apiClipFilm/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPaperworkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPaperworkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPaperworkLogic {
	return &AddPaperworkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPaperworkLogic) AddPaperwork(req *types.AddPaperworkReq) (resp *types.AddPaperworkRes, errMsg string, err error) {
	resp = &types.AddPaperworkRes{}
	paperwork, err := l.svcCtx.RpcClient.PaperworkRpcClient.AddPaperwork(l.ctx, &paperworkservice.AddPaperworkRequest{
		Content: req.Content,
		Type:    req.Type,
		UUID:    l.ctx.Value("uuid").(string),
	})
	if err != nil {
		return nil, errMsg, err
	}
	if paperwork.GetErrMsg() == "" {
		return nil, paperwork.GetErrMsg(), nil
	}
	if paperwork.Id != nil {
		resp.Status = 1
	}
	return
}
