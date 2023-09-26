package biz

import (
	"context"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/data/schema"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/livekit/protocol/livekit"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type RoomVod struct {
	ID           uint64    `json:"-"`
	UUID         string    `json:"uuid"`
	Name         string    `json:"name"`
	Sid          string    `json:"sid"`
	EgressID     string    `json:"egress_id"`
	Status       uint8     `json:"status"`   // vod status 0.init 1.start 2.complete 3.fail
	Platform     uint8     `json:"platform"` // vod storage platform 1.azure 2.gcp 3.s3
	VodType      uint8     `json:"vod_type"` // vod type 1.file 2.steam data
	VodPath      string    `json:"vod_path"`
	VodURL       string    `json:"vod_url"`
	StartTime    time.Time `json:"start_time"`
	CompleteTime time.Time `json:"complete_time"`
	Duration     uint64    `json:"duration"`
}

func (rv *RoomVod) IsComplete() bool {
	return rv.Status == schema.RoomVodStatusComplete
}

//go:generate mockgen -source room_vod.go -destination ../mock/biz/room_vod_mock.go -package=biz
type RoomVodRepo interface {
	Save(ctx context.Context, vod *RoomVod) error
	GetByEgressID(ctx context.Context, egressID string) (*RoomVod, error)
	GetBySid(ctx context.Context, sid string) (*RoomVod, error)
	UpdateStatus(ctx context.Context, egressID string, fromStatus []uint8, toStatus uint8, vod *RoomVod) error
}

type RoomVodUseCase struct {
	log         *log.Helper
	confStorage *conf.Storage
	roomRepo    RoomRepo
	roomVodRepo RoomVodRepo
}

func NewRoomVodUseCase(
	logger log.Logger,
	confStorage *conf.Storage,
	roomRepo RoomRepo,
	roomVodRepo RoomVodRepo,
) *RoomVodUseCase {
	return &RoomVodUseCase{
		log:         log.NewHelper(logger),
		confStorage: confStorage,
		roomRepo:    roomRepo,
		roomVodRepo: roomVodRepo,
	}
}

func (rv *RoomVodUseCase) GetBySid(ctx context.Context, sid string) (*RoomVod, error) {
	return rv.roomVodRepo.GetBySid(ctx, sid)
}

func (rv *RoomVodUseCase) GetByEgressID(ctx context.Context, egressID string) (*RoomVod, error) {
	return rv.roomVodRepo.GetByEgressID(ctx, egressID)
}

func (rv *RoomVodUseCase) CreateWithEgressInfoAzure(ctx context.Context, egress *livekit.EgressInfo) (*RoomVod, error) {
	if egress == nil {
		return nil, nil
	}
	vodPath := egress.GetFileResults()[0].GetFilename()
	if vodPath == "" {
		return nil, nil
	}
	urlslice := []string{
		strings.TrimRight(rv.confStorage.AzureBlob.GetCdnHost(), "/"),
		strings.Trim(rv.confStorage.AzureBlob.GetContainerName(), "/"),
		strings.TrimLeft(vodPath, "/"),
	}
	return rv.Create(ctx, &RoomVod{
		Name:     egress.GetRoomName(),
		Sid:      egress.GetRoomId(),
		EgressID: egress.GetEgressId(),
		Status:   schema.RoomVodStatusReady,
		Platform: schema.VodPlatFormAzure,
		VodType:  schema.VodTypeFile,
		VodPath:  vodPath,
		VodURL:   strings.Join(urlslice, "/"),
	})
}

func (rv *RoomVodUseCase) Create(ctx context.Context, vod *RoomVod) (*RoomVod, error) {
	if vod == nil || vod.EgressID == "" {
		return nil, nil
	}

	roomVod, err := rv.roomVodRepo.GetByEgressID(ctx, vod.EgressID)
	if err != nil {
		if errors.Is(err, ErrRoomVodNotFound) {
			if err := rv.roomVodRepo.Save(ctx, vod); err != nil {
				return nil, errors.Wrap(err, "room vod save err")
			}
		}
		return nil, errors.Wrap(err, "get by egress id err")
	}

	// already complete
	if roomVod.Status == schema.RoomVodStatusComplete {
		return nil, nil
	}

	var statusErr error
	switch vod.Status {
	case schema.RoomVodStatusReady:

	case schema.RoomVodStatusStarting:
		statusErr = rv.roomVodRepo.UpdateStatus(
			ctx, vod.EgressID,
			[]uint8{schema.RoomVodStatusReady, schema.RoomVodStatusFail},
			schema.RoomVodStatusStarting,
			vod)
	case schema.RoomVodStatusComplete:
		statusErr = rv.roomVodRepo.UpdateStatus(
			ctx, vod.EgressID,
			[]uint8{schema.RoomVodStatusReady, schema.RoomVodStatusStarting, schema.RoomVodStatusFail},
			schema.RoomVodStatusComplete,
			vod)
	case schema.RoomVodStatusFail:
		statusErr = rv.roomVodRepo.UpdateStatus(
			ctx, vod.EgressID,
			[]uint8{schema.RoomVodStatusReady, schema.RoomVodStatusStarting},
			schema.RoomVodStatusFail,
			vod)
	}
	if statusErr != nil {
		return nil, errors.Wrap(err, "set room vod status err")
	}

	return nil, nil
}
