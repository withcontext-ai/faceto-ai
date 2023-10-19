package liveGPT

import (
	"bytes"
	"context"
	"faceto-ai/internal/data/schema"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/pion/webrtc/v3"
	"io"
	"sync"
	"time"
)

// transcriberTimeTicker set time ticker
func (p *GPTParticipant) transcriberTimeTicker(track *webrtc.TrackRemote, publication *lksdk.RemoteTrackPublication, rp *lksdk.RemoteParticipant, transcriber *Transcriber) {
	languageText, ok := NoticeText[p.language.Code]
	if !ok {
		languageText = NoticeText[DefaultLanguage.Code]
	}

	// init api config
	apiConfig := &schema.RoomConfig{
		Duration: p.roomConfig.GetConfigDuration(),
		Greeting: p.roomConfig.GetConfigGreeting(),
	}
	if apiConfig.Greeting != "" {
		languageText["greeting"] = apiConfig.Greeting
	}

	limitDuration := schema.Duration
	if apiConfig.Duration > 0 {
		limitDuration = int(apiConfig.Duration)
	}

	// time ticker
	for _ = range transcriber.timeTicker.C {
		// the conversation limit duration
		switch int(time.Since(p.subBeginTime).Seconds()) {
		case limitDuration - 30: // Give a 30-second warning
			p.initiativeSpeak(p.ctx, &InitiativeSpeak{
				Sid:      rp.SID(),
				Text:     "This video conversation will end in 30 seconds",
				Language: p.language,
			}, nil)
		case limitDuration:
			p.initiativeSpeak(p.ctx, &InitiativeSpeak{
				Sid:      rp.SID(),
				Text:     languageText["end"],
				Language: p.language,
			}, func() {
				go func() {
					time.Sleep(3 * time.Second)
					p.trackUnsubscribed(track, publication, rp)
					p.disconnected()
				}()
			})
		}

		if p.lastActivity.IsZero() {
			switch int(time.Since(p.subBeginTime).Seconds()) {
			case 3:
				// Remove AI greeting
				//greeting := &SpeechEvent{
				//	ParticipantName: BotIdentity,
				//	IsBot:           true,
				//	Text:            languageText["greeting"],
				//	Timestamp:       uint64(time.Now().Unix()),
				//}
				//p.lock.Lock()
				//p.events = append(p.events, &MeetingEvent{
				//	Speech: greeting,
				//})
				//p.lock.Unlock()
				//
				//// record user msg
				//userRoomMessage := &biz.RoomMessage{
				//	Sid:             p.room.SID(),
				//	ParticipantSID:  rp.SID(),
				//	ParticipantName: BotIdentity,
				//	Type:            biz.RoomMessageTypeBot,
				//	EventTime:       time.Now(),
				//	Text:            languageText["greeting"],
				//}
				//if err := p.roomMsgUC.Record(context.Background(), userRoomMessage); err != nil {
				//	p.log.Errorf("failed to record user msg, err:%v, roomSID:%s, text:%s", err, p.room.SID(), languageText["greeting"])
				//}
				//
				//p.initiativeSpeak(p.ctx, &InitiativeSpeak{
				//	Sid:      rp.SID(),
				//	Text:     languageText["greeting"],
				//	Language: p.language,
				//}, nil)
			case 30:
				p.initiativeSpeak(p.ctx, &InitiativeSpeak{
					Sid:      rp.SID(),
					Text:     languageText["notice"],
					Language: p.language,
				}, nil)
			case 60:
				p.initiativeSpeak(p.ctx, &InitiativeSpeak{
					Sid:      rp.SID(),
					Text:     languageText["end"],
					Language: p.language,
				}, func() {
					go func() {
						time.Sleep(3 * time.Second)
						p.trackUnsubscribed(track, publication, rp)
						p.disconnected()
					}()
				})
			}
			continue
		}

		switch int(time.Since(p.lastActivity).Seconds()) {
		case 30:
			p.initiativeSpeak(p.ctx, &InitiativeSpeak{
				Sid:      rp.SID(),
				Text:     languageText["notice"],
				Language: p.language,
			}, nil)
		case 60:
			p.initiativeSpeak(p.ctx, &InitiativeSpeak{
				Sid:      rp.SID(),
				Text:     languageText["end"],
				Language: p.language,
			}, func() {
				go func() {
					time.Sleep(3 * time.Second)
					p.trackUnsubscribed(track, publication, rp)
					p.disconnected()
				}()
			})
		}
	}
}

func (p *GPTParticipant) initiativeSpeak(
	ctx context.Context,
	initSpeak *InitiativeSpeak,
	f func(),
) {
	var sg sync.WaitGroup

	_ = p.sendStatePacket(state_Loading)

	p.gptTrack.OnComplete(func(err error) {
		if f != nil && err == io.EOF {
			f()
		}
	})
	p.gptTrack.OnStop(false)
	p.gptTrack.OnSpeaking(func(data []byte, err error) {

	})

	sg.Add(1)
	go func() {
		defer sg.Done()

		resp, err := p.synthesizer.Synthesize(p.ctx, initSpeak.Text, initSpeak.Language)
		if err != nil {
			p.log.Errorf("initiativeSpeak, failed to synthesize, err:%v, sentence:%s", err, initSpeak.Text)
			_ = p.sendErrorPacket("Sorry, an error occured while synthesizing voice data using Google TTS")
			return
		}

		p.log.Debugf("initiativeSpeak, finished synthesizing, queuing sentence, text:%s", initSpeak.Text)
		_ = p.gptTrack.QueueReaderString(initSpeak.Text)
		err = p.gptTrack.QueueReader(bytes.NewReader(resp.AudioContent))
		if err != nil {
			p.log.Errorf("initiativeSpeak, failed to queue reader, err:%v, sentence:%s", err, initSpeak.Text)
			return
		}

		return
	}()

	sg.Wait()
	_ = p.sendStatePacket(state_Idle)
	p.log.Debugf("p.sendStatePacket(state_Idle)")
	return
}
