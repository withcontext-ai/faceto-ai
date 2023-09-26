package data

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/data/ent"
	"faceto-ai/internal/data/ent/auth"
	"faceto-ai/internal/pkg/middleware"
	"faceto-ai/internal/pkg/utils/helper"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"strings"
)

type authRepo struct {
	data *Data
	log  *log.Helper
}

func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *authRepo) GetAuthByClientID(ctx context.Context, clientID string) (*biz.Auth, error) {
	result, err := r.data.DB(ctx).Auth.Query().
		Where(auth.ClientID(clientID)).
		Where(auth.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrAuthFound
		}
		return nil, err
	}
	return &biz.Auth{
		UUID:         result.UUID,
		ClientID:     result.ClientID,
		ClientSecret: result.ClientSecret,
		GrantScope:   result.GrantScope,
	}, nil
}

func (r *authRepo) GetAuthByUUID(ctx context.Context, uuid string) (*biz.Auth, error) {
	result, err := r.data.DB(ctx).Auth.Query().
		Where(auth.UUID(uuid)).
		Where(auth.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrAuthFound
		}
		return nil, err
	}
	return &biz.Auth{
		UUID:         result.UUID,
		ClientID:     result.ClientID,
		ClientSecret: result.ClientSecret,
		GrantScope:   result.GrantScope,
	}, nil
}

func (r *authRepo) Save(ctx context.Context, auth *biz.Auth) error {
	db := r.data.DB(ctx).Auth.Create().
		SetClientID(auth.ClientID).
		SetClientSecret(auth.ClientSecret).
		SetGrantScope(auth.GrantScope)

	result, err := db.Save(ctx)
	if err != nil {
		return err
	}

	auth.UUID = result.UUID
	return nil
}

func (r *authRepo) CreateAuthToken(ctx context.Context, authToken *biz.AuthToken) error {
	db := r.data.DB(ctx).AuthToken.Create().
		SetClientID(authToken.ClientID).
		SetAccessToken(authToken.AccessToken).
		SetRefreshToken(authToken.RefreshToken).
		SetExpiresIn(authToken.ExpiresIn)

	result, err := db.Save(ctx)
	if err != nil {
		return err
	}

	authToken.UUID = result.UUID
	return nil
}

func (r *authRepo) CreateAuthGrantByUUID(ctx context.Context, uuid string) (*middleware.AuthGrant, error) {
	result, err := r.GetAuthByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	fmt.Println(result.GrantScope)
	if result.GrantScope == "" {
		return &middleware.AuthGrant{}, nil
	}

	existScope := helper.Intersect(middleware.GrantScope, strings.Split(result.GrantScope, ":"))
	fmt.Println(existScope)
	CanGetLink, _ := helper.IndexOfStrSlice(existScope, "link")
	CanGetTranscript, _ := helper.IndexOfStrSlice(existScope, "transcript")
	CanGetVideo, _ := helper.IndexOfStrSlice(existScope, "video")
	CanGetAudio, _ := helper.IndexOfStrSlice(existScope, "audio")
	return &middleware.AuthGrant{
		ClientID:         result.ClientID,
		CanGetLink:       CanGetLink,
		CanGetTranscript: CanGetTranscript,
		CanGetVideo:      CanGetVideo,
		CanGetAudio:      CanGetAudio,
	}, nil
}
