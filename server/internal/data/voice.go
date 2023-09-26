package data

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/biz/liveGPT/elevenlabs"
	"faceto-ai/internal/biz/liveGPT/elevenlabs/types"
	"faceto-ai/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"os"
	"sync"
)

type voiceRepo struct {
	client   elevenlabs.Client
	thirdApi *conf.ThirdApi
	log      *log.Helper
}

func NewVoiceRepo(thirdApi *conf.ThirdApi, logger log.Logger) biz.VoiceRepo {
	loggerObj := log.NewHelper(logger)
	return &voiceRepo{
		client: elevenlabs.New(thirdApi.Eleventlabs.Key, loggerObj),
		log:    loggerObj,
	}
}

func (v *voiceRepo) GetVoices(ctx context.Context) (*biz.VoicesList, error) {
	voiceModelList, err := v.client.GetVoices(ctx)
	if err != nil {
		v.log.Errorf("get voices err:%v", err)
		return nil, errors.Wrap(err, "get voices err")
	}

	resp := &biz.VoicesList{Voices: make([]*biz.VoiceConfig, 0)}
	for _, voice := range voiceModelList {
		config := &biz.VoiceConfig{
			Name:       voice.Name,
			VoiceId:    voice.VoiceID,
			PreviewUrl: voice.PreviewURL,
			Category:   voice.Category,
			Settings: &biz.VoiceSetting{
				SimilarityBoost: float32(voice.Settings.SimilarityBoost),
				Stability:       float32(voice.Settings.Stability),
			},
			Labels: &biz.VoiceLabel{
				Accent: voice.Labels["Accent"],
				Age:    voice.Labels["Age"],
				Gender: voice.Labels["Gender"],
			},
		}
		resp.Voices = append(resp.Voices, config)
	}

	return resp, nil
}

func (v *voiceRepo) GetVoice(ctx context.Context, voiceID string) (*biz.VoiceConfig, error) {
	voice, err := v.client.GetVoice(ctx, voiceID)
	if err != nil {
		v.log.Errorf("get voice err:%v", err)
		return nil, errors.Wrap(err, "get voice err")
	}

	resp := &biz.VoiceConfig{
		Name:       voice.Name,
		VoiceId:    voice.VoiceID,
		PreviewUrl: voice.PreviewURL,
		Category:   voice.Category,
		Settings: &biz.VoiceSetting{
			SimilarityBoost: float32(voice.Settings.SimilarityBoost),
			Stability:       float32(voice.Settings.Stability),
		},
		Labels: &biz.VoiceLabel{
			Accent: voice.Labels["Accent"],
			Age:    voice.Labels["Age"],
			Gender: voice.Labels["Gender"],
		},
	}
	return resp, nil
}

func (v *voiceRepo) EditVoice(ctx context.Context, editVoice *biz.EditVoice) error {
	setting := types.SynthesisOptions{
		Stability:       float64(editVoice.Stability),
		SimilarityBoost: float64(editVoice.SimilarityBoost),
	}
	if err := v.client.EditVoiceSettings(ctx, editVoice.VoiceId, setting); err != nil {
		v.log.Errorf("client edit voice err:%v", err)
		return errors.Wrap(err, "edit voice err")
	}
	return nil
}

func (v *voiceRepo) AddVoice(ctx context.Context, addVoice *biz.AddVoice) error {
	var sg sync.WaitGroup
	files := make([]*os.File, 0, len(addVoice.Files))
	// get voice link
	for _, file := range addVoice.Files {
		sg.Add(1)
		go func(link string) {
			// 发送 HTTP GET 请求
			response, err := http.Get(link)
			if err != nil {
				v.log.Errorf("http.Get(link), link:%s, err:%v", link, err)
				return
			}
			defer response.Body.Close()

			tempFile, err := os.CreateTemp("", "tempfile")
			if err != nil {
				v.log.Errorf("os.CreateTemp, link:%s, err:%v", link, err)
				return
			}
			defer tempFile.Close()
			v.log.Debugf("os.CreateTempy, link:%s, tempName:%s", link, tempFile.Name())

			// 将响应的内容写入临时文件
			_, err = io.Copy(tempFile, response.Body)
			if err != nil {
				v.log.Errorf("io.Copy, link:%s, err:%v", link, err)
				return
			}

			// 打开临时文件并返回 os.File 结构
			file, err := os.Open(tempFile.Name())
			if err != nil {
				v.log.Errorf("os.Open(tempFile.Name()), link:%s, tempName:%s, err:%v", link, tempFile.Name(), err)
				return
			}

			// 最后记得删除临时文件
			defer os.Remove(tempFile.Name())

			files = append(files, file)
		}(file)
	}
	sg.Wait()

	if err := v.client.CreateVoice(ctx, addVoice.Name, addVoice.Description, []string{addVoice.Labels}, files); err != nil {
		v.log.Errorf("client create voice err:%v", err)
		return errors.Wrap(err, "create voice err")
	}

	return nil
}
