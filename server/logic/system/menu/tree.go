package menu

import (
	"context"
	"server/server/svc"
	types "server/server/types/system/menu"

	"github.com/zeromicro/go-zero/core/logx"
)

type Tree struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTree(ctx context.Context, svcCtx *svc.ServiceContext) *Tree {
	return &Tree{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Tree) Tree(req *types.TreeRequest) (resp []types.TreeResponse, err error) {
	resp = []types.TreeResponse{}
	return
}
