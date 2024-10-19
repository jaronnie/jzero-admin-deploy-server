package route

import (
	"context"
	"server/server/svc"
	types "server/server/types/route"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRoutes struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserRoutes(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRoutes {
	return &GetUserRoutes{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserRoutes) GetUserRoutes(req *types.GetUserRoutesRequest) (resp *types.GetUserRoutesResponse, err error) {

	return
}
