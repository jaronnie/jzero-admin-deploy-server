package api

import (
	"net/http"
	"os"

	"github.com/jzero-io/jzero-admin/server/server/custom"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/server/server/config"
	"github.com/jzero-io/jzero-admin/server/server/handler"
	"github.com/jzero-io/jzero-admin/server/server/middleware"
	"github.com/jzero-io/jzero-admin/server/server/svc"
)

var (
	Serverless *rest.Serverless
)

type EnvConfigurator struct {
	Key string
}

func (e *EnvConfigurator) GetConfig() (config.Config, error) {
	logx.Infof("get config from env: key: %s, value: %s", e.Key, os.Getenv(e.Key))

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

	server := rest.MustNewServer(c.Rest.RestConf, rest.WithCustomCors(func(header http.Header) {
		header.Set("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Headers", "X-Request-Id")
		header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	}, nil, "*"))

	ctm := custom.New(server)
	ctm.Init()

	middleware.Register(server)

	svcCtx := svc.NewServiceContext(cc, handler.Route2Code)
	handler.RegisterHandlers(server, svcCtx)

	Serverless, err = rest.NewServerless(server)
	logx.Must(err)

	group := service.NewServiceGroup()
	group.Add(ctm)
	group.Start()
}

func Index(w http.ResponseWriter, r *http.Request) {
	Serverless.Serve(w, r)
}
