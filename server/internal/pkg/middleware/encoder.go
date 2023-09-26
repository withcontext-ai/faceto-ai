package middleware

import (
	"context"
	"fmt"
	netHttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"

	v1 "faceto-ai/api_gen/error/v1"
	"faceto-ai/internal/conf"
)

const (
	baseContentType = "application"
)

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

type ErrorLogParams struct {
	Error     error
	Code      int32
	Reason    string
	Operation string
	Message   string
}

func (p *ErrorLogParams) Tag() string {
	return "error_encoder"
}

func (p *ErrorLogParams) KeyValues() []interface{} {
	return []interface{}{
		tagKey,
		p.Tag(),
		p.ErrorKey(),
		fmt.Sprintf("%s", p.Error),
		p.CodeKey(),
		p.Code,
		p.ReasonKey(),
		p.Reason,
		p.OperationKey(),
		p.Operation,
		"stack_trace",
		fmt.Sprintf("%+v", p.Error),
		"message",
		p.Message,
	}
}

func (p *ErrorLogParams) ErrorKey() string {
	return log.DefaultMessageKey
}

func (p *ErrorLogParams) CodeKey() string {
	return "code"
}

func (p *ErrorLogParams) ReasonKey() string {
	return "reason"
}

func (p *ErrorLogParams) OperationKey() string {
	return "operation"
}

func WrapDefaultErrorEncoder(c *conf.Server) func(w netHttp.ResponseWriter, r *netHttp.Request, err error) {

	// DefaultErrorEncoder encodes the error to the HTTP response.
	return func(w netHttp.ResponseWriter, r *netHttp.Request, err error) {
		if errors.Is(err, context.DeadlineExceeded) && !v1.IsGatewayTimeout(err) {
			err = v1.ErrorGatewayTimeout("").WithCause(err)
		}

		se := errors.FromError(err)

		codec, _ := http.CodecForRequest(r, "Accept")
		body, err := codec.Marshal(se)
		if err != nil {
			w.WriteHeader(netHttp.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", ContentType(codec.Name()))
		w.WriteHeader(int(se.Code))
		_, _ = w.Write(body)

		var operation string
		tr, ok := transport.FromServerContext(r.Context())
		if ok {
			operation = tr.Operation()
		}

		params := ErrorLogParams{
			Error:     se.Unwrap(),
			Code:      se.Code,
			Reason:    se.Reason,
			Operation: operation,
			Message:   se.Message,
		}

		log.Context(r.Context()).Log(log.LevelError, params.KeyValues()...)

		if c.Debug == conf.ServerDebug_true {
			fmt.Printf("%+v", se.Unwrap())
		}
	}
}
