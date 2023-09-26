// Package helper
// ---------------------------------
// @file      : string_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/5/17 18:03
// @desc      : file description
// ---------------------------------
package helper

import (
	"fmt"
	"testing"
)

func TestGenerateRoomID(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRoomID()
			fmt.Println(got)
		})
	}
}

func TestTruncateString(t *testing.T) {
	fmt.Println(TruncateString("hello, worddmdadadadada", 10))
	fmt.Println(TruncateString("hello,你好世界,23232", 10))
}
