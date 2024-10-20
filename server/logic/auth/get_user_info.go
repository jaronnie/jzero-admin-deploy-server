package auth

import (
	"context"
	"server/server/svc"
	types "server/server/types/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfo struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfo(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfo {
	return &GetUserInfo{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfo) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {

	return
}
