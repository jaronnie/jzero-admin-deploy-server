package handler

import (
	"net/http"
	auth "server/server/handler/auth"
	route "server/server/handler/route"
	systemmenu "server/server/handler/system/menu"
	systemrole "server/server/handler/system/role"
	systemuser "server/server/handler/system/user"
	version "server/server/handler/version"
	"server/server/svc"
	"time"

	"github.com/zeromicro/go-zero/rest"
)

var (
	_ = http.StatusOK
	_ = time.Now()
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/auth/code-login",
					Handler: auth.CodeLogin(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/auth/pwd-login",
					Handler: auth.PwdLogin(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/auth/refreshToken",
					Handler: auth.RefreshToken(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/auth/register",
					Handler: auth.Register(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/auth/resetPassword",
					Handler: auth.ResetPassword(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/auth/sendVerificationCode",
					Handler: auth.SendVerificationCode(serverCtx),
				},
			},
		)

		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/auth/error",
					Handler: auth.Error(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/auth/getUserInfo",
					Handler: auth.GetUserInfo(serverCtx),
				},
			},
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/route/getConstantRoutes",
					Handler: route.GetConstantRoutes(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/route/getUserRoutes",
					Handler: route.GetUserRoutes(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/route/isRouteExist",
					Handler: route.IsRouteExist(serverCtx),
				},
			},
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getAllPages",
					Handler: systemmenu.GetAllPages(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getMenuList/v2",
					Handler: systemmenu.List(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getMenuTree",
					Handler: systemmenu.Tree(serverCtx),
				},
			},
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/system/addRole",
					Handler: systemrole.Add(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/system/deleteRole",
					Handler: systemrole.Delete(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/system/editRole",
					Handler: systemrole.Edit(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/system/getAllRoles",
					Handler: systemrole.GetAll(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/system/getRoleList",
					Handler: systemrole.List(serverCtx),
				},
			},
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/system/addUser",
					Handler: systemuser.Add(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/system/deleteUser",
					Handler: systemuser.Delete(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/system/editUser",
					Handler: systemuser.Edit(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/system/getUserList",
					Handler: systemuser.List(serverCtx),
				},
			},
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/version",
					Handler: version.Get(serverCtx),
				},
			},
			rest.WithPrefix("/api/v1"),
		)
	}

}
