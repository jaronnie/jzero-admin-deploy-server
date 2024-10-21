package model

import (
	"server/server/model/casbin_rule"
	"server/server/model/system_menu"
	"server/server/model/system_role"
	"server/server/model/system_role_menu"
	"server/server/model/system_user"
	"server/server/model/system_user_role"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Model struct {
	CasbinRule     casbin_rule.CasbinRuleModel
	SystemMenu     system_menu.SystemMenuModel
	SystemRole     system_role.SystemRoleModel
	SystemRoleMenu system_role_menu.SystemRoleMenuModel
	SystemUser     system_user.SystemUserModel
	SystemUserRole system_user_role.SystemUserRoleModel
}

func NewModel(conn sqlx.SqlConn) Model {
	return Model{
		CasbinRule:     casbin_rule.NewCasbinRuleModel(conn),
		SystemMenu:     system_menu.NewSystemMenuModel(conn),
		SystemRole:     system_role.NewSystemRoleModel(conn),
		SystemRoleMenu: system_role_menu.NewSystemRoleMenuModel(conn),
		SystemUser:     system_user.NewSystemUserModel(conn),
		SystemUserRole: system_user_role.NewSystemUserRoleModel(conn),
	}
}
