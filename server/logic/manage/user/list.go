package user

import (
	"context"
	"net/http"
	"time"

	"github.com/guregu/null/v5"
	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"

	"github.com/jzero-io/jzero-admin/server/server/model/manage_role"
	types "github.com/jzero-io/jzero-admin/server/server/types/manage/user"
	"github.com/jzero-io/jzero-admin/server/server/svc"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_user_role"
	"github.com/jzero-io/jzero-admin/server/server/model/manage_user"
)

type List struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
	r	*http.Request
}

func NewList(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *List {
	return &List{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx, r: r,
	}
}

func (l *List) List(req *types.ListRequest) (resp *types.ListResponse, err error) {
	users, total, err := l.svcCtx.Model.ManageUser.PageByCondition(l.ctx, nil, condition.Condition{
		Operator:	condition.Limit,
		Value:		req.Size,
	}, condition.Condition{
		Operator:	condition.Offset,
		Value:		(req.Current - 1) * req.Size,
	}, condition.Condition{
		Skip:		req.Username == "",
		Field:		manage_user.Username,
		Operator:	condition.Like,
		Value:		"%" + req.Username + "%",
	}, condition.Condition{
		Skip:		req.UserGender == "",
		Field:		manage_user.Gender,
		Operator:	condition.Equal,
		Value:		req.UserGender,
	}, condition.Condition{
		Skip:		req.NickName == "",
		Field:		manage_user.Nickname,
		Operator:	condition.Like,
		Value:		"%" + req.NickName + "%",
	}, condition.Condition{
		Skip:		req.UserPhone == "",
		Field:		manage_user.Phone,
		Operator:	condition.Like,
		Value:		"%" + req.UserPhone + "%",
	}, condition.Condition{
		Skip:		req.UserEmail == "",
		Field:		manage_user.Email,
		Operator:	condition.Like,
		Value:		"%" + req.UserEmail + "%",
	}, condition.Condition{
		Skip:		req.Status == "",
		Field:		manage_user.Status,
		Operator:	condition.Equal,
		Value:		req.Status,
	})

	var records []types.ManageUser
	for _, user := range users {
		records = append(records, types.ManageUser{
			Id:		user.Id,
			Username:	user.Username,
			UserGender:	user.Gender,
			NickName:	user.Nickname,
			UserPhone:	null.NewString(user.Phone.String, user.Phone.Valid).ValueOrZero(),
			UserEmail:	null.NewString(user.Email.String, user.Email.Valid).ValueOrZero(),
			Status:		user.Status,
			CreateTime:	user.CreateTime.Format(time.DateTime),
			UpdateTime:	user.UpdateTime.Format(time.DateTime),
		})
	}

	err = mr.MapReduceVoid(func(source chan<- int) {
		for index := range records {
			source <- index
		}
	}, func(item int, writer mr.Writer[types.ManageUser], cancel func(error)) {
		userRoles, err := l.svcCtx.Model.ManageUserRole.FindByCondition(l.ctx, nil, condition.Condition{
			Field:		manage_user_role.UserId,
			Operator:	condition.Equal,
			Value:		records[item].Id,
		})
		if err != nil {
			cancel(err)
			return
		}
		var roleIds []int
		for _, userRole := range userRoles {
			roleIds = append(roleIds, int(userRole.RoleId))
		}
		if len(roleIds) == 0 {
			return
		}

		roles, err := l.svcCtx.Model.ManageRole.FindByCondition(l.ctx, nil, condition.Condition{
			Field:		manage_role.Id,
			Operator:	condition.In,
			Value:		roleIds,
		})
		if err != nil {
			cancel(err)
			return
		}
		var roleCodes []string
		for _, role := range roles {
			roleCodes = append(roleCodes, role.Code)
		}
		records[item].UserRoles = roleCodes
	}, func(pipe <-chan types.ManageUser, cancel func(error)) {})

	resp = &types.ListResponse{
		Records:	records,
		PageResponse: types.PageResponse{
			Current:	req.Current,
			Size:		req.Size,
			Total:		total,
		},
	}
	return
}
