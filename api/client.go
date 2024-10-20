package api

import (
	"server/server/config"
	"server/server/handler"
	"server/server/middleware"
	"server/server/svc"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var (
	server *rest.Server
)

func init() {
	var c config.Config

	if err := conf.LoadFromJsonBytes([]byte(os.Getenv("CONFIG")), &c); err != nil {
		panic(err)
	}
	config.C = c

	if err := logx.SetUp(c.Log.LogConf); err != nil {
		logx.Must(err)
	}
	if c.Log.LogConf.Mode != "console" {
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}

	server = rest.MustNewServer(c.Rest.RestConf)
	middleware.Register(server)

	svcCtx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, svcCtx)

	svcCtx.Custom.AddRoutes(server)

	logx.Infof("Starting rest server at %s:%d...", svcCtx.Config.Rest.Host, svcCtx.Config.Rest.Port)
}

func Index(w http.ResponseWriter, r *http.Request) {
	server.ServeHTTP(w, r)
}
