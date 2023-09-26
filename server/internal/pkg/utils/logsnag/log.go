package logsnag

import (
	"context"
	"encoding/json"
	"faceto-ai/internal/pkg/utils/helper"
	pkgLog "faceto-ai/internal/pkg/utils/log"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
	"os"
	"strings"
)

func init() {
	if os.Getenv("FACETOAI_ENV") == "dev" {
		PROJECT = "facetoai"
		CHANNEL = "api"
	}
}

func logger(ctx context.Context) *log.Helper {
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)
	return log.NewHelper(logger).WithContext(ctx)
}

func Event(ctx context.Context, event EventLog) {
	bgctx := helper.NewWithParentReqID(ctx)
	go func(ctx context.Context) {
		traceID := helper.GetReqIDFromContext(ctx)
		event.SetTraceID(traceID)

		payload := new(Payload)
		payload.Project = PROJECT
		payload.Channel = CHANNEL
		payload.Event = event.Event
		payload.Icon = event.Icon
		payload.Notify = event.Notify
		payload.Description = event.Description
		payload.Tags = event.Tags
		// notify
		if payload.Tags.NotifyElapsedTime > 0 && payload.Tags.ElapsedTime >= payload.Tags.NotifyElapsedTime {
			payload.Notify = true
		}
		pb, _ := json.Marshal(payload)

		httpBody := strings.NewReader(string(pb))
		client := &http.Client{}
		req, err := http.NewRequest("POST", LOGAPI, httpBody)
		if err != nil {
			logger(ctx).Errorf("logsnag.Event http.NewRequest err:%v", err)
			return
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+os.Getenv("LOGSNAG_API_KEY"))
		req.Header.Add("Trace-ID", traceID)

		res, err := client.Do(req)
		if err != nil {
			logger(ctx).Errorf("logsnag.Event client.Do err:%v", err)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			logger(ctx).Errorf("logsnag.Event io.ReadAll err:%v", err)
			return
		}
		logger(ctx).Infof("logsnag.Event http.NewRequest body:%v", string(body))
	}(bgctx)
}

func Insight(ctx context.Context, inslog InsightLog) {
	bgctx := helper.NewWithParentReqID(ctx)
	go func(ctx context.Context) {
		traceID := helper.GetReqIDFromContext(ctx)

		ins := new(InsightValue)
		ins.Project = PROJECT
		ins.Title = inslog.Title
		ins.Value = inslog.Value
		ins.Icon = inslog.Icon
		pb, _ := json.Marshal(ins)

		httpBody := strings.NewReader(string(pb))
		client := &http.Client{}
		req, err := http.NewRequest("POST", INSIGHTAPI, httpBody)
		if err != nil {
			logger(ctx).Errorf("logsnag.Insight http.NewRequest err:%v", err)
			return
		}

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+os.Getenv("LOGSNAG_API_KEY"))
		req.Header.Add("Trace-ID", traceID)

		res, err := client.Do(req)
		if err != nil {
			logger(ctx).Errorf("logsnag.Insight client.Do err:%v", err)
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			logger(ctx).Errorf("logsnag.Insight io.ReadAll err:%v", err)
			return
		}
		logger(ctx).Infof("logsnag.Insight http.NewRequest body:%v", string(body))
	}(bgctx)
}
