// Package middleware
// ---------------------------------
// @file      : cors.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/7/4 16:27
// @desc      : file description
// ---------------------------------
package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	nethttp "net/http"
)

func CorsHandler() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if ht, ok := tr.(http.Transporter); ok {
					origin := ht.RequestHeader().Get("Origin")
					method := ht.Request().Method
					if method == nethttp.MethodOptions {
						ht.ReplyHeader().Set("Access-Control-Allow-Origin", origin)
						ht.ReplyHeader().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH,DELETE")
						ht.ReplyHeader().Set("Access-Control-Allow-Credentials", "true")
						ht.ReplyHeader().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Origin, Accept")
					}
				}
			}
			return handler(ctx, req)
		}
	}
}
