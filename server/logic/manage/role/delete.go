package role

import (
	"context"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"server/server/svc"
	types "server/server/types/manage/role"
)

type Delete struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext) *Delete {
	return &Delete{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx,
	}
}

func (l *Delete) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	if len(req.Ids) == 0 {
		return nil, nil
	}

	userRoles, err := l.svcCtx.Model.ManageUserRole.FindByCondition(l.ctx, nil, condition.Condition{
		Field:		"role_id",
		Operator:	condition.In,
		Value:		req.Ids,
	})
	if err != nil {
		return nil, err
	}
	if len(userRoles) > 0 {
		return nil, errors.New("角色被绑定, 请解除绑定后再进行删除")
	}

	err = l.svcCtx.Model.ManageRole.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:		"id",
		Operator:	condition.In,
		Value:		req.Ids,
	})
	return
}
