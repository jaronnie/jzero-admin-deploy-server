package user

import (
	"context"
	"server/server/svc"
	types "server/server/types/system/user"
	"time"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/jzero-io/jzero-contrib/nullx"
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
	users, total, err := l.svcCtx.Model.SystemUser.PageByCondition(l.ctx, condition.Condition{
		Operator: condition.Limit,
		Value:    req.Size,
	}, condition.Condition{
		Operator: condition.Offset,
		Value:    (req.Current - 1) * req.Size,
	})

	var records []types.SystemUser
	for _, user := range users {
		records = append(records, types.SystemUser{
			Id:         user.Id,
			Username:   user.Username,
			UserGender: user.Gender,
			NickName:   user.Nickname,
			UserPhone:  nullx.NewString(user.Phone).ValueOrZero(),
			UserEmail:  nullx.NewString(user.Email).ValueOrZero(),
			UserRoles:  []string{"R_super"},
			Status:     user.Status,
			CreateTime: user.CreateTime.Format(time.DateTime),
			UpdateTime: user.UpdateTime.Format(time.DateTime),
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
