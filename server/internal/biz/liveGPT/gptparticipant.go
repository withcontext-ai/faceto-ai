package liveGPT

import (
	"context"
	"encoding/json"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/conf"
	"faceto-ai/internal/pkg/utils/helper"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	stt "cloud.google.com/go/speech/apiv1"
	tts "cloud.google.com/go/texttospeech/apiv1"
	"github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/pion/webrtc/v3"
	"github.com/sashabaranov/go-openai"
)

// GPTParticipant GPT Participant Service
type GPTParticipant struct {
	ctx    context.Context
	cancel context.CancelFunc
	log    *log.Helper

	room         *lksdk.Room
	sttClient    *stt.Client
	ttsClient    *tts.Client
	gptClient    *openai.Client
	confThirdApi *conf.ThirdApi

	gptTrack *GPTTrack

	transcribers map[string]*Transcriber
	synthesizer  *Synthesizer
	completion   ChatCompletionAPI

	lock           sync.Mutex
	onDisconnected func()
	events         []*MeetingEvent
	language       *Language
	roomConfig     *biz.Link
	linkUC         *biz.LinkUseCase
	roomUC         *biz.RoomUseCase
	roomMsgUC      *biz.RoomMessageUseCase

	stopSpeak    chan struct{}
	stopAndLeave chan struct{}

	// Current active participant
	isBusy            atomic.Bool
	isStop            atomic.Bool
	activeInterim     atomic.Bool // True when KITT has been activated using an interim result
	activeId          uint64
	activeParticipant *lksdk.RemoteParticipant // If set, answer his next sentence/question
	subBeginTime      time.Time
	lastActivity      time.Time
	launchTime        time.Time
}

func ConnectGPTParticipant(
	url, token string,
	sttClient *stt.Client, ttsClient *tts.Client, gptClient *openai.Client,
	logger *log.Helper,
	confThirdApi *conf.ThirdApi,
	linkConfig *biz.Link,
	linkUC *biz.LinkUseCase,
	roomUC *biz.RoomUseCase,
	roomMsgUC *biz.RoomMessageUseCase,
) (*GPTParticipant, error) {
	ctx, cancel := context.WithCancel(context.Background())
	log := logger.WithContext(ctx)

	p := &GPTParticipant{
		ctx:    ctx,
		cancel: cancel,
		log:    log,

		sttClient:    sttClient,
		ttsClient:    ttsClient,
		gptClient:    gptClient,
		confThirdApi: confThirdApi,

		roomConfig: linkConfig,
		linkUC:     linkUC,
		roomUC:     roomUC,
		roomMsgUC:  roomMsgUC,

		transcribers: make(map[string]*Transcriber),
		stopSpeak:    make(chan struct{}),
		stopAndLeave: make(chan struct{}),

		synthesizer: NewSynthesizer(ttsClient, log, confThirdApi, linkConfig),
		completion:  NewChatCompletion(gptClient, log),
	}
	if linkConfig.ChatAPI != "" {
		p.completion = NewChatWithAPI(gptClient, log, linkConfig)
	}

	BotIdentity = "KITT"
	if linkConfig.GetConfigBotName() != "" {
		BotIdentity = linkConfig.GetConfigBotName()
	}

	roomCallback := &lksdk.RoomCallback{
		ParticipantCallback: lksdk.ParticipantCallback{
			OnTrackPublished:    p.trackPublished,
			OnTrackSubscribed:   p.trackSubscribed,
			OnTrackUnsubscribed: p.trackUnsubscribed,
		},
		OnParticipantDisconnected: p.participantDisconnected,
		OnDisconnected:            p.disconnected,
		OnActiveSpeakersChanged:   p.OnActiveSpeakersChanged,
	}

	room, err := lksdk.ConnectToRoomWithToken(url, token, roomCallback, lksdk.WithAutoSubscribe(false))
	if err != nil {
		return nil, err
	}

	track, err := NewGPTTrack(p.log)
	if err != nil {
		return nil, err
	}

	_, err = track.Publish(room.LocalParticipant)
	if err != nil {
		return nil, err
	}

	p.gptTrack = track
	p.room = room

	go func() {
		// Check if there's no participant when KITT joins.
		// It can happen when the participant who created the room directly leaves.
		time.Sleep(5 * time.Second)
		if len(room.GetParticipants()) == 0 {
			p.Disconnect()
		}
	}()

	return p, nil
}

// SetSynthesizerVoiceID Set AI VoiceID
func (p *GPTParticipant) SetSynthesizerVoiceID(ctx context.Context, roomName, voiceID string) error {
	if p.synthesizer != nil {
		// set link config
		if p.synthesizer.SetVoiceID(voiceID) {
			// update db room config
			return p.linkUC.SetRoomVoiceID(ctx, roomName, voiceID)
		}
		p.log.Infof("p.synthesizer.SetVoiceID false, roomName:%s, voiceID:%s", roomName, voiceID)
	}
	return nil
}

