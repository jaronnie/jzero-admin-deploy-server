package role

import (
	"context"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"

	"server/server/svc"
	types "server/server/types/manage/role"
)

type GetHome struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
}

func NewGetHome(ctx context.Context, svcCtx *svc.ServiceContext) *GetHome {
	return &GetHome{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx,
	}
}

func (l *GetHome) GetHome(req *types.GetHomeRequest) (resp string, err error) {
	roleHomeMenu, err := l.svcCtx.Model.ManageRoleMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal("role_id", req.RoleId).
		Equal("is_home", true).
		Build()...)
	if err != nil {
		return "", err
	}
	one, err := l.svcCtx.Model.ManageMenu.FindOne(l.ctx, nil, uint64(roleHomeMenu.MenuId))
	if err != nil {
		return "", err
	}
	return one.RouteName, nil
}
