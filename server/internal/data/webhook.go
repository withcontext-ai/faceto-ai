package data

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/data/schema"
	"faceto-ai/internal/pkg/utils/helper"
	"faceto-ai/internal/pkg/utils/logsnag"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	"time"
)

const version = "1.0"

type WebhookRequest struct {
	ID           string         `json:"id"`
	Source       *WebhookSource `json:"source"`
	Target       *WebhookTarget `json:"target"`
	Object       string         `json:"object"`
	Type         string         `json:"type"`
	Data         *WebhookData   `json:"data"`
	NeedCallback bool           `json:"needCallback"`
	CreatedAt    uint64         `json:"created_at"`
}

func (w *WebhookRequest) GetData() *WebhookData {
	if w.Data != nil {
		return w.Data
	}
	return nil
}

type WebhookReply struct {
}

type WebhookSource struct {
	Platform string `json:"platform"`
	Other    string `json:"other"`
}

type WebhookTarget struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Other   string `json:"other"`
}

type WebhookData struct {
	Room       *Room       `json:"room"`
	Transcript *Transcript `json:"transcript,omitempty"`
	Vod        *Vod        `json:"vod,omitempty"`
}

func (w *WebhookData) GetRoom() *Room {
	if w.Room != nil {
		return w.Room
	}
	return &Room{}
}

func (w *WebhookData) GetTranscript() *Transcript {
	if w.Transcript != nil {
		return w.Transcript
	}
	return &Transcript{}
}

func (w *WebhookData) GetVod() *Vod {
	if w.Vod != nil {
		return w.Vod
	}
	return &Vod{}
}

type Room struct {
	Name string `json:"name"`
	Sid  string `json:"sid"`
}

type Transcript struct {
	Total int32             `json:"total"`
	List  []*TranscriptItem `json:"list"`
}

type TranscriptItem struct {
	IsBot     bool   `json:"is_bot"`
	Name      string `json:"name"`
	Text      string `json:"text"`
	Timestamp uint64 `json:"timestamp"`
}

type Vod struct {
	EgressID     string `json:"egress_id,omitempty"`
	URL          string `json:"url,omitempty"`
	Status       uint8  `json:"status,omitempty"`
	StartTime    uint64 `json:"start_time,omitempty"`
	CompleteTime uint64 `json:"complete_time,omitempty"`
	Duration     uint64 `json:"duration"`
}

type WebhookQueue struct {
	Ctx  context.Context
	Api  *biz.WebhookURLConfig
	Data *WebhookRequest
}

// WebhookService Handle webhook business
type liveGPTWebhook struct {
	data        *Data
	ctx         context.Context
	cancel      context.CancelFunc
	log         *log.Helper
	linkRepo    biz.LinkRepo
	roomRepo    biz.RoomRepo
	roomMsgRepo biz.RoomMessageRepo
	roomVodResp biz.RoomVodRepo
	authRepo    biz.AuthRepo
	queue       chan *WebhookQueue
}

func NewLiveGPTWebhook(
	data *Data,
	logger log.Logger,
	linkRepo biz.LinkRepo,
	roomRepo biz.RoomRepo,
	roomMsgRepo biz.RoomMessageRepo,
	roomVodResp biz.RoomVodRepo,
	authRepo biz.AuthRepo,
) biz.LiveGPTWebhook {
	ctx, cancel := context.WithCancel(context.Background())
	w := &liveGPTWebhook{
		ctx:         ctx,
		cancel:      cancel,
		data:        data,
		log:         log.NewHelper(logger),
		linkRepo:    linkRepo,
		roomRepo:    roomRepo,
		roomMsgRepo: roomMsgRepo,
		roomVodResp: roomVodResp,
		authRepo:    authRepo,
		queue:       make(chan *WebhookQueue),
	}
	w.Consumer() // start consumer
	return w
}

func (w *liveGPTWebhook) Consumer() {
	for i := 0; i < 5; i++ {
		go func() {
			for event := range w.queue {
				if event.Data == nil {
					continue
				}

				api := event.Api
				if api == nil {
					continue
				}
				reqHeader := make(map[string]string)
				if api.Key != "" {
					reqHeader["Authorization"] = "Bearer " + api.Key
				}

				eventRoom := &Room{}
				eventData := event.Data.GetData()
				if eventData != nil {
					eventRoom = eventData.GetRoom()
				}

				saveToDb := func(resp *resty.Response) {
					var (
						statusCode uint16
						respData   string
					)
					if resp == nil {
						statusCode = 500
						respData = ""
					} else {
						statusCode = uint16(resp.StatusCode())
						respData = resp.String()
					}
					_ = w.Save(w.ctx, &biz.ServiceWebhookLog{
						UUID: event.Data.ID,
						Name: eventRoom.Name,
						Sid:  eventRoom.Sid,
						Url: &biz.WebhookURLConfig{
							API: api.API,
							Key: api.Key,
						},
						Request:    &schema.WebhookRequest{Data: event.Data},
						StatusCode: statusCode,
						Times:      0,
						Resp:       &schema.WebhookResp{Data: respData},
					})
				}

				// send request
				resp, err := helper.RestyRequest(w.ctx, &helper.RestyOptions{
					Url:              api.API,
					Req:              event.Data,
					Response:         &WebhookReply{},
					IsRetry:          true,
					Headers:          reqHeader,
					Timeout:          time.Duration(3) * time.Second,
					RetryCount:       3, // retry 3 times
					RetryWaitTime:    time.Duration(200) * time.Millisecond,
					RetryMaxWaitTime: time.Duration(3) * time.Second,
				})
				if err != nil {
					w.log.Errorf("Webhook.Consumer helper.RestyRequest, err:%v", err)
					go saveToDb(resp)
					return
				}

				go saveToDb(resp)
				w.log.Debugf("Webhook.Consumer helper.RestyRequest, done. api:%s", api.API)
			}
		}()
	}
}

