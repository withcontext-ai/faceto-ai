package service

import (
	stt "cloud.google.com/go/speech/apiv1"
	tts "cloud.google.com/go/texttospeech/apiv1"
	"context"
	"encoding/base64"
	"encoding/json"
	"faceto-ai/internal/biz/liveGPT"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/data/schema"
	"faceto-ai/internal/pkg/middleware"
	"faceto-ai/internal/pkg/utils/helper"
	pkgLog "faceto-ai/internal/pkg/utils/log"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/livekit/protocol/auth"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/webhook"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/pkg/errors"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/option"
	"net/http"
	"strconv"
	"sync"
	"time"

	errorV1 "faceto-ai/api_gen/error/v1"
	v1 "faceto-ai/api_gen/room/v1"
	"faceto-ai/internal/biz"
)

type ActiveParticipant struct {
	Connecting  bool
	Participant *liveGPT.GPTParticipant
}

type RoomCache struct {
	Room       *livekit.Room          `json:"room"`
	Transcript []*liveGPT.SpeechEvent `json:"transcript"`
	Vod        *RoomVod               `json:"vod"`
}

type RoomVod struct {
	Url          string `json:"url"`
	Status       int32  `json:"status"`
	CompleteTime int64  `json:"complete_time"`
}

type CheckRoomConfig struct {
	UserName string `json:"username,omitempty"`
	BotName  string `json:"botname,omitempty"`
}

// RoomService is a room service.
type RoomService struct {
	v1.UnimplementedRoomServer

	log          *log.Helper
	confThirdApi *conf.ThirdApi
	confStorage  *conf.Storage
	confLiveKit  *conf.LiveKit
	confGcp      *conf.GcpCredentials

	linkUC    *biz.LinkUseCase
	roomUC    *biz.RoomUseCase
	roomMsgUC *biz.RoomMessageUseCase
	roomVodUC *biz.RoomVodUseCase
	webhookUC *biz.RoomWebhookUseCase
	authUC    *biz.AuthUseCase
	voiceUC   *biz.VoiceUseCase

	roomService  *lksdk.RoomServiceClient
	egressClient *lksdk.EgressClient
	egressinfo   map[string]*livekit.EgressInfo
	keyProvider  *auth.SimpleKeyProvider
	gptClient    *openai.Client
	sttClient    *stt.Client
	ttsClient    *tts.Client
	participants map[string]*ActiveParticipant
	cacheRoom    map[string]*RoomCache

	lock sync.Mutex
}

// NewRoomService new a room service.
func NewRoomService(
	logger log.Logger,

	confThirdApi *conf.ThirdApi,
	confStorage *conf.Storage,
	confLiveKit *conf.LiveKit,
	confGcp *conf.GcpCredentials,

	linkUC *biz.LinkUseCase,
	webhookUC *biz.RoomWebhookUseCase,
	roomUC *biz.RoomUseCase,
	roomMsgUC *biz.RoomMessageUseCase,
	roomVodUC *biz.RoomVodUseCase,
	authUC *biz.AuthUseCase,
	voiceUC *biz.VoiceUseCase,
) *RoomService {
	ctx := context.Background()

	// init sst tts client
	//gcpCred := option.WithCredentialsFile(confGcp.Path)
	gcpBytes, err := conf.GetGcpCredBytes(confGcp.GetPath())
	if err != nil {
		panic(err)
	}
	gcpCred := option.WithCredentialsJSON(gcpBytes)

	// stt client
	sttClient, err := stt.NewClient(ctx, gcpCred)
	if err != nil {
		panic(err)
	}
	// tts client
	ttsClient, err := tts.NewClient(ctx, gcpCred)
	if err != nil {
		panic(err)
	}

	// init gpt config
	gptConfig := openai.DefaultConfig(confThirdApi.Openai.Key)
	gptConfig.BaseURL = confThirdApi.Openai.Host

	return &RoomService{
		log: log.NewHelper(logger),

		// init usecase
		linkUC:    linkUC,
		webhookUC: webhookUC,
		roomUC:    roomUC,
		roomMsgUC: roomMsgUC,
		roomVodUC: roomVodUC,
		authUC:    authUC,
		voiceUC:   voiceUC,

		// init config
		confLiveKit:  confLiveKit,
		confStorage:  confStorage,
		confGcp:      confGcp,
		confThirdApi: confThirdApi,

		// init service
		roomService:  lksdk.NewRoomServiceClient(confLiveKit.Url, confLiveKit.ApiKey, confLiveKit.SecretKey),
		egressClient: lksdk.NewEgressClient(confLiveKit.Url, confLiveKit.ApiKey, confLiveKit.SecretKey),
		keyProvider:  auth.NewSimpleKeyProvider(confLiveKit.ApiKey, confLiveKit.SecretKey),
		gptClient:    openai.NewClientWithConfig(gptConfig),
		participants: make(map[string]*ActiveParticipant),
		sttClient:    sttClient,
		ttsClient:    ttsClient,

		egressinfo: make(map[string]*livekit.EgressInfo, 1),
		cacheRoom:  make(map[string]*RoomCache, 0),
	}
}

