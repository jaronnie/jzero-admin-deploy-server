package handler

import (
	"net/http"
	auth "server/server/handler/auth"
	managemenu "server/server/handler/manage/menu"
	managerole "server/server/handler/manage/role"
	manageuser "server/server/handler/manage/user"
	route "server/server/handler/route"
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
			rest.WithPrefix("/api"),
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
				{
					Method:  http.MethodPost,
					Path:    "/auth/refreshToken",
					Handler: auth.RefreshToken(serverCtx),
				},
			},
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
			rest.WithPrefix("/api"),
		)
	}
	{
		server.AddRoutes(
			rest.WithMiddlewares(
				[]rest.Middleware{serverCtx.Authx},
				[]rest.Route{
					{
						Method:  http.MethodPost,
						Path:    "/manage/addMenu",
						Handler: managemenu.Add(serverCtx),
					},
					{
						Method:  http.MethodPost,
						Path:    "/manage/deleteMenu",
						Handler: managemenu.Delete(serverCtx),
					},
					{
						Method:  http.MethodPost,
						Path:    "/manage/editMenu",
						Handler: managemenu.Edit(serverCtx),
					},
					{
						Method:  http.MethodGet,
						Path:    "/manage/getAllPages",
						Handler: managemenu.GetAllPages(serverCtx),
					},
					{
						Method:  http.MethodGet,
						Path:    "/manage/getMenuList/v2",
						Handler: managemenu.List(serverCtx),
					},
					{
						Method:  http.MethodGet,
						Path:    "/manage/getMenuTree",
						Handler: managemenu.Tree(serverCtx),
					},
				}...,
			),
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
			rest.WithPrefix("/api"),
		)
	}
	{
		server.AddRoutes(
			rest.WithMiddlewares(
				[]rest.Middleware{serverCtx.Authx},
				[]rest.Route{
					{
						Method:  http.MethodPost,
						Path:    "/manage/addRole",
						Handler: managerole.Add(serverCtx),
					},
					{
						Method:  http.MethodPost,
						Path:    "/manage/deleteRole",
						Handler: managerole.Delete(serverCtx),
					},
					{
						Method:  http.MethodPost,
						Path:    "/manage/editRole",
						Handler: managerole.Edit(serverCtx),
					},
					{
						Method:  http.MethodGet,
						Path:    "/manage/getAllRoles",
						Handler: managerole.GetAll(serverCtx),
					},
					{
						Method:  http.MethodGet,
						Path:    "/manage/getRoleList",
						Handler: managerole.List(serverCtx),
					},
					{
						Method:  http.MethodGet,
						Path:    "/manage/getRoleMenus",
						Handler: managerole.GetMenus(serverCtx),
					},
					{
						Method:  http.MethodPost,
						Path:    "/manage/setRoleMenus",
						Handler: managerole.SetMenus(serverCtx),
					},
				}...,
			),
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
			rest.WithPrefix("/api"),
		)
	}
	{
		server.AddRoutes(
			rest.WithMiddlewares(
				[]rest.Middleware{serverCtx.Authx},
				[]rest.Route{
					{
						Method:  http.MethodPost,
						Path:    "/manage/addUser",
						Handler: manageuser.Add(serverCtx),
					},
					{
						Method:  http.MethodPost,
						Path:    "/manage/deleteUser",
						Handler: manageuser.Delete(serverCtx),
					},
					{
						Method:  http.MethodPost,
						Path:    "/manage/editUser",
						Handler: manageuser.Edit(serverCtx),
					},
					{
						Method:  http.MethodGet,
						Path:    "/manage/getUserList",
						Handler: manageuser.List(serverCtx),
					},
				}...,
			),
			rest.WithJwt(serverCtx.Config.Jwt.AccessSecret),
			rest.WithPrefix("/api"),
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
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
			rest.WithPrefix("/api"),
		)

		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/route/getConstantRoutes",
					Handler: route.GetConstantRoutes(serverCtx),
				},
			},
			rest.WithPrefix("/api"),
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
			rest.WithPrefix("/api"),
		)
	}

}
