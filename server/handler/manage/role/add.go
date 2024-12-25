package role

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"server/server/logic/manage/role"
	types "server/server/types/manage/role"
	"server/server/svc"
)

func Add(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewAdd(r.Context(), svcCtx, r)
		resp, err := l.Add(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
