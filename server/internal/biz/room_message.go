package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"time"
)

const (
	RoomMessageTypeUser = iota + 1
	RoomMessageTypeBot
)

const (
	HistorySize = 20
)

type RoomMessage struct {
	ID              uint64    `json:"-"`
	UUID            string    `json:"uuid"`
	RoomName        string    `json:"room_name"`
	Sid             string    `json:"sid"`
	ParticipantSID  string    `json:"participant_sid"`
	ParticipantName string    `json:"participant_name"`
	Type            uint32    `json:"type"`
	IsReply         uint32    `json:"is_reply"`
	ReplyID         uint64    `json:"reply_id"`
	EventTime       time.Time `json:"event_time"`
	Text            string    `json:"text"`
}

type RoomMessageRepo interface {
	Save(ctx context.Context, msg *RoomMessage) error
	Reply(ctx context.Context, msg *RoomMessage) error
	History(ctx context.Context, sid string, page int) (int, []*RoomMessage, error)
}

type RoomMessageUseCase struct {
	log         *log.Helper
	roomRepo    RoomRepo
	roomMsgRepo RoomMessageRepo
}

func NewRoomMsgUseCase(
	logger log.Logger,
	roomRepo RoomRepo,
	roomMsgRepo RoomMessageRepo,
) *RoomMessageUseCase {
	return &RoomMessageUseCase{
		log:         log.NewHelper(logger),
		roomRepo:    roomRepo,
		roomMsgRepo: roomMsgRepo,
	}
}

func (ru *RoomMessageUseCase) Record(ctx context.Context, roomMsg *RoomMessage) error {
	return ru.roomMsgRepo.Save(ctx, roomMsg)
}

func (ru *RoomMessageUseCase) Reply(ctx context.Context, roomMsg *RoomMessage) error {
	return ru.roomMsgRepo.Reply(ctx, roomMsg)
}

func (ru *RoomMessageUseCase) Transcript(ctx context.Context, sid string, page uint32) (int, []*RoomMessage, error) {
	total, history, err := ru.roomMsgRepo.History(ctx, sid, int(page))
	if err != nil {
		if errors.Is(err, ErrRoomHistoryFound) {
			return 0, nil, nil
		}
		return 0, nil, err
	}
	return total, history, nil
}
