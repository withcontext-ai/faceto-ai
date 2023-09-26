// Package liveGPT
// ---------------------------------
// @file      : reader_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/9/5 18:59
// @desc      : file description
// ---------------------------------
package liveGPT

import (
	"fmt"
	"testing"
)

func Test_GetStream(t *testing.T) {
	stream, fc, err := getSteam()
	if err != nil {
		t.Errorf("get stream err:%v", err)
	}
	defer fc()
	for {
		text, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Println(text)
	}
}