// OnActiveSpeakersChanged
func (p *GPTParticipant) OnActiveSpeakersChanged(participant []lksdk.Participant) {

}

func (p *GPTParticipant) OnDisconnected(f func()) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.onDisconnected = f
}

func (p *GPTParticipant) Disconnect() {
	p.log.Debugf("disconnecting gpt participant, room:%s", p.room.Name())
	p.room.Disconnect()
	p.roomConfig = nil

	for _, transcriber := range p.transcribers {
		transcriber.Close()
	}

	p.cancel()
	close(p.stopSpeak)

	p.lock.Lock()
	onDisconnected := p.onDisconnected
	p.lock.Unlock()

	if onDisconnected != nil {
		onDisconnected()
	}
}

func (p *GPTParticipant) trackPublished(publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	if publication.Source() != livekit.TrackSource_MICROPHONE {
		return
	}

	err := publication.SetSubscribed(true)
	if err != nil {
		p.log.Errorf("failed to subscribe to the track, err:%v, track:%s, participant:%s", err, publication.SID(), rp.SID())
		return
	}
}

func (p *GPTParticipant) trackSubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if _, ok := p.transcribers[rp.SID()]; ok {
		return
	}

	metadata := ParticipantMetadata{}
	if rp.Metadata() != "" {
		err := json.Unmarshal([]byte(rp.Metadata()), &metadata)
		if err != nil {
			p.log.Warnf("error unmarshalling participant metadata, err:%v, track:%s, participant:%s", err, publication.SID(), rp.SID())
		}
	}

	language, ok := Languages[metadata.LanguageCode]
	if !ok {
		language = DefaultLanguage
	}
	p.language = language

	p.log.Debugf("starting to transcribe, participant:%s, language:%s", rp.Identity(), language.Code)
	transcriber, err := NewTranscriber(track.Codec(), p.sttClient, language, p.log)
	if err != nil {
		p.log.Errorf("failed to create the transcriber, err:%v", err)
		return
	}
	p.subBeginTime = time.Now()

	p.transcribers[rp.SID()] = transcriber

	// consumer define func
	p.gptTrack.OnConsumer(func(text string) {
		_ = p.sendStatePacket(state_Active)
		_ = p.sendPacket(&packet{
			Type: packet_Transcript,
			Data: &transcriptPacket{
				Sid:     rp.SID(),
				Name:    BotIdentity,
				Text:    text,
				IsFinal: false,
			},
		})
		//p.sendPacketByWord(text, rp.SID(), time.Duration(2000))
	})

	// on break function
	p.gptTrack.OnBreak(func(text string) {
		// last AI answer
		lastAnswer := p.events[len(p.events)-1]
		if !lastAnswer.Speech.IsBot {
			return
		}
		if strings.Contains(lastAnswer.Speech.Text, "---") {
			return
		}

		p.lock.Lock()
		// split by text
		p.events = p.events[:len(p.events)-1]
		botAnswer := &SpeechEvent{
			ParticipantName: BotIdentity,
			IsBot:           true,
			Text:            text + " ---",
			Timestamp:       lastAnswer.Speech.Timestamp,
		}
		p.events = append(p.events, &MeetingEvent{
			Speech: botAnswer,
		})
		p.lock.Unlock()
	})

	go func() {
		for _ = range transcriber.quitSpeak {
			p.log.Debugf("<-------------- transcriber.quitSpeak -------------->")
			p.isBusy.Store(false)
			p.isStop.Store(true)
			p.gptTrack.OnStop(true) // stop read queue data -> gpttrack.go line 133
			transcriber.isSpeaking.Store(false)
		}
	}()

	go func(ctx context.Context) {
		defer helper.RecoverPanic(p.ctx, p.log, "onTranscriptionReceived Panic")
		// time queue
		tq := NewTimeQueue(3*time.Second, &QueueCallback{TimeTicker: func(text string) {
			if text == "" {
				return
			}
			p.launchTime = time.Now()
			_ = p.sendStatePacket(state_Loading) // add loading state
			p.onTranscriptionReceived(ctx, RecognizeResult{
				Text:    text,
				IsFinal: true,
			}, rp, transcriber)
		}})
		defer tq.Close()

		for result := range transcriber.Results() {
			if result.Error != nil {
				_ = p.sendErrorPacket(fmt.Sprintf("Sorry, an error occured while transcribing %s's speech using Google STT", rp.Identity()))
				return
			}
			p.log.Debugf("onTranscriptionReceived, sid:%s, text:%s, IsFinal:%v", rp.SID(), result.Text, result.IsFinal)

			_ = p.sendPacket(&packet{
				Type: packet_Transcript,
				Data: &transcriptPacket{
					Sid:     rp.SID(),
					Name:    rp.Name(),
					Text:    result.Text,
					IsFinal: result.IsFinal,
				},
			})

			tq.SetLastTime()
			if result.IsFinal {
				tq.Append(result.Text, false)
			}

			p.lastActivity = time.Now()
		}
	}(p.ctx)

	// transcriber time ticker
	go p.transcriberTimeTicker(track, publication, rp, transcriber)

	go func() {
		for _ = range p.stopAndLeave {
			p.trackUnsubscribed(track, publication, rp)
			p.disconnected()
		}
	}()

	// Forward track packets to the transcriber
	go func() {
		for {
			pkt, _, err := track.ReadRTP()
			if err != nil {
				if err != io.EOF {
					p.log.Errorf("failed to read track, err:%v, participant:%s", err, rp.SID())
				}
				return
			}

			err = transcriber.WriteRTP(pkt)
			if err != nil {
				if err != io.EOF {
					p.log.Errorf("failed to forward pkt to the transcriber, err:%v, participant:%s", err, rp.SID())
				}
				return
			}
		}
	}()
}

