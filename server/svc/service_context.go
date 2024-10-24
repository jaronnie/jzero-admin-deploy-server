package svc

import (
	"server/server/config"
	"server/server/custom"
	"server/server/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	SqlxConn sqlx.SqlConn
	Model    model.Model

	Custom *custom.Custom
}

func NewServiceContext(c config.Config) *ServiceContext {
	svcCtx := &ServiceContext{
		Config:   c,
		SqlxConn: MustSqlConn(c),

		Custom: custom.New(),
	}
	svcCtx.Model = model.NewModel(svcCtx.SqlxConn)
	return svcCtx
}
