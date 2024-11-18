package model

import (
	"server/server/model/casbin_rule"
	"server/server/model/system_email"
	"server/server/model/system_menu"
	"server/server/model/system_role"
	"server/server/model/system_role_menu"
	"server/server/model/system_user"
	"server/server/model/system_user_role"

	"github.com/eddieowens/opts"
	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type Model struct {
	CasbinRule     casbin_rule.CasbinRuleModel
	SystemEmail    system_email.SystemEmailModel
	SystemMenu     system_menu.SystemMenuModel
	SystemRole     system_role.SystemRoleModel
	SystemRoleMenu system_role_menu.SystemRoleMenuModel
	SystemUser     system_user.SystemUserModel
	SystemUserRole system_user_role.SystemUserRoleModel
}

func NewModel(conn sqlx.SqlConn, op ...opts.Opt[modelx.ModelOpts]) Model {
	return Model{
		CasbinRule:     casbin_rule.NewCasbinRuleModel(conn, op...),
		SystemEmail:    system_email.NewSystemEmailModel(conn, op...),
		SystemMenu:     system_menu.NewSystemMenuModel(conn, op...),
		SystemRole:     system_role.NewSystemRoleModel(conn, op...),
		SystemRoleMenu: system_role_menu.NewSystemRoleMenuModel(conn, op...),
		SystemUser:     system_user.NewSystemUserModel(conn, op...),
		SystemUserRole: system_user_role.NewSystemUserRoleModel(conn, op...),
	}
}
