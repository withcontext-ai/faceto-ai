package service

import (
	"context"
	errorV1 "faceto-ai/api_gen/error/v1"
	v1 "faceto-ai/api_gen/faceto/v1"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/data/schema"
	"faceto-ai/internal/pkg/middleware"
	"faceto-ai/internal/pkg/utils/helper"
	"faceto-ai/internal/pkg/utils/logsnag"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	netHttp "net/http"
	"strings"
	"time"
)

const defaultKey = "with-context-face-to-ai-key"

// FaceToService is a faceto api service.
type FaceToService struct {
	v1.UnimplementedRoomServer
	log *log.Helper

	linkUC *biz.LinkUseCase
	authUC *biz.AuthUseCase
}

func NewFaceToService(
	logger log.Logger,

	linkUC *biz.LinkUseCase,
	authUC *biz.AuthUseCase,
) *FaceToService {
	return &FaceToService{
		log: log.NewHelper(logger),

		linkUC: linkUC,
		authUC: authUC,
	}
}

func (s *FaceToService) Token(ctx context.Context, in *v1.TokenRequest) (*v1.TokenReply, error) {
	if in.Key != defaultKey {
		return nil, errorV1.ErrorBadRequest("")
	}

	expireAt := time.Second * 3600 * 24 // default 24 hours
	if in.ExpireAt > 0 {
		// max value 1 year 31536000
		if in.ExpireAt > 3600*24*365 {
			return nil, errorV1.ErrorBadRequest("in.ExpireAt > 1 year")
		}
		expireAt = time.Second * time.Duration(in.ExpireAt)
	}

	// create token
	token, err := middleware.CreateToken(expireAt, &middleware.AuthGrant{
		CanGetLink:       true,
		CanGetTranscript: true,
		CanGetAudio:      true,
		CanGetVideo:      true,
	})
	if err != nil {
		return nil, errorV1.ErrorBadRequest("create token err:%v", err)
	}

	s.log.WithContext(ctx).Debugf("Token:%s, Expire:%s", token, time.Now().Add(expireAt).String())
	return &v1.TokenReply{
		Token:    token,
		ExpireAt: time.Now().Add(expireAt).String(),
	}, nil
}

func (s *FaceToService) Link(ctx context.Context, in *v1.RoomLinkRequest) (*v1.RoomLinkReply, error) {
	s.log.Debugw("link request", in)
	authGrant, err := middleware.AuthJWT(ctx)
	if err != nil {
		return nil, err
	}
	if !authGrant.CanGetLink {
		s.log.WithContext(ctx).Errorf("Link().authGrant.CanGetLink false")
		logsnag.Event(ctx, logsnag.EventApplyLinkFailed.SetNotify().SetError("token authorization is not allowed."))
		return nil, errorV1.ErrorForbidden("token authorization is not allowed.")
	}

	authConfig, err := s.authUC.GetAuthByClientID(ctx, authGrant.ClientID)
	if err != nil {
		s.log.WithContext(ctx).Errorf("s.authUC.GetAuthByClientID err:%v", err)
		logsnag.Event(ctx, logsnag.EventApplyLinkFailed.SetNotify().SetError("token authorization error."))
		return nil, errorV1.ErrorForbidden("token authorization error.")
	}

	roomConfig := in.GetConfig()
	schemaConfig := new(schema.RoomConfig)
	if roomConfig != nil {
		schemaConfig.Duration = uint32(roomConfig.GetDuration())
		if schemaConfig.Duration > 0 && (schemaConfig.Duration < 180 || schemaConfig.Duration > 600) {
			logsnag.Event(ctx, logsnag.EventApplyLinkFailed.SetNotify().SetError("Config.Duration are limited to between 180 and 600."))
			return nil, errorV1.ErrorBadRequest("Config.Duration are limited to between 180 and 600.")
		}
		schemaConfig.Greeting = roomConfig.GetGreeting()
		schemaConfig.VoiceID = roomConfig.GetVoiceId()
		schemaConfig.BotName = roomConfig.GetBotname()
		schemaConfig.UserName = roomConfig.GetUsername()
	}

	var chatAPI, chatAPIKey string
	chatAPICfg := in.GetChatapi()
	if chatAPICfg != nil {
		chatAPI = chatAPICfg.GetApi()
		chatAPIKey = chatAPICfg.GetKey()
	}

	webhookCfg := in.GetWebhook()
	var schemaWebhookCfg *schema.Webhook
	if webhookCfg != nil {
		schemaWebhookCfg = &schema.Webhook{
			Api: webhookCfg.GetApi(),
			Key: webhookCfg.GetKey(),
		}
	}

	promptCfg := in.GetPrompt()
	var schemaPromptCfg *schema.Prompt
	if promptCfg != nil {
		schemaPromptCfg = &schema.Prompt{
			Role:     promptCfg.GetRole(),
			Question: promptCfg.GetQuestions(),
		}
	}

	bizLink := &biz.Link{
		ChatAPI:    chatAPI,
		ChatAPIKey: chatAPIKey,
		Config:     schemaConfig,
		Webhook:    schemaWebhookCfg,
		Token:      authConfig.UUID,
		Prompt:     schemaPromptCfg,
	}

	if err := s.linkUC.Create(ctx, bizLink); err != nil {
		s.log.WithContext(ctx).Errorf("Link().s.linkUC.Create, err:%v", err)
		logsnag.Event(ctx, logsnag.EventApplyLinkFailed.SetNotify().SetError("s.linkUC.Create err,"+err.Error()))
		return nil, errorV1.ErrorInternalServerError("Failed to get link.")
	}

	return &v1.RoomLinkReply{
		Token: bizLink.UUID,
		Name:  bizLink.RoomName,
		Link:  bizLink.Link,
	}, nil
}

