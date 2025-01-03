package route

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/jzero-io/jzero-admin/server/server/logic/route"
	types "github.com/jzero-io/jzero-admin/server/server/types/route"
	"github.com/jzero-io/jzero-admin/server/server/svc"
)

func IsRouteExist(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IsRouteExistRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := route.NewIsRouteExist(r.Context(), svcCtx, r)
		resp, err := l.IsRouteExist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
