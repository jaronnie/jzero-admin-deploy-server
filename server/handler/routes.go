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
					Path:    "/auth/login",
					Handler: auth.Login(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/auth/refreshToken",
					Handler: auth.RefreshToken(serverCtx),
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
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getAllRoles",
					Handler: systemrole.GetAll(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getRoleList",
					Handler: systemrole.List(serverCtx),
				},
			},
		)
	}
	{
		server.AddRoutes(
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/systemManage/getUserList",
					Handler: systemuser.List(serverCtx),
				},
			},
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
