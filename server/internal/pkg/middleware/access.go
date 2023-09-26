package middleware

import (
	"net"
	netHttp "net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const tagKey = "tag"

type AccessLogParams struct {
	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time
	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// BodySize is the size of the Response Body
	BodySize  int
	RequestID string
	UserAgent string
}

func (p *AccessLogParams) Tag() string {
	return "access"
}

func (p *AccessLogParams) KeyValues() []interface{} {
	if p.Latency > time.Minute {
		p.Latency = p.Latency.Truncate(time.Second)
	}

	clientIP, _, _ := net.SplitHostPort(strings.TrimSpace(p.ClientIP))

	return []interface{}{
		tagKey,
		p.Tag(),
		p.TimeStampKey(),
		p.TimeStamp,
		p.StatusCodeKey(),
		p.StatusCode,
		p.LatencyKey(),
		// fmt.Sprintf("%13v", p.Latency),
		p.Latency,
		p.ClientIPKey(),
		clientIP,
		p.MethodKey(),
		p.Method,
		p.PathKey(),
		p.Path,
		p.BodySizeKey(),
		p.BodySize,
		p.RequestIDKey(),
		p.RequestID,
		p.UserAgentKey(),
		p.UserAgent,
	}
}

func (p *AccessLogParams) TimeStampKey() string {
	return "time_stamp"
}

func (p *AccessLogParams) StatusCodeKey() string {
	return "status_code"
}

func (p *AccessLogParams) LatencyKey() string {
	return "latency"
}

func (p *AccessLogParams) ClientIPKey() string {
	return "client_ip"
}

func (p *AccessLogParams) MethodKey() string {
	return "method"
}

func (p *AccessLogParams) PathKey() string {
	return "path"
}

func (p *AccessLogParams) BodySizeKey() string {
	return "body_size"
}

func (p *AccessLogParams) RequestIDKey() string {
	return "request_id"
}

func (p *AccessLogParams) UserAgentKey() string {
	return "user_agent"
}

func AccessLogFilter(next netHttp.Handler) netHttp.Handler {
	return netHttp.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := newResponseWriter(w)
		next.ServeHTTP(writer, r)

		path := r.URL.Path
		raw := r.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		param := AccessLogParams{
			TimeStamp:  start,
			StatusCode: writer.Status(),
			Latency:    time.Since(start),
			ClientIP:   r.RemoteAddr,
			UserAgent:  r.UserAgent(),
			Method:     r.Method,
			Path:       path,
			BodySize:   writer.Size(),
			RequestID:  w.Header().Get(HeaderXRequestID),
		}

		log.Context(r.Context()).Log(log.LevelInfo, param.KeyValues()...)
	})
}
