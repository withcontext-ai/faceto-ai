package liveGPT

import (
	"context"
	"faceto-ai/internal/biz"
	"faceto-ai/internal/pkg/utils/logsnag"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"strings"
	"time"

	lksdk "github.com/livekit/server-sdk-go"
	openai "github.com/sashabaranov/go-openai"
)

// A sentence in the conversation (Used for the history)
type SpeechEvent struct {
	ParticipantName string     `json:"participant_name"`
	IsBot           bool       `json:"is_bot"`
	Text            string     `json:"text"`
	History         [][]string `json:"history"`
	Questions       string     `json:"questions"`
	SessionID       string     `json:"session_id"`
	Conclusion      string     `json:"conclusion"`
	Timestamp       uint64     `json:"timestamp"`
}

type JoinLeaveEvent struct {
	Leave           bool
	ParticipantName string
	Time            time.Time
}

type MeetingEvent struct {
	Speech *SpeechEvent
	Join   *JoinLeaveEvent
}

type ChatCompletionAPI interface {
	Complete(ctx context.Context, events []*MeetingEvent, prompt *SpeechEvent,
		participant *lksdk.RemoteParticipant, room *lksdk.Room, roomConfig *biz.Link, language *Language) (StreamReader, func(), error)
}

type ChatCompletion struct {
	client *openai.Client
	log    *log.Helper
}

func NewChatCompletion(client *openai.Client, logger *log.Helper) *ChatCompletion {
	return &ChatCompletion{
		client: client,
		log:    logger,
	}
}

func (c *ChatCompletion) Complete(ctx context.Context, events []*MeetingEvent, prompt *SpeechEvent,
	participant *lksdk.RemoteParticipant, room *lksdk.Room, roomConfig *biz.Link, language *Language) (StreamReader, func(), error) {
	//return getSteam()
	launchTime := time.Now()

	var sb strings.Builder
	participants := room.GetParticipants()
	for i, participant := range participants {
		sb.WriteString(participant.Name())
		if i != len(participants)-1 {
			sb.WriteString(", ")
		}
	}
	participantNames := sb.String()
	sb.Reset()

	messages := make([]openai.ChatCompletionMessage, 0, len(events)+3)

	systemRoleContent := fmt.Sprintf("You are %s, a role player of voice assistant. ", BotIdentity)

	// prompt config
	if roomConfig.GetPromptConfig() != nil {
		role := roomConfig.GetPromptConfig().Role
		if role != "" {
			systemRoleContent += fmt.Sprintf("Now, your role is %s, and your task is to communicate with users according to the current role. If there is a list of questions provided, you need to ask questions one by one according to these questions. \n", role)
		}
		qas := roomConfig.GetPromptConfig().Question
		if len(qas) > 0 {
			systemRoleContent += fmt.Sprintf("Here are %d questions. ", len(qas)) +
				"After the user asks the first question, you need to guide the user to start the question and answer, please ask each question one by one and wait for user's response. " +
				"If user indicates that they don't know how to answer the question, give them some hints. " +
				fmt.Sprintf("Once you have asked all %d questions, politely end the meeting. Throughout the process, please maintain the role of an %s and avoid revealing that you are an AI model. ", len(qas), role) +
				fmt.Sprintf("The list of questions is as follows: \n %s \n", strings.Join(qas, "\n"))
		}
	}

	systemRoleContent += "Keep your responses concise while still being friendly and personable. " +
		"If your response is a question, please append a question mark symbol to the end of it. " + // Used for auto-trigger
		fmt.Sprintf("There are actually %d participants in the meeting: %s. ", len(participants), participantNames) +
		fmt.Sprintf("Current language: %s Current date: %s. ", language.Label, time.Now().Format("January 2, 2006 3:04pm"))

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemRoleContent,
	})

	for _, e := range events {
		if e.Speech != nil {
			if e.Speech.IsBot {
				messages = append(messages, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleAssistant,
					Content: e.Speech.Text,
					//Name:    BotIdentity,
				})
			} else {
				messages = append(messages, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("%s said %s", e.Speech.ParticipantName, e.Speech.Text),
					//Name:    e.Speech.ParticipantName,
				})
			}
		}

		if e.Join != nil {
			if e.Join.Leave {
				messages = append(messages, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleSystem,
					Content: fmt.Sprintf("%s left the meeting at %s", e.Join.ParticipantName, e.Join.Time.Format("3:04pm")),
				})
			} else {
				messages = append(messages, openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleSystem,
					Content: fmt.Sprintf("%s joined the meeting at %s", e.Join.ParticipantName, e.Join.Time.Format("3:04pm")),
				})
			}
		}
	}

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: fmt.Sprintf("You are currently talking to %s", participant.Name()),
	})

	// prompt
	if prompt.Text != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt.Text,
			//Name:    prompt.ParticipantName,
		})
	}

	c.log.Debugw("CreateChatCompletionStream", messages)
	stream, err := c.client.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
		Stream:   true,
	})

	if err != nil {
		c.log.Errorf("error creating chat completion stream, err:%v", err)
		return nil, func() {}, err
	}
	c.log.Debugw("p.launchTime", launchTime.String(), "now", time.Now(), "chat_openai_elapsed_time", time.Since(launchTime).Seconds(), "uint", "s")

	// logsnag
	logsnag.Event(ctx, logsnag.Event_CHATGPT_ElAPSED_TIME.
		SetRoomName(room.Name()).
		SetUID(roomConfig.UUID).
		SetElapsedTime(float32(time.Since(launchTime).Seconds())).
		SetMessage(prompt.Text))

	return &ChatStream{
		stream: stream,
	}, func() {}, nil
}

// Wrapper around openai.ChatCompletionStream to return only complete sentences
type ChatStream struct {
	stream *openai.ChatCompletionStream
}

func (c *ChatStream) Recv() (string, error) {
	sb := strings.Builder{}

	for {
		response, err := c.stream.Recv()
		if err != nil {
			content := sb.String()
			if err == io.EOF && len(strings.TrimSpace(content)) != 0 {
				return content, nil
			}
			return "", err
		}

		if len(response.Choices) == 0 {
			continue
		}

		delta := response.Choices[0].Delta.Content
		sb.WriteString(delta)

		if strings.HasSuffix(strings.TrimSpace(delta), ",") {
			return sb.String(), nil
		}
		if strings.HasSuffix(strings.TrimSpace(delta), "?") {
			return sb.String(), nil
		}
		if strings.HasSuffix(strings.TrimSpace(delta), ".") {
			return sb.String(), nil
		}
	}
}

func (c *ChatStream) Close() {
	c.stream.Close()
}
