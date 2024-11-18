package svc

import (
	"server/pkg/localcache"
	"server/server/config"
	"server/server/custom"
	"server/server/model"

	"github.com/jzero-io/jzero-contrib/modelx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/syncx"
)

type ServiceContext struct {
	Config   config.Config
	SqlxConn sqlx.SqlConn
	Model    model.Model
	Cache    cache.Cache

	Custom *custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		SqlxConn: MustSqlConn(c),

		Custom: custom.New(),
	}
	if c.CacheType == "local" {
		svcCtx.Cache = &localcache.Cache{
			Vals: make(map[string][]byte),
		}
	} else {

		singleFlights := syncx.NewSingleFlight()
		stats := cache.NewStat("redis-cache")
		svcCtx.Cache = cache.NewNode(&redis.Redis{
			Addr: svcCtx.Config.Redis.Host,
			Type: svcCtx.Config.Redis.Type,
			Pass: svcCtx.Config.Redis.Pass,
		}, singleFlights, stats, errors.New("no cache"))
	}
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn, modelx.WithCachedConn(sqlc.NewConnWithCache(svcCtx.SqlxConn, svcCtx.Cache)))
	return svcCtx
}
