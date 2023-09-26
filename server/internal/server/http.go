package server

import (
	"context"
	v1 "faceto-ai/api_gen/error/v1"
	faceToV1 "faceto-ai/api_gen/faceto/v1"
	roomV1 "faceto-ai/api_gen/room/v1"
	voiceV1 "faceto-ai/api_gen/voice/v1"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/pkg/middleware"
	"faceto-ai/internal/server/handler"
	"faceto-ai/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	netHttp "net/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(
	c *conf.Server,
	roomService *service.RoomService,
	faceToService *service.FaceToService,
	voiceService *service.VoiceService,
	logger log.Logger,
) *http.Server {
	var opts = []http.ServerOption{
		http.ErrorEncoder(middleware.WrapDefaultErrorEncoder(c)),
		http.Filter(
			middleware.RequestIDFilter,
			handlers.CORS(
				handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Origin", "Accept"}),
				handlers.AllowedOrigins([]string{"*"}),
				handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"}),
				handlers.AllowCredentials(),
				handlers.ExposedHeaders([]string{
					"Content-Type",
					"Access-Control-Allow-Headers",
					"Access-Control-Allow-Origin",
					"Content-Length",
				}),
			),
			middleware.AccessLogFilter,
		),
		http.Middleware(
			middleware.HandlerAccessLog(),
			middleware.CorsHandler(),
			validate.Validator(),
			recovery.Recovery(recovery.WithHandler(recoveryHandler)),
			middleware.AccessTokenValidate([]string{
				faceToV1.OperationRoomLink,
				roomV1.OperationRoomRoomTranscript,

				//roomV1.OperationRoomJoinRoom,
				//roomV1.OperationRoomCheckRoom,
			}),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	// init well-knonw router
	sh := netHttp.StripPrefix("/.well-known", netHttp.FileServer(netHttp.Dir("./.well-known")))
	srv.HandlePrefix("/.well-known/", sh)

	// init HTTP Server
	roomV1.RegisterRoomHTTPServer(srv, roomService)

	// faceto API HTTP Server
	faceToV1.RegisterRoomHTTPServer(srv, faceToService)

	// voice API HTTP Server
	voiceV1.RegisterVoiceHTTPServer(srv, voiceService)

	// init special handler
	RegisterSpecialHandler(srv, roomService, faceToService, voiceService)

	return srv

}

func RegisterSpecialHandler(
	srv *http.Server,
	roomSVC *service.RoomService,
	faceToService *service.FaceToService,
	voiceService *service.VoiceService,
) {
	r := srv.Route("/")
	// webhook
	r.POST("/webhook", handler.WebhookHandler(roomSVC))

	// ChatGPT Plugin Auth Redirect
	r.GET("/v1/auth/redirect", handler.RedirectUriHandler(faceToService))

	// voice file upload
	r.POST("/v1/voice/upload", handler.VoiceUploadHandler(voiceService))
}

func recoveryHandler(_ context.Context, _, err interface{}) error {
	if e, ok := err.(error); ok {
		return v1.ErrorInternalServerError("").WithCause(e)
	}
	return v1.ErrorInternalServerError("recovery")
}
