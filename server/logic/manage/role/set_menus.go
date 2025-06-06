package role

import (
	"context"
	"net/http"
	"time"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/jzero-io/jzero-admin/server/server/logic/manage/menu"
	menu_types "github.com/jzero-io/jzero-admin/server/server/types/manage/menu"
	types "github.com/jzero-io/jzero-admin/server/server/types/manage/role"
	"github.com/jzero-io/jzero-admin/server/server/svc"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_menu"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_role_menu"
)

type SetMenus struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
	r	*http.Request
}

func NewSetMenus(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *SetMenus {
	return &SetMenus{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx, r: r,
	}
}

func (l *SetMenus) SetMenus(req *types.SetMenusRequest) (resp *types.SetMenusResponse, err error) {
	if err = l.svcCtx.SqlxConn.TransactCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 找到该角色的首页
		roleHomeMenu, err := l.svcCtx.Model.ManageRoleMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
			Equal(manage_role_menu.RoleId, req.RoleId).
			Equal(manage_role_menu.IsHome, cast.ToInt(true)).
			Build()...)
		if err != nil {
			return errors.New("该角色无首页路由")
		}
		var datas []*manage_role_menu.ManageRoleMenu

		for _, v := range req.MenuIds {
			data := &manage_role_menu.ManageRoleMenu{
				RoleId:		int64(req.RoleId),
				MenuId:		int64(v),
				CreateTime:	time.Now(),
				UpdateTime:	time.Now(),
			}
			if data.MenuId == roleHomeMenu.MenuId {
				data.IsHome = cast.ToInt64(true)
			}
			datas = append(datas, data)
		}

		if err = l.svcCtx.Model.ManageRoleMenu.DeleteByCondition(l.ctx, session, condition.Condition{
			Field:		manage_role_menu.RoleId,
			Operator:	condition.Equal,
			Value:		req.RoleId,
		}); err != nil {
			return err
		}
		if len(datas) > 0 {
			if err = l.svcCtx.Model.ManageRoleMenu.BulkInsert(l.ctx, session, datas); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// update casbin_rule
	_, err = l.svcCtx.CasbinEnforcer.RemoveFilteredPolicy(0, cast.ToString(req.RoleId))
	if err != nil {
		return nil, errors.New("fail to remove filtered policy: " + err.Error())
	}
	// load policies
	err = l.svcCtx.CasbinEnforcer.LoadPolicy()
	if err != nil {
		return nil, errors.New("fail to load policy: " + err.Error())
	}

	// add casbin_rule
	var newPolicies [][]string
	// get menu perms
	menus, err := l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.New(condition.Condition{
		Field:		manage_menu.Id,
		Operator:	condition.In,
		Value:		req.MenuIds,
	})...)
	if err != nil {
		return nil, err
	}
	for _, v := range menus {
		var permissions []menu_types.Permission
		menu.Unmarshal(v.Permissions.String, &permissions)
		for _, perm := range permissions {
			newPolicies = append(newPolicies, []string{cast.ToString(req.RoleId), perm.Code})
		}
	}

	var b bool
	if len(newPolicies) > 0 {
		b, _ = l.svcCtx.CasbinEnforcer.AddPolicies(newPolicies)
		if !b {
			return nil, errors.New("fail to add policies")
		}
		// load policies
		err = l.svcCtx.CasbinEnforcer.LoadPolicy()
	}
	return
}
