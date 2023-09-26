package elevenlabs

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
)

const apiEndpoint = "https://api.elevenlabs.io"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrUnspecified  = errors.New("unspecified error")
)

type Client struct {
	apiKey   string
	log      *log.Helper
	endpoint string
}

func New(apiKey string, log *log.Helper) Client {
	return Client{
		apiKey:   apiKey,
		log:      log,
		endpoint: apiEndpoint,
	}
}

func (c Client) WithEndpoint(endpoint string) Client {
	c.endpoint = endpoint
	return c
}
