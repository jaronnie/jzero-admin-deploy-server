package menu

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/status"
	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/server/errcodes/manage"
	types "github.com/jzero-io/jzero-admin/server/server/types/manage/menu"
	"github.com/jzero-io/jzero-admin/server/server/svc"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_menu"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_role_menu"
)

type Delete struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
	r	*http.Request
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Delete {
	return &Delete{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx, r: r,
	}
}

func (l *Delete) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	if len(req.Ids) == 0 {
		return nil, nil
	}
	// whether it has submenu
	subMenus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.Condition{
		Field:		manage_menu.ParentId,
		Operator:	condition.In,
		Value:		req.Ids,
	})
	if err == nil && len(subMenus) > 0 {
		return nil, status.ErrorMessage(manage.ExistSubMenuCode, l.svcCtx.Trans.Trans(l.ctx, "manage.menu.existSubMenu"))
	}
	// remove permissions
	menus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.Condition{
		Field:		manage_menu.Id,
		Operator:	condition.In,
		Value:		req.Ids,
	})
	if err == nil {
		for _, menu := range menus {
			var permissions []types.Permission
			Unmarshal(menu.Permissions.String, &permissions)
			if len(permissions) > 0 {
				roles, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.Condition{
					Field:		manage_role_menu.MenuId,
					Operator:	condition.Equal,
					Value:		menu.Id,
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
		Field:		manage_menu.Id,
		Operator:	condition.In,
		Value:		req.Ids,
	})
	return
}
