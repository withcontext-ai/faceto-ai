// Package liveGPT
// ---------------------------------
// @file      : subtitleReader_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/5/19 12:05
// @desc      : file description
// ---------------------------------
package liveGPT

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"time"
)

func TestNewSubTitleSteamReader(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "case",
			args: args{str: "As technological advancements continue?" +
				" to reshape our world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			steam, clean, err := NewSubTitleSteamReader(tt.args.str)
			defer clean()
			if err != nil {
				t.Errorf("NewSubTitleSteamReader() error = %v", err)
				return
			}
			sb := strings.Builder{}
			for {
				word, err := steam.Recv()
				if err != nil {
					if err == io.EOF {
						sb.WriteString(word)
						text := sb.String()
						sb.Reset()
						fmt.Println("=========")
						fmt.Println(text)
					}
					break
				}
				sb.WriteString(word)
				fmt.Println(sb.String())
				time.Sleep(100 * time.Millisecond)
			}
		})
	}
}
