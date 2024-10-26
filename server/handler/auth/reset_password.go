package auth

import (
	"net/http"
	"server/server/logic/auth"
	"server/server/svc"
	types "server/server/types/auth"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ResetPassword(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResetPasswordRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewResetPassword(r.Context(), svcCtx)
		resp, err := l.ResetPassword(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
