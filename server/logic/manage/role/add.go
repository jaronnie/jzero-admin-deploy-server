package role

import (
	"context"
	"time"

	"github.com/guregu/null/v5"
	"github.com/zeromicro/go-zero/core/logx"

	"server/server/auth"
	types "server/server/types/manage/role"
	"server/server/svc"
	"server/server/model/manage_role"
)

type Add struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
}

func NewAdd(ctx context.Context, svcCtx *svc.ServiceContext) *Add {
	return &Add{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx,
	}
}

func (l *Add) Add(req *types.AddRequest) (resp *types.AddResponse, err error) {
	authInfo, err := auth.Info(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.Model.ManageRole.Insert(l.ctx, nil, &manage_role.ManageRole{
		Code:		req.RoleCode,
		Name:		req.RoleName,
		Desc:		req.RoleDesc,
		CreateTime:	time.Now(),
		UpdateTime:	time.Now(),
		CreateBy:	null.IntFrom(int64(authInfo.Id)).NullInt64,
		Status:		req.Status,
	})
	return
}
