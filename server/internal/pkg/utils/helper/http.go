package helper

import (
	"context"
	pkgLog "faceto-ai/internal/pkg/utils/log"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

func logger(ctx context.Context) *log.Helper {
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)
	return log.NewHelper(logger).WithContext(ctx)
}

type RestyOptions struct {
	Url      string            `json:"url"`
	Req      interface{}       `json:"req"`
	Response interface{}       `json:"response"`
	IsRetry  bool              `json:"isRetry"`
	Headers  map[string]string `json:"headers"`
	Timeout  time.Duration     `json:"timeout"`

	// retry config
	RetryCount       int                                                `json:"retryCount"`
	RetryWaitTime    time.Duration                                      `json:"retryWaitTime"`
	RetryMaxWaitTime time.Duration                                      `json:"retryMaxWaitTime"`
	RetryCondition   func(response *resty.Response, err error) (b bool) `json:"retryCondition"`
}

func RestyRequest(ctx context.Context, restyOptions *RestyOptions) (*resty.Response, error) {
	if restyOptions == nil || restyOptions.Url == "" {
		return nil, nil
	}

	client := resty.New()

	client.EnableTrace()
	if restyOptions.Timeout > 0 {
		client.SetTimeout(restyOptions.Timeout)
	}

	reqID := uuid.New().String()

	logger(ctx).Debugw("RestyRequest", "trace", "x-request-id", reqID, "request", restyOptions.Req)

	// retry
	if restyOptions.IsRetry {
		// default config
		RetryCount := 3
		RetryWaitTime := time.Duration(100) * time.Millisecond
		RetryMaxWaitTime := time.Second * time.Duration(15)
		RetryCondition := func(response *resty.Response, err error) (b bool) {
			if response == nil || response.StatusCode() == 0 || (response.StatusCode() >= http.StatusLocked && response.StatusCode() < http.StatusNotExtended) {
				return true
			} else {
				return false
			}
		}
		if restyOptions.RetryCount > 0 {
			RetryCount = restyOptions.RetryCount
		}
		if restyOptions.RetryWaitTime > 0 {
			RetryWaitTime = restyOptions.RetryWaitTime
		}
		if restyOptions.RetryMaxWaitTime > 0 {
			RetryMaxWaitTime = restyOptions.RetryMaxWaitTime
		}
		if restyOptions.RetryCondition != nil {
			RetryCondition = restyOptions.RetryCondition
		}
		client.SetRetryCount(RetryCount)
		client.SetRetryWaitTime(RetryWaitTime)
		client.SetRetryMaxWaitTime(RetryMaxWaitTime)
		client.AddRetryCondition(RetryCondition)
	}

	clientR := client.R()
	if len(restyOptions.Headers) > 0 {
		for key, val := range restyOptions.Headers {
			clientR = clientR.SetHeader(key, val)
		}
	}

	clientR = clientR.
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader("x-request-id", reqID)

	resp, err := clientR.
		//SetDoNotParseResponse(true).
		SetContext(ctx).
		SetBody(restyOptions.Req).
		SetResult(restyOptions.Response).
		Post(restyOptions.Url)

	logger(ctx).Debugw("RestyRequest", "trace", "x-request-id", reqID, "traceInfo", resp.Request.TraceInfo())
	if err != nil {
		logger(ctx).Errorw("RestyRequest", "end", "err", err)
		return resp, err
	}
	logger(ctx).Debugw("RestyRequest", "trace", "x-request-id", reqID, "status", resp.StatusCode(), "response", resp.String())

	if resp.IsError() {
		logger(ctx).Errorw("RestyRequest", "end", "err", err)
		return resp, err
	}
	return resp, nil
}
