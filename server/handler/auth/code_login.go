package auth

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"server/server/logic/auth"
	types "server/server/types/auth"
	"server/server/svc"
)

func CodeLogin(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CodeLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewCodeLogin(r.Context(), svcCtx, r)
		resp, err := l.CodeLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
