package service

import (
	"context"
	errorV1 "faceto-ai/api_gen/error/v1"
	v1 "faceto-ai/api_gen/voice/v1"
	"faceto-ai/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// VoiceService is voice api service.
type VoiceService struct {
	v1.UnimplementedVoiceServer
	log *log.Helper

	voiceUC *biz.VoiceUseCase
}

func NewVoiceService(
	logger log.Logger,
	voiceUC *biz.VoiceUseCase,
) *VoiceService {
	return &VoiceService{
		log:     log.NewHelper(logger),
		voiceUC: voiceUC,
	}
}

func (v *VoiceService) Voices(ctx context.Context, in *v1.VoicesRequest) (*v1.VoicesReply, error) {
	voiceList, err := v.voiceUC.GetVoices(ctx, in.Category)
	if err != nil {
		return nil, errorV1.ErrorBadRequest("")
	}
	resp := &v1.VoicesReply{
		Voices: make([]*v1.VoiceConfig, 0, len(voiceList.Voices)),
	}
	for _, voice := range voiceList.Voices {
		config := &v1.VoiceConfig{
			Name:       voice.Name,
			VoiceId:    voice.VoiceId,
			PreviewUrl: voice.PreviewUrl,
			Category:   voice.Category,
			Settings: &v1.VoiceSetting{
				SimilarityBoost: voice.Settings.SimilarityBoost,
				Stability:       voice.Settings.Stability,
			},
			Labels: &v1.VoiceLabel{
				Accent: voice.Labels.Accent,
				Age:    voice.Labels.Age,
				Gender: voice.Labels.Gender,
			},
		}
		resp.Voices = append(resp.Voices, config)
	}
	return resp, nil
}

func (v *VoiceService) GetVoice(ctx context.Context, in *v1.GetVoiceRequest) (*v1.VoiceConfig, error) {
	voice, err := v.voiceUC.GetVoice(ctx, in.VoiceId)
	if err != nil {
		return nil, errorV1.ErrorBadRequest("")
	}
	resp := &v1.VoiceConfig{
		Name:       voice.Name,
		VoiceId:    voice.VoiceId,
		PreviewUrl: voice.PreviewUrl,
		Category:   voice.Category,
		Settings: &v1.VoiceSetting{
			SimilarityBoost: voice.Settings.SimilarityBoost,
			Stability:       voice.Settings.Stability,
		},
		Labels: &v1.VoiceLabel{
			Accent: voice.Labels.Accent,
			Age:    voice.Labels.Age,
			Gender: voice.Labels.Gender,
		},
	}
	return resp, nil
}

func (v *VoiceService) UploadFileToStorage(ctx context.Context, file *biz.File) (*v1.FileUploadReply, error) {
	if err := v.voiceUC.UploadVoice(ctx, file); err != nil {
		return nil, errorV1.ErrorBadRequest("").WithCause(err)
	}
	return &v1.FileUploadReply{
		Url: file.URL,
	}, nil
}

func (v *VoiceService) EditVoice(ctx context.Context, in *v1.EditVoiceRequest) (*v1.Nil, error) {
	req := &biz.EditVoice{
		VoiceId: in.VoiceId,
		VoiceSetting: &biz.VoiceSetting{
			SimilarityBoost: in.SimilarityBoost,
			Stability:       in.Stability,
		},
	}
	if err := v.voiceUC.EditVoice(ctx, req); err != nil {
		return nil, errorV1.ErrorBadRequest("")
	}
	return &v1.Nil{}, nil
}

func (v *VoiceService) AddVoice(ctx context.Context, in *v1.AddVoiceRequest) (*v1.Nil, error) {
	req := &biz.AddVoice{
		Name:        in.Name,
		Labels:      in.Labels,
		Description: in.Description,
		Files:       in.Files,
	}
	if err := v.voiceUC.AddVoice(ctx, req); err != nil {
		return nil, errorV1.ErrorBadRequest("")
	}
	return &v1.Nil{}, nil
}
