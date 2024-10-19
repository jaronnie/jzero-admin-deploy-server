package auth

import (
	"context"
	"server/server/svc"
	types "server/server/types/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshToken struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshToken(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshToken {
	return &RefreshToken{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshToken) RefreshToken(req *types.RefreshTokenRequest) (resp *types.RefreshTokenResponse, err error) {

	return
}
