package liveGPT

import (
	stt "cloud.google.com/go/speech/apiv1"
	sttpb "cloud.google.com/go/speech/apiv1/speechpb"
	"context"
	"errors"
	"faceto-ai/internal/pkg/middleware"
	"faceto-ai/internal/pkg/utils/helper"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/livekit/protocol/logger"
	"github.com/pion/rtp"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media/oggwriter"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type Transcriber struct {
	ctx    context.Context
	cancel context.CancelFunc
	log    *log.Helper

	speechClient *stt.Client
	language     *Language

	rtpCodec webrtc.RTPCodecParameters
	//sb       *samplebuilder.SampleBuilder

	lock          sync.Mutex
	oggWriter     *io.PipeWriter
	oggReader     *io.PipeReader
	oggSerializer *oggwriter.OggWriter

	results   chan RecognizeResult
	quitSpeak chan struct{}
	closeCh   chan struct{}

	timeTicker *time.Ticker
	isSpeaking atomic.Bool
}

type RecognizeResult struct {
	Error   error
	Text    string
	IsFinal bool
}

func NewTranscriber(rtpCodec webrtc.RTPCodecParameters, speechClient *stt.Client, language *Language, logger *log.Helper) (*Transcriber, error) {
	if !strings.EqualFold(rtpCodec.MimeType, "audio/opus") {
		return nil, errors.New("only opus is supported")
	}
	logger.Debugf("NewTranscriber start, rtpCodec.MimeType:%s, rtpCodec.ClockRate:%d, rtpCodec.Channels:%d", rtpCodec.MimeType, rtpCodec.ClockRate, rtpCodec.Channels)

	oggReader, oggWriter := io.Pipe()
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, middleware.TraceID, helper.Generator())

	t := &Transcriber{
		ctx:      ctx,
		cancel:   cancel,
		rtpCodec: rtpCodec,
		//sb:           samplebuilder.New(200, &codecs.OpusPacket{}, rtpCodec.ClockRate),
		oggReader:    oggReader,
		oggWriter:    oggWriter,
		language:     language,
		speechClient: speechClient,
		results:      make(chan RecognizeResult),
		quitSpeak:    make(chan struct{}),
		closeCh:      make(chan struct{}),
		timeTicker:   time.NewTicker(1 * time.Second),
		log:          logger,
	}

	go t.start()
	//go func() {
	//	time.Sleep(5 * time.Second)
	//	t.results <- RecognizeResult{
	//		Text:    "who are you?",
	//		IsFinal: true,
	//	}
	//	fmt.Println("============= who are you?")
	//
	//	timeTicker := time.NewTicker(time.Second * 15)
	//	startTime := time.Now()
	//	for {
	//		select {
	//		case <-timeTicker.C:
	//			if time.Since(startTime).Seconds() >= 45 {
	//				fmt.Println("============= timeTicker.Stop()")
	//				timeTicker.Stop()
	//				return
	//			}
	//			if t.isSpeaking.Load() {
	//				t.StopSpeak()
	//				time.Sleep(50 * time.Millisecond)
	//			}
	//
	//			t.results <- RecognizeResult{
	//				Text:    "please tell me a another joke",
	//				IsFinal: true,
	//			}
	//			fmt.Println("============= please tell me a joke")
	//		}
	//	}
	//}()
	return t, nil
}

func (t *Transcriber) Language() *Language {
	return t.language
}

func (t *Transcriber) WriteRTP(pkt *rtp.Packet) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	if t.oggSerializer == nil {
		oggSerializer, err := oggwriter.NewWith(t.oggWriter, t.rtpCodec.ClockRate, t.rtpCodec.Channels)
		if err != nil {
			logger.Errorw("failed to create ogg serializer", err)
			return err
		}
		t.oggSerializer = oggSerializer
	}

	//t.sb.Push(pkt)
	//for _, p := range t.sb.PopPackets() {
	if err := t.oggSerializer.WriteRTP(pkt); err != nil {
		return err
	}
	//}

	return nil
}

