// Package liveGPT
// ---------------------------------
// @file      : synthesizer_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/9/6 10:26
// @desc      : file description
// ---------------------------------
package liveGPT

import (
	"bytes"
	tts "cloud.google.com/go/texttospeech/apiv1"
	ttspb "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/data/schema"
	pkgLog "faceto-ai/internal/pkg/utils/log"
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/go-kratos/kratos/v2/log"
	c "github.com/smartystreets/goconvey/convey"
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/option"
	"io"
	"testing"
	"time"
)

var (
	ttsClient *tts.Client
	thirdConf *conf.ThirdApi
	ctx       context.Context
)

func init() {
	ctx = context.Background()
	bcConf := conf.NewConf()
	// third conf
	thirdConf = bcConf.ThirdApi

	// tts client
	gcpBytes, err := conf.GetGcpCredBytes(bcConf.GcpCredentials.GetPath())
	if err != nil {
		return
	}
	gcpCred := option.WithCredentialsJSON(gcpBytes)
	ttsc, err := tts.NewClient(ctx, gcpCred)
	if err != nil {
		fmt.Println("tts.NewClient err", err)
		return
	}
	ttsClient = ttsc
}

func Test_SynthesizeGoogle(t *testing.T) {
	req := &ttspb.SynthesizeSpeechRequest{
		Input: &ttspb.SynthesisInput{
			InputSource: &ttspb.SynthesisInput_Text{
				Text: "test",
			},
		},
		Voice: &ttspb.VoiceSelectionParams{
			LanguageCode: "en-US",
			Name:         "en-US-Wavenet-D",
		},
		AudioConfig: &ttspb.AudioConfig{
			AudioEncoding:   ttspb.AudioEncoding_OGG_OPUS,
			SampleRateHertz: 48000,
		},
	}
	if ttsClient == nil {
		fmt.Println("ttsClient nil")
		return
	}
	resp, err := ttsClient.SynthesizeSpeech(ctx, req)
	if err != nil {
		fmt.Println("ttsClient.SynthesizeSpeech err", err)
		return
	}
	fmt.Println(resp.String())
	fmt.Println("ttsClient.SynthesizeSpeech", len(resp.AudioContent))
}

func TestSynthesizer_SetVoiceID(t *testing.T) {
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)

	linkConfig := &biz.Link{
		UUID:     "test",
		RoomName: "test",
		Config: &schema.RoomConfig{
			Duration: 600,
			Greeting: "",
			VoiceID:  "",
			UserName: "",
			BotName:  "",
		},
	}

	if ttsClient == nil {
		return
	}
	ns := NewSynthesizer(ttsClient, log.NewHelper(logger), thirdConf, linkConfig)

	c.Convey("Synthesizer SetVoiceID", t, func() {
		got1 := ns.SetVoiceID("aa")
		got2 := ns.SetVoiceID("8QAi78THegBm75BpJ4f5")

		c.So(got1, c.ShouldEqual, false) // 断言
		c.So(got2, c.ShouldEqual, true)  // 断言
	})
}

func TestSynthesizer_SynthesizeGoogle(t *testing.T) {
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)

	if ttsClient == nil {
		return
	}

	type fields struct {
		client     *tts.Client
		log        *log.Helper
		thirdConf  *conf.ThirdApi
		linkConfig *biz.Link
	}
	type args struct {
		ctx      context.Context
		text     string
		language *Language
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "google tts",
			fields: fields{
				client:    ttsClient,
				log:       log.NewHelper(logger),
				thirdConf: thirdConf,
				linkConfig: &biz.Link{
					UUID:     "test",
					RoomName: "test",
					Config: &schema.RoomConfig{
						Duration: 600,
						Greeting: "",
						VoiceID:  "",
						UserName: "",
						BotName:  "",
					},
				},
			},
			args: args{
				ctx:      ctx,
				text:     "hello, world!",
				language: DefaultLanguage,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSynthesizer(tt.fields.client, tt.fields.log, tt.fields.thirdConf, tt.fields.linkConfig)
			resp, err := s.SynthesizeGoogle(tt.args.ctx, tt.args.text, tt.args.language)
			if err != nil {
				t.Errorf("SynthesizeGoogle() error = %v", err)
				return
			}
			fmt.Println(len(resp.AudioContent))
		})
	}
}

func TestSynthesizer_SynthesizeElevenlabs(t *testing.T) {
	logger := pkgLog.InitProductStdLogger(zapcore.DebugLevel)
	if ttsClient == nil {
		return
	}

	type fields struct {
		client     *tts.Client
		log        *log.Helper
		thirdConf  *conf.ThirdApi
		linkConfig *biz.Link
	}
	type args struct {
		ctx      context.Context
		text     string
		language *Language
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "elevenlabs tts",
			fields: fields{
				client:    ttsClient,
				log:       log.NewHelper(logger),
				thirdConf: thirdConf,
				linkConfig: &biz.Link{
					UUID:     "test",
					RoomName: "test",
					Config: &schema.RoomConfig{
						Duration: 600,
						Greeting: "Hello, I am Elon Musk",
						VoiceID:  "8QAi78THegBm75BpJ4f5",
						UserName: "test",
						BotName:  "Elon Musk",
					},
				},
			},
			args: args{
				ctx:      ctx,
				text:     "hello, world!",
				language: DefaultLanguage,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSynthesizer(tt.fields.client, tt.fields.log, tt.fields.thirdConf, tt.fields.linkConfig)
			resp, err := s.Synthesize(tt.args.ctx, tt.args.text, tt.args.language)
			if err != nil {
				t.Errorf("SynthesizeGoogle() error = %v", err)
				return
			}
			fmt.Println(len(resp.AudioContent))
		})
	}
}

func beepPlay(data []byte) {

	readCloser := io.NopCloser(bytes.NewReader(data))

	// pipe reader
	fmt.Println("pipe reader...")
	streamer, format, err := wav.Decode(readCloser)
	if err != nil {
		fmt.Println("mp3.Decode, err", err)
		return
	}
	defer streamer.Close()

	_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
