package liveGPT

import (
	"github.com/pkg/errors"
	"time"
)

var (
	ErrCodecNotSupported = errors.New("this codec isn't supported")
	ErrBusy              = errors.New("the gpt participant is already used")

	BotIdentity = "KITT"

	// Naive trigger/activation implementation
	GreetingWords = []string{"hi", "hello", "hey", "hallo", "salut", "bonjour", "hola", "eh", "ey"}
	NameWords     = []string{"kit", "gpt", "kitt", "livekit", "live-kit", "kid"}

	ActivationWordsLen = 2
	ActivationTimeout  = 4 * time.Second // If the participant didn't say anything for this duration, stop listening

	Languages = map[string]*Language{
		"en-US": {
			Code:             "en-US",
			Label:            "English",
			TranscriberCode:  "en-US",
			SynthesizerModel: "en-US-Wavenet-D",
		},
		"fr-FR": {
			Code:             "fr-FR",
			Label:            "Français",
			TranscriberCode:  "fr-FR",
			SynthesizerModel: "fr-FR-Wavenet-B",
		},
		"de-DE": {
			Code:             "de-DE",
			Label:            "German",
			TranscriberCode:  "de-DE",
			SynthesizerModel: "de-DE-Wavenet-B",
		},
		"es-ES": {
			Code:             "es-ES",
			Label:            "Spanish",
			TranscriberCode:  "es-ES",
			SynthesizerModel: "es-ES-Wavenet-B",
		},
		"cmn-CN": {
			Code:             "cmn-CN",
			Label:            "Mandarin",
			TranscriberCode:  "cmn-CN",
			SynthesizerModel: "cmn-CN-Wavenet-C",
		},
		"ja-JP": {
			Code:             "ja-JP",
			Label:            "Japanese",
			TranscriberCode:  "ja-JP",
			SynthesizerModel: "ja-JP-Wavenet-C",
		},
	}

	DefaultLanguage = Languages["en-US"]

	NoticeText = map[string]map[string]string{
		"en-US": {
			"greeting": "Hello, I am the interviewer of this time",
			"notice":   "Hello, are you still there?",
			"end":      "I'm leaving in 3 seconds.",
		},
		"cmn-CN": {
			"greeting": "你好，我是本次的面试官，你准备好了吗？",
			"notice":   "你好，你还在听吗？",
			"end":      "我会在3秒之后退出，谢谢你的参与。",
		},
		"ja-JP": {
			"greeting": "こんにちは、今回の面接官です、準備はできていますか？",
			"notice":   "こんにちは、まだ聞いていますか？",
			"end":      "3秒以内に終了します、ご参加いただきありがとうございます。",
		},
	}
)

type Language struct {
	Code             string
	Label            string
	TranscriberCode  string
	SynthesizerModel string
}

type ParticipantMetadata struct {
	LanguageCode string `json:"languageCode,omitempty"`
}

type InitiativeSpeak struct {
	Sid      string
	Text     string
	Language *Language
}
