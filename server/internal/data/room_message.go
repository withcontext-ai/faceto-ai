package data

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/data/ent"
	"faceto-ai/internal/data/ent/roommessage"
	"faceto-ai/internal/pkg/utils/helper"
	"faceto-ai/internal/pkg/utils/logsnag"
	"github.com/pkg/errors"

	"github.com/go-kratos/kratos/v2/log"
)

type roomMessageRepo struct {
	data *Data
	log  *log.Helper
}

// NewRoomMsgRepo .
func NewRoomMsgRepo(data *Data, logger log.Logger) biz.RoomMessageRepo {
	return &roomMessageRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *roomMessageRepo) Save(ctx context.Context, msg *biz.RoomMessage) error {
	result, err := r.data.DB(ctx).RoomMessage.Create().
		SetSid(msg.Sid).
		SetParticipantSid(msg.ParticipantSID).
		SetParticipantName(msg.ParticipantName).
		SetType(msg.Type).
		SetIsReply(msg.IsReply).
		SetReplyID(msg.ReplyID).
		SetEventTime(msg.EventTime).
		SetText(msg.Text).
		Save(ctx)
	if err != nil {
		return err
	}

	msg.ID = result.ID
	msg.UUID = result.UUID
	msg.Sid = result.Sid
	msg.ParticipantSID = result.ParticipantSid
	msg.ParticipantName = result.ParticipantName
	msg.Type = result.Type
	msg.EventTime = result.EventTime
	msg.Text = result.Text

	// logsnag
	go func() {
		if msg.Type == biz.RoomMessageTypeBot {
			logsnag.Event(ctx, logsnag.EventRoomMessageAI.SetRoomName(msg.RoomName).SetUID(msg.Sid).SetMessage(msg.Text))
		} else {
			logsnag.Event(ctx, logsnag.EventRoomMessageUser.SetRoomName(msg.RoomName).SetUID(msg.Sid).SetMessage(msg.Text))
		}

		bgctx := helper.NewWithParentReqID(ctx)
		total, err := r.Count(bgctx)
		if err != nil {
			return
		}
		logsnag.Insight(bgctx, logsnag.InsightRoomMsgCount.SetValue(total))
	}()

	return nil
}

func (r *roomMessageRepo) Count(ctx context.Context) (int, error) {
	return r.data.DB(ctx).RoomMessage.Query().Where(roommessage.DeletedAtIsNil()).Count(ctx)
}

func (r *roomMessageRepo) Reply(ctx context.Context, msg *biz.RoomMessage) error {

	if err := r.data.WithTx(ctx, func(ctx context.Context) error {
		if err := r.Save(ctx, &biz.RoomMessage{
			Sid:             msg.Sid,
			ParticipantSID:  msg.ParticipantSID,
			ParticipantName: msg.ParticipantName,
			Type:            msg.Type,
			IsReply:         1,
			ReplyID:         msg.ReplyID,
			EventTime:       msg.EventTime,
			Text:            msg.Text,
		}); err != nil {
			return errors.Wrap(err, "save err")
		}

		if err := r.data.DB(ctx).RoomMessage.
			UpdateOneID(msg.ReplyID).
			SetIsReply(1).
			Exec(ctx); err != nil {
			return errors.Wrap(err, "update is_reply err")
		}

		return nil
	}); err != nil {
		r.log.Errorf("data.room_message.Reply withTx err:%v", err)
		return err
	}

	// logsnag
	logsnag.Event(ctx, logsnag.EventRoomMessageAI.SetRoomName(msg.RoomName).SetUID(msg.Sid).SetMessage(msg.Text))

	return nil
}

func (r *roomMessageRepo) History(ctx context.Context, sid string, page int) (int, []*biz.RoomMessage, error) {
	var (
		total  int
		result []*ent.RoomMessage
		err    error
	)

	total, err = r.data.DB(ctx).RoomMessage.Query().
		Where(roommessage.Sid(sid)).
		Where(roommessage.DeletedAtIsNil()).
		Count(ctx)
	if err != nil {
		return 0, nil, err
	}

	if page <= 0 {
		result, err = r.data.DB(ctx).RoomMessage.Query().
			Where(roommessage.Sid(sid)).
			Where(roommessage.DeletedAtIsNil()).
			Order(ent.Asc(roommessage.FieldCreatedAt)).
			All(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return 0, nil, biz.ErrRoomHistoryFound

			}
			return 0, nil, err
		}

	} else {
		offset := (page - 1) * biz.HistorySize
		result, err = r.data.DB(ctx).RoomMessage.Query().
			Where(roommessage.Sid(sid)).
			Where(roommessage.DeletedAtIsNil()).
			Order(ent.Asc(roommessage.FieldCreatedAt)).
			Offset(offset).Limit(biz.HistorySize).
			All(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return 0, nil, biz.ErrRoomHistoryFound
			}
			return 0, nil, err
		}
	}

	resp := make([]*biz.RoomMessage, 0, len(result))
	for _, v := range result {
		resp = append(resp, &biz.RoomMessage{
			ID:              v.ID,
			UUID:            v.UUID,
			Sid:             v.Sid,
			ParticipantSID:  v.ParticipantSid,
			ParticipantName: v.ParticipantName,
			Type:            v.Type,
			IsReply:         v.IsReply,
			ReplyID:         v.ReplyID,
			EventTime:       v.EventTime,
			Text:            v.Text,
		})
	}
	return total, resp, nil
}
