package biz

import (
	"context"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/data/schema"
	"faceto-ai/internal/pkg/utils/helper"
	"faceto-ai/internal/pkg/utils/logsnag"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/webhook"
	"net/http"
	"strings"
	"time"
)

type RoomWebhook struct {
	UUID         string
	Name         string
	Sid          string
	Event        string
	EventTime    time.Time
	Room         *livekit.Room
	Participant  *livekit.ParticipantInfo
	Track        *livekit.TrackInfo
	EgressInfo   *livekit.EgressInfo
	InEgressInfo *livekit.IngressInfo
}

type ServiceWebhookLog struct {
	ID         uint64                 `json:"-"`
	UUID       string                 `json:"uuid"`
	Name       string                 `json:"name"`
	Sid        string                 `json:"sid"`
	Url        *WebhookURLConfig      `json:"url"`
	Request    *schema.WebhookRequest `json:"request"`
	StatusCode uint16                 `json:"status_code"`
	Times      uint8                  `json:"times"`
	Resp       *schema.WebhookResp    `json:"resp"`
}

type WebhookURLConfig struct {
	API string
	Key string
}

func (swl *ServiceWebhookLog) IsSuccess() bool {
	return swl.StatusCode == http.StatusOK
}

type RoomWebhookRepo interface {
	Save(ctx context.Context, room *RoomWebhook) error
}

//go:generate mockgen -source room_webhook.go -destination ../mock/biz/room_webhook_mock.go -package=biz
type LiveGPTWebhook interface {
	Save(ctx context.Context, swlog *ServiceWebhookLog) error

	// webhook event func
	EventRoomStarted(ctx context.Context, roomName, roomSid string) error
	EventParticipantJoined(ctx context.Context, roomName, roomSid string) error
	EventParticipantLeft(ctx context.Context, roomName, roomSid string) error
	EventRoomEgressEnded(ctx context.Context, roomName, roomSid string) error
	EventRoomFinished(ctx context.Context, roomName, roomSid string) error
}

type RoomWebhookUseCase struct {
	log            *log.Helper
	confLiveKit    *conf.LiveKit
	confStorage    *conf.Storage
	linkUC         *LinkUseCase
	roomVodUC      *RoomVodUseCase
	webhookRepo    RoomWebhookRepo
	roomRepo       RoomRepo
	liveGPTWebhook LiveGPTWebhook
}

func NewRoomWebhookUseCase(
	logger log.Logger,
	confLiveKit *conf.LiveKit,
	confStorage *conf.Storage,
	linkUC *LinkUseCase,
	roomVodUC *RoomVodUseCase,
	webhookRepo RoomWebhookRepo,
	roomRepo RoomRepo,
	liveGPTWebhook LiveGPTWebhook,
) *RoomWebhookUseCase {
	return &RoomWebhookUseCase{
		log:            log.NewHelper(logger),
		confLiveKit:    confLiveKit,
		confStorage:    confStorage,
		linkUC:         linkUC,
		roomVodUC:      roomVodUC,
		webhookRepo:    webhookRepo,
		roomRepo:       roomRepo,
		liveGPTWebhook: liveGPTWebhook,
	}
}

func (ru *RoomWebhookUseCase) Save(ctx context.Context, event *livekit.WebhookEvent) error {
	backgroundCtx := helper.NewWithParentReqID(ctx)
	go func() {
		defer helper.RecoverPanic(backgroundCtx, ru.log, "RoomWebhookUseCase.Save, Panic, sid:%s", event.Room.GetSid())

		var (
			name string
			sid  string
		)
		if event.Room != nil {
			name = event.Room.GetName()
			sid = event.Room.GetSid()
		}
		if event.EgressInfo != nil {
			name = event.EgressInfo.GetRoomName()
			sid = event.EgressInfo.GetRoomId()
		}
		if err := ru.webhookRepo.Save(backgroundCtx, &RoomWebhook{
			UUID:         event.GetId(),
			Name:         name,
			Sid:          sid,
			Event:        event.GetEvent(),
			EventTime:    time.Unix(event.GetCreatedAt(), 0),
			Room:         event.GetRoom(),
			Participant:  event.GetParticipant(),
			Track:        event.GetTrack(),
			EgressInfo:   event.GetEgressInfo(),
			InEgressInfo: event.GetIngressInfo(),
		}); err != nil {
			ru.log.WithContext(ctx).Errorf("webhook save err:%v", err)
		}

		ru.log.WithContext(ctx).Debug("webhook save done.")
		return
	}()
	return nil
}

