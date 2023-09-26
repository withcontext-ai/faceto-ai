package event

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var (
	_ Event = (*Message)(nil)
	_ Event = (*IndexMessage)(nil)
)

type Message struct {
	key   string
	value []byte
}

func (m *Message) Key() string {
	return m.key
}

func (m *Message) Value() []byte {
	return m.value
}

func NewMessage(key string, value []byte) *Message {
	return &Message{
		key:   key,
		value: value,
	}
}

type IndexValue struct {
	UUID   string `json:"uuid"`
	Status uint32 `json:"status"`
}

type IndexMessage struct {
	Message
	Index IndexValue
}

// func NewIndexMessage(key string, Index IndexValue) (*IndexMessage, error) {
// 	b, err := json.Marshal(Index)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "Index json marshal error")
// 	}
//
// 	return &IndexMessage{
// 		Message: Message{
// 			key:   key,
// 			value: b,
// 		},
// 		Index: Index,
// 	}, nil
// }

func NewIndexMessageFromEvent(event Event) (*IndexMessage, error) {
	if len(event.Value()) == 0 {
		return nil, errors.New("event value is empty")
	}
	value := IndexValue{}
	if err := json.Unmarshal(event.Value(), &value); err != nil {
		return nil, errors.Wrap(err, "Index json unmarshal error")
	}

	return &IndexMessage{
		Message: Message{
			key:   event.Key(),
			value: event.Value(),
		},
		Index: value,
	}, nil
}

type FLowValue struct {
	ChatID string `json:"chat_id"`
	FlowID string `json:"flow_id"`
	Status string `json:"status"`
}

type FLowMessage struct {
	Message
	Flow FLowValue
}

func NewFLowMessageFromFlowValue(value FLowValue) (*FLowMessage, error) {
	b, err := json.Marshal(value)
	if err != nil {
		return nil, errors.Wrap(err, "FLow json marshal error")
	}
	return &FLowMessage{
		Message: Message{
			key:   value.ChatID,
			value: b,
		},
		Flow: value,
	}, nil
}