func (w *liveGPTWebhook) GetWebhookConfig(ctx context.Context, name string) (*schema.Webhook, error) {
	linkConfig, err := w.linkRepo.GetLinkByName(ctx, name)
	if err != nil {
		w.log.WithContext(ctx).Errorf("w.linkRepo.GetLinkByName, name:%s, err:%v", name, err)
		return nil, errors.Wrap(err, "get link config err")
	}

	return linkConfig.GetWebHookConfig(), nil
}

func (w *liveGPTWebhook) Save(ctx context.Context, swlog *biz.ServiceWebhookLog) error {
	result, err := w.data.DB(ctx).ServiceWebhookLog.Create().
		SetUUID(swlog.UUID).
		SetName(swlog.Name).SetSid(swlog.Sid).
		SetURL(swlog.Url.API).
		SetRequest(swlog.Request).
		SetStatusCode(swlog.StatusCode).
		SetTimes(swlog.Times).
		SetResp(swlog.Resp).
		Save(ctx)
	if err != nil {
		return err
	}

	swlog.UUID = result.UUID
	return nil
}

func (w *liveGPTWebhook) EventRoomStarted(ctx context.Context, roomName, roomSid string) error {
	go func() {
		defer helper.RecoverPanic(w.ctx, w.log, "Webhook.EventParticipantLeft roomSid:%s", roomSid)

		webhookConfig, err := w.GetWebhookConfig(w.ctx, roomName)
		if err != nil {
			return
		}
		if webhookConfig == nil {
			w.log.Debugf("EventRoomStarted.webhookConfig == nil")
			return
		}

		webhookEvent := NewWebhookEvent(roomName, roomSid)
		webhookData := &WebhookData{
			Room: &Room{
				Name: roomName,
				Sid:  roomSid,
			},
		}

		w.queue <- &WebhookQueue{
			Ctx: w.ctx,
			Api: &biz.WebhookURLConfig{
				API: webhookConfig.GetAPI(),
				Key: webhookConfig.GetKey(),
			},
			Data: webhookEvent.CreateEventRoomStarted(webhookData),
		}

		// logsnag
		logsnag.Event(ctx, logsnag.EventRoomWebhook_Push_RoomStarted.SetRoomName(roomName).SetUID(roomSid))
	}()

	return nil
}

func (w *liveGPTWebhook) EventParticipantJoined(ctx context.Context, roomName, roomSid string) error {

	return nil
}

func (w *liveGPTWebhook) EventParticipantLeft(ctx context.Context, roomName, roomSid string) error {
	go func() {
		defer helper.RecoverPanic(w.ctx, w.log, "Webhook.EventParticipantLeft roomSid:%s", roomSid)

		webhookConfig, err := w.GetWebhookConfig(w.ctx, roomName)
		if err != nil {
			return
		}
		if webhookConfig == nil {
			w.log.Debugf("EventParticipantLeft.webhookConfig == nil")
			return
		}

		linkConfig, err := w.linkRepo.GetLinkByName(w.ctx, roomName)
		if err != nil {
			w.log.Errorf("EventParticipantLeft, w.linkRepo.GetLinkByName err:%v", err)
			return
		}
		authGrant, err := w.authRepo.CreateAuthGrantByUUID(w.ctx, linkConfig.Token)
		if err != nil {
			w.log.Errorf("EventParticipantLeft, w.authRepo.CreateAuthGrantByUUID err:%v", err)
			return
		}

		webhookData := &WebhookData{
			Room: &Room{
				Name: roomName,
				Sid:  roomSid,
			},
		}

		if authGrant.CanGetVideo || authGrant.CanGetAudio {
			// get room info
			room, err := w.roomRepo.GetBySID(w.ctx, roomSid)
			if err != nil {
				w.log.Errorf("EventParticipantLeft, w.roomRepo.GetBySID err:%v", err)
				return
			}
			webhookData.Vod = &Vod{
				StartTime:    uint64(room.StartTime.Unix()),
				CompleteTime: uint64(room.LeftTime.Unix()),
				Duration:     uint64(room.LeftTime.Sub(room.StartTime).Seconds()),
			}
		}

		webhookEvent := NewWebhookEvent(roomName, roomSid)
		w.queue <- &WebhookQueue{
			Ctx: w.ctx,
			Api: &biz.WebhookURLConfig{
				API: webhookConfig.GetAPI(),
				Key: webhookConfig.GetKey(),
			},
			Data: webhookEvent.CreateEventParticipantLeft(webhookData),
		}

		// logsnag
		logsnag.Event(ctx, logsnag.EventRoomWebhook_Push_ParticipantLeft.SetRoomName(roomName).SetUID(roomSid))
	}()

	return nil
}

