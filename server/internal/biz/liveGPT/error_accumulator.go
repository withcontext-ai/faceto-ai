package liveGPT

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type unmarshaler interface {
	unmarshal(data []byte, v any) error
}

type jsonUnmarshaler struct{}

func (jm *jsonUnmarshaler) unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

type errorAccumulator interface {
	write(p []byte) error
	unmarshalError() *ErrorResponse
}

type errorBuffer interface {
	io.Writer
	Len() int
	Bytes() []byte
}

type defaultErrorAccumulator struct {
	buffer      errorBuffer
	unmarshaler unmarshaler
}

func newErrorAccumulator() errorAccumulator {
	return &defaultErrorAccumulator{
		buffer:      &bytes.Buffer{},
		unmarshaler: &jsonUnmarshaler{},
	}
}

func (e *defaultErrorAccumulator) write(p []byte) error {
	_, err := e.buffer.Write(p)
	if err != nil {
		return fmt.Errorf("error accumulator write error, %w", err)
	}
	return nil
}

func (e *defaultErrorAccumulator) unmarshalError() (errResp *ErrorResponse) {
	if e.buffer.Len() == 0 {
		return
	}

	err := e.unmarshaler.unmarshal(e.buffer.Bytes(), &errResp)
	if err != nil {
		errResp = nil
	}

	return
}
