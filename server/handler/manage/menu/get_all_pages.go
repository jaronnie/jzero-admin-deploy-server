package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/jzero-io/jzero-admin/server/server/logic/manage/menu"
	types "github.com/jzero-io/jzero-admin/server/server/types/manage/menu"
	"github.com/jzero-io/jzero-admin/server/server/svc"
)

func GetAllPages(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetAllPagesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := menu.NewGetAllPages(r.Context(), svcCtx, r)
		resp, err := l.GetAllPages(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
