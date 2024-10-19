package route

import (
	"net/http"
	"server/server/logic/route"
	"server/server/svc"
	types "server/server/types/route"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func IsRouteExist(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IsRouteExistRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := route.NewIsRouteExist(r.Context(), svcCtx)
		resp, err := l.IsRouteExist(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
