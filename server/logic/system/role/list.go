package role

import (
	"context"
	"server/server/svc"
	types "server/server/types/system/role"
	"time"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/zeromicro/go-zero/core/logx"
)

type List struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewList(ctx context.Context, svcCtx *svc.ServiceContext) *List {
	return &List{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *List) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	roles, total, err := l.svcCtx.Model.SystemRole.PageByCondition(l.ctx, condition.Condition{
		Operator: condition.Limit,
		Value:    req.Size,
	}, condition.Condition{
		Operator: condition.Offset,
		Value:    (req.Current - 1) * req.Size,
	}, condition.Condition{
		Skip:     req.RoleName == "",
		Field:    "name",
		Operator: condition.Like,
		Value:    "%" + req.RoleName + "%",
	}, condition.Condition{
		Skip:     req.RoleCode == "",
		Field:    "code",
		Operator: condition.Like,
		Value:    "%" + req.RoleCode + "%",
	}, condition.Condition{
		Skip:     req.Status == "",
		Field:    "status",
		Operator: condition.Equal,
		Value:    req.Status,
	})

	var records []types.SystemRole
	for _, role := range roles {
		records = append(records, types.SystemRole{
			Id:         role.Id,
			RoleCode:   role.Code,
			RoleName:   role.Name,
			RoleDesc:   role.Desc,
			Status:     role.Status,
			CreateTime: role.CreateTime.Format(time.DateTime),
			UpdateTime: role.UpdateTime.Format(time.DateTime),
		})
	}

	resp = &types.ListResponse{
		Records: records,
		PageResponse: types.PageResponse{
			Current: req.Current,
			Size:    req.Size,
			Total:   total,
		},
	}
	return
}
