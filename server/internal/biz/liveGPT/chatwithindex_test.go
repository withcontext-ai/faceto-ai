// package biz
// ---------------------------------
// @file      : chatwithindex_test.go
// @project   : kitt
// @author    : zhangxiubo
// @time      : 2023/5/6 22:17
// @desc      : file description
// ---------------------------------
package liveGPT

import (
	"context"
	"faceto-ai/internal/biz"
	pkgLog "faceto-ai/internal/pkg/utils/log"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	openaiClient *openai.Client
)

func InitSetup() {
	fmt.Println("InitSetup")
}

func init() {
	if os.Getenv("OPENAI_KEY") == "" {
		return
	}
	openaiClient = openai.NewClient(os.Getenv("OPENAI_KEY"))
}

func TestChatWithIndex_doChatStreamRequest(t *testing.T) {
	if openaiClient == nil {
		return
	}
	type fields struct {
		client *openai.Client
	}
	type args struct {
		ctx     context.Context
		request *ChatRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "case",
			fields: fields{client: openaiClient},
			args: args{
				ctx: context.Background(),
				request: &ChatRequest{
					Messages: []openai.ChatCompletionMessage{
						{
							Role:    "user",
							Content: "hello",
						},
					},
					Stream: true,
				},
			},
		},
	}

	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatWithIndex{
				client: tt.fields.client,
				log:    log.NewHelper(logger),
				linkConfig: &biz.Link{
					ChatAPI: "https://ai-interview.withcontext.ai/v1/chat/completions",
				},
			}
			stream, got1, err := c.doChatStreamRequest(tt.args.ctx, tt.args.request, tt.args.request.Messages[0].Content)
			defer got1()
			if err != nil {
				t.Errorf("doChatStreamRequest() error = %v", err)
				return
			}
			text, err := reader(stream)
			if err != nil {
				t.Errorf("reader(stream) error = %v", err)
				return
			}
			fmt.Println(text)
		})
	}
}

func reader(stream StreamReader) (string, error) {
	sb := strings.Builder{}
	for {
		sentence, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
				break
			}
			fmt.Println("stream.Recv() err", err)
			return "", err
		}
		sb.WriteString(sentence)
		fmt.Println(sentence)
	}
	return strings.TrimSpace(sb.String()), nil
}

func TestChatWithIndex_NewChatWithAPI(t *testing.T) {
	ctx := context.Background()
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)
	if openaiClient == nil {
		return
	}

	testLinkConfig := &biz.Link{
		UUID:       "test",
		RoomName:   "test",
		Link:       "",
		ChatAPI:    "https://ai-interview.withcontext.ai/v1/chat/completions",
		ChatAPIKey: "",
	}

	room := lksdk.CreateRoom(nil)

	type fields struct {
		client     *openai.Client
		log        *log.Helper
		linkConfig *biz.Link
	}
	type args struct {
		ctx         context.Context
		events      []*MeetingEvent
		prompt      *SpeechEvent
		participant *lksdk.RemoteParticipant
		room        *lksdk.Room
		roomConfig  *biz.Link
		language    *Language
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			fields: fields{
				client:     openaiClient,
				log:        log.NewHelper(logger),
				linkConfig: testLinkConfig,
			},
			args: args{
				ctx: ctx,
				events: []*MeetingEvent{
					{
						Speech: &SpeechEvent{},
						Join: &JoinLeaveEvent{
							Leave:           false,
							ParticipantName: "",
						},
					},
				},
				prompt: &SpeechEvent{
					ParticipantName: "test",
					IsBot:           false,
					Text:            "hello",
					Timestamp:       uint64(time.Now().Unix()),
				},
				participant: &lksdk.RemoteParticipant{},
				room:        room,
				roomConfig:  testLinkConfig,
				language:    DefaultLanguage,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ChatWithIndex{
				client:     tt.fields.client,
				log:        tt.fields.log,
				linkConfig: tt.fields.linkConfig,
			}
			streamResp, streamCancel, err := c.Complete(tt.args.ctx, tt.args.events, tt.args.prompt, tt.args.participant, tt.args.room, tt.args.roomConfig, tt.args.language)
			if err != nil {
				t.Errorf("Complete() error = %v \n", err)
				return
			}
			defer streamCancel()
			text, err := reader(streamResp)
			if err != nil {
				t.Errorf("reader(stream) error = %v", err)
				return
			}
			fmt.Println(text)
		})
	}
}
