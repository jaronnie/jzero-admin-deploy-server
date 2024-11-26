package menu

import (
	"net/http"
	"server/server/logic/manage/menu"
	"server/server/svc"
	types "server/server/types/manage/menu"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Tree(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TreeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewTree(r.Context(), svcCtx)
		resp, err := l.Tree(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
