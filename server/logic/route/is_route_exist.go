package route

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"server/server/svc"
	types "server/server/types/route"
)

type IsRouteExist struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
}

func NewIsRouteExist(ctx context.Context, svcCtx *svc.ServiceContext) *IsRouteExist {
	return &IsRouteExist{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx,
	}
}

func (l *IsRouteExist) IsRouteExist(req *types.IsRouteExistRequest) (resp bool, err error) {
	return true, nil
}
