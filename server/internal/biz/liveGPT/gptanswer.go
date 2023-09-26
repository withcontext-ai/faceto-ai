package liveGPT

import (
	"bytes"
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/biz/liveGPT/elevenlabs"
	"faceto-ai/internal/pkg/utils/helper"
	"faceto-ai/internal/pkg/utils/logsnag"
	"fmt"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/pkg/errors"
	"golang.org/x/exp/slices"
	"io"
	"strings"
	"sync"
	"time"
)

func (p *GPTParticipant) onTranscriptionReceived(ctx context.Context, result RecognizeResult, rp *lksdk.RemoteParticipant, transcriber *Transcriber) {

	p.gptTrack.OnComplete(func(err error) {
		//wg.Done()
	})

	// When there's only one participant in the meeting, no activation/trigger is needed
	// The bot will answer directly.
	//
	// When there are multiple participants, activation is required.
	// 1. Wait for activation sentence (Hey Kitt!)
	// 2. If the participant stop speaking after the activation, ignore the next "isFinal" result
	// 3. If activated, anwser the next sentence

	p.lock.Lock()
	activeParticipant := p.activeParticipant
	if activeParticipant == rp {
		p.lastActivity = time.Now()
	}
	p.lock.Unlock()

	shouldAnswer := false
	if len(p.room.GetParticipants()) == 1 {
		// Always answer when we're alone with KITT
		if activeParticipant == nil {
			activeParticipant = rp
			p.activateParticipant(rp)
		}

		shouldAnswer = result.IsFinal
	} else {
		// language case
		// Check if the participant is activating the KITT
		justActivated := false
		words := strings.Split(strings.ToLower(strings.TrimSpace(result.Text)), " ")
		if len(words) >= 2 { // No max length but only check the first 3 words
			limit := len(words)
			if limit > ActivationWordsLen {
				limit = ActivationWordsLen
			}
			activationWords := words[:limit]

			// Check if text contains at least one GreentingWords
			greetIndex := -1
			for _, greet := range GreetingWords {
				if greetIndex = slices.Index(activationWords, greet); greetIndex != -1 {
					break
				}
			}

			nameIndex := -1
			for _, name := range NameWords {
				if nameIndex = slices.Index(activationWords, name); nameIndex != -1 {
					break
				}
			}

			if greetIndex < nameIndex && greetIndex != -1 {
				justActivated = true
				p.activeInterim.Store(!result.IsFinal)
				if activeParticipant != rp {
					activeParticipant = rp
					p.log.Debugf("activating KITT for participant, activationText:%s, participant:%s", strings.Join(activationWords, " "), rp.Identity())
					p.activateParticipant(rp)
				}
			}

			p.log.Debugf("Check justActivated:%v, greetIndex:%d, nameIndex:%d", justActivated, greetIndex, nameIndex)
		}
		shouldAnswer = result.IsFinal
		//if result.IsFinal {
		//	shouldAnswer = activeParticipant == rp
		//	if (justActivated || p.activeInterim.Load()) && len(words) <= ActivationWordsLen+1 {
		//		// Ignore if the participant stopped speaking after the activation, answer his next sentence
		//		logger.Debugw("Ignore if the participant stopped speaking after the activation, answer his next sentence")
		//		shouldAnswer = false
		//	}
		//}
	}

	if shouldAnswer {
		prompt := &SpeechEvent{
			ParticipantName: rp.Name(),
			IsBot:           false,
			Text:            result.Text,
			Timestamp:       uint64(time.Now().Unix()),
		}

		var userSid, userIdentity string
		if p.activeParticipant != nil {
			userSid = p.activeParticipant.SID()
			userIdentity = p.activeParticipant.Name()
		} else {
			userSid = rp.SID()
			userIdentity = rp.Name()
		}

		p.lock.Lock()
		// Don't include the current prompt in the history when answering
		events := make([]*MeetingEvent, len(p.events))
		copy(events, p.events)
		p.events = append(p.events, &MeetingEvent{
			Speech: prompt,
		})
		p.activeParticipant = nil
		p.lock.Unlock()

		p.log.Debugf("answering before, shouldAnswer:%v, p.isBusy:%v", shouldAnswer, p.isBusy.Load())
		if shouldAnswer && p.isBusy.CompareAndSwap(false, true) {
			go func(ctx context.Context) {
				defer func() {
					helper.RecoverPanic(p.ctx, p.log, "Answer Panic")
					p.log.Debugf("..................answer defer done..................")
					p.isBusy.Store(false)
					transcriber.isSpeaking.Store(false)
				}()

				_ = p.sendStatePacket(state_Loading)

				// record user msg
				userRoomMessage := &biz.RoomMessage{
					RoomName:        p.roomConfig.RoomName,
					Sid:             p.room.SID(),
					ParticipantSID:  userSid,
					ParticipantName: userIdentity,
					Type:            biz.RoomMessageTypeUser,
					EventTime:       time.Now(),
					Text:            result.Text,
				}
				if err := p.roomMsgUC.Record(ctx, userRoomMessage); err != nil {
					p.log.Errorf("failed to record user msg, err:%v, roomSID:%s, text:%s", err, p.room.SID(), result.Text)
				}

				p.log.Debugf("answering begin, participant:%s, text:%s", rp.SID(), result.Text)
				answer, err := p.answer(events, prompt, rp, transcriber) // Will send state_Speaking
				if err != nil {
					p.log.Errorf("failed to answer, err:%v, roomSID:%s, text:%s", err, p.room.SID(), result.Text)
					p.sendStatePacket(state_Idle)
					return
				}
				p.log.Debugf("answering over, roomSID:%s, answer:%s", p.room.SID(), answer)

				// KITT finished speaking, check if the last sentence was a question.
				// If so, auto activate the current participant
				//if strings.HasSuffix(answer, "?") {
				//	// Checking this suffix should be enough
				//	p.activateParticipant(rp)
				//} else {
				//	p.sendStatePacket(state_Idle)
				//}
				go func() {
					childCtx, cancel := context.WithCancel(ctx)
					defer cancel()
					// record bot msg
					if err := p.roomMsgUC.Reply(childCtx, &biz.RoomMessage{
						RoomName:        p.roomConfig.RoomName,
						Sid:             p.room.SID(),
						ParticipantSID:  rp.SID(),
						ParticipantName: BotIdentity,
						Type:            biz.RoomMessageTypeBot,
						ReplyID:         userRoomMessage.ID,
						EventTime:       time.Now(),
						Text:            answer,
					}); err != nil {
						p.log.Errorf("failed to record bot msg, err:%v, roomSID:%s, text:%s", err, p.room.SID(), result.Text)
					}
				}()

				p.sendStatePacket(state_Idle)

				botAnswer := &SpeechEvent{
					ParticipantName: BotIdentity,
					IsBot:           true,
					Text:            answer,
					Timestamp:       uint64(time.Now().Unix()),
				}

				p.lock.Lock()
				p.events = append(p.events, &MeetingEvent{
					Speech: botAnswer,
				})
				p.lock.Unlock()
			}(ctx)
		}
	}
}

