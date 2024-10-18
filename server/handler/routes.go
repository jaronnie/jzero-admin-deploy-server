package handler

import (
	"net/http"
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
					Method:  http.MethodGet,
					Path:    "/version",
					Handler: version.Get(serverCtx),
				},
			},
			rest.WithPrefix("/api/v1"),
		)
	}

}
