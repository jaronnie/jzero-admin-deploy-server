package auth

import (
	"net/http"
	"server/server/logic/auth"
	"server/server/svc"
	types "server/server/types/auth"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendVerificationCode(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendVerificationCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewSendVerificationCode(r.Context(), svcCtx)
		resp, err := l.SendVerificationCode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
