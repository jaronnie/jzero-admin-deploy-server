package auth

import (
	"context"
	"server/server/svc"
	types "server/server/types/auth"

	"github.com/zeromicro/go-zero/core/logx"
)

type Error struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewError(ctx context.Context, svcCtx *svc.ServiceContext) *Error {
	return &Error{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Error) Error(req *types.ErrorRequest) (resp *types.ErrorResponse, err error) {

	return
}
