package role

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/server/model/manage_role"
	types "github.com/jzero-io/jzero-admin/server/server/types/manage/role"
	"github.com/jzero-io/jzero-admin/server/server/svc"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_user_role"
)

type Delete struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
	r	*http.Request
}

func NewDelete(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *Delete {
	return &Delete{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx, r: r,
	}
}

func (l *Delete) Delete(req *types.DeleteRequest) (resp *types.DeleteResponse, err error) {
	if len(req.Ids) == 0 {
		return nil, nil
	}

	userRoles, err := l.svcCtx.Model.ManageUserRole.FindByCondition(l.ctx, nil, condition.Condition{
		Field:		manage_user_role.RoleId,
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
		Field:		manage_role.Id,
		Operator:	condition.In,
		Value:		req.Ids,
	})
	return
}
