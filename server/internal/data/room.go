package data

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/data/ent"
	"faceto-ai/internal/data/ent/room"
	"faceto-ai/internal/data/schema"
	"faceto-ai/internal/pkg/utils/helper"
	"faceto-ai/internal/pkg/utils/logsnag"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type roomRepo struct {
	data *Data
	log  *log.Helper
}

// NewRoomRepo .
func NewRoomRepo(data *Data, logger log.Logger) biz.RoomRepo {
	return &roomRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *roomRepo) Save(ctx context.Context, room *biz.Room) error {
	getRoom, err := r.GetBySID(ctx, room.Sid)
	if err != nil {
		if ent.IsNotFound(err) {
			result, err := r.data.DB(ctx).Room.Create().
				SetName(room.Name).
				SetSid(room.Sid).
				SetStartTime(time.Now()).
				Save(ctx)
			if err != nil {
				return err
			}
			room.UUID = result.UUID
			room.Name = result.Name
			room.Sid = result.Sid
			room.StartTime = result.StartTime
			room.Status = result.Status
			go func() {
				bgctx := helper.NewWithParentReqID(ctx)
				total, err := r.Count(bgctx)
				if err != nil {
					return
				}
				logsnag.Insight(bgctx, logsnag.InsightRoomCount.SetValue(total))
			}()
			return nil
		}
		return err
	}

	room.UUID = getRoom.UUID
	room.Name = getRoom.Name
	room.Sid = getRoom.Sid
	room.Status = getRoom.Status
	room.StartTime = getRoom.StartTime
	room.LeftTime = getRoom.LeftTime
	room.EndTime = getRoom.EndTime
	room.VodStatus = getRoom.VodStatus
	return nil
}

func (r *roomRepo) Count(ctx context.Context) (int, error) {
	return r.data.DB(ctx).Room.Query().Where(room.DeletedAtIsNil()).Count(ctx)
}

func (r *roomRepo) UpdateStatus(ctx context.Context, sid string, fromStatus []uint8, toStatus uint8, up *biz.Room) (int, error) {
	db := r.data.DB(ctx).Room.Update().Where(room.Sid(sid))
	if len(fromStatus) > 0 {
		db.Where(room.StatusIn(fromStatus...))
	}
	switch toStatus {
	case schema.RoomStatusParticipantLeft: // user left
		return db.SetLeftTime(time.Now()).SetStatus(toStatus).Save(ctx)
	case schema.RoomStatusEnd: // room end
		return db.SetEndTime(time.Now()).SetStatus(toStatus).Save(ctx)
	default:
		return db.SetStatus(toStatus).Save(ctx)
	}
}

func (r *roomRepo) GetBySID(ctx context.Context, sid string) (*biz.Room, error) {
	result, err := r.data.DB(ctx).Room.Query().
		Where(room.Sid(sid)).
		Where(room.DeletedAtIsNil()).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, err
		}
		return nil, err
	}

	return &biz.Room{
		UUID:      result.UUID,
		Name:      result.Name,
		Sid:       result.Sid,
		Status:    result.Status,
		StartTime: result.StartTime,
		LeftTime:  result.LeftTime,
		EndTime:   result.EndTime,
		VodStatus: result.VodStatus,
	}, nil
}

func (r *roomRepo) GetByID(ctx context.Context, ID uint64) (*biz.Room, error) {
	result, err := r.data.DB(ctx).Room.Get(ctx, ID)
	if ent.IsNotFound(err) {
		return nil, err
	}
	return &biz.Room{
		UUID:      result.UUID,
		Name:      result.Name,
		Sid:       result.Sid,
		Status:    result.Status,
		StartTime: result.StartTime,
		LeftTime:  result.LeftTime,
		EndTime:   result.EndTime,
		VodStatus: result.VodStatus,
	}, nil
}

func (r *roomRepo) GetByUUID(ctx context.Context, uuid string) (*biz.Room, error) {
	result, err := r.data.DB(ctx).Room.Query().Where(room.UUID(uuid)).First(ctx)
	if ent.IsNotFound(err) {
		return nil, err
	}
	return &biz.Room{
		UUID:      result.UUID,
		Name:      result.Name,
		Sid:       result.Sid,
		Status:    result.Status,
		StartTime: result.StartTime,
		LeftTime:  result.LeftTime,
		EndTime:   result.EndTime,
		VodStatus: result.VodStatus,
	}, nil
}

func (r *roomRepo) GetByName(ctx context.Context, name string) (*biz.Room, error) {
	result, err := r.data.DB(ctx).Room.Query().
		Where(room.Name(name)).
		Order(ent.Desc(room.FieldCreatedAt)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, biz.ErrRoomNotFound
		}
		return nil, err
	}

	return &biz.Room{
		UUID:      result.UUID,
		Name:      result.Name,
		Sid:       result.Sid,
		Status:    result.Status,
		StartTime: result.StartTime,
		LeftTime:  result.LeftTime,
		EndTime:   result.EndTime,
		VodStatus: result.VodStatus,
	}, nil
}

func (r *roomRepo) List(ctx context.Context, page uint) ([]*biz.Room, error) {
	return nil, nil
}