func (p *GPTParticipant) answer(events []*MeetingEvent, prompt *SpeechEvent,
	rp *lksdk.RemoteParticipant, transcriber *Transcriber) (string, error) {
	language := transcriber.Language()

	// chat with gpt
	stream, streamCancel, err := p.completion.Complete(p.ctx, events, prompt, rp, p.room, p.roomConfig, language)
	defer streamCancel()
	if err != nil {
		if errors.Is(err, context.Canceled) {
			p.log.Error("p.completion.Complete, errors.Is(err, context.Canceled)")
			return "", nil
		}
		p.log.Errorf("p.completion.Complete, err:%v", err)
		_ = p.sendErrorPacket("Sorry, an error occured while communicating with OpenAI. Max context length reached?")
		return "", err
	}

	p.isStop.Store(false)
	p.gptTrack.OnStop(false)
	p.gptTrack.OnComplete(func(err error) {
		//wg.Done()
	})
	p.gptTrack.OnSpeaking(func(data []byte, err error) {
		p.lastActivity = time.Now()
		transcriber.isSpeaking.Store(true)
	})

	// statistical response time
	p.log.Debugw("p.launchTime", p.launchTime.String(), "now", time.Now(), "elapsed_time", time.Since(p.launchTime).Seconds(), "uint", "s")

	// logsnag
	logsnag.Event(p.ctx, logsnag.Event_STT_TO_TTS_ElAPSED_TIME.
		SetRoomName(p.roomConfig.RoomName).
		SetUID(p.roomConfig.UUID).
		SetElapsedTime(float32(time.Since(p.launchTime).Seconds())).
		SetNotifyElapsedTime(3.00).
		SetMessage(prompt.Text))

	// ai tts by elevenlabs socket
	//if p.ttsElevenlabsSocketCheck(language) {
	//	return p.synthesizerBySocket(stream, language)
	//}

	var last chan struct{} // Used to order the goroutines (See QueueReader bellow)
	var wg sync.WaitGroup

	sb := strings.Builder{}
	limiter := make(chan struct{}, 10)

	for {
		sentence, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
				wg.Wait()
				return strings.TrimSpace(sb.String()), nil
			}

			_ = p.sendErrorPacket("Sorry, an error occured while communicating with OpenAI. It can happen when the servers are overloaded")
			return "", err
		}
		if sentence == "" {
			continue
		}

		if p.isStop.Load() {
			return sb.String(), nil
		}

		// controller goroutine
		limiter <- struct{}{}

		// Try to parse the language from the sentence (ChatGPT can provide <en-US>, en-US as a prefix)
		trimSentence := strings.TrimSpace(sentence)
		lowerSentence := strings.ToLower(trimSentence)
		for code, lang := range Languages {
			prefix1 := strings.ToLower(fmt.Sprintf("<%s>", code))
			prefix2 := strings.ToLower(code)

			if strings.HasPrefix(lowerSentence, prefix1) {
				trimSentence = trimSentence[len(prefix1):]
			} else if strings.HasPrefix(lowerSentence, prefix2) {
				trimSentence = trimSentence[len(prefix2):]
			} else {
				continue
			}

			language = lang
			break
		}

		sb.WriteString(trimSentence)
		sb.WriteString(" ")

		tmpLast := last
		tmpLang := language

		currentCh := make(chan struct{})
		last = currentCh

		wg.Add(1)
		go func() {
			defer close(currentCh)
			defer wg.Done()

			resp, err := p.synthesizer.Synthesize(p.ctx, trimSentence, tmpLang)
			<-limiter // consume
			if err != nil {
				p.log.Errorf("failed to synthesize, err:%v, sentence:%s", err, trimSentence)
				_ = p.sendErrorPacket("Sorry, an error occured while synthesizing voice data using Google TTS")
				return
			}

			if tmpLast != nil {
				<-tmpLast // Reorder outputs
			}

			_ = p.gptTrack.QueueReaderString(trimSentence)
			p.log.Debugf("finished synthesizing, queuing sentence:%s", trimSentence)
			err = p.gptTrack.QueueReader(bytes.NewReader(resp.AudioContent))
			if err != nil {
				p.log.Errorf("failed to queue reader, err:%v, sentence:%s", err, trimSentence)
				return
			}

			_ = p.sendStatePacket(state_Speaking)

			//wg.Add(1)
			p.lastActivity = time.Now()
		}()
	}
}

