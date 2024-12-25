package route

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"

	"server/server/svc"
	types "server/server/types/route"
)

type IsRouteExist struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
	r	*http.Request
}

func NewIsRouteExist(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *IsRouteExist {
	return &IsRouteExist{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx, r: r,
	}
}

func (l *IsRouteExist) IsRouteExist(req *types.IsRouteExistRequest) (resp bool, err error) {
	return true, nil
}
