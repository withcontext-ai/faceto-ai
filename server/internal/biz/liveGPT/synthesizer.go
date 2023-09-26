package liveGPT

import (
	"bytes"
	tts "cloud.google.com/go/texttospeech/apiv1"
	ttspb "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/biz/liveGPT/elevenlabs"
	"faceto-ai/internal/biz/liveGPT/elevenlabs/types"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/pkg/utils/logsnag"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"io"
	"os/exec"
	"strings"
	"time"
)

type Synthesizer struct {
	client           *tts.Client
	elevenlabsClient elevenlabs.Client
	log              *log.Helper
	thirdConf        *conf.ThirdApi
	linkConfig       *biz.Link
	Voices           []string
}

type SynthesizeSpeechResponse struct {
	AudioContent []byte `json:"audio_content,omitempty"`
}

func NewSynthesizer(
	client *tts.Client,
	log *log.Helper,
	thirdConf *conf.ThirdApi,
	linkConfig *biz.Link,
) *Synthesizer {
	synthesizer := &Synthesizer{
		client:           client,
		elevenlabsClient: elevenlabs.New(thirdConf.Eleventlabs.Key, log), // os.Getenv("ELEVEN_API_KEY")
		log:              log,
		thirdConf:        thirdConf,
		linkConfig:       linkConfig,
	}
	go getVoices(synthesizer)
	return synthesizer
}

func getVoices(s *Synthesizer) error {
	ctx := context.Background()
	ids, err := s.elevenlabsClient.GetVoiceIDs(ctx)
	if err != nil {
		s.log.WithContext(ctx).Errorf("Synthesize.client.GetVoiceIDs err:%v", err)
		return err
	}
	s.log.Debugw("Synthesize.GetVoiceIDs", ids)
	s.Voices = ids
	return nil
}

func (s *Synthesizer) SetVoiceID(voiceID string) bool {
	if len(s.Voices) == 0 {
		if err := getVoices(s); err != nil {
			return false
		}
	}
	flag := false
	// check voice id
	for _, vid := range s.Voices {
		if vid == voiceID {
			flag = true
			break
		}
	}
	if flag {
		s.linkConfig.SetConfigVoiceID(voiceID)
		return true
	}
	return false
}

func (s *Synthesizer) SynthesizeGoogle(ctx context.Context, text string, language *Language) (*SynthesizeSpeechResponse, error) {
	req := &ttspb.SynthesizeSpeechRequest{
		Input: &ttspb.SynthesisInput{
			InputSource: &ttspb.SynthesisInput_Text{
				Text: text,
			},
		},
		Voice: &ttspb.VoiceSelectionParams{
			LanguageCode: language.Code,
			Name:         language.SynthesizerModel,
		},
		AudioConfig: &ttspb.AudioConfig{
			AudioEncoding:   ttspb.AudioEncoding_OGG_OPUS,
			SampleRateHertz: 48000,
		},
	}

	launchTime := time.Now()
	if s.client == nil {
		return nil, errors.New("s.client is nil")
	}
	resp, err := s.client.SynthesizeSpeech(ctx, req)
	if err != nil {
		s.log.Errorf("Synthesizer s.client.SynthesizeSpeech err:%v", err)
		return nil, errors.Wrap(err, "s.client.SynthesizeSpeech err")
	}
	s.log.Debugw("p.launchTime", launchTime.String(), "now", time.Now(), "tts_google_elapsed_time", time.Since(launchTime).Seconds(), "uint", "s")

	// logsnag
	logsnag.Event(ctx, logsnag.Event_API_TTS_GOOGLE_ElAPSED_TIME.
		SetRoomName(s.linkConfig.RoomName).
		SetUID(s.linkConfig.UUID).
		SetElapsedTime(float32(time.Since(launchTime).Seconds())).
		SetNotifyElapsedTime(3.00).
		SetMessage(text))

	return &SynthesizeSpeechResponse{AudioContent: resp.AudioContent}, nil
}

