package api

import (
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"github.com/jzero-io/jzero-admin/server/cmd"
	"github.com/jzero-io/jzero-admin/server/plugins"
	"github.com/jzero-io/jzero-admin/server/server/config"
	"github.com/jzero-io/jzero-admin/server/server/global"
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
	key := "CONFIG"
	if os.Getenv(key) == "" {
		_ = os.Setenv(key, `{"log":{"encoding":"plain"},"rest":{"host":"0.0.0.0","name":"server-api","port":8001,"timeout": 10000}, "sqlx":{"driverName":"pgx","dataSource": "postgres://neondb_owner:npg_le3oEzmNMr9u@ep-nameless-bar-a4xuvs05-pooler.us-east-1.aws.neon.tech/neondb?sslmode=require"},"cacheType":"local"}`)
	}

	cc := &EnvConfigurator{Key: key}
	c, err := cc.GetConfig()
	logx.Must(err)

	logx.Must(logx.SetUp(c.Log.LogConf))

	if c.Log.LogConf.Mode != "console" {
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}

	s := cmd.NewServer(c.Rest.RestConf)
	logx.Must(s.Custom.Init(cc))

	svcCtx := svc.NewServiceContext(cc, handler.Route2Code)
	global.ServiceContext = *svcCtx
	middleware.Register(s.Rest)
	handler.RegisterHandlers(s.Rest, svcCtx)

	plugins.LoadPlugins(s.Rest, *svcCtx)

	Serverless, err = rest.NewServerless(s.Rest)
	logx.Must(err)

	group := service.NewServiceGroup()
	group.Add(s.Custom)
	group.Start()
}

func Index(w http.ResponseWriter, r *http.Request) {
	Serverless.Serve(w, r)
}
