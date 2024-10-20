package menu

import (
	"context"
	"server/server/svc"
	types "server/server/types/system/menu"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPages struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPages(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPages {
	return &GetAllPages{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPages) GetAllPages(req *types.GetAllPagesRequest) (resp *types.GetAllPagesResponse, err error) {

	return
}