func (t *Transcriber) start() error {
	defer func() {
		close(t.closeCh)
	}()

	for {
		stream, err := t.newStream()
		if err != nil {
			if status, ok := status.FromError(err); ok && status.Code() == codes.Canceled {
				t.log.Errorf("start.t.newStream() Cancel, err:%v", err)
				return nil
			}

			t.log.Errorf("failed to create a new speech stream, err:%v", err)
			t.StopSpeak()
			t.results <- RecognizeResult{
				Error: err,
			}
			return err
		}

		endStreamCh := make(chan struct{})
		nextCh := make(chan struct{})

		// Forward oggreader to the speech stream
		go func() {
			defer close(nextCh)
			buf := make([]byte, 1024)

			for {
				select {
				case <-endStreamCh:
					t.log.Debugf("<------------- <-endStreamCh")
					return
				default:
					n, err := t.oggReader.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.log.Errorf("failed to read from ogg reader, err:%v", err)
						}
						return
					}

					if n <= 0 {
						continue // No data
					}

					if err := stream.Send(&sttpb.StreamingRecognizeRequest{
						StreamingRequest: &sttpb.StreamingRecognizeRequest_AudioContent{
							AudioContent: buf[:n],
						},
					}); err != nil {
						if stream.Context().Err() != nil {
							t.log.Errorf("stream.Context().Err(), err:%v", stream.Context().Err())
							return
						}
						if err != io.EOF {
							t.log.Errorf("failed to forward audio data to speech stream, err:%v", err)
							t.StopSpeak()
							t.results <- RecognizeResult{
								Error: err,
							}
						}
						return
					}
					//time.Sleep(time.Millisecond)
				}
			}
		}()

		// Read transcription results
		for {
			resp, err := stream.Recv()
			if err != nil {
				t.log.Errorf("Transcriber.stream.Recv(), err:%v", err)
				if status, ok := status.FromError(err); ok {
					if status.Code() == codes.OutOfRange {
						t.log.Debugf("stream.Recv().err.codes.OutOfRange")
						break // Create a new speech stream (maximum speech length exceeded)
					} else if status.Code() == codes.Canceled {
						t.log.Debugf("stream.Recv().err.codes.Canceled")
						return nil // Context canceled (Stop)
					}
				}

				t.log.Errorf("failed to receive response from speech stream, err:%v", err)
				t.StopSpeak()
				t.results <- RecognizeResult{
					Error: err,
				}

				return err
			}

			if resp.Error != nil {
				t.log.Errorf("transcriber.stream.Recv(), err:%v", errors.New(resp.Error.GetMessage()))
				break
			}

			//t.log.Debugw("transcriber.stream.Recv()", resp)

			// Read the whole transcription and put inside one string
			// We don't need to process each part individually (atm?)
			var sb strings.Builder
			final := false
			for _, result := range resp.Results {
				alt := result.Alternatives[0]
				text := alt.Transcript
				sb.WriteString(text)

				if result.IsFinal {
					sb.Reset()
					sb.WriteString(text)
					final = true
					break
				}
			}

			// quit last msg
			if t.isSpeaking.Load() {
				t.StopSpeak()
				time.Sleep(50 * time.Millisecond)
			}

			if sb.String() != "" {
				t.results <- RecognizeResult{
					Text:    sb.String(),
					IsFinal: final,
				}
			}
		}

		close(endStreamCh)
		t.log.Debugf("close(endStreamCh)")

		// When nothing is written on the transcriber (The track is muted), this will block because the oggReader is waiting for data.
		// It avoids to create useless speech streams.
		// (Also we end up here because Google automatically close the previous stream when there's no "activity")
		//
		// Otherwise (When we have data) it is used to wait for the end of the current stream,
		// so we can create the next one and reset the oggSerializer
		<-nextCh

		// Create a new oggSerializer each time we open a new SpeechStream
		// This is required because the stream requires ogg headers to be sent again
		t.lock.Lock()
		t.oggSerializer = nil
		t.log.Debugf("t.oggSerializer = nil")
		t.lock.Unlock()
	}
}

