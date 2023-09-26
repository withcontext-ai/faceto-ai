package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport"
	"strings"
	"time"
)

const authHeader = "Authorization"
const AuthType = "bearer"
const AuthSecret = "IhKk3IiytdGzfLYr4VOpu2oJHKRwyDmFYSIWAD4pEaL"

type ContextAuthKey struct{}

func AccessTokenValidate(path []string) middleware.Middleware {
	middlewareFunc := func() middleware.Middleware {
		return func(handler middleware.Handler) middleware.Handler {
			return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
				t, ok := transport.FromServerContext(ctx)
				if !ok {
					return nil, ErrWrongContext
				}

				authToken := t.RequestHeader().Get(authHeader)
				_, authGrant, valid := authorization(ctx, authToken)
				if !valid {
					return nil, ErrWrongAuthorization
				}

				ctx = context.WithValue(ctx, ContextAuthKey{}, authGrant)
				reply, err = handler(ctx, req)
				return
			}
		}
	}
	return selector.Server(middlewareFunc()).Path(path...).
		// Regex(`/test.hello/Get[0-9]+`).
		// Prefix("/kratos.", "/go-kratos.", "/helloworld.Greeter/").
		Build()
}

func authorization(ctx context.Context, authToken string) (string, *AuthGrant, bool) {
	log.Context(ctx).Log(log.LevelInfo, "token", authToken)
	if authToken == "" {
		return "", nil, false
	}
	headerList := strings.Split(authToken, " ")
	if len(headerList) != 2 {
		log.Context(ctx).Log(log.LevelError, "token", authToken, "err", "split != 2")
		return "", nil, false
	}
	t := headerList[0]
	value := headerList[1]
	if t != "Bearer" {
		log.Context(ctx).Log(log.LevelError, "token", authToken, "err", "Bearer error")
		return "", nil, false
	}

	token, err := jwt.ParseWithClaims(value, &ClaimGrants{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(AuthSecret), nil
	})
	if err != nil {
		log.Context(ctx).Log(log.LevelError, "token", authToken, "err", err)
		return "", nil, false
	}

	if claims, ok := token.Claims.(*ClaimGrants); ok && token.Valid {
		return value, claims.AuthGrant, true
	} else {
		return "", nil, false
	}
}

// AuthJWT Get Auth Grant From Ctx
func AuthJWT(ctx context.Context) (*AuthGrant, error) {
	authGrant, ok := ctx.Value(ContextAuthKey{}).(*AuthGrant)
	if !ok {
		return nil, ErrWrongAuthorization
	}
	return authGrant, nil
}

var GrantScope = []string{"link", "transcript", "video", "audio"}

// AuthGrant auth grant struct
type AuthGrant struct {
	ClientID string `json:"client_id,omitempty"`
	// get video link
	CanGetLink bool `json:"GetLink,omitempty"`
	// get transcript
	CanGetTranscript bool `json:"GetTranscript,omitempty"`
	// get video
	CanGetVideo bool `json:"CanGetVideo,omitempty"`
	// get audio
	CanGetAudio bool `json:"CanGetAudio,omitempty"`
}

type ClaimGrants struct {
	AuthGrant      *AuthGrant `json:"authGrant,omitempty"`
	StandardClaims jwt.StandardClaims
}

func (c *ClaimGrants) Valid() error {
	return c.StandardClaims.Valid()
}

func CreateToken(expireDuration time.Duration, authGrant *AuthGrant) (string, error) {
	expire := time.Now().Add(expireDuration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &ClaimGrants{
		AuthGrant: authGrant,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})
	return token.SignedString([]byte(AuthSecret))
}
