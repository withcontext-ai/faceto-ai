package biz

import (
	"context"
	"faceto-ai/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

type VoicesList struct {
	Voices []*VoiceConfig `json:"voice"`
}

type VoiceConfig struct {
	Name       string        `json:"name"`
	VoiceId    string        `json:"voice_id"`
	PreviewUrl string        `json:"preview_url"`
	Category   string        `json:"category"`
	Settings   *VoiceSetting `json:"settings"`
	Labels     *VoiceLabel   `json:"labels"`
}

type EditVoice struct {
	VoiceId string `json:"voice_id"`
	*VoiceSetting
}

type AddVoice struct {
	Name        string   `json:"name"`
	Labels      string   `json:"labels"`
	Description string   `json:"description"`
	Files       []string `json:"files"`
}

type VoiceSetting struct {
	SimilarityBoost float32 `json:"similarity_boost"`
	Stability       float32 `json:"stability"`
}

type VoiceLabel struct {
	Accent string `json:"accent"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

type VoiceRepo interface {
	GetVoices(ctx context.Context) (*VoicesList, error)
	GetVoice(ctx context.Context, voiceID string) (*VoiceConfig, error)
	EditVoice(ctx context.Context, voice *EditVoice) error
	AddVoice(ctx context.Context, voice *AddVoice) error
}

type VoiceUseCase struct {
	log         *log.Helper
	thirdApi    *conf.ThirdApi
	voiceRepo   VoiceRepo
	storageRepo StorageRepo
}

func NewVoiceUseCase(
	logger log.Logger,
	thirdApi *conf.ThirdApi,
	voiceRepo VoiceRepo,
	storageRepo StorageRepo,
) *VoiceUseCase {
	return &VoiceUseCase{
		log:         log.NewHelper(logger),
		thirdApi:    thirdApi,
		voiceRepo:   voiceRepo,
		storageRepo: storageRepo,
	}
}

func (v *VoiceUseCase) GetVoices(ctx context.Context, category string) (*VoicesList, error) {
	voiceList, err := v.voiceRepo.GetVoices(ctx)
	if err != nil {
		return nil, err
	}
	// all
	if category == "" {
		return voiceList, nil
	}
	if voiceList.Voices != nil && len(voiceList.Voices) == 0 {
		return voiceList, nil
	}

	// category: premade, generated
	resp := &VoicesList{Voices: make([]*VoiceConfig, 0, len(voiceList.Voices))}
	for _, voice := range voiceList.Voices {
		if voice.Category == category {
			resp.Voices = append(resp.Voices, voice)
		}
	}
	return resp, nil
}

func (v *VoiceUseCase) GetVoice(ctx context.Context, voiceID string) (*VoiceConfig, error) {
	return v.voiceRepo.GetVoice(ctx, voiceID)
}

func (v *VoiceUseCase) EditVoice(ctx context.Context, editVoice *EditVoice) error {
	return v.voiceRepo.EditVoice(ctx, editVoice)
}

func (v *VoiceUseCase) AddVoice(ctx context.Context, addVoice *AddVoice) error {
	return v.voiceRepo.AddVoice(ctx, addVoice)
}

func (v *VoiceUseCase) UploadVoice(ctx context.Context, file *File) error {
	if err := v.storageRepo.UploadFile(ctx, file); err != nil {
		return errors.Wrap(err, "storage upload err")
	}
	return nil
}