func (t *Transcriber) StopSpeak() {
	t.log.Debugf("<-------------- t.quitSpeak <- struct{}{} -------------->")
	t.quitSpeak <- struct{}{}
}

func (t *Transcriber) Close() {
	logger.Debugw("Transcriber.Close()")
	t.timeTicker.Stop()
	t.cancel()
	t.oggReader.Close()
	t.oggWriter.Close()
	<-t.closeCh
	close(t.results)
	close(t.quitSpeak)
}

func (t *Transcriber) Results() <-chan RecognizeResult {
	return t.results
}

func (t *Transcriber) newStream() (sttpb.Speech_StreamingRecognizeClient, error) {
	stream, err := t.speechClient.StreamingRecognize(t.ctx)
	if err != nil {
		return nil, err
	}

	config := &sttpb.RecognitionConfig{
		Model:             "command_and_search",
		UseEnhanced:       true,
		ProfanityFilter:   true, // replacing all but the initial character in each filtered word with asterisks, e.g. "f***".
		Encoding:          sttpb.RecognitionConfig_OGG_OPUS,
		SampleRateHertz:   int32(t.rtpCodec.ClockRate),
		AudioChannelCount: int32(t.rtpCodec.Channels),
		LanguageCode:      t.language.TranscriberCode,
		//EnableAutomaticPunctuation: true,
		//Adaptation: &sttpb.SpeechAdaptation{
		//	PhraseSets: []*sttpb.PhraseSet{
		//		{
		//			Phrases: []*sttpb.PhraseSet_Phrase{
		//				{Value: "${hello} ${gpt}"},
		//				{Value: "${gpt}"},
		//				{Value: "Hey ${gpt}"},
		//				{Value: "Kitt"},
		//				{Value: "Kit-t"},
		//				{Value: "Kit"},
		//				{Value: "你好"},
		//			},
		//			Boost: 16,
		//		},
		//	},
		//	CustomClasses: []*sttpb.CustomClass{
		//		{
		//			CustomClassId: "hello",
		//			Items: []*sttpb.CustomClass_ClassItem{
		//				{Value: "Hi"},
		//				{Value: "Hello"},
		//				{Value: "Hey"},
		//				{Value: "你好"},
		//				{Value: "你在吗"},
		//				{Value: "在吗"},
		//			},
		//		},
		//		{
		//			CustomClassId: "gpt",
		//			Items: []*sttpb.CustomClass_ClassItem{
		//				{Value: "Kit"},
		//				{Value: "Kite"},
		//				{Value: "KITT"},
		//				{Value: "GPT"},
		//				{Value: "Live Kit"},
		//				{Value: "Live GPT"},
		//				{Value: "LiveKit"},
		//				{Value: "LiveGPT"},
		//				{Value: "Live-Kit"},
		//				{Value: "Live-GPT"},
		//			},
		//		},
		//	},
		//},
	}

	if err := stream.Send(&sttpb.StreamingRecognizeRequest{
		StreamingRequest: &sttpb.StreamingRecognizeRequest_StreamingConfig{
			// https://cloud.google.com/speech-to-text/docs/reference/rpc/google.cloud.speech.v1#google.cloud.speech.v1.StreamingRecognitionConfig
			StreamingConfig: &sttpb.StreamingRecognitionConfig{
				//SingleUtterance: true,
				InterimResults: true,
				Config:         config,
				//EnableVoiceActivityEvents: true,
				//VoiceActivityTimeout: &sttpb.StreamingRecognitionConfig_VoiceActivityTimeout{
				//SpeechStartTimeout: &durationpb.Duration{
				//	Seconds: 30,
				//	Nanos:   0,
				//},
				//SpeechEndTimeout: &durationpb.Duration{
				//	Seconds: 1,
				//	Nanos:   0,
				//},
				//},
			},
		},
	}); err != nil {
		return nil, err
	}

	return stream, nil
}
