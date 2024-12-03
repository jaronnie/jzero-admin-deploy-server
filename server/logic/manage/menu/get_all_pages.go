package menu

import (
	"context"
	"server/server/model/manage_menu"
	"server/server/svc"
	types "server/server/types/manage/menu"

	"github.com/jzero-io/jzero-contrib/condition"
	"github.com/samber/lo"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllPages struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPages(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPages {
	return &GetAllPages{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPages) GetAllPages(req *types.GetAllPagesRequest) (resp []string, err error) {
	var pages []*manage_menu.ManageMenu
	if req.RoleId != 0 {
		roleMenus, err := l.svcCtx.Model.ManageRoleMenu.FindByCondition(l.ctx, nil, condition.NewChain().
			In("role_id", req.RoleId).
			Build()...)
		if err != nil {
			return nil, err
		}

		var menuIds []uint64
		for _, roleMenu := range roleMenus {
			menuIds = append(menuIds, uint64(roleMenu.MenuId))
		}
		uniqMenuIds := lo.Uniq(menuIds)

		if len(uniqMenuIds) == 0 {
			return resp, nil
		}

		pages, err = l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
			In("id", uniqMenuIds).
			Equal("menu_type", "2").
			Equal("status", "1").
			NotEqual("hide_in_menu", true).
			Build()...)
		if err != nil {
			return nil, err
		}
	} else {
		pages, err = l.svcCtx.Model.ManageMenu.FindByCondition(l.ctx, nil, condition.NewChain().
			Equal("menu_type", "2").
			Equal("status", "1").
			NotEqual("hide_in_menu", true).
			Build()...)
		if err != nil {
			return nil, err
		}
	}
	for _, page := range pages {
		resp = append(resp, page.RouteName)
	}
	return resp, nil
}
