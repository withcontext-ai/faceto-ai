package data

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/data/ent"
	"faceto-ai/internal/data/ent/link"
	"faceto-ai/internal/data/schema"
	"faceto-ai/internal/pkg/utils/helper"
	"faceto-ai/internal/pkg/utils/logsnag"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
)

const LinkUrl = "https://faceto.withcontext.ai/rooms/%s"

type linkRepo struct {
	data *Data
	log  *log.Helper
}

func NewLinkRepo(data *Data, logger log.Logger) biz.LinkRepo {
	return &linkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *linkRepo) Save(ctx context.Context, view *biz.Link) error {
	uuid := schema.MustUID()
	roomName := helper.GenerateRoomID()
	linkUrl := fmt.Sprintf(LinkUrl, roomName)

	db := r.data.DB(ctx).Link.Create().
		SetUUID(uuid).
		SetRoomName(roomName).
		SetLink(linkUrl).
		SetChatAPI(view.ChatAPI).
		SetChatAPIKey(view.ChatAPIKey).
		SetToken(view.Token)

	if view.Config != nil {
		db.SetConfig(view.Config)
	}
	if view.Webhook != nil {
		db.SetWebhook(view.Webhook)
	}
	if view.Prompt != nil {
		db.SetPrompt(view.Prompt)
	}

	result, err := db.Save(ctx)
	if err != nil {
		return err
	}

	view.UUID = result.UUID
	view.RoomName = roomName
	view.Link = linkUrl

	// logsnag
	logsnag.Event(ctx, logsnag.EventApplyLink.SetUID(view.UUID).SetRoomName(roomName).SetNotify())
	go func() {
		bgctx := helper.NewWithParentReqID(ctx)
		total, err := r.Count(bgctx)
		if err != nil {
			return
		}
		logsnag.Insight(bgctx, logsnag.InsightLinkCount.SetValue(total))
	}()

	return nil
}

func (r *linkRepo) Count(ctx context.Context) (int, error) {
	return r.data.DB(ctx).Link.Query().Where(link.DeletedAtIsNil()).Count(ctx)
}

func (r *linkRepo) GetLinkByName(ctx context.Context, roomName string) (*biz.Link, error) {
	result, err := r.data.DB(ctx).Link.Query().
		Where(link.RoomName(roomName)).
		Where(link.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrInterviewNotFound
		}
		return nil, err
	}
	return &biz.Link{
		UUID:       result.UUID,
		RoomName:   result.RoomName,
		Link:       result.Link,
		ChatAPI:    result.ChatAPI,
		ChatAPIKey: result.ChatAPIKey,
		Config:     result.Config,
		Webhook:    result.Webhook,
		Prompt:     result.Prompt,
		Token:      result.Token,
	}, nil
}

func (r *linkRepo) SetRoomVoiceID(ctx context.Context, roomName, voiceID string) error {
	room, err := r.GetLinkByName(ctx, roomName)
	if err != nil {
		if ent.IsNotFound(err) {
			return biz.ErrInterviewNotFound
		}
		return err
	}
	if room.Config != nil {
		room.Config.VoiceID = voiceID
	}
	if err := r.data.DB(ctx).Link.Update().
		Where(link.RoomName(roomName)).
		SetConfig(room.Config).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (r *linkRepo) SetConfigByUUID(ctx context.Context, uuid string, config *schema.RoomConfig) error {
	if err := r.data.DB(ctx).Link.Update().
		Where(link.UUID(uuid)).
		SetConfig(config).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}
