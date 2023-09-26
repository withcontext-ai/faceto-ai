// Package biz
// ---------------------------------
// @file      : link_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/9/11 15:51
// @desc      : file description
// ---------------------------------
package biz

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/data/schema"
	pkgLog "faceto-ai/internal/pkg/utils/log"
	"github.com/golang/mock/gomock"
	"github.com/livekit/protocol/livekit"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"testing"
	"time"
)

func Test_Link(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)
	thirdApi := &conf.ThirdApi{}

	linkRepo := NewMockLinkRepo(ctl)
	linkUc := biz.NewLinkUseCase(logger, thirdApi, linkRepo)

	l := &biz.Link{
		UUID:       "ABC",
		RoomName:   "room-test",
		Link:       "https://faceto.withcontext.ai/rooms/room-test",
		ChatAPI:    "",
		ChatAPIKey: "",
		Token:      "",
		Config:     nil,
		Webhook:    nil,
		Prompt:     nil,
		VoiceID:    "",
	}
	cfg := &schema.RoomConfig{
		Duration: 600,
		Greeting: "hello",
		VoiceID:  "voice-id",
		BotName:  "kitt",
	}
	retErr := errors.New("link: not found")

	// 调用顺序（InOrder）
	gomock.InOrder(
		linkRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).AnyTimes(),
		linkRepo.EXPECT().GetLinkByName(gomock.Any(), "room-test").Return(l, nil),
		linkRepo.EXPECT().SetRoomVoiceID(gomock.Any(), "room-test", "voice-id").Return(nil),
		linkRepo.EXPECT().SetConfigByUUID(gomock.Any(), "ABCD", cfg).Return(retErr),
		//linkRepo.EXPECT().SetConfigByUUID(gomock.Any(), "ABC", cfg).DoAndReturn(func(ctx context.Context, id uint) (*Link, error) {
		//	fmt.Println(id)
		//	return nil, nil
		//}),
	)

	err := linkUc.Create(ctx, l)
	require.NoError(t, err)
	require.Equal(t, l, l)

	_link, err := linkUc.GetLinkByName(ctx, "room-test")
	require.NoError(t, err)
	require.Equal(t, l, _link)

	err = linkUc.SetRoomVoiceID(ctx, "room-test", "voice-id")
	require.NoError(t, err)

	err = linkUc.SetConfigByUUID(ctx, "ABCD", cfg)
	require.Equal(t, retErr, err)
}