func (s *RoomService) Health(ctx context.Context, in *v1.HealthRequest) (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "hello world!"}, nil
}

func (s *RoomService) CheckRoom(ctx context.Context, in *v1.CheckRoomRequest) (*v1.CheckRoomReply, error) {
	link, err := s.linkUC.GetLinkByName(ctx, in.GetName())
	if err != nil {
		return &v1.CheckRoomReply{
			Valid: false,
		}, nil
	}

	// set room link config
	if in.GetC() != "" {
		go func() {
			bgctx := helper.NewWithParentReqID(ctx)
			// base64 decode
			// {"username":"", "botname":""}
			decoded, err := base64.StdEncoding.DecodeString(in.GetC())
			if err != nil {
				s.log.WithContext(bgctx).Errorf("base64.StdEncoding.DecodeString, config:%s, err:%v", in.GetC(), err)
			} else {
				roomConfig := &CheckRoomConfig{}
				if err := json.Unmarshal(decoded, roomConfig); err == nil {
					if link.Config == nil {
						link.Config = &schema.RoomConfig{}
					}
					if roomConfig.UserName != "" || roomConfig.BotName != "" {
						link.Config.UserName = roomConfig.UserName
						link.Config.BotName = roomConfig.BotName
						if err := s.linkUC.SetConfigByUUID(bgctx, link.UUID, link.Config); err != nil {
							s.log.WithContext(bgctx).Errorf("s.linkUC.SetConfigByUUID, err:%v", err)
						}
					}
				} else {
					s.log.WithContext(bgctx).Errorf("json.Unmarshal, config:%s, err:%v", in.GetC(), err)
				}
			}
		}()
	}

	room, err := s.roomUC.GetByName(ctx, in.GetName())
	if err != nil {
		if errors.Is(err, biz.ErrRoomNotFound) {
			return &v1.CheckRoomReply{
				Valid: true,
			}, nil
		}
		return &v1.CheckRoomReply{
			Valid: false,
		}, nil
	}

	return &v1.CheckRoomReply{
		Valid: room.AccessOrNot(),
	}, nil
}

func (s *RoomService) JoinRoom(ctx context.Context, in *v1.JoinRoomRequest) (*v1.JoinRoomReply, error) {
	listRes, err := s.roomService.ListRooms(ctx, &livekit.ListRoomsRequest{
		Names: []string{
			in.Name,
		},
	})
	if err != nil {
		s.log.WithContext(ctx).Errorf("error listing rooms, err:%v", err)
		return nil, errorV1.ErrorBadRequest("list room err:%v", err)
	}

	if len(listRes.Rooms) == 0 {
		return nil, errorV1.ErrorNotFound("room not found")
	}

	room := listRes.Rooms[0]

	// room set db
	if err := s.roomUC.JoinRoom(ctx, room); err != nil {
		s.log.WithContext(ctx).Errorf("error biz join room, name:%s, sid:%s, err:%v", room.Name, room.Sid, err)
	}

	s.log.WithContext(ctx).Debugf("JoinRoom, name:%s, sid:%s", room.Name, room.Sid)
	if err := s.joinRoomHandler(ctx, room); err != nil {
		return nil, errorV1.ErrorInternalServerError("join room err:%v", err)
	}

	return &v1.JoinRoomReply{
		Name: room.Name,
		Sid:  room.Sid,
	}, nil
}

