package helper

import (
	"context"
	"faceto-ai/internal/pkg/middleware"
)

func NewWithParentReqID(ctx context.Context) context.Context {
	reqID := middleware.RequestIDFromContext(ctx)
	if reqID == "" {
		id, ok := ctx.Value(middleware.TraceID).(string)
		if !ok {
			reqID = Generator()
		} else {
			reqID = id
		}
	}
	bgCtx := context.Background()
	return context.WithValue(bgCtx, middleware.TraceID, reqID)
}

func GetReqIDFromContext(ctx context.Context) string {
	reqID := middleware.RequestIDFromContext(ctx)
	if reqID == "" {
		reqID = ctx.Value(middleware.TraceID).(string)
	}
	return reqID
}
