package middleware

import (
	"net/http"
	"server/server/config"
	"time"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func Register(server *rest.Server) {
	httpx.SetOkHandler(ResponseMiddleware)
	httpx.SetErrorHandler(ErrorMiddleware)
	httpx.SetValidator(NewValidator())

	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {

			time.Sleep(time.Second * time.Duration(config.C.DelaySecond))
			next(writer, request)
		}
	})
}
