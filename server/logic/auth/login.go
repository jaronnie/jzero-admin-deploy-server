package auth

import (
	"context"
	"server/server/svc"
	types "server/server/types/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type Login struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogin(ctx context.Context, svcCtx *svc.ServiceContext) *Login {
	return &Login{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Login) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {

	return
}
