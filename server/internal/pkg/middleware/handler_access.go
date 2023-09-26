package middleware

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

type HandlerAccessLogParams struct {
	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time
	Latency   time.Duration
	Operation string
}

func (p *HandlerAccessLogParams) Tag() string {
	return "handler_access"
}

func (p *HandlerAccessLogParams) KeyValues() []interface{} {
	if p.Latency > time.Minute {
		p.Latency = p.Latency.Truncate(time.Second)
	}

	return []interface{}{
		tagKey,
		p.Tag(),
		p.TimeStampKey(),
		p.TimeStamp,
		p.LatencyKey(),
		// fmt.Sprintf("%13v", p.Latency),
		p.Latency,
		p.OperationKey(),
		p.Operation,
	}
}

func (p *HandlerAccessLogParams) TimeStampKey() string {
	return "time_stamp"
}

func (p *HandlerAccessLogParams) LatencyKey() string {
	return "latency"
}

func (p *HandlerAccessLogParams) OperationKey() string {
	return "operation"
}

func HandlerAccessLog() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			start := time.Now()
			t, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			reply, err = handler(ctx, req)

			{
				param := HandlerAccessLogParams{
					TimeStamp: start,
					Latency:   time.Since(start),
					Operation: t.Operation(),
				}

				log.Context(ctx).Log(log.LevelInfo, param.KeyValues()...)
			}
			return
		}
	}
}
