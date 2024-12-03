package role

import (
	"context"
	"server/server/svc"
	types "server/server/types/manage/role"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHome struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHome(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHome {
	return &UpdateHome{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateHome) UpdateHome(req *types.UpdateHomeRequest) (resp *types.Empty, err error) {
	menu, err := l.svcCtx.Model.ManageMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal("route_name", req.Home).
		Build()...)
	if err != nil {
		return nil, err
	}

	oldRoleMenu, err := l.svcCtx.Model.ManageRoleMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal("role_id", req.RoleId).
		Equal("is_home", true).
		Build()...)
	if err != nil {
		return nil, err
	}
	oldRoleMenu.IsHome = cast.ToInt64(false)

	roleMenu, err := l.svcCtx.Model.ManageRoleMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal("role_id", req.RoleId).
		Equal("menu_id", menu.Id).
		Build()...)
	if err != nil {
		return nil, err
	}
	roleMenu.IsHome = cast.ToInt64(true)
	err = l.svcCtx.Model.ManageRoleMenu.Update(l.ctx, nil, roleMenu)
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.Model.ManageRoleMenu.Update(l.ctx, nil, oldRoleMenu)
	if err != nil {
		return nil, err
	}
	return
}
