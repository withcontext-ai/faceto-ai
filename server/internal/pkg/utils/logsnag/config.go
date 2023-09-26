package logsnag

import (
	"faceto-ai/internal/pkg/utils/helper"
	"fmt"
)

var (
	PROJECT    = "builder"
	CHANNEL    = "faceto"
	LOGAPI     = "https://api.logsnag.com/v1/log"
	INSIGHTAPI = "https://api.logsnag.com/v1/insight"
)

// ======= Event Defined =======

type Payload struct {
	Project string `json:"project"`
	Channel string `json:"channel"`
	EventLog
}

type EventLog struct {
	Event       string `json:"event"`
	Icon        string `json:"icon"`
	Notify      bool   `json:"notify"`
	Description string `json:"description"`
	Tags        Tags   `json:"tags"`
}

type Tags struct {
	TraceID           string  `json:"trace-id,omitempty"`
	UID               string  `json:"uid,omitempty"`
	RoomName          string  `json:"room-name,omitempty"`
	ElapsedTime       float32 `json:"elapsed-time,omitempty"`
	NotifyElapsedTime float32 `json:"notify-elapsed-time,omitempty"`
	Message           string  `json:"message"`
	Error             string  `json:"error"`
}

func (e EventLog) SetNotify() EventLog {
	e.Notify = true
	return e
}

func (e EventLog) SetTraceID(traceID string) EventLog {
	e.Tags.TraceID = traceID
	return e
}

func (e EventLog) SetUID(uid string) EventLog {
	e.Tags.UID = uid
	return e
}

func (e EventLog) SetRoomName(name string) EventLog {
	e.Tags.RoomName = name
	e.Description = fmt.Sprintf(e.Description+" \nRoom Name:%s", name)
	return e
}

func (e EventLog) SetElapsedTime(t float32) EventLog {
	e.Tags.ElapsedTime = t
	return e
}

func (e EventLog) SetNotifyElapsedTime(t float32) EventLog {
	e.Tags.NotifyElapsedTime = t
	return e
}

func (e EventLog) SetMessage(msg string) EventLog {
	e.Tags.Message = helper.TruncateString(msg, 50)
	e.Description = fmt.Sprintf(e.Description+" \nText:%s", helper.TruncateString(msg, 10))
	return e
}

func (e EventLog) SetError(err string) EventLog {
	e.Tags.Error = err
	return e
}

// Link Event
var (
	EventApplyLink = EventLog{
		Event:       "Apply Link",
		Icon:        "🔗",
		Description: "Request event for applying video link",
	}
	EventApplyLinkFailed = EventLog{
		Event:       "Apply Link Failed",
		Icon:        "❌",
		Description: "Request event for applying video link, Failed.",
	}
)

// Room Webhook Event
var (
	EventRoomWebhook_Room_Started = EventLog{
		Event:       "🚩Webhook Room Started",
		Icon:        "🚀",
		Description: "This is the webhook event for Room Started.",
	}
	EventRoomWebhook_Track_Published = EventLog{
		Event:       "🚩Webhook Track Published",
		Icon:        "⏳",
		Description: "This is the webhook event for Track Published.",
	}
	EventRoomWebhook_Participant_Joined = EventLog{
		Event:       "🚩Webhook Participant Joined",
		Icon:        "⏩",
		Description: "This is the webhook event for Participant Joined.",
	}
	EventRoomWebhook_Participant_Left = EventLog{
		Event:       "🚩Webhook Participant Left",
		Icon:        "👋",
		Description: "This is the webhook event for Participant Left.",
	}
	EventRoomWebhook_Egress_Ended = EventLog{
		Event:       "🚩Webhook Egress Ended",
		Icon:        "📽️",
		Description: "This is the webhook event for Egress Ended.",
	}
	EventRoomWebhook_Finish = EventLog{
		Event:       "🚩Webhook Room Finish",
		Icon:        "🔚",
		Description: "This is the webhook event for Room Finish.",
	}
	EventRoomWebhook_Push_RoomStarted = EventLog{
		Event:       "🚩Webhook Push Room Started",
		Icon:        "⏩",
		Description: "This is the push webhook event for Room Started.",
	}
	EventRoomWebhook_Push_ParticipantLeft = EventLog{
		Event:       "🚩Webhook Push Participant Left",
		Icon:        "⏩",
		Description: "This is the push webhook event for Participant Left.",
	}
)

