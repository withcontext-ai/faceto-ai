package biz

import (
	"context"
	"faceto-ai/internal/data/schema"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/livekit/protocol/livekit"
	"github.com/pkg/errors"
	"time"
)

type Room struct {
	UUID      string    `json:"-,omitempty"`
	Name      string    `json:"name,omitempty"`
	Sid       string    `json:"sid,omitempty"`
	Status    uint8     `json:"status"` // room status 0.ready 1.starting 2.participant left 3.room end
	StartTime time.Time `json:"start_time"`
	LeftTime  time.Time `json:"left_time"`
	EndTime   time.Time `json:"end_time"`
	VodStatus uint8     `json:"vod_status"` // room vod status:0.init 1.starting 2.complete
}

func (r *Room) IsInterviewing() bool {
	return r.Status == schema.RoomStatusInterviewing && time.Since(r.StartTime) > 0
}

func (r *Room) IsParticipantLeft() bool {
	return r.Status == schema.RoomStatusParticipantLeft
}

func (r *Room) IsComplete() bool {
	return r.Status == schema.RoomStatusEnd && !r.EndTime.IsZero()
}

func (r *Room) AccessOrNot() bool {
	return r.Status == schema.RoomStatusReady
}

//go:generate mockgen -source room.go -destination ../mock/biz/room_mock.go -package=biz
type RoomRepo interface {
	Save(ctx context.Context, room *Room) error
	UpdateStatus(ctx context.Context, sid string, fromStatus []uint8, toStatus uint8, up *Room) (int, error)
	GetByID(ctx context.Context, ID uint64) (*Room, error)
	GetByUUID(ctx context.Context, uuid string) (*Room, error)
	GetBySID(ctx context.Context, sid string) (*Room, error)
	GetByName(ctx context.Context, name string) (*Room, error)
	List(ctx context.Context, page uint) ([]*Room, error)
}

type RoomUseCase struct {
	log      *log.Helper
	roomRepo RoomRepo
	linkRepo LinkRepo
}

func NewRoomUseCase(logger log.Logger, roomRepo RoomRepo, linkRepo LinkRepo) *RoomUseCase {
	return &RoomUseCase{
		log:      log.NewHelper(logger),
		roomRepo: roomRepo,
		linkRepo: linkRepo,
	}
}

func (ru *RoomUseCase) JoinRoom(ctx context.Context, room *livekit.Room) error {
	ru.log.WithContext(ctx).Infof("Biz.JoinRoom: sid:%s, name:%s", room.Sid, room.Name)
	return ru.roomRepo.Save(ctx, &Room{
		Name: room.Name,
		Sid:  room.Sid,
	})
}

func (ru *RoomUseCase) GetByName(ctx context.Context, name string) (*Room, error) {
	return ru.roomRepo.GetByName(ctx, name)
}

func (ru *RoomUseCase) GetBySid(ctx context.Context, sid string) (*Room, error) {
	return ru.roomRepo.GetBySID(ctx, sid)
}

func (ru *RoomUseCase) GetLinkInfo(ctx context.Context, sid string) (*Link, error) {
	room, err := ru.roomRepo.GetBySID(ctx, sid)
	if err != nil {
		return nil, errors.Wrap(err, "get room err")
	}
	return ru.linkRepo.GetLinkByName(ctx, room.Name)
}
