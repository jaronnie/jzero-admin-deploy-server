package role

import (
	"context"
	"server/server/auth"
	"server/server/model/system_role"
	"server/server/svc"
	types "server/server/types/system/role"
	"time"

	"github.com/guregu/null/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type Add struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdd(ctx context.Context, svcCtx *svc.ServiceContext) *Add {
	return &Add{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Add) Add(req *types.AddRequest) (resp *types.AddResponse, err error) {
	authInfo, err := auth.Info(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.Model.SystemRole.Insert(l.ctx, &system_role.SystemRole{
		Code:       req.RoleCode,
		Name:       req.RoleName,
		Desc:       req.RoleDesc,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		CreateBy:   null.IntFrom(int64(authInfo.Id)).NullInt64,
		Status:     req.Status,
	})
	return
}
