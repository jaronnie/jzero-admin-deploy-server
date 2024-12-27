package api

import (
	"github.com/zeromicro/go-zero/core/service"
	"net/http"
	"os"

	"server/server/config"
	"server/server/handler"
	"server/server/middleware"
	"server/server/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var (
	server *rest.Server
)

type EnvConfigurator struct {
	Key string
}

func (e *EnvConfigurator) GetConfig() (config.Config, error) {
	var c config.Config
	if err := conf.LoadFromJsonBytes([]byte(os.Getenv(e.Key)), &c); err == nil {
		return c, nil
	} else {
		return config.Config{}, err
	}
}

func (e *EnvConfigurator) AddListener(listener func()) {
	listener()
}

func init() {
	cc := &EnvConfigurator{Key: "CONFIG"}
	c, err := cc.GetConfig()
	logx.Must(err)

	logx.Must(logx.SetUp(c.Log.LogConf))

	if c.Log.LogConf.Mode != "console" {
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}

	server = rest.MustNewServer(c.Rest.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Headers", "X-Request-Id")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	}, nil, "*"))
	middleware.Register(server)

	svcCtx := svc.NewServiceContext(cc, handler.Route2Code)
	handler.RegisterHandlers(server, svcCtx)
	svcCtx.Custom.AddRoutes(server)

	group := service.NewServiceGroup()
	group.Add(svcCtx.Custom)
	group.Start()
}

func Index(w http.ResponseWriter, r *http.Request) {
	server.ServeHTTP(w, r)
}
