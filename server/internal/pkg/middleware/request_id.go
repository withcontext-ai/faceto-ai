package middleware

import (
	"context"
	netHttp "net/http"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/uuid"
)

const HeaderXRequestID = "x-request-id"
const TraceID = "trace.id"
const SpanID = "span.id"

func ClientRequestID() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			t, ok := transport.FromClientContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			t.RequestHeader().Set(HeaderXRequestID, GetRequestOrTraceIDFromContext(ctx))
			return handler(ctx, req)
		}
	}
}

func RequestIDFilter(next netHttp.Handler) netHttp.Handler {
	return netHttp.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get(HeaderXRequestID)
		if rid == "" {
			rid = generator()
		}
		w.Header().Set(HeaderXRequestID, rid)
		next.ServeHTTP(w, r)
	})
}

func generator() string {
	return uuid.New().String()
}

func RequestID() log.Valuer {
	return func(ctx context.Context) interface{} {
		t, ok := transport.FromServerContext(ctx)
		if !ok {
			return ""
		}
		return t.ReplyHeader().Get(HeaderXRequestID)
	}
}

func RequestIDFromContext(ctx context.Context) string {
	t, ok := transport.FromServerContext(ctx)
	if !ok {
		return ""
	}
	return t.ReplyHeader().Get(HeaderXRequestID)
}

func GetTraceID(key string) log.Valuer {
	return func(ctx context.Context) interface{} {
		keyValue, ok := ctx.Value(key).(string)
		if !ok {
			return ""
		}
		return keyValue
	}
}

func GetTraceIDFromContext(ctx context.Context, key string) string {
	keyValue, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return keyValue
}

func GetRequestOrTraceIDFromContext(ctx context.Context) string {
	requestID := RequestIDFromContext(ctx)
	if requestID == "" {
		requestID = GetTraceIDFromContext(ctx, TraceID)
	}
	return requestID
}