func (ru *RoomWebhookUseCase) EventHandler(ctx context.Context, event *livekit.WebhookEvent) error {
	ru.log.WithContext(ctx).Debugf("EventHandler, event:%+v", event)
	// init event
	_ = ru.Save(ctx, event)

	switch event.Event {
	case webhook.EventRoomStarted: // room_started
		return ru.EventRoomStarted(ctx, event)
	case webhook.EventTrackPublished: // track_published
		return ru.EventTrackPublished(ctx, event)
	case webhook.EventParticipantJoined: // participant_joined
		return ru.EventParticipantJoined(ctx, event)
	case webhook.EventParticipantLeft: // participant_left
		return ru.EventParticipantLeft(ctx, event)
	case webhook.EventEgressEnded: // egress_ended
		return ru.EventEgressEnded(ctx, event)
	case webhook.EventRoomFinished: // finish
		return ru.EventRoomFinished(ctx, event)
	default:
		ru.log.WithContext(ctx).Debugf("EventHandler, default, event.Event:%s", event.Event)
	}

	return nil
}

func (ru *RoomWebhookUseCase) EventRoomStarted(ctx context.Context, event *livekit.WebhookEvent) error {
	// logsnag
	logsnag.Event(ctx, logsnag.EventRoomWebhook_Room_Started.SetRoomName(event.Room.GetName()).SetUID(event.Room.GetSid()))

	// webhook push
	if err := ru.liveGPTWebhook.EventRoomStarted(ctx, event.Room.GetName(), event.Room.GetSid()); err != nil {
		ru.log.WithContext(ctx).Debugf("EventParticipantLeft, ru.liveGPTWebhook.RoomEndPush, err:%v", err)
	}

	return nil
}

func (ru *RoomWebhookUseCase) EventTrackPublished(ctx context.Context, event *livekit.WebhookEvent) error {
	if event.Participant.Identity != ru.confLiveKit.BotIdentity {
		ru.log.WithContext(ctx).Debugf("EventTrackPublished, event.Participant.Identity != BotIdentity")
		return nil
	}

	// logsnag
	logsnag.Event(ctx, logsnag.EventRoomWebhook_Track_Published.SetRoomName(event.Room.GetName()).SetUID(event.Room.GetSid()))

	// RoomStatusInterviewing
	affected, err := ru.roomRepo.UpdateStatus(ctx,
		event.Room.GetSid(),
		[]uint8{schema.RoomStatusReady},
		schema.RoomStatusInterviewing,
		&Room{},
	)
	if err != nil {
		ru.log.WithContext(ctx).Errorf("EventTrackPublished, update status, err:%v", err)
	}
	ru.log.WithContext(ctx).Debugf("EventTrackPublished, u.roomRepo.UpdateStatus, affected == %d", affected)

	return nil
}

func (ru *RoomWebhookUseCase) EventParticipantJoined(ctx context.Context, event *livekit.WebhookEvent) error {
	backCtx := helper.NewWithParentReqID(ctx)
	go func(ctx context.Context) {
		defer helper.RecoverPanic(ctx, ru.log, "EventParticipantJoined Panic, sid:%s", event.Room.GetSid())

		room := event.GetRoom()
		if err := ru.roomRepo.Save(ctx, &Room{
			Name: room.Name,
			Sid:  room.Sid,
		}); err != nil {
			ru.log.Errorf("EventParticipantJoined, name:%s, sid:%s, err:%v", event.Room.Name, event.Room.Sid, err)
		}
		return
	}(backCtx)

	// logsnag
	logsnag.Event(ctx, logsnag.EventRoomWebhook_Participant_Joined.SetRoomName(event.Room.GetName()).SetUID(event.Room.GetSid()))

	// webhook push
	if err := ru.liveGPTWebhook.EventParticipantJoined(ctx, event.Room.GetName(), event.Room.GetSid()); err != nil {
		ru.log.WithContext(ctx).Debugf("EventParticipantLeft, ru.liveGPTWebhook.RoomEndPush, err:%v", err)
	}

	return nil
}

