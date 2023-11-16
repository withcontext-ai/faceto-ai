package liveGPT

// Packets sent over the datachannels
type packetType int32

const (
	packet_Transcript    packetType = 0
	packet_State         packetType = 1
	packet_Error         packetType = 2 // Show an error message to the user screen
	packet_EventStopRoom packetType = 3 // Stop Room Alert
)

type gptState int32

const (
	state_Idle     gptState = 0
	state_Loading  gptState = 1
	state_Speaking gptState = 2
	state_Active   gptState = 3
)

type packet struct {
	Type packetType  `json:"type"`
	Data interface{} `json:"data"`
}

type transcriptPacket struct {
	Sid     string `json:"sid"`
	Name    string `json:"name"`
	Text    string `json:"text"`
	IsFinal bool   `json:"isFinal"`
}

type statePacket struct {
	State gptState `json:"state"`
}

type errorPacket struct {
	Message string `json:"message"`
}

type eventPacket struct {
	Sid   string `json:"sid"`
	Name  string `json:"name"`
	Event string `json:"event"`
	Text  string `json:"text"`
}