func (s *RoomService) joinRoomHandler(ctx context.Context, room *livekit.Room) error {
	// get room api config
	link, err := s.linkUC.GetLinkByName(ctx, room.GetName())
	if err != nil {
		s.log.WithContext(ctx).Errorf("error get link config by name, err:%v", err)
		return errors.Wrap(err, "error get link config by name")
	}

	// If the GPT participant is not connected, connect it
	s.lock.Lock()
	if _, ok := s.participants[room.Sid]; ok {
		s.lock.Unlock()
		s.log.WithContext(ctx).Debugf("gpt participant already connected, room:%s, participantCount:%d", room.Name, room.NumParticipants)
		return nil
	}

	s.participants[room.Sid] = &ActiveParticipant{
		Connecting: true,
	}
	s.lock.Unlock()

	botIdentity := link.GetConfigBotName()
	if botIdentity == "" {
		botIdentity = s.confLiveKit.BotIdentity
	}
	s.log.WithContext(ctx).Debugf("joinRoomHandler.botIdentity,%s, room.GetName():%s, room.GetSid():%s", botIdentity, room.GetName(), room.GetSid())

	token := s.roomService.CreateToken().
		SetIdentity(botIdentity).
		AddGrant(&auth.VideoGrant{
			Room:       room.Name,
			RoomJoin:   true,
			RoomRecord: true,
		})

	jwt, err := token.ToJWT()
	if err != nil {
		s.log.WithContext(ctx).Errorf("error creating jwt, err:%v", err)
		return errors.Wrap(err, "error creating jwt")
	}

	// start egress service
	s.startRoomEgress(ctx, room)

	s.log.WithContext(ctx).Debugf("connecting gpt participant, name:%s, sid:%s", room.Name, room.Sid)
	// Connect Participant Client
	p, err := liveGPT.ConnectGPTParticipant(
		s.confLiveKit.Url,
		jwt,
		s.sttClient, s.ttsClient, s.gptClient,
		s.log, s.confThirdApi, link, s.linkUC, s.roomUC, s.roomMsgUC,
	)
	if err != nil {
		s.log.WithContext(ctx).Errorf("error connecting gpt participant, room:%s, sid:%s", room.Name, room.Sid)
		s.lock.Lock()
		delete(s.participants, room.Sid)
		s.lock.Unlock()
		return errors.Wrap(err, "error connecting gpt participant")
	}

	s.lock.Lock()
	s.participants[room.Sid] = &ActiveParticipant{
		Connecting:  false,
		Participant: p,
	}
	s.lock.Unlock()

	p.OnDisconnected(func() {
		s.log.WithContext(ctx).Infof("gpt participant disconnected, room:%s, sid:%s", room.Name, room.Sid)

		// stop egress
		s.stopRoomEgress(ctx, room)

		s.lock.Lock()
		delete(s.participants, room.Sid)
		s.lock.Unlock()
	})

	return nil
}

func (s *RoomService) startRoomEgress(ctx context.Context, room *livekit.Room) {
	fileRequest := &livekit.RoomCompositeEgressRequest{
		RoomName: room.GetName(),
		Layout:   "grid",
		FileOutputs: []*livekit.EncodedFileOutput{
			{
				FileType: livekit.EncodedFileType_MP4,
				Filepath: "egress/" + room.GetName() + "-" + strconv.Itoa(int(time.Now().Unix())) + ".mp4",
				Output: &livekit.EncodedFileOutput_Azure{
					Azure: &livekit.AzureBlobUpload{
						AccountName:   s.confStorage.AzureBlob.GetAccountName(),
						AccountKey:    s.confStorage.AzureBlob.GetAccountKey(),
						ContainerName: s.confStorage.AzureBlob.GetContainerName(),
					},
				},
			},
		},
	}
	egressInfo, err := s.egressClient.StartRoomCompositeEgress(ctx, fileRequest)
	if err != nil {
		s.log.Errorf("s.egressClient.StartRoomCompositeEgress err:%v", err)
		return
	}
	// insert into mysql
	if _, err := s.roomVodUC.CreateWithEgressInfoAzure(ctx, egressInfo); err != nil {
		s.log.Errorf("s.roomVodUC.CreateWithEgressInfo err:%v", err)
	}
	s.egressinfo[room.GetSid()] = egressInfo
	s.log.Debugw("startRoomEgress.egressInfo", egressInfo)
	return
}

