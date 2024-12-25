package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"server/server/logic/manage/user"
	types "server/server/types/manage/user"
	"server/server/svc"
)

func Edit(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EditRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewEdit(r.Context(), svcCtx, r)
		resp, err := l.Edit(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