func (ru *RoomWebhookUseCase) EventParticipantLeft(ctx context.Context, event *livekit.WebhookEvent) error {

	// logsnag
	logsnag.Event(ctx, logsnag.EventRoomWebhook_Participant_Left.SetRoomName(event.Room.GetName()).SetUID(event.Room.GetSid()))

	room, err := ru.roomRepo.GetBySID(ctx, event.Room.GetSid())
	if err != nil {
		ru.log.WithContext(ctx).Errorf("EventParticipantLeft, roomRepo.GetBySID, err:%v", err)
		return nil
	}
	// has left
	if room.IsParticipantLeft() {
		ru.log.WithContext(ctx).Debugf("EventParticipantLeft, room.IsParticipantLeft()")
		return nil
	}

	// RoomStatusParticipantLeft set status
	affected, err := ru.roomRepo.UpdateStatus(ctx,
		event.Room.GetSid(),
		[]uint8{schema.RoomStatusReady, schema.RoomStatusInterviewing},
		schema.RoomStatusParticipantLeft,
		&Room{},
	)
	if err != nil {
		ru.log.WithContext(ctx).Errorf("EventParticipantLeft, update status, err:%v", err)
		return nil
	}
	if affected == 0 {
		ru.log.WithContext(ctx).Debugf("EventParticipantLeft, u.roomRepo.UpdateStatus, affected == %d", affected)
		return nil
	}

	// webhook push
	if err := ru.liveGPTWebhook.EventParticipantLeft(ctx, event.Room.GetName(), event.Room.GetSid()); err != nil {
		ru.log.WithContext(ctx).Debugf("EventParticipantLeft, ru.liveGPTWebhook.RoomEndPush, err:%v", err)
	}

	return nil
}

