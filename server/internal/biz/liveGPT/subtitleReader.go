package liveGPT

import (
	"bufio"
	"io"
	"strings"
)

type subTitleSteamReader struct {
	reader *bufio.Reader
}

func (qsr *subTitleSteamReader) Recv() (string, error) {
	sb := strings.Builder{}
	for {
		//buf := make([]byte, 1)
		char, size, err := qsr.reader.ReadRune()
		if err != nil {
			if err == io.EOF && sb.String() != "" {
				return sb.String(), err
			}
			return "", err
		}
		//char := string(buf[:n])
		//sb.WriteString(char)
		sb.WriteRune(char)
		if size == 3 || char == ' ' || char == '\n' || char == '\r' || char == '，' || char == '。' {
			if sb.String() != "" {
				text := sb.String()
				sb.Reset()
				return text, nil
			}
		}
	}
}

func (qsr *subTitleSteamReader) Close() {

}

func NewSubTitleSteamReader(str string) (StreamReader, func(), error) {
	stream := &subTitleSteamReader{
		reader: bufio.NewReader(strings.NewReader(str)),
	}
	return stream, func() {
		stream.Close()
	}, nil
}
