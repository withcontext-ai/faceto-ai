package elevenlabs

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const socketEndpoint = "wss://api.elevenlabs.io"

type SocketCallback struct {
	ReadMessage func(data []byte) error
}
type ElevenlabsSocket struct {
	ctx  context.Context
	conn *websocket.Conn
	log  *log.Helper

	done     chan struct{}
	callback *SocketCallback
}

type SocketOption struct {
	ApiKey   string
	VoiceID  string `json:"voice_id"`
	ModelID  string `json:"model_id"`
	Callback *SocketCallback
}

func NewWebSocket(ctx context.Context, log *log.Helper, sop *SocketOption) (*ElevenlabsSocket, error) {
	headers := http.Header{}
	headers.Set("xi-api-key", sop.ApiKey)

	// socket dial
	dialUrl := fmt.Sprintf(socketEndpoint+"/v1/text-to-speech/%s/stream-input?model_type=%s", sop.VoiceID, sop.ModelID)
	log.WithContext(ctx).Debugf("NewWebSocket, dialUrl:%s", dialUrl)
	conn, _, err := websocket.DefaultDialer.Dial(dialUrl, headers)
	if err != nil {
		log.WithContext(ctx).Errorf("NewWebSocket, Failed to connect:%v", err)
		return nil, err
	}

	e := &ElevenlabsSocket{
		ctx:      ctx,
		conn:     conn,
		log:      log,
		done:     make(chan struct{}),
		callback: sop.Callback,
	}

	go e.readMessage()
	//go e.ping()
	return e, nil
}

func (e *ElevenlabsSocket) Ping(ticker bool) error {

	data := map[string]interface{}{
		"text":                   "",
		"try_trigger_generation": true,
	}
	dataJSON, _ := json.Marshal(data)

	if ticker {
		// time ping
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := e.conn.WriteMessage(websocket.PingMessage, dataJSON)
				if err != nil {
					e.log.Errorf("Socket.ping, Failed to WriteMessage err:%v", err)
					return err
				}
			}
		}
	} else {
		err := e.conn.WriteMessage(websocket.PingMessage, dataJSON)
		if err != nil {
			e.log.Errorf("Socket.ping, Failed to WriteMessage err:%v", err)
			return err
		}
	}

	return nil
}

func (e *ElevenlabsSocket) readMessage() {
	defer func() {
		close(e.done)
	}()

	for {
		_, message, err := e.conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				e.log.Infof("Socket.ReadMessage, websocket.IsCloseError:%v", err)
				break
			}
			e.log.Errorf("Socket.ReadMessage, Failed to read message:%v", err)
			continue
		}

		var responseData map[string]interface{}
		err = json.Unmarshal(message, &responseData)
		if err != nil {
			e.log.Errorf("Socket.ReadMessage, Failed to parse JSON:%v", err)
			continue
		}

		if audio, ok := responseData["audio"].(string); ok {
			audioData, _ := base64.StdEncoding.DecodeString(audio)
			if e.callback != nil {
				if e.callback.ReadMessage != nil {
					if err := e.callback.ReadMessage(audioData); err != nil {
						e.log.Errorf("Socket.ReadMessage, e.callback.ReadMessage() err:%v", err)
					}
				}
			}
			//io.Copy(e.w, bytes.NewReader(audioData))
		}
	}
}

func (e *ElevenlabsSocket) Start() error {
	BOS := map[string]interface{}{
		"text":                   " ",
		"try_trigger_generation": true,
		"voice_settings": map[string]interface{}{
			"stability":        0.8,
			"similarity_boost": true,
		},
		"generation_config": map[string]interface{}{
			"chunk_length_schedule": []int{50, 120, 160, 250, 290},
		},
	}
	BOSData, _ := json.Marshal(BOS)
	// Send beginning of stream
	if err := e.conn.WriteMessage(websocket.TextMessage, BOSData); err != nil {
		e.log.Errorf("Socket.Start, Failed to WriteMessage err:%v", err)
		return err
	}
	return nil
}

func (e *ElevenlabsSocket) End() error {
	EOS := map[string]interface{}{
		"text": "",
	}
	EOSData, _ := json.Marshal(EOS)
	// Send beginning of stream
	if err := e.conn.WriteMessage(websocket.TextMessage, EOSData); err != nil {
		e.log.Errorf("Socket.End, Failed to WriteMessage err:%v", err)
		return err
	}

	// wait read message done
	<-e.done
	return nil
}

func (e *ElevenlabsSocket) Write(text string) error {
	// Stream text chunks and receive audio
	data := map[string]interface{}{
		"text":                   text,
		"try_trigger_generation": true,
	}
	dataJSON, _ := json.Marshal(data)
	if err := e.conn.WriteMessage(websocket.TextMessage, dataJSON); err != nil {
		e.log.Errorf("Socket.Write, Failed to WriteMessage err:%v", err)
		return err
	}

	return nil
}

func (e *ElevenlabsSocket) Close() {
	defer e.conn.Close()
	// Cleanly close the connection by sending a close message and then
	// waiting (with timeout) for the server to close the connection.
	if err := e.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
		e.log.Infof("Socket.Close, Failed to WriteMessage err:%v", err)
	}
}
