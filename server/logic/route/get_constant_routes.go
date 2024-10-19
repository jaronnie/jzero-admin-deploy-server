package route

import (
	"context"
	"server/server/svc"
	types "server/server/types/route"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConstantRoutes struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConstantRoutes(ctx context.Context, svcCtx *svc.ServiceContext) *GetConstantRoutes {
	return &GetConstantRoutes{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConstantRoutes) GetConstantRoutes(req *types.GetConstantRoutesRequest) (resp *types.GetConstantRoutesResponse, err error) {

	return
}
