package handler

import (
	"context"
	"faceto-ai/internal/service"
	"github.com/go-kratos/kratos/v2/transport/http"
	netHttp "net/http"
)

const OperationRoomWebhook = "/interview.v1.Room/Webhook"

func WebhookHandler(srv *service.RoomService) func(ctx http.Context) error {
	return func(ctx http.Context) error {

		http.SetOperation(ctx, OperationRoomWebhook)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WebHookHandler(ctx, req.(*http.Request))
		})

		reply, err := h(ctx, ctx.Request())
		if err != nil {
			return err
		}

		return ctx.Result(netHttp.StatusOK, reply)
	}
}
