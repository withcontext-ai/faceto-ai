package liveGPT

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrMessageTokenTooLong = errors.New("message token too long")
)

const ContextLengthExceeded = "context_length_exceeded"

// APIError provides error information returned by the OpenAI API.
type APIError struct {
	Code       *string `json:"code,omitempty"`
	Message    string  `json:"message"`
	StatusCode int     `json:"-"`
}

// RequestError provides informations about generic request errors.
type RequestError struct {
	StatusCode int
	Err        error
}

type ErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

func (e *ErrorResponse) IsContextLengthExceeded() bool {
	return e.Error != nil && e.Error.Code != nil && *e.Error.Code == ContextLengthExceeded
}

func (e *APIError) Error() string {
	return e.Message
}

func (e *RequestError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return fmt.Sprintf("status code %d", e.StatusCode)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}