func (p *GPTParticipant) trackUnsubscribed(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant) {
	p.lock.Lock()
	if transcriber, ok := p.transcribers[rp.SID()]; ok {
		p.log.Debugf("trackUnsubscribed, participant:%s", rp.SID())
		p.lock.Unlock()
		transcriber.Close()
		p.lock.Lock()
		delete(p.transcribers, rp.SID())
		p.log.Debugf("delete(p.transcribers), participant:%s", rp.SID())
	}
	p.lock.Unlock()
}

func (p *GPTParticipant) participantDisconnected(rp *lksdk.RemoteParticipant) {
	participants := p.room.GetParticipants()
	p.log.Debugf("participantDisconnected, numParticipants:%d", len(participants))
	if len(participants) == 0 {
		p.Disconnect()
	}
}

func (p *GPTParticipant) disconnected() {
	p.Disconnect()
}

// In a multi-user meeting, the bot will only answer when it is activated.
// Activate the participant rp
func (p *GPTParticipant) activateParticipant(rp *lksdk.RemoteParticipant) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.activeParticipant != rp {
		p.activeId++
		p.activeParticipant = rp
		p.lastActivity = time.Now()
		_ = p.sendStatePacket(state_Active)

		tmpActiveId := p.activeId
		go func() {
			time.Sleep(ActivationTimeout)
			for {
				p.lock.Lock()
				if p.activeId != tmpActiveId {
					p.lock.Unlock()
					return
				}

				if time.Since(p.lastActivity) >= ActivationTimeout {
					p.activeParticipant = nil
					_ = p.sendStatePacket(state_Idle)
					p.lock.Unlock()
					return
				}

				p.lock.Unlock()
				time.Sleep(1 * time.Second)
			}
		}()
	}
}

func (p *GPTParticipant) sendPacket(packet *packet) error {
	data, err := json.Marshal(packet)
	if err != nil {
		return err
	}
	return p.room.LocalParticipant.PublishData(data, livekit.DataPacket_RELIABLE, []string{})
}

func (p *GPTParticipant) sendStatePacket(state gptState) error {
	return p.sendPacket(&packet{
		Type: packet_State,
		Data: &statePacket{
			State: state,
		},
	})
}

func (p *GPTParticipant) sendErrorPacket(message string) error {
	return p.sendPacket(&packet{
		Type: packet_Error,
		Data: &errorPacket{
			Message: message,
		},
	})
}

func (p *GPTParticipant) GetEvents() []*MeetingEvent {
	return p.events
}

func (p *GPTParticipant) RoomEvent() []*MeetingEvent {
	return p.events
}

func (p *GPTParticipant) SendStopEvent(ctx context.Context) error {
	if err := p.sendPacket(&packet{
		Type: packet_EventStopRoom,
		Data: &eventPacket{
			Sid:   p.room.SID(),
			Name:  p.room.Name(),
			Event: "CloseRoom",
			Text:  "Video Interaction Over.",
		},
	}); err != nil {
		p.log.WithContext(ctx).Errorf("SendStopEvent p.sendPacket.err:%v", err)
		return err
	}
	p.stopAndLeave <- struct{}{}
	p.log.WithContext(ctx).Debugf("SendStopEvent roomName:%s, p.stopAndLeave push...", p.room.Name())
	return nil
}
