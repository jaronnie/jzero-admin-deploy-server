package role

import (
	"net/http"
	"server/server/logic/manage/role"
	"server/server/svc"
	types "server/server/types/manage/role"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func Edit(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewEdit(r.Context(), svcCtx)
		resp, err := l.Edit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