func (s *Synthesizer) Synthesize(ctx context.Context, text string, language *Language) (*SynthesizeSpeechResponse, error) {
	// if not set voiceID then use Google cloud tts
	voiceID := s.linkConfig.GetConfigVoiceID()
	if voiceID == "" {
		return s.SynthesizeGoogle(ctx, text, language)
	}

	// English only in third-party voice
	if language.Code != "en-US" {
		return s.SynthesizeGoogle(ctx, text, language)
	}

	s.log.WithContext(ctx).Debugw("s.VoiceID", voiceID, "Synthesize.voiceID", voiceID)
	launchTime := time.Now()

	// get tts voice byte data
	data, err := s.elevenlabsClient.TTS(ctx, text, voiceID, "eleven_monolingual_v1", types.SynthesisOptions{Stability: 0.75, SimilarityBoost: 0.75})
	if err != nil {
		// logsnag
		logsnag.Event(ctx, logsnag.Event_API_TTS_ELEVENLABS_ERROR.
			SetRoomName(s.linkConfig.RoomName).
			SetUID(s.linkConfig.UUID).
			SetNotify().
			SetError(err.Error()).
			SetMessage(text))
		s.log.WithContext(ctx).Errorf("Synthesize.client.TTS err:%v", err)
		return nil, errors.Wrap(err, "Synthesize.client.TTS err")
	}

	// data to ogg data byte
	audioByte, err := s.FfmpegToOggData(ctx, data)
	if err != nil {
		return nil, errors.Wrap(err, "ffmpeg to ogg data err")
	}

	s.log.Debugw("p.launchTime", launchTime.String(), "now", time.Now(), "tts_elevenlabs_elapsed_time", time.Since(launchTime).Seconds(), "uint", "s")

	// logsnag
	logsnag.Event(ctx, logsnag.Event_API_TTS_ELEVENLABS_ElAPSED_TIME.
		SetRoomName(s.linkConfig.RoomName).
		SetUID(s.linkConfig.UUID).
		SetElapsedTime(float32(time.Since(launchTime).Seconds())).
		SetNotifyElapsedTime(3.00).
		SetMessage(text))

	return &SynthesizeSpeechResponse{AudioContent: audioByte}, nil
}

func (s *Synthesizer) FfmpegToOggData(ctx context.Context, data []byte) ([]byte, error) {
	s.log.WithContext(ctx).Info("Synthesize exec.Command start...")
	// tts data transfer ogg opus byte
	// ffmpeg -f mp3 -i pipe:0 -c:a libopus -ar 48000 -f ogg pipe:1
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-c:a", "libopus", "-ar", "48000", "-f", "ogg", "pipe:1")

	stdin, err := cmd.StdinPipe()
	defer stdin.Close()
	if err != nil {
		s.log.WithContext(ctx).Errorf("cmd.StdinPipe err:%v", err)
		return nil, errors.Wrap(err, "cmd stdin err")
	}
	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()
	if err != nil {
		s.log.WithContext(ctx).Errorf("cmd.StdoutPipe err:%v", err)
		return nil, errors.Wrap(err, "cmd stdout err")
	}

	// set stdin
	s.populateStdin(data)(stdin)

	err = cmd.Start()
	if err != nil {
		s.log.WithContext(ctx).Errorf("cmd.Start err:%v", err)
		return nil, errors.Wrap(err, "cmd start err")
	}

	byteBuf := bytes.Buffer{}
	done := make(chan bool)

	go func() {
		defer func() {
			s.log.WithContext(ctx).Info("Synthesize done <- true")
			done <- true
		}()
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if n > 0 {
				byteBuf.Write(buf[:n])
			}
			if err != nil {
				if err == io.EOF {
					s.log.WithContext(ctx).Infof("stdout.Read.err == io.EOF")
					break
				}
				if strings.Contains(err.Error(), " file already closed") {
					break
				}
				s.log.WithContext(ctx).Errorf("stdout.Read.err:%v", err)
			}
		}
	}()

	// 等待命令执行完成
	if err := cmd.Wait(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			s.log.WithContext(ctx).Infof("cmd failed with error:%v", string(exitErr.Stderr))
		} else {
			s.log.WithContext(ctx).Errorf("Failed to wait for cat command, err:%v", err)
		}
		return nil, errors.Wrap(err, "cmd wait err")
	}

	// wait stdout write over
	s.log.WithContext(ctx).Info("Synthesize.Done")
	<-done

	return byteBuf.Bytes(), nil
}

func (s *Synthesizer) populateStdin(file []byte) func(io.WriteCloser) {
	return func(stdin io.WriteCloser) {
		defer stdin.Close()
		io.Copy(stdin, bytes.NewReader(file))
	}
}
