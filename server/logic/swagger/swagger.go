package swagger

import (
	"context"
	"net/http"

	"github.com/jzero-io/jzero/core/swaggerv2"
	"github.com/jzero-io/jzero-admin/server/server/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type Swagger struct {
	logx.Logger
	ctx	context.Context
	svcCtx	*svc.ServiceContext
	r	*http.Request
	w	http.ResponseWriter
}

func NewSwagger(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request, w http.ResponseWriter) *Swagger {
	return &Swagger{
		Logger:	logx.WithContext(ctx),
		ctx:	ctx,
		svcCtx:	svcCtx,
		r:	r,
		w:	w,
	}
}

func (l *Swagger) Swagger() error {
	opts := new(swaggerv2.Swaggerv2Opts).DefaultOptions()
	swaggerv2.SwaggerHandler(opts, l.w, l.r)
	return nil
}
