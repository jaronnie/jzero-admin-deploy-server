package menu

import (
	"context"
	"server/server/errcodes/manage"
	"server/server/svc"
	types "server/server/types/manage/menu"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/jzero-io/jzero-contrib/status"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
)

type Delete struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext) *Delete {
	return &Delete{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	if len(req.Ids) == 0 {
		return nil, nil
	}

	subMenus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.Condition{
		Field:    "parent_id",
		Operator: condition.In,
		Value:    req.Ids,
	})
	if err == nil && len(subMenus) > 0 {
		return nil, status.Error(manage.ExistSubMenuCode)
	}

	menus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.Condition{
		Field:    "id",
		Operator: condition.In,
		Value:    req.Ids,
	})
	if err == nil {
		for _, menu := range menus {
			var permissions []types.Permission
			Unmarshal(menu.Permissions.String, &permissions)
			if len(permissions) > 0 {
				roles, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.Condition{
					Field:    "menu_id",
					Operator: condition.Equal,
					Value:    menu.Id,
				})
				if err != nil {
					return nil, err
				}
				for _, role := range roles {
					for _, permission := range permissions {
						_, _ = l.svcCtx.CasbinEnforcer.RemovePolicy(cast.ToString(role.RoleId), permission.Code)
					}
				}
			}
		}
	}
	err = l.svcCtx.Model.ManageMenu.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:    "id",
		Operator: condition.In,
		Value:    req.Ids,
	})
	return
}