func (s *RoomService) stopRoomEgress(ctx context.Context, room *livekit.Room) {
	if egress, ok := s.egressinfo[room.GetSid()]; ok {
		s.log.WithContext(ctx).Infof("s.stopRoomEgress, egressId:%s", egress.GetEgressId())
		_, err := s.egressClient.StopEgress(context.Background(), &livekit.StopEgressRequest{EgressId: egress.GetEgressId()})
		if err != nil {
			s.log.Errorf("s.egressClient.StopEgress, err:%v", err)
		}
	}
}

// WebHookHandler https://ai-interview-mmxbwgwwaq-uw.a.run.app/webhook
func (s *RoomService) WebHookHandler(ctx context.Context, req *http.Request) (*v1.JoinRoomReply, error) {
	s.log.WithContext(ctx).Debugf("WebHookHandler, req:%+v", req)

	// get event struct
	event, err := webhook.ReceiveWebhookEvent(req, s.keyProvider)
	if err != nil {
		s.log.WithContext(ctx).Errorf("error receiving webhook event, err:%v", err)
		return nil, err
	}
	s.log.WithContext(ctx).Debugf("WebHookHandler, event:%+v", event)

	// test
	//event := &livekit.WebhookEvent{
	//	Id:        "EV_EANZaEhu8TMU",
	//	Event:     "egress_ended",
	//	CreatedAt: 1691553790,
	//	EgressInfo: &livekit.EgressInfo{
	//		EgressId:  "EG_LrczPB2E29gk",
	//		RoomId:    "RM_2BKBwdsPGApB",
	//		RoomName:  "a7hx-fn5t",
	//		Status:    3,
	//		StartedAt: 1691553761343377109,
	//		EndedAt:   1691553789616161606,
	//		FileResults: []*livekit.FileInfo{
	//			{
	//				Filename:  "egress/a7hx-fn5t-1691553760.mp4",
	//				StartedAt: 1691553761343377109,
	//				EndedAt:   1691553789616161606,
	//				Duration:  24210839037,
	//				Size:      11690339,
	//			},
	//		},
	//	},
	//}

	// event handler
	if err := s.webhookUC.EventHandler(ctx, event); err != nil {
		s.log.WithContext(ctx).Errorf("WebHookHandler, EventHandler, err:%v", err)
		return nil, err
	}

	// join AI client
	if event.Event == webhook.EventParticipantJoined {
		if event.Participant.Identity == s.confLiveKit.BotIdentity {
			s.log.WithContext(ctx).Debugf("event.Participant.Identity == BotIdentity")
			return nil, nil
		}

		// join room
		s.log.WithContext(ctx).Debugf("JoinRoom, name:%s, sid:%s", event.Room.Name, event.Room.Sid)
		if err := s.joinRoomHandler(ctx, event.GetRoom()); err != nil {
			s.log.WithContext(ctx).Errorf("error join room, name:%s, sid:%s, err:%v", event.Room.Name, event.Room.Sid, err)
			return nil, errors.Wrap(err, "join room handler err")
		}
	}

	return &v1.JoinRoomReply{
		Name: event.Room.GetName(),
		Sid:  event.Room.GetSid(),
	}, nil
}

func getJoinToken(apiKey, apiSecret, room, identity string) (string, error) {
	at := auth.NewAccessToken(apiKey, apiSecret)
	grant := &auth.VideoGrant{
		RoomCreate: true,
		RoomList:   true,
		RoomRecord: true,

		RoomAdmin: true,
		RoomJoin:  true,
		Room:      room,
	}
	at.AddGrant(grant).SetIdentity(identity).SetValidFor(time.Hour)

	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)

	jwt, _ := at.ToJWT()
	log.NewHelper(logger).Debugf("at.ToJWT:%s", jwt)
	return jwt, nil
}

