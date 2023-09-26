package data

import (
	"context"
	"faceto-ai/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type roomWebhookRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoomWebhookRepo(data *Data, logger log.Logger) biz.RoomWebhookRepo {
	return &roomWebhookRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *roomWebhookRepo) Save(ctx context.Context, webhook *biz.RoomWebhook) error {
	result, err := r.data.DB(ctx).RoomWebhook.Create().
		SetSid(webhook.Sid).
		SetName(webhook.Name).
		SetEvent(webhook.Event).
		SetEventTime(webhook.EventTime).
		SetParticipant(webhook.Participant).
		SetRoom(webhook.Room).
		SetTrack(webhook.Track).
		SetEgressInfo(webhook.EgressInfo).
		SetIngressInfo(webhook.InEgressInfo).
		SetTrack(webhook.Track).
		Save(ctx)
	if err != nil {
		return err
	}

	webhook.UUID = result.UUID
	return nil
}
