package user

import (
	"context"
	"server/server/model/system_user"
	"server/server/svc"
	types "server/server/types/system/user"
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
	_, err = l.svcCtx.Model.SystemUser.Insert(l.ctx, &system_user.SystemUser{
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Username:   req.Username,
		Password:   req.Password,
		Nickname:   req.NickName,
		Gender:     req.UserGender,
		Phone:      null.StringFrom(req.UserPhone).NullString,
		Status:     req.Status,
		Email:      null.StringFrom(req.UserEmail).NullString,
	})
	return
}
