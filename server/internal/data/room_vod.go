package data

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/data/ent"
	"faceto-ai/internal/data/ent/room"
	"faceto-ai/internal/data/ent/roomvod"
	"faceto-ai/internal/data/schema"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

type roomVodRepo struct {
	data *Data
	log  *log.Helper
}

func NewRoomVodRepo(data *Data, logger log.Logger) biz.RoomVodRepo {
	return &roomVodRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *roomVodRepo) Save(ctx context.Context, vod *biz.RoomVod) error {
	db := r.data.DB(ctx).RoomVod.Create().
		SetSid(vod.Sid).
		SetName(vod.Name).
		SetEgressID(vod.EgressID).
		SetStatus(vod.Status).
		SetPlatform(vod.Platform).
		SetVodType(vod.VodType).
		SetVodPath(vod.VodPath).
		SetVodURL(vod.VodURL).
		SetDuration(vod.Duration)

	if !vod.StartTime.IsZero() {
		db.SetStartTime(vod.StartTime)
	}
	if !vod.CompleteTime.IsZero() {
		db.SetCompleteTime(vod.CompleteTime)
	}

	result, err := db.Save(ctx)
	if err != nil {
		return err
	}

	vod.UUID = result.UUID
	return nil
}

func (r *roomVodRepo) GetByEgressID(ctx context.Context, egressID string) (*biz.RoomVod, error) {
	result, err := r.data.DB(ctx).RoomVod.Query().
		Where(roomvod.EgressID(egressID)).
		Where(roomvod.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrRoomVodNotFound
		}
		return nil, err
	}
	return &biz.RoomVod{
		ID:           result.ID,
		UUID:         result.UUID,
		Name:         result.Name,
		Sid:          result.Sid,
		EgressID:     result.EgressID,
		Status:       result.Status,
		Platform:     result.Platform,
		VodType:      result.VodType,
		VodPath:      result.VodPath,
		VodURL:       result.VodURL,
		StartTime:    result.StartTime,
		CompleteTime: result.CompleteTime,
		Duration:     result.Duration,
	}, nil
}

func (r *roomVodRepo) GetBySid(ctx context.Context, sid string) (*biz.RoomVod, error) {
	result, err := r.data.DB(ctx).RoomVod.Query().
		Where(roomvod.Sid(sid)).
		Where(roomvod.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrRoomVodNotFound
		}
		return nil, err
	}
	return &biz.RoomVod{
		ID:           result.ID,
		UUID:         result.UUID,
		Name:         result.Name,
		Sid:          result.Sid,
		EgressID:     result.EgressID,
		Status:       result.Status,
		Platform:     result.Platform,
		VodType:      result.VodType,
		VodPath:      result.VodPath,
		VodURL:       result.VodURL,
		StartTime:    result.StartTime,
		CompleteTime: result.CompleteTime,
		Duration:     result.Duration,
	}, nil
}

func (r *roomVodRepo) UpdateStatus(ctx context.Context, egressID string, fromStatus []uint8, toStatus uint8, vod *biz.RoomVod) error {
	if err := r.data.WithTx(ctx, func(ctx context.Context) error {
		// 1.room vod_status
		_, err := r.data.DB(ctx).Room.Update().Where(room.Sid(vod.Sid)).SetVodStatus(toStatus).Save(ctx)
		if err != nil {
			return err
		}

		// 2.vod status
		db := r.data.DB(ctx).RoomVod.Update().Where(roomvod.EgressID(egressID))
		if len(fromStatus) > 0 {
			db = db.Where(roomvod.StatusIn(fromStatus...))
		}
		if toStatus == schema.RoomVodStatusComplete {
			db = db.SetStartTime(vod.StartTime).
				SetCompleteTime(vod.CompleteTime).
				SetVodURL(vod.VodURL).
				SetVodPath(vod.VodPath).
				SetDuration(vod.Duration)
		}
		_, err = db.SetStatus(toStatus).Save(ctx)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		r.log.WithContext(ctx).Errorf("roomVodRepo.UpdateStatus r.data.WithTx err:%v", err)
		return errors.Wrap(err, "update status error")
	}
	return nil
}
