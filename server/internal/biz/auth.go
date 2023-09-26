package biz

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/pkg/middleware"
	"faceto-ai/internal/pkg/utils/crypt"
	"faceto-ai/internal/pkg/utils/helper"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type Auth struct {
	UUID         string `json:"-"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantScope   string `json:"grant_scope"`
}

func (a *Auth) GetGrantScope() []string {
	if a.GrantScope != "" {
		return strings.Split(a.GrantScope, ":")
	}
	return []string{}
}

type AuthToken struct {
	UUID         string    `json:"-"`
	ClientID     string    `json:"client_id"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    time.Time `json:"expires_in"`
}

type AuthRepo interface {
	Save(ctx context.Context, auth *Auth) error
	GetAuthByClientID(ctx context.Context, clientID string) (*Auth, error)
	GetAuthByUUID(ctx context.Context, uuid string) (*Auth, error)

	CreateAuthToken(ctx context.Context, authToken *AuthToken) error
	CreateAuthGrantByUUID(ctx context.Context, uuid string) (*middleware.AuthGrant, error)
}

type AuthUseCase struct {
	log      *log.Helper
	thirdApi *conf.ThirdApi
	authRepo AuthRepo
}

func NewAuthUseCase(
	logger log.Logger,
	thirdApi *conf.ThirdApi,
	authRepo AuthRepo,
) *AuthUseCase {
	return &AuthUseCase{
		log:      log.NewHelper(logger),
		thirdApi: thirdApi,
		authRepo: authRepo,
	}
}

func (auc *AuthUseCase) GetAuthByClientID(ctx context.Context, clientID string) (*Auth, error) {
	return auc.authRepo.GetAuthByClientID(ctx, clientID)
}

func (auc *AuthUseCase) Create(ctx context.Context, grantScope []string) (string, string, error) {
	// generate secret
	clientSecret := helper.Generator()
	secret, err := bcrypt.GenerateFromPassword([]byte(clientSecret), bcrypt.DefaultCost)
	if err != nil {
		return "", "", errors.Wrap(err, "bcrypt err")
	}

	clientID := ksuid.New().String()
	auth := &Auth{
		ClientID:     clientID,
		ClientSecret: string(secret),
		GrantScope:   strings.Join(grantScope, ":"),
	}
	if err := auc.authRepo.Save(ctx, auth); err != nil {
		return "", "", errors.Wrap(err, "create err")
	}
	return clientID, clientSecret, nil
}

func (auc *AuthUseCase) CreateAuthToken(ctx context.Context, clientID, clientSecret string) (*AuthToken, error) {
	// get auth info
	auth, err := auc.authRepo.GetAuthByClientID(ctx, clientID)
	if err != nil {
		return nil, errors.Wrap(err, "get auth")
	}

	// secret
	if err := bcrypt.CompareHashAndPassword([]byte(auth.ClientSecret), []byte(clientSecret)); err != nil {
		auc.log.Errorf("bcrypt.CompareHashAndPassword err:%v", err)
		return nil, errors.Wrap(err, "failed to hash secret")
	}

	var canGetLink, canGetTranscript, canGetAudio, canGetVideo bool
	for _, scope := range auth.GetGrantScope() {
		if scope == "link" {
			canGetLink = true
		}
		if scope == "transcript" {
			canGetTranscript = true
		}
		if scope == "audio" {
			canGetAudio = true
		}
		if scope == "video" {
			canGetVideo = true
		}
	}
	if !canGetLink && !canGetTranscript && !canGetVideo && !canGetAudio {
		auc.log.Errorf("auth grant scope err:%v", err)
		return nil, errors.New("auth grant scope err")
	}

	// create token
	expireAt := time.Second * 3600 * 24 // default 24 hours
	accessToken, err := middleware.CreateToken(expireAt, &middleware.AuthGrant{
		ClientID:         auth.ClientID,
		CanGetLink:       canGetLink,
		CanGetTranscript: canGetTranscript,
		CanGetAudio:      canGetAudio,
		CanGetVideo:      canGetVideo,
	})
	if err != nil {
		auc.log.Errorf("generate access token err:%v", err)
		return nil, errors.Wrap(err, "generate access token err")
	}

	hash := md5.Sum([]byte(middleware.AuthSecret))
	refreshToken, err := crypt.EncryptByAes(clientID+"@"+clientSecret, hex.EncodeToString(hash[:]))
	if err != nil {
		auc.log.Errorf("crypt.EncryptByAes err:%v", err)
		return nil, errors.Wrap(err, "generate refresh token err")
	}

	// create auth token
	authToken := &AuthToken{
		ClientID:     clientID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    time.Now().Add(expireAt),
	}
	if err := auc.authRepo.CreateAuthToken(ctx, authToken); err != nil {
		auc.log.Errorf("auc.authRepo.CreateAuthToken err:%v", err)
		return nil, errors.Wrap(err, "create access token")
	}

	return authToken, nil
}

func (auc *AuthUseCase) RefreshAuthToken(ctx context.Context, refreshToken string) (*AuthToken, error) {
	hash := md5.Sum([]byte(middleware.AuthSecret))
	clientIDWithSecret, err := crypt.DecryptByAes(refreshToken, hex.EncodeToString(hash[:]))
	if err != nil {
		auc.log.Errorf("crypt.DecryptByAes err:%v", err)
		return nil, errors.Wrap(err, "token decrypt err")
	}

	s := strings.Split(clientIDWithSecret, "@")
	if len(s) != 2 {
		auc.log.Errorf("clientIDWithSecret err, len != 2")
		return nil, errors.Wrap(err, "token err")
	}

	return auc.CreateAuthToken(ctx, s[0], s[1])
}
