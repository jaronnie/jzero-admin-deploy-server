package role

import (
	"context"
	"server/server/svc"
	types "server/server/types/system_manage/role"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAll struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAll(ctx context.Context, svcCtx *svc.ServiceContext) *GetAll {
	return &GetAll{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAll) GetAll(req *types.GetAllRequest) (resp *types.GetAllResponse, err error) {

	return
}
