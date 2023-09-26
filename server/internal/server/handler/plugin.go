package handler

import (
	"context"
	v1 "faceto-ai/api_gen/faceto/v1"
	"faceto-ai/internal/service"
	netHttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	OperationAuthRedirectUri = "/faceto.v1.Room/AuthRedirect"
)

func RedirectUriHandler(srv *service.FaceToService) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.AuthRedirectRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}

		http.SetOperation(ctx, OperationAuthRedirectUri)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RedirectUriHandler(ctx, req.(*v1.AuthRedirectRequest))
		})

		reply, err := h(ctx, &in)
		if err != nil {
			return err
		}

		return ctx.Result(netHttp.StatusFound, reply)
	}
}
