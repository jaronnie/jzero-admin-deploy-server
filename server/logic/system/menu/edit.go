package menu

import (
	"context"
	"server/server/svc"
	types "server/server/types/system/menu"

	"github.com/zeromicro/go-zero/core/logx"
)

type Edit struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEdit(ctx context.Context, svcCtx *svc.ServiceContext) *Edit {
	return &Edit{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Edit) Edit(req *types.EditRequest) (resp *types.EditResponse, err error) {

	return
}
