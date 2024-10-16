package version

import (
	"context"
	"server/server/svc"
	types "server/server/types/version"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVersion struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVersion(ctx context.Context, svcCtx *svc.ServiceContext) *GetVersion {
	return &GetVersion{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVersion) GetVersion(req *types.GetVersionRequest) (resp *types.GetVersionResponse, err error) {

	return
}
