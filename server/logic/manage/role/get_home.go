package role

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/stores/condition"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jzero-io/jzero-admin/server/server/model/manage_role_menu"
	types "github.com/jzero-io/jzero-admin/server/server/types/manage/role"
	"github.com/jzero-io/jzero-admin/server/server/svc"
)

type GetHome struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
	r	*http.Request
}

func NewGetHome(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetHome {
	return &GetHome{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx, r: r,
	}
}

func (l *GetHome) GetHome(req *types.GetHomeRequest) (resp string, err error) {
	roleHomeMenu, err := l.svcCtx.Model.ManageRoleMenu.FindOneByCondition(l.ctx, nil, condition.NewChain().
		Equal(manage_role_menu.RoleId, req.RoleId).
		Equal(manage_role_menu.IsHome, cast.ToInt(true)).
		Build()...)
	if err != nil {
		return "", err
	}
	one, err := l.svcCtx.Model.ManageMenu.FindOne(l.ctx, nil, uint64(roleHomeMenu.MenuId))
	if err != nil {
		return "", err
	}
	return one.RouteName, nil
}
