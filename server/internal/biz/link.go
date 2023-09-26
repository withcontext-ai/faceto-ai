package biz

import (
	"context"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/data/schema"
	"github.com/go-kratos/kratos/v2/log"
)

type Link struct {
	UUID       string             `json:"-"`
	RoomName   string             `json:"room_name"`
	Link       string             `json:"link"`
	ChatAPI    string             `json:"chat_api"`
	ChatAPIKey string             `json:"chat_api_key"`
	Token      string             `json:"token"`
	Config     *schema.RoomConfig `json:"config"`
	Webhook    *schema.Webhook    `json:"webhook"`
	Prompt     *schema.Prompt     `json:"prompt"`
	VoiceID    string             `json:"voice_id"`
}

func (l *Link) GetConfigGreeting() string {
	if l.Config != nil {
		return l.Config.Greeting
	}
	return ""
}

func (l *Link) GetConfigDuration() uint32 {
	if l.Config != nil {
		return l.Config.Duration
	}
	return 0
}

func (l *Link) GetConfigVoiceID() string {
	if l.Config != nil {
		return l.Config.VoiceID
	}
	return ""
}

func (l *Link) GetConfigUserName() string {
	if l.Config != nil {
		return l.Config.UserName
	}
	return ""
}

func (l *Link) GetConfigBotName() string {
	if l.Config != nil {
		return l.Config.BotName
	}
	return ""
}

func (l *Link) SetConfigVoiceID(voiceID string) {
	if l.Config != nil {
		l.Config.VoiceID = voiceID
	}
}

func (l *Link) GetWebHookConfig() *schema.Webhook {
	if l.Webhook != nil {
		return l.Webhook
	}
	return nil
}

func (l *Link) GetPromptConfig() *schema.Prompt {
	if l.Prompt != nil {
		return l.Prompt
	}
	return nil
}

//go:generate mockgen -source link.go -destination ../mock/biz/link_mock.go -package=biz
type LinkRepo interface {
	Save(ctx context.Context, view *Link) error
	Count(ctx context.Context) (int, error)
	GetLinkByName(ctx context.Context, roomName string) (*Link, error)
	SetRoomVoiceID(ctx context.Context, roomName, voiceID string) error
	SetConfigByUUID(ctx context.Context, uuid string, config *schema.RoomConfig) error
}

type LinkUseCase struct {
	log      *log.Helper
	thirdApi *conf.ThirdApi
	linkRepo LinkRepo
}

func NewLinkUseCase(
	logger log.Logger,
	thirdApi *conf.ThirdApi,
	linkRepo LinkRepo,
) *LinkUseCase {
	return &LinkUseCase{
		log:      log.NewHelper(logger),
		thirdApi: thirdApi,
		linkRepo: linkRepo,
	}
}

func (ru *LinkUseCase) Create(ctx context.Context, view *Link) error {
	// config init
	roomConfig := new(schema.RoomConfig)
	roomConfig.Duration = schema.Duration

	if view.Config != nil {
		if view.Config.Duration > 0 {
			roomConfig.Duration = view.Config.Duration
		}
		if view.Config.Greeting != "" {
			roomConfig.Greeting = view.Config.Greeting
		}
		if view.Config.VoiceID != "" {
			roomConfig.VoiceID = view.Config.VoiceID
		}
		if view.Config.UserName != "" {
			roomConfig.UserName = view.Config.UserName
		}
		if view.Config.BotName != "" {
			roomConfig.BotName = view.Config.BotName
		}
	}
	view.Config = roomConfig

	return ru.linkRepo.Save(ctx, view)
}

func (ru *LinkUseCase) GetLinkByName(ctx context.Context, name string) (*Link, error) {
	return ru.linkRepo.GetLinkByName(ctx, name)
}

func (ru *LinkUseCase) SetRoomVoiceID(ctx context.Context, name, voiceID string) error {
	return ru.linkRepo.SetRoomVoiceID(ctx, name, voiceID)
}

func (ru *LinkUseCase) SetConfigByUUID(ctx context.Context, uuid string, config *schema.RoomConfig) error {
	if config == nil {
		return nil
	}
	return ru.linkRepo.SetConfigByUUID(ctx, uuid, config)
}
