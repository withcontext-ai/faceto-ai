// Package liveGPT
// ---------------------------------
// @file      : time_queue_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/9/6 11:43
// @desc      : file description
// ---------------------------------
package liveGPT

import (
	"fmt"
	c "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestNewTimeQueue(t *testing.T) {
	var resp string

	tq := NewTimeQueue(time.Second*3, &QueueCallback{
		TimeTicker: func(text string) {
			fmt.Println(text)
			resp = text
		},
	})
	defer tq.Close()

	c.Convey("new time queue", t, func() {
		tq.Append("hello", false)
		tq.Append("world!", false)

		time.Sleep(4 * time.Second)
		tq.Append("new word.", false)

		c.So(resp, c.ShouldEqual, "hello world!") // 断言
	})
}