func (s *FaceToService) AuthCreate(ctx context.Context, in *v1.AuthCreateRequest) (*v1.AuthCreateReply, error) {
	if in.GetKey() != defaultKey {
		return nil, errorV1.ErrorBadRequest("")
	}
	s.log.Debugw("FaceToService.AuthCreate in", in)

	// grant format: link:transcript:voice:video
	grant := in.GetGrantScope()
	scope := strings.Split(grant, ":")
	grantScope := helper.Intersect(middleware.GrantScope, scope)
	if len(grantScope) == 0 {
		grantScope = []string{"link"}
	}

	// create
	clientID, clientSecret, err := s.authUC.Create(ctx, grantScope)
	if err != nil {
		return nil, errorV1.ErrorInternalServerError("create auth")
	}

	return &v1.AuthCreateReply{
		ClientId:     clientID,
		ClientSecret: clientSecret,
	}, nil
}

func (s *FaceToService) RedirectUriHandler(ctx context.Context, req *v1.AuthRedirectRequest) (http.Redirector, error) {
	s.log.WithContext(ctx).Debugf("RedirectUriHandler, req:%v", req)
	if req.GetRedirectUri() == "" {
		return nil, errorV1.ErrorBadRequest("REDIRECT URI ERROR")
	}
	if req.GetClientId() == "" {
		return nil, errorV1.ErrorBadRequest("CLIENT ID ERROR")
	}
	// get auth info
	_, err := s.authUC.GetAuthByClientID(ctx, req.GetClientId())
	if err != nil {
		return nil, errorV1.ErrorBadRequest("AUTH ERROR")
	}

	// /v1/auth/redirect?
	// response_type=code
	// &client_id=2SXxHieyn8SEkB7QgQ0BXruomlb
	// &redirect_uri=https%3A%2F%2Fchat.openai.com%2Faip%2Fplugin-6a48296c-cf55-4508-8eec-84e212c9053e%2Foauth%2Fcallback
	// &scope=link%3Atranscript%3Avideo%3Aaudio
	// &state=94bcc5d0-a16d-4de1-bdb4-b820b6edd0fe
	redirectUri := fmt.Sprintf(
		req.GetRedirectUri()+"?client_id=%s&scope=%s&code=%s&state=%s",
		req.GetClientId(),
		req.GetScope(),
		"faceto-ai-code",
		req.GetState(),
	)
	return http.NewRedirect(redirectUri, netHttp.StatusFound), nil
}

func (s *FaceToService) Auth(ctx context.Context, in *v1.AuthRequest) (*v1.AuthReply, error) {
	s.log.Debugw("FaceToService.Auth in", in)
	if in.GetRefreshToken() != "" {
		return s.AuthExchange(ctx, &v1.AuthExchangeRequest{
			RefreshToken: in.GetRefreshToken(),
		})
	}
	if in.GetGrantType() != "authorization_code" {
		return nil, errorV1.ErrorBadRequest("param authorization_code err")
	}
	if in.GetCode() != "faceto-ai-code" {
		return nil, errorV1.ErrorBadRequest("param code err")
	}
	authToken, err := s.authUC.CreateAuthToken(ctx, in.GetClientId(), in.GetClientSecret())
	if err != nil {
		return nil, errorV1.ErrorBadRequest("create auth token")
	}
	return &v1.AuthReply{
		AccessToken:  authToken.AccessToken,
		TokenType:    middleware.AuthType,
		RefreshToken: authToken.RefreshToken,
		ExpiresIn:    int32(authToken.ExpiresIn.Unix()),
	}, nil
}

func (s *FaceToService) AuthExchange(ctx context.Context, in *v1.AuthExchangeRequest) (*v1.AuthReply, error) {
	// refresh token
	authToken, err := s.authUC.RefreshAuthToken(ctx, in.GetRefreshToken())
	if err != nil {
		return nil, errorV1.ErrorBadRequest("refresh token err")
	}

	return &v1.AuthReply{
		AccessToken:  authToken.AccessToken,
		TokenType:    middleware.AuthType,
		RefreshToken: authToken.RefreshToken,
		ExpiresIn:    int32(authToken.ExpiresIn.Unix()),
	}, nil
}
