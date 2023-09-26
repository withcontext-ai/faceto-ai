package liveGPT

import (
	"bufio"
	"fmt"
	"github.com/livekit/protocol/logger"
	"io"
	"strings"
	"sync"
	"time"
)

type questionSteamReader struct {
	isFinished bool
	reader     *bufio.Reader
	lock       sync.Mutex
}

func (qsr *questionSteamReader) Reader() (text string, err error) {
	if qsr.isFinished {
		err = io.EOF
		return
	}

	qsr.lock.Lock()

	buf := make([]byte, 1)
	n, err := qsr.reader.Read(buf)
	if err != nil {
		if err != io.EOF {
			fmt.Printf("读取出错: %v\n", err)
		}
		qsr.isFinished = true
	}

	text = string(buf[:n])

	qsr.lock.Unlock()

	time.Sleep(time.Millisecond * 50)
	return
}

func (c *questionSteamReader) Recv() (string, error) {
	sb := strings.Builder{}

	for {
		response, err := c.Reader()
		if err != nil {
			content := sb.String()
			if err == io.EOF && len(strings.TrimSpace(content)) != 0 {
				return content, nil
			}
			return "", err
		}

		delta := response
		sb.WriteString(delta)

		logger.Debugw("questionSteamReader,Recv", "delta", strings.TrimSpace(delta))
		if strings.HasSuffix(strings.TrimSpace(delta), ".") {
			return sb.String(), nil
		}
	}
}

func (qsr *questionSteamReader) Close() {

}

// test get steam
func getSteam() (StreamReader, func(), error) {
	//str := "Some promising use cases for LLMs"
	str := "As technological advancements continue to reshape our world. " +
		"we are witnessing in virtually every aspect of our lives. " +
		"from the way we communicate and work to how we consume entertainment and learn new things. " +
		"and these changes, driven by emerging technologies such as artificial intelligence. " +
		"blockchain, and the Internet of Things. " +
		"are only accelerating, with profound implications for our economy, society, and the environment." +
		"as we struggle to navigate this rapidly changing landscape and confront the challenges and opportunities that lie ahead."
	stream := &questionSteamReader{
		reader: bufio.NewReader(strings.NewReader(str)),
	}
	return stream, func() {
		stream.Close()
	}, nil
}