func (w *liveGPTWebhook) EventRoomEgressEnded(ctx context.Context, roomName, roomSid string) error {

	return nil
}

func (w *liveGPTWebhook) EventRoomFinished(ctx context.Context, roomName, roomSid string) error {

	return nil
}

// const room event type
const (
	EventTypeRoomStarted   = "Event.RoomStarted"       // room start
	EventParticipantJoined = "Event.ParticipantJoined" // room participant joined
	EventParticipantLeft   = "Event.ParticipantLeft"   // room participant left
	EventTypeRoomEgressEnd = "Event.RoomEgressEnd"     // video created end
	EventRoomFinished      = "Event.RoomFinished"      // room finished
)

// WebhookEvent Webhook Event Define
type WebhookEvent struct {
	roomName string
	roomSid  string
}

func NewWebhookEvent(name, sid string) *WebhookEvent {
	return &WebhookEvent{roomName: name, roomSid: sid}
}

func (we *WebhookEvent) GetTargetName() string {
	if we.roomName == "" || we.roomSid == "" {
		return "biz.default"
	}
	return we.roomName + "_" + we.roomSid
}

func (we *WebhookEvent) CreateEventRoomStarted(data *WebhookData) *WebhookRequest {
	return &WebhookRequest{
		ID: schema.MustUID(),
		Source: &WebhookSource{
			Platform: "FaceToAI",
			Other:    "",
		},
		Target: &WebhookTarget{
			Name:    we.GetTargetName(),
			Version: version,
			Other:   "",
		},
		Object:       "Event",
		Type:         EventTypeRoomStarted,
		Data:         data,
		NeedCallback: false,
		CreatedAt:    uint64(time.Now().Unix()),
	}
}

func (we *WebhookEvent) CreateEventParticipantJoined(data *WebhookData) *WebhookRequest {
	return &WebhookRequest{
		ID: schema.MustUID(),
		Source: &WebhookSource{
			Platform: "FaceToAI",
			Other:    "",
		},
		Target: &WebhookTarget{
			Name:    we.GetTargetName(),
			Version: version,
			Other:   "",
		},
		Object:       "Event",
		Type:         EventParticipantJoined,
		Data:         data,
		NeedCallback: false,
		CreatedAt:    uint64(time.Now().Unix()),
	}
}

func (we *WebhookEvent) CreateEventParticipantLeft(data *WebhookData) *WebhookRequest {
	return &WebhookRequest{
		ID: schema.MustUID(),
		Source: &WebhookSource{
			Platform: "FaceToAI",
			Other:    "",
		},
		Target: &WebhookTarget{
			Name:    we.GetTargetName(),
			Version: version,
			Other:   "",
		},
		Object:       "Event",
		Type:         EventParticipantLeft,
		Data:         data,
		NeedCallback: false,
		CreatedAt:    uint64(time.Now().Unix()),
	}
}

func (we *WebhookEvent) CreateEventRoomFinished(data *WebhookData) *WebhookRequest {
	return &WebhookRequest{
		ID: schema.MustUID(),
		Source: &WebhookSource{
			Platform: "FaceToAI",
			Other:    "",
		},
		Target: &WebhookTarget{
			Name:    we.GetTargetName(),
			Version: version,
			Other:   "",
		},
		Object:       "Event",
		Type:         EventRoomFinished,
		Data:         data,
		NeedCallback: false,
		CreatedAt:    uint64(time.Now().Unix()),
	}
}

func (we *WebhookEvent) CreateEventRoomEgressEnd(data *WebhookData) *WebhookRequest {
	return &WebhookRequest{
		ID: schema.MustUID(),
		Source: &WebhookSource{
			Platform: "FaceToAI",
			Other:    "",
		},
		Target: &WebhookTarget{
			Name:    we.GetTargetName(),
			Version: version,
			Other:   "",
		},
		Object:       "Event",
		Type:         EventTypeRoomEgressEnd,
		Data:         data,
		NeedCallback: false,
		CreatedAt:    uint64(time.Now().Unix()),
	}
}
