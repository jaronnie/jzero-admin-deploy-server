package model

import (
	"server/server/model/casbin_rule"
	"server/server/model/manage_email"
	"server/server/model/manage_menu"
	"server/server/model/manage_menu_permission"
	"server/server/model/manage_menu_role_permission"
	"server/server/model/manage_role"
	"server/server/model/manage_role_menu"
	"server/server/model/manage_user"
	"server/server/model/manage_user_role"

	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Model struct {
	CasbinRule               casbin_rule.CasbinRuleModel
	ManageEmail              manage_email.ManageEmailModel
	ManageMenu               manage_menu.ManageMenuModel
	ManageMenuPermission     manage_menu_permission.ManageMenuPermissionModel
	ManageMenuRolePermission manage_menu_role_permission.ManageMenuRolePermissionModel
	ManageRole               manage_role.ManageRoleModel
	ManageRoleMenu           manage_role_menu.ManageRoleMenuModel
	ManageUser               manage_user.ManageUserModel
	ManageUserRole           manage_user_role.ManageUserRoleModel
}

func NewModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) Model {
	return Model{
		CasbinRule:               casbin_rule.NewCasbinRuleModel(conn, op...),
		ManageEmail:              manage_email.NewManageEmailModel(conn, op...),
		ManageMenu:               manage_menu.NewManageMenuModel(conn, op...),
		ManageMenuPermission:     manage_menu_permission.NewManageMenuPermissionModel(conn, op...),
		ManageMenuRolePermission: manage_menu_role_permission.NewManageMenuRolePermissionModel(conn, op...),
		ManageRole:               manage_role.NewManageRoleModel(conn, op...),
		ManageRoleMenu:           manage_role_menu.NewManageRoleMenuModel(conn, op...),
		ManageUser:               manage_user.NewManageUserModel(conn, op...),
		ManageUserRole:           manage_user_role.NewManageUserRoleModel(conn, op...),
	}
}
