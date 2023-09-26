package liveGPT

import (
	"bufio"
	"bytes"
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/pkg/utils/helper"
	"faceto-ai/internal/pkg/utils/logsnag"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/pkg/errors"
	openai "github.com/sashabaranov/go-openai"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"strings"
	"time"
)

const HeaderXRequestID = "x-request-id"
const AuthHeader = "Authorization"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

type ChatRequest struct {
	Messages []openai.ChatCompletionMessage `json:"messages"`
	Stream   bool                           `json:"steam"`
}

type StreamResponse struct {
	Content string `json:"content"`
}

type StreamReader interface {
	Recv() (text string, err error)
	Close()
}

type streamReader struct {
	isFinished     bool
	reader         *bufio.Reader
	response       *http.Response
	errAccumulator errorAccumulator
	unmarshaler    unmarshaler
}

func (stream *streamReader) Reader() (response *StreamResponse, err error) {
	if stream.isFinished {
		err = io.EOF
		return
	}

waitForData:
	line, err := stream.reader.ReadBytes('\n')
	if err != nil {
		respErr := stream.errAccumulator.unmarshalError()
		if respErr != nil {
			if respErr.IsContextLengthExceeded() {
				err = ErrMessageTokenTooLong
			} else {
				err = fmt.Errorf("error, %w", respErr.Error)
			}
		}
		return
	}

	// read empty line
	_, _ = stream.reader.ReadBytes('\n')

	var headerData = []byte("data: ")
	line = bytes.TrimSpace(line)
	if !bytes.HasPrefix(line, headerData) {
		if writeErr := stream.errAccumulator.write(line); writeErr != nil {
			err = writeErr
			return
		}
		goto waitForData
	}

	line = bytes.TrimPrefix(line, headerData)
	if string(line) == "[DONE]" {
		stream.isFinished = true
		err = io.EOF
		return
	}

	err = stream.unmarshaler.unmarshal(line, &response)
	return
}

func (stream *streamReader) Recv() (string, error) {
	sb := strings.Builder{}
	for {
		response, err := stream.Reader()
		if err != nil {
			content := sb.String()
			if err == io.EOF && len(strings.TrimSpace(content)) != 0 {
				return content, nil
			}
			return "", err
		}

		delta := response.Content
		sb.WriteString(delta)

		if len(sb.String()) > 2 {
			if strings.HasSuffix(strings.TrimSpace(delta), ",") {
				return sb.String(), nil
			}
			if strings.HasSuffix(strings.TrimSpace(delta), "?") {
				return sb.String(), nil
			}
			if strings.HasSuffix(strings.TrimSpace(delta), ".") {
				return sb.String(), nil
			}

			if strings.HasSuffix(strings.TrimSpace(delta), "，") {
				return sb.String(), nil
			}
			if strings.HasSuffix(strings.TrimSpace(delta), "？") {
				return sb.String(), nil
			}
			if strings.HasSuffix(strings.TrimSpace(delta), "。") {
				return sb.String(), nil
			}
		}
	}
}

func (stream *streamReader) Close() {
	stream.response.Body.Close()
}

// ChatWithIndex service
type ChatWithIndex struct {
	client     *openai.Client
	log        *log.Helper
	linkConfig *biz.Link
}

func NewChatWithAPI(client *openai.Client, logger *log.Helper, linkConfig *biz.Link) *ChatWithIndex {
	return &ChatWithIndex{
		client:     client,
		log:        logger,
		linkConfig: linkConfig,
	}
}

func (c *ChatWithIndex) Complete(ctx context.Context, events []*MeetingEvent, prompt *SpeechEvent,
	participant *lksdk.RemoteParticipant, room *lksdk.Room, roomConfig *biz.Link, language *Language) (StreamReader, func(), error) {

	messages := make([]openai.ChatCompletionMessage, 0, len(events))
	for _, e := range events {
		if e.Speech != nil {
			if e.Speech.IsBot {
				messages = append(messages, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleAssistant,
					Content: e.Speech.Text,
					Name:    BotIdentity,
				})
			} else {
				messages = append(messages, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("%s said %s", e.Speech.ParticipantName, e.Speech.Text),
					Name:    e.Speech.ParticipantName,
				})
			}
		}
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: fmt.Sprintf("You are currently talking to %s", participant.Identity()),
	})

	// prompt
	if prompt.Text != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt.Text,
			Name:    prompt.ParticipantName,
		})
	}

	// request
	request := &ChatRequest{
		Messages: messages,
		Stream:   true,
	}
	return c.doChatStreamRequest(ctx, request, prompt.Text)
}

