package data

import (
	"context"
	"errors"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/pkg/middleware"
	"faceto-ai/internal/pkg/utils/helper"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type httpRepo struct {
	data *Data
	log  *log.Helper
}

func NewHttpRepo(data *Data, logger log.Logger) biz.HttpRepo {
	return &httpRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (h *httpRepo) Request(ctx context.Context, url string, req interface{}, response interface{}, isRetry bool) error {
	client := resty.New()

	client.EnableTrace()

	reqID := helper.GetReqIDFromContext(ctx)
	h.log.WithContext(ctx).Debugw("httpRepo, Request", "begin", "x-request-id", reqID, "request", req)

	// retry
	if isRetry {
		client.SetRetryCount(3)
		client.SetRetryWaitTime(time.Duration(100) * time.Millisecond)
		client.SetRetryMaxWaitTime(time.Second * time.Duration(15))
		client.AddRetryCondition(func(response *resty.Response, err error) (b bool) {
			if response == nil || response.StatusCode() == 0 || (response.StatusCode() >= http.StatusLocked && response.StatusCode() < http.StatusNotExtended) {
				return true
			} else {
				return false
			}
		})
	}

	resp, err := client.R().
		SetHeader("Accept", "*/*").
		SetHeader("Content-Type", "application/json").
		SetHeader(middleware.HeaderXRequestID, reqID).
		//SetDoNotParseResponse(true).
		SetContext(ctx).
		SetBody(req).
		SetResult(response).
		Post(url)

	h.log.WithContext(ctx).Debugw("httpRepo, Request", "trace", "x-request-id", reqID, "traceInfo", resp.Request.TraceInfo())
	if err != nil {
		h.log.WithContext(ctx).Errorw("httpRepo, Request", "end", "x-request-id", reqID, "err", err)
		return err
	}

	h.log.WithContext(ctx).Debugw("httpRepo, Request", "end", "x-request-id", reqID, "status", resp.StatusCode(), "response", resp.String())
	if resp.IsError() {
		h.log.WithContext(ctx).Errorw("httpRepo, Request", "end", "x-request-id", reqID, "status", resp.StatusCode())
		return errors.New(fmt.Sprintf("http status [%d]", resp.StatusCode()))
	}
	return nil
}
