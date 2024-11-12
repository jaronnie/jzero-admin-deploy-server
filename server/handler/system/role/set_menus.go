package role

import (
	"net/http"
	"server/server/logic/system/role"
	"server/server/svc"
	types "server/server/types/system/role"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetMenus(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetMenusRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewSetMenus(r.Context(), svcCtx)
		resp, err := l.SetMenus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