func (c *ChatWithIndex) generator() string {
	return helper.Generator()
}

func (c *ChatWithIndex) doChatStreamRequest(ctx context.Context, request *ChatRequest, text string) (StreamReader, func(), error) {
	reqURL := c.linkConfig.ChatAPI
	launchTime := time.Now()

	client := resty.New()
	client.EnableTrace()
	// max timeout 120s for return whole stream
	client.SetTimeout(120 * time.Second)

	// get first response byte timeout
	firstResponseByteTimeout := 30 * time.Second

	// cancel context when first response byte timeout
	cancelCtx, cancel := context.WithCancel(ctx)

	var respChan = make(chan *resty.Response, 1)
	reqID := c.generator()

	eg, eCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		c.log.Debugw("doChatStreamRequest", "start", HeaderXRequestID, reqID, "request", request)

		// retry
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

		resp, err := client.R().
			SetHeader(AuthHeader, "Bearer "+c.linkConfig.ChatAPIKey).
			SetHeader(HeaderXRequestID, reqID).
			SetHeader("Accept", "*/*").
			SetHeader("Content-Type", "application/json").
			SetBody(request).
			SetDoNotParseResponse(true).
			SetContext(cancelCtx).
			Post(reqURL)

		c.log.Debugw("doChatStreamRequest", "trace", HeaderXRequestID, reqID, "traceInfo", resp.Request.TraceInfo())
		if err != nil {
			// logsnag
			logsnag.Event(ctx, logsnag.Event_CHATGPT_THIRDAPI_ERROR.
				SetRoomName(c.linkConfig.RoomName).
				SetUID(c.linkConfig.UUID).
				SetNotify().
				SetError(err.Error()).
				SetMessage(text))
			c.log.Errorw("doChatStreamRequest", "end", HeaderXRequestID, reqID, "err", err)
			return errors.Wrap(err, "doChatStreamRequest")
		}
		c.log.Debugw("doChatStreamRequest", "end", HeaderXRequestID, reqID, "status", resp.StatusCode(), "response", resp.String())

		respChan <- resp
		return nil
	})

	var resp *resty.Response

	eg.Go(func() error {
		select {
		case resp = <-respChan:
			return nil
		case <-eCtx.Done():
			return nil
		case <-time.After(firstResponseByteTimeout):
			cancel()
			c.log.Errorf("doChatStreamRequest request, x_request_id:%s, err:%v", reqID, context.DeadlineExceeded)
			// logsnag
			logsnag.Event(ctx, logsnag.Event_CHATGPT_THIRDAPI_TIMEOUT.
				SetRoomName(c.linkConfig.RoomName).
				SetUID(c.linkConfig.UUID).
				SetElapsedTime(float32(firstResponseByteTimeout)).
				SetNotify().
				SetMessage(text))
			return errors.Wrap(context.DeadlineExceeded, "get first response byte timeout")
		}
	})

	if err := eg.Wait(); err != nil {
		cancel()
		return nil, func() {}, err
	}

	if resp == nil {
		c.log.Errorf("doChatStreamRequest request, x_request_id:%s, empty response", reqID)
		return nil, func() {}, errors.New("empty response")
	}

	stream := &streamReader{
		reader:         bufio.NewReader(resp.RawBody()),
		response:       resp.RawResponse,
		errAccumulator: newErrorAccumulator(),
		unmarshaler:    &jsonUnmarshaler{},
	}
	c.log.Debugw("p.launchTime", launchTime.String(), "now", time.Now(), "chat_third_api_elapsed_time", time.Since(launchTime).Seconds(), "uint", "s")

	// logsnag
	logsnag.Event(ctx, logsnag.Event_CHATGPT_THIRDAPI_ElAPSED_TIME.
		SetRoomName(c.linkConfig.RoomName).
		SetUID(c.linkConfig.UUID).
		SetElapsedTime(float32(time.Since(launchTime).Seconds())).
		SetMessage(text))

	return stream, func() {
		stream.Close()
		cancel()
	}, nil
}
