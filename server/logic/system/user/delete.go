package user

import (
	"context"
	"server/server/svc"
	types "server/server/types/system/user"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"
)

type Delete struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext) *Delete {
	return &Delete{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Delete) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	if len(req.Ids) == 0 {
		return nil, nil
	}

	err = l.svcCtx.Model.SystemUser.DeleteByCondition(l.ctx, nil, condition.Condition{
		Field:    "id",
		Operator: condition.In,
		Value:    req.Ids,
	})
	return
}
