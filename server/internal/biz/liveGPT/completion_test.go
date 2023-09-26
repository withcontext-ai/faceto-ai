package liveGPT

import (
	"context"
	"faceto-ai/internal/biz"
	pkgLog "faceto-ai/internal/pkg/utils/log"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
	"time"
)

func setup() {
	fmt.Println("setup print")
}

func teardown() {
	fmt.Println("teardown print")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	fmt.Println(code)
	os.Exit(code)
}

func TestCompletion_NewChatWithAPI(t *testing.T) {
	ctx := context.Background()
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)
	if openaiClient == nil {
		return
	}
	testLinkConfig := &biz.Link{
		UUID:       "test",
		RoomName:   "test",
		Link:       "",
		ChatAPI:    "",
		ChatAPIKey: "",
	}

	room := lksdk.CreateRoom(nil)

	type fields struct {
		client *openai.Client
		log    *log.Helper
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
				client: openaiClient,
				log:    log.NewHelper(logger),
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
			c := &ChatCompletion{
				client: tt.fields.client,
				log:    tt.fields.log,
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
