package route

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/jzero-io/jzero-admin/server/server/logic/route"
	types "github.com/jzero-io/jzero-admin/server/server/types/route"
	"github.com/jzero-io/jzero-admin/server/server/svc"
)

func GetUserRoutes(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserRoutesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := route.NewGetUserRoutes(r.Context(), svcCtx, r)
		resp, err := l.GetUserRoutes(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
