package svc

import (
	"net/http"
	"server/server/middleware"

	"github.com/zeromicro/go-zero/rest"
)

type Middleware struct {
	Authx rest.Middleware
}

func NewMiddleware(svcCtx *ServiceContext, route2Code func(r *http.Request) string) Middleware {
	return Middleware{
		Authx: middleware.NewAuthxMiddleware(svcCtx.CasbinEnforcer, route2Code).Handle,
	}
}