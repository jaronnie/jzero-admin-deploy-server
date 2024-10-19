package route

import (
	"context"
	"server/server/svc"
	types "server/server/types/route"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsRouteExist struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIsRouteExist(ctx context.Context, svcCtx *svc.ServiceContext) *IsRouteExist {
	return &IsRouteExist{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsRouteExist) IsRouteExist(req *types.IsRouteExistRequest) (resp *types.IsRouteExistResponse, err error) {

	return
}
