package route

import (
	"net/http"
	"server/server/logic/route"
	"server/server/svc"
	types "server/server/types/route"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetConstantRoutes(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetConstantRoutesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := route.NewGetConstantRoutes(r.Context(), svcCtx)
		resp, err := l.GetConstantRoutes(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