// API Event
var (
	Event_CHATGPT_ElAPSED_TIME = EventLog{
		Event:       "🕙 ChatGPT Default Elapsed Time",
		Icon:        "🤖",
		Description: "This is the response time of ChatGPT's api interface.",
		Tags: Tags{
			ElapsedTime: 0.00,
		},
	}
	Event_CHATGPT_THIRDAPI_ElAPSED_TIME = EventLog{
		Event:       "🕙 ChatGPT For Third API Elapsed Time",
		Icon:        "🤖",
		Description: "This is the response time of Third ChatAPI.",
		Tags: Tags{
			ElapsedTime: 0.00,
		},
	}
	Event_CHATGPT_THIRDAPI_ERROR = EventLog{
		Event:       "❌ ChatGPT For Third API ERROR",
		Icon:        "🤖",
		Description: "The third ChatAPI Error.",
		Tags: Tags{
			ElapsedTime: 0.00,
		},
	}
	Event_CHATGPT_THIRDAPI_TIMEOUT = EventLog{
		Event:       "❌ ChatGPT For Third API TimeOut",
		Icon:        "🤖",
		Description: "The third ChatAPI timeout. Error.",
		Tags: Tags{
			ElapsedTime: 0.00,
		},
	}
	Event_API_TTS_GOOGLE_ElAPSED_TIME = EventLog{
		Event:       "🕙 API TTS Google Elapsed Time",
		Icon:        "🔈",
		Description: "This is the interface response time of google cloud text to speech.",
		Tags: Tags{
			ElapsedTime: 0.00,
		},
	}
	Event_API_TTS_ELEVENLABS_ElAPSED_TIME = EventLog{
		Event:       "🕙 API TTS Elevenlabs Elapsed Time",
		Icon:        "🔈",
		Description: "This is the interface response time of elevenlabs text to speech.",
		Tags: Tags{
			ElapsedTime: 0.00,
		},
	}
	Event_API_TTS_ELEVENLABS_ERROR = EventLog{
		Event:       "❌ API TTS Elevenlabs ERROR",
		Icon:        "🤖",
		Description: "This is the interface response time of elevenlabs Error.",
		Tags: Tags{
			ElapsedTime: 0.00,
		},
	}
	Event_STT_TO_TTS_ElAPSED_TIME = EventLog{
		Event:       "🕙 STT To TTS Elapsed Time",
		Icon:        "⏳",
		Description: "This is the elapsed time of user stop talk until AI talk.",
		Tags: Tags{
			ElapsedTime: 0.00,
		},
	}
)

// Room Message Event
var (
	EventRoomMessageUser = EventLog{
		Event:       "💬 Room Message For User.",
		Icon:        "👤",
		Description: "this is what people say.",
	}
	EventRoomMessageAI = EventLog{
		Event:       "💬 Room Message For AI.",
		Icon:        "🤖",
		Description: "this is what AI say.",
	}
)

// ======= Insight Defined =======

type InsightValue struct {
	Project string `json:"project"`
	InsightLog
}

type InsightLog struct {
	Title string `json:"title"`
	Value int    `json:"value,omitempty"`
	Icon  string `json:"icon,omitempty"`
}

func (ins InsightLog) SetValue(v int) InsightLog {
	ins.Value = v
	return ins
}

// Room Insight
var (
	InsightLinkCount = InsightLog{
		Title: "Link Count",
		Icon:  "🔗",
	}
	InsightRoomCount = InsightLog{
		Title: "Room Count",
		Icon:  "🏠",
	}
	InsightRoomMsgCount = InsightLog{
		Title: "Room Message Count",
		Icon:  "💬",
	}
)