func (ru *RoomWebhookUseCase) EventEgressEnded(ctx context.Context, event *livekit.WebhookEvent) error {
	// EventHandler,
	// event:event:"egress_ended"
	// id:"EV_EANZaEhu8TMU" created_at:1691553790
	// egress_info:{
	// 	egress_id:"EG_LrczPB2E29gk" room_id:"RM_2BKBwdsPGApB" room_name:"a7hx-fn5t" status:EGRESS_COMPLETE started_at:1691553761343377109 ended_at:1691553789616161606
	//	room_composite:{
	//		room_name:"a7hx-fn5t" layout:"grid" file_outputs:{file_type:MP4 filepath:"egress/a7hx-fn5t-1691553760.mp4" azure:{account_name:"{account_name}" account_key:"{account_key}" container_name:"facetoai"}}
	//	}
	//	file:{filename:"egress/a7hx-fn5t-1691553760.mp4" started_at:1691553765092241493 ended_at:1691553789303080530 duration:24210839037 size:11690339 location:"https://withcontext.blob.core.windows.net/facetoai/egress/a7hx-fn5t-1691553760.mp4"}
	//	file_results:{filename:"egress/a7hx-fn5t-1691553760.mp4" started_at:1691553765092241493 ended_at:1691553789303080530 duration:24210839037 size:11690339 location:"https://withcontext.blob.core.windows.net/facetoai/egress/a7hx-fn5t-1691553760.mp4"}
	// }
	if event.GetEgressInfo() == nil {
		return nil
	}
	egressInfo := event.GetEgressInfo()
	// logsnag
	logsnag.Event(ctx, logsnag.EventRoomWebhook_Egress_Ended.SetRoomName(egressInfo.GetRoomName()).SetUID(egressInfo.GetRoomId()))

	ru.log.WithContext(ctx).Debugw("sid", egressInfo.GetRoomId(), "EventEgressEnded.egressInfo", egressInfo)
	status := 0
	switch egressInfo.GetStatus() {
	case livekit.EgressStatus_EGRESS_STARTING:
		status = schema.RoomVodStatusReady
	case livekit.EgressStatus_EGRESS_ACTIVE, livekit.EgressStatus_EGRESS_ENDING:
		status = schema.RoomVodStatusStarting
	case livekit.EgressStatus_EGRESS_COMPLETE:
		status = schema.RoomVodStatusComplete
	case livekit.EgressStatus_EGRESS_FAILED, livekit.EgressStatus_EGRESS_ABORTED:
		status = schema.RoomVodStatusFail
	}
	if len(egressInfo.GetFileResults()) == 0 {
		ru.log.WithContext(ctx).Errorf("len(egressInfo.GetFileResults()) == 0")
		return nil
	}

	urlslice := []string{
		strings.TrimRight(ru.confStorage.AzureBlob.GetCdnHost(), "/"),
		strings.Trim(ru.confStorage.AzureBlob.GetContainerName(), "/"),
		strings.TrimLeft(egressInfo.GetFileResults()[0].GetFilename(), "/"),
	}
	// time value is noa sec
	startTime := egressInfo.GetStartedAt() / 1e9 // 1691553765092241493
	startTimeNsec := egressInfo.GetStartedAt() % 1e9
	completeTime := egressInfo.GetEndedAt() / 1e9 // 1691553789303080530
	completeTimeNsec := egressInfo.GetEndedAt() % 1e9
	duration := egressInfo.GetFileResults()[0].GetDuration() / 1e9 // 24210839037
	if _, err := ru.roomVodUC.Create(ctx, &RoomVod{
		Name:         egressInfo.GetRoomName(),
		Sid:          egressInfo.GetRoomId(),
		EgressID:     egressInfo.GetEgressId(),
		Status:       uint8(status),
		Platform:     schema.VodPlatFormAzure,
		VodType:      schema.VodTypeFile,
		VodPath:      egressInfo.GetFileResults()[0].GetFilename(),
		VodURL:       strings.Join(urlslice, "/"),
		StartTime:    time.Unix(startTime, startTimeNsec),
		CompleteTime: time.Unix(completeTime, completeTimeNsec),
		Duration:     uint64(duration),
	}); err != nil {
		ru.log.WithContext(ctx).Errorf("EventEgressEnded, save.err:%v", err)
		return nil
	}

	// webhook push
	if err := ru.liveGPTWebhook.EventRoomEgressEnded(ctx, egressInfo.GetRoomName(), egressInfo.GetRoomId()); err != nil {
		ru.log.WithContext(ctx).Debugf("EventParticipantLeft, ru.liveGPTWebhook.RoomEndPush, err:%v", err)
	}

	return nil
}

func (ru *RoomWebhookUseCase) EventRoomFinished(ctx context.Context, event *livekit.WebhookEvent) error {

	// logsnag
	logsnag.Event(ctx, logsnag.EventRoomWebhook_Finish.SetRoomName(event.Room.GetName()).SetUID(event.Room.GetSid()))

	room, err := ru.roomRepo.GetBySID(ctx, event.Room.GetSid())
	if err != nil {
		ru.log.WithContext(ctx).Errorf("EventRoomFinished, roomRepo.GetBySID, err:%v", err)
		return nil
	}
	// has complete
	if room.IsComplete() {
		ru.log.WithContext(ctx).Debugf("EventRoomFinished, room.IsComplete()")
		return nil
	}

	// RoomStatusParticipantLeft set status
	affected, err := ru.roomRepo.UpdateStatus(ctx,
		event.Room.GetSid(),
		[]uint8{schema.RoomStatusInterviewing, schema.RoomStatusParticipantLeft},
		schema.RoomStatusEnd,
		&Room{},
	)
	if err != nil {
		ru.log.WithContext(ctx).Errorf("EventRoomFinished, update status, err:%v", err)
		return nil
	}
	if affected == 0 {
		ru.log.WithContext(ctx).Debugf("EventRoomFinished, u.roomRepo.UpdateStatus, affected == %d", affected)
		return nil
	}

	// webhook push
	if err := ru.liveGPTWebhook.EventRoomFinished(ctx, event.Room.GetName(), event.Room.GetSid()); err != nil {
		ru.log.WithContext(ctx).Debugf("EventParticipantLeft, ru.liveGPTWebhook.RoomEndPush, err:%v", err)
	}

	return nil
}