func Test_RoomWebhook(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	ctx := context.Background()
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)
	thirdApi := &conf.ThirdApi{}
	livekitCfg := &conf.LiveKit{
		BotIdentity: "KITT",
	}
	storageCfg := &conf.Storage{}

	linkRepo := NewMockLinkRepo(ctl)
	roomRepo := NewMockRoomRepo(ctl)
	roomVodRepo := NewMockRoomVodRepo(ctl)
	roomWebhookRepo := NewMockRoomWebhookRepo(ctl)
	liveGPTWebhookRepo := NewMockLiveGPTWebhook(ctl)

	linkUc := biz.NewLinkUseCase(logger, thirdApi, linkRepo)
	roomVodUc := biz.NewRoomVodUseCase(logger, storageCfg, roomRepo, roomVodRepo)
	webhookUc := biz.NewRoomWebhookUseCase(logger, livekitCfg, storageCfg, linkUc, roomVodUc, roomWebhookRepo, roomRepo, liveGPTWebhookRepo)

	l := &biz.Link{
		UUID:       "ABC",
		RoomName:   "room-test",
		Link:       "https://faceto.withcontext.ai/rooms/room-test",
		ChatAPI:    "",
		ChatAPIKey: "",
		Token:      "",
		Config:     nil,
		Webhook:    nil,
		Prompt:     nil,
		VoiceID:    "",
	}
	cfg := &schema.RoomConfig{
		Duration: 600,
		Greeting: "hello",
		VoiceID:  "voice-id",
		BotName:  "kitt",
	}
	r := &biz.Room{
		UUID:      "room-uid",
		Name:      "room-test",
		Sid:       "room-test-sid",
		Status:    0,
		StartTime: time.Time{},
		LeftTime:  time.Time{},
		EndTime:   time.Time{},
		VodStatus: 0,
	}
	//roomWebhook := &biz.RoomWebhook{
	//	UUID:         "webhook-uid",
	//	Name:         "webhook-name",
	//	Sid:          "webhook-sid",
	//	Event:        "webhook-event",
	//	EventTime:    time.Time{},
	//	Room:         nil,
	//	Participant:  nil,
	//	Track:        nil,
	//	EgressInfo:   nil,
	//	InEgressInfo: nil,
	//}
	//wlog := &biz.ServiceWebhookLog{
	//	Name:       "room-test",
	//	Sid:        "room-test-sid",
	//	Url:        nil,
	//	Request:    nil,
	//	StatusCode: 0,
	//	Times:      0,
	//	Resp:       nil,
	//}
	roomvod := &biz.RoomVod{
		ID:           1,
		UUID:         "room-vod-uid",
		Name:         "room-test",
		Sid:          "room-test-sid",
		EgressID:     "egress-id",
		Status:       0,
		Platform:     0,
		VodType:      0,
		VodPath:      "",
		VodURL:       "",
		StartTime:    time.Time{},
		CompleteTime: time.Time{},
		Duration:     0,
	}

	retErr := errors.New("link: not found")

	// link repo expect
	linkRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	linkRepo.EXPECT().GetLinkByName(gomock.Any(), "room-test").Return(l, nil).AnyTimes()
	linkRepo.EXPECT().SetRoomVoiceID(gomock.Any(), "room-test", "voice-id").Return(nil).AnyTimes()
	linkRepo.EXPECT().SetConfigByUUID(gomock.Any(), "ABCD", cfg).Return(retErr).AnyTimes()
	//linkRepo.EXPECT().SetConfigByUUID(gomock.Any(), "ABC", cfg).DoAndReturn(func(ctx context.Context, id uint) (*Link, error) {
	//	fmt.Println(id)
	//	return nil, nil
	//})

	// room repo expect
	roomRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
	roomRepo.EXPECT().UpdateStatus(gomock.Any(), "room-test-sid", gomock.Any(), gomock.Any(), &biz.Room{}).Return(1, nil).AnyTimes()
	roomRepo.EXPECT().GetByID(gomock.Any(), 1).Return(r, nil).AnyTimes()
	roomRepo.EXPECT().GetByUUID(gomock.Any(), "room-uid").Return(r, nil).AnyTimes()
	roomRepo.EXPECT().GetBySID(gomock.Any(), "room-test-sid").Return(r, nil).AnyTimes()
	roomRepo.EXPECT().GetByName(gomock.Any(), "room-test").Return(r, nil).AnyTimes()
	roomRepo.EXPECT().List(gomock.Any(), 0).Return([]*biz.Room{r}, nil).AnyTimes()

	// room vod repo expect
	roomVodRepo.EXPECT().Save(gomock.Any(), roomvod).Return(nil).AnyTimes()
	roomVodRepo.EXPECT().GetByEgressID(gomock.Any(), "egress-id").Return(roomvod, nil).AnyTimes()
	roomVodRepo.EXPECT().GetBySid(gomock.Any(), "room-test-sid").Return(roomvod, nil).AnyTimes()
	roomVodRepo.EXPECT().UpdateStatus(gomock.Any(), "egress-id", gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	// room webhook repo expect
	roomWebhookRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	// liveGPT Webhook repo expect
	//liveGPTWebhookRepo.EXPECT().Save(gomock.Any(), wlog).Return(nil),
	liveGPTWebhookRepo.EXPECT().EventRoomStarted(gomock.Any(), "room-test", "room-test-sid").Return(nil).AnyTimes()
	liveGPTWebhookRepo.EXPECT().EventParticipantJoined(gomock.Any(), "room-test", "room-test-sid").Return(nil).AnyTimes()
	liveGPTWebhookRepo.EXPECT().EventParticipantLeft(gomock.Any(), "room-test", "room-test-sid").Return(nil).AnyTimes()
	liveGPTWebhookRepo.EXPECT().EventRoomEgressEnded(gomock.Any(), "room-test", "room-test-sid").Return(nil).AnyTimes()
	liveGPTWebhookRepo.EXPECT().EventRoomFinished(gomock.Any(), "room-test", "room-test-sid").Return(nil).AnyTimes()

	// 调用顺序（InOrder）
	//gomock.InOrder()

	event := &livekit.WebhookEvent{
		Id:        "EV_EANZaEhu8TMU",
		Event:     "room_started",
		CreatedAt: 1691553790,
		Room: &livekit.Room{
			Sid:  "room-test-sid",
			Name: "room-test",
		},
	}

	err := webhookUc.EventRoomStarted(ctx, event)
	require.NoError(t, err)

	err = webhookUc.Save(ctx, event)
	require.NoError(t, err)

	event = &livekit.WebhookEvent{
		Id:        "EV_EANZaEhu8TMU",
		Event:     "participant_joined",
		CreatedAt: 1691553790,
		Room: &livekit.Room{
			Sid:  "room-test-sid",
			Name: "room-test",
		},
	}
	err = webhookUc.EventParticipantJoined(ctx, event)
	require.NoError(t, err)

	event = &livekit.WebhookEvent{
		Id:        "EV_EANZaEhu8TMU",
		Event:     "participant_left",
		CreatedAt: 1691553790,
		Room: &livekit.Room{
			Sid:  "room-test-sid",
			Name: "room-test",
		},
	}
	err = webhookUc.EventParticipantLeft(ctx, event)
	require.NoError(t, err)

	event = &livekit.WebhookEvent{
		Id:        "EV_EANZaEhu8TMU",
		Event:     "track_published",
		CreatedAt: 1691553790,
		Participant: &livekit.ParticipantInfo{
			Identity: "KITT",
		},
		Room: &livekit.Room{
			Sid:  "room-test-sid",
			Name: "room-test",
		},
	}
	err = webhookUc.EventTrackPublished(ctx, event)
	require.NoError(t, err)

	event = &livekit.WebhookEvent{
		Id:        "EV_EANZaEhu8TMU",
		Event:     "egress_ended",
		CreatedAt: 1691553790,
		EgressInfo: &livekit.EgressInfo{
			EgressId:  "egress-id",
			RoomId:    "room-test-sid",
			RoomName:  "room-test",
			Status:    3,
			StartedAt: 1691553761343377109,
			EndedAt:   1691553789616161606,
			FileResults: []*livekit.FileInfo{
				{
					Filename:  "egress/a7hx-fn5t-1691553760.mp4",
					StartedAt: 1691553761343377109,
					EndedAt:   1691553789616161606,
					Duration:  24210839037,
					Size:      11690339,
				},
			},
		},
	}
	err = webhookUc.EventEgressEnded(ctx, event)
	require.NoError(t, err)

	event = &livekit.WebhookEvent{
		Id:        "EV_EANZaEhu8TMU",
		Event:     "room_finished",
		CreatedAt: 1691553790,
		Room: &livekit.Room{
			Sid:  "room-test-sid",
			Name: "room-test",
		},
	}
	err = webhookUc.EventRoomFinished(ctx, event)
	require.NoError(t, err)
}
