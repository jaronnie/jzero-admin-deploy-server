package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/jzero-io/jzero-admin/server/server/logic/manage/role"
	types "github.com/jzero-io/jzero-admin/server/server/types/manage/role"
	"github.com/jzero-io/jzero-admin/server/server/svc"
)

func GetMenus(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMenusRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewGetMenus(r.Context(), svcCtx)
		resp, err := l.GetMenus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