func (s *RoomService) RoomTranscriptOnline(ctx context.Context, in *v1.RoomTranscriptRequest) (*v1.RoomTranscriptReply, error) {
	s.log.WithContext(ctx).Debugf("RoomTranscriptOnline, name:%s, page:%d", in.GetName(), in.GetPage())

	room, err := s.roomUC.GetByName(ctx, in.GetName())
	if err != nil {
		s.log.WithContext(ctx).Errorf("s.roomUC.GetByName failed, err:%v", err)
		return nil, errorV1.ErrorBadRequest("")
	}

	total, transcript, err := s.roomMsgUC.Transcript(ctx, room.Sid, in.GetPage())
	if err != nil {
		s.log.WithContext(ctx).Errorf("s.roomMsgUC.Transcript failed, err:%v", err)
		return nil, errorV1.ErrorBadRequest("")
	}

	// resp
	resp := &v1.RoomTranscriptReply{
		Version: "1.0",
		Room: &v1.RoomTranscriptReply_Room{
			Name: room.Name,
			Sid:  room.Sid,
		},
		Transcript: &v1.RoomTranscriptReply_Transcript{
			Total: uint32(total),
			List:  []*v1.RoomTranscriptReply_Transcript_List{},
		},
		Vod: &v1.RoomTranscriptReply_Vod{},
	}

	if len(transcript) == 0 {
		return resp, nil
	}

	for _, ev := range transcript {
		resp.Transcript.List = append(resp.Transcript.List, &v1.RoomTranscriptReply_Transcript_List{
			IsBot:     ev.Type == biz.RoomMessageTypeBot,
			Name:      ev.ParticipantName,
			Text:      ev.Text,
			Timestamp: uint64(ev.EventTime.Unix()),
		})
	}

	return resp, nil
}

func (s *RoomService) RoomTranscript(ctx context.Context, in *v1.RoomTranscriptRequest) (*v1.RoomTranscriptReply, error) {
	s.log.WithContext(ctx).Debugf("RoomTranscript, name:%s, page:%d", in.GetName(), in.GetPage())

	authGrant, err := middleware.AuthJWT(ctx)
	if err != nil {
		return nil, err
	}
	if !authGrant.CanGetTranscript {
		s.log.WithContext(ctx).Errorf("Link().authGrant.CanGetTranscript false")
		return nil, errorV1.ErrorForbidden("token authorization is not allowed.")
	}

	if _, err := s.authUC.GetAuthByClientID(ctx, authGrant.ClientID); err != nil {
		s.log.WithContext(ctx).Errorf("s.authUC.GetAuthByClientID err:%v", err)
		return nil, errorV1.ErrorForbidden("token authorization error.")
	}

	resp, err := s.RoomTranscriptOnline(ctx, in)
	if err != nil {
		s.log.WithContext(ctx).Errorf("s.RoomTranscriptOnline err:%v", err)
		return nil, errorV1.ErrorBadRequest("")
	}

	if authGrant.CanGetVideo || authGrant.CanGetAudio {
		// get vod info
		roomVod, err := s.roomVodUC.GetBySid(ctx, resp.Room.Sid)
		if err != nil {
			s.log.WithContext(ctx).Errorf("s.roomVodUC.GetBySid failed, err:%v", err)
		} else {
			if roomVod.IsComplete() {
				resp.Vod = &v1.RoomTranscriptReply_Vod{
					EgressId:     roomVod.EgressID,
					Url:          roomVod.VodURL,
					Status:       uint32(roomVod.Status),
					StartTime:    uint32(roomVod.StartTime.Unix()),
					CompleteTime: uint32(roomVod.CompleteTime.Unix()),
					Duration:     uint32(roomVod.Duration),
				}
			} else {
				resp.Vod = &v1.RoomTranscriptReply_Vod{
					EgressId: roomVod.EgressID,
					Status:   uint32(roomVod.Status),
				}
			}
		}
	}

	return resp, nil
}

func (s *RoomService) SetRoomVoice(ctx context.Context, in *v1.SetRoomVoiceRequest) (*v1.NilReply, error) {
	if in.RoomName == "" || in.VoiceId == "" {
		return nil, nil
	}

	voices, err := s.voiceUC.GetVoices(ctx, "")
	if err != nil {
		return nil, errorV1.ErrorBadRequest("").WithCause(err)
	}
	flag := false
	// check voice id
	for _, voice := range voices.Voices {
		if voice.VoiceId == in.GetVoiceId() {
			flag = true
			break
		}
	}
	if !flag {
		return nil, errorV1.ErrorBadRequest("bad params")
	}
	if err := s.linkUC.SetRoomVoiceID(ctx, in.GetRoomName(), in.GetVoiceId()); err != nil {
		return nil, errorV1.ErrorBadRequest("").WithCause(err)
	}
	return &v1.NilReply{}, nil
}
