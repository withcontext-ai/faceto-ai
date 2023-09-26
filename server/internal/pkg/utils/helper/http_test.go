// Package helper
// ---------------------------------
// @file      : http_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/5/22 19:58
// @desc      : file description
// ---------------------------------
package helper

import (
	"context"
	"fmt"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

type Room struct {
	Name string `json:"name"`
}
type RoomReq struct {
	Rooms []*Room `json:"rooms"`
}

type RoomReply struct {
	Msg string `json:"msg"`
}

func TestRestyRequest(t *testing.T) {

	defer gock.Off() // 测试执行后刷新挂起的mock

	type args struct {
		ctx      context.Context
		url      string
		req      interface{}
		response interface{}
		isRetry  bool
	}

	req := &RoomReq{
		Rooms: make([]*Room, 10),
	}
	req.Rooms = append(req.Rooms, &Room{Name: "gvn9-gguv"})

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: Generator(),
			args: args{
				ctx:      context.Background(),
				url:      "https://www.baidu.com",
				req:      req,
				response: RoomReply{},
				isRetry:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// mock 请求外部api时传参 {"name": "zhangsan", "age": 30} 返回 hello golang
			gock.New("https://www.baidu.com").
				Post("/").
				MatchType("json").
				//JSON().
				Reply(200).
				JSON(map[string]interface{}{"msg": "hello golang"})

			option := &RestyOptions{
				Url:      tt.args.url,
				Req:      tt.args.req,
				Response: tt.args.response,
				IsRetry:  tt.args.isRetry,
				Headers:  nil,
			}

			_, err := RestyRequest(tt.args.ctx, option)
			if err != nil {
				t.Errorf("RestyRequest() error = %v", err)
			}
			fmt.Println("done.")

			//var resp RoomReply
			//if err := json.Unmarshal(body.Body(), &resp); err != nil {
			//	t.Errorf("json.Unmarshal error = %v", err)
			//}

			// 校验返回结果是否符合预期
			//assert.Equal(t, resp.Msg, "hello golang")

			//assert.True(t, gock.IsDone()) // 断言mock被触发
		})
	}
}