func (p *GPTParticipant) ttsElevenlabsSocketCheck(lan *Language) bool {
	// English only in third-party voice
	if lan.Code != "en-US" {
		return false
	}

	// if set voice, then use elevenlabs
	voiceID := p.roomConfig.GetConfigVoiceID()
	return voiceID != ""
}

// synthesizerBySocket notice: temporarily unused
func (p *GPTParticipant) synthesizerBySocket(stream StreamReader, language *Language) (string, error) {
	// new socket client
	socket, err := elevenlabs.NewWebSocket(p.ctx, p.log, &elevenlabs.SocketOption{
		ApiKey:  p.confThirdApi.Eleventlabs.Key,
		VoiceID: p.roomConfig.GetConfigVoiceID(),
		ModelID: "eleven_monolingual_v1",
		Callback: &elevenlabs.SocketCallback{
			ReadMessage: func(data []byte) error {
				audio, err := p.synthesizer.FfmpegToOggData(p.ctx, data)
				if err != nil {
					p.log.Errorf("p.synthesizer.FfmpegToOggData, err:%v", err)
					return err
				}

				err = p.gptTrack.QueueReader(bytes.NewReader(audio))
				if err != nil {
					p.log.Errorf("failed to queue reader, err:%v", err)
					return errors.Wrap(err, "failed to queue reader")
				}

				_ = p.sendStatePacket(state_Speaking)
				p.lastActivity = time.Now()
				return nil
			},
		}},
	)
	defer socket.Close()
	if err != nil {
		return "", errors.Wrap(err, "new socket err")
	}

	sb := strings.Builder{}

	socket.Start()
	for {
		sentence, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) || errors.Is(err, context.Canceled) {
				break
			}
			_ = p.sendErrorPacket("Sorry, an error occured while communicating with OpenAI. It can happen when the servers are overloaded")
			break
		}
		if sentence == "" {
			continue
		}

		if p.isStop.Load() {
			break
		}

		// Try to parse the language from the sentence (ChatGPT can provide <en-US>, en-US as a prefix)
		trimSentence := strings.TrimSpace(sentence)
		lowerSentence := strings.ToLower(trimSentence)
		for code, lang := range Languages {
			prefix1 := strings.ToLower(fmt.Sprintf("<%s>", code))
			prefix2 := strings.ToLower(code)

			if strings.HasPrefix(lowerSentence, prefix1) {
				trimSentence = trimSentence[len(prefix1):]
			} else if strings.HasPrefix(lowerSentence, prefix2) {
				trimSentence = trimSentence[len(prefix2):]
			} else {
				continue
			}

			language = lang
			break
		}

		sb.WriteString(trimSentence)
		sb.WriteString(" ")

		// socket write text
		socket.Write(trimSentence + " ")

		_ = p.gptTrack.QueueReaderString(trimSentence)
		p.log.Debugf("finished synthesizing, queuing sentence:%s", trimSentence)
	}
	socket.End()

	return sb.String(), nil
}

func (p *GPTParticipant) sendPacketByWord(str, sid string, sleep time.Duration) {
	// subtitle steam reader
	strStream, clean, err := NewSubTitleSteamReader(str)
	defer clean()
	if err != nil {
		p.log.Errorf("NewSubTitleSteamReader, err:%v, sentence:%s", err, str)
		return
	}

	sb := strings.Builder{}
	for {
		word, err := strStream.Recv()
		if err != nil {
			if err == io.EOF {
				sb.WriteString(word)
				text := sb.String()
				sb.Reset()
				_ = p.sendPacket(&packet{
					Type: packet_Transcript,
					Data: &transcriptPacket{
						Sid:     sid,
						Name:    BotIdentity,
						Text:    text,
						IsFinal: true,
					},
				})
			}
			break
		}

		_ = p.sendStatePacket(state_Active)
		sb.WriteString(word)

		_ = p.sendPacket(&packet{
			Type: packet_Transcript,
			Data: &transcriptPacket{
				Sid:     sid,
				Name:    BotIdentity,
				Text:    sb.String(),
				IsFinal: false,
			},
		})
	}
	return
}
