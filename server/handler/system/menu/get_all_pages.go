package menu

import (
	"net/http"
	"server/server/logic/system/menu"
	"server/server/svc"
	types "server/server/types/system/menu"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllPages(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllPagesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetAllPages(r.Context(), svcCtx)
		resp, err := l.GetAllPages(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
