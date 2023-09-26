// Package logsnag
// ---------------------------------
// @file      : log_test.go
// @project   : faceto-ai
// @author    : zhangxiubo
// @time      : 2023/8/25 16:00
// @desc      : file description
// ---------------------------------
package logsnag

import (
	"context"
	"faceto-ai/internal/pkg/middleware"
	"faceto-ai/internal/pkg/utils/helper"
	"fmt"
	"os"
	"testing"
	"time"
)

var ctx = context.Background()

func init() {

	os.Setenv("FACETOAI_ENV", "dev")
	// dev
	os.Setenv("LOGSNAG_API_KEY", "f0963d40f010505006e6ca698cc7975a")
	// prod
	//os.Setenv("LOGSNAG_API_KEY", "0ea16c92372a78b62a4e29e3e81f5bb0")
	PROJECT = "facetoai"
	CHANNEL = "api"

	traceID := helper.Generator()
	fmt.Println("init", traceID)
	ctx = context.WithValue(ctx, middleware.TraceID, traceID)
}

func TestEvent(t *testing.T) {
	Event(ctx, EventApplyLink.SetRoomName("sw72-udri").SetUID("01H7YGWJ19VH03XWXTVTMQRRG7"))
	Event(ctx, EventApplyLinkFailed)
	time.Sleep(time.Second * 3)
}

func TestEventWebhook(t *testing.T) {
	traceID := "f1f32afc-6222-409c-b3e6-a31e0d7032ce"
	name := "sw72-udri"

	Event(ctx, EventRoomWebhook_Room_Started.SetTraceID(traceID).SetRoomName(name))
	Event(ctx, EventRoomWebhook_Track_Published.SetTraceID(traceID).SetRoomName(name))
	Event(ctx, EventRoomWebhook_Participant_Joined.SetTraceID(traceID).SetRoomName(name))
	Event(ctx, EventRoomWebhook_Participant_Left.SetTraceID(traceID).SetRoomName(name))
	Event(ctx, EventRoomWebhook_Egress_Ended.SetTraceID(traceID).SetRoomName(name))
	Event(ctx, EventRoomWebhook_Finish.SetTraceID(traceID).SetRoomName(name))
	Event(ctx, EventRoomWebhook_Push_RoomStarted.SetTraceID(traceID).SetRoomName(name))
	Event(ctx, EventRoomWebhook_Push_ParticipantLeft.SetTraceID(traceID).SetRoomName(name))
	time.Sleep(time.Second * 3)
}

func TestAPIEvent(t *testing.T) {
	Event(ctx, Event_CHATGPT_ElAPSED_TIME.SetElapsedTime(1.88))
	Event(ctx, Event_CHATGPT_THIRDAPI_ElAPSED_TIME.SetElapsedTime(2.13))
	Event(ctx, Event_API_TTS_GOOGLE_ElAPSED_TIME.SetElapsedTime(2.44))
	Event(ctx, Event_API_TTS_ELEVENLABS_ElAPSED_TIME.SetElapsedTime(1.88))

	Event(ctx, Event_STT_TO_TTS_ElAPSED_TIME.SetElapsedTime(1.88).SetNotifyElapsedTime(2.00))

	time.Sleep(time.Second * 3)
}

func TestEventMessage(t *testing.T) {
	Event(ctx, EventRoomMessageUser.SetRoomName("sw72-udri").SetUID("01H7YGWJ19VH03XWXTVTMQRRG7").SetMessage("hello, world"))
	Event(ctx, EventRoomMessageAI.SetRoomName("sw72-udri").SetUID("01H7YGWJ19VH03XWXTVTMQRRG7").SetMessage("hello, worldfssfdsfdsfsdsdfds"))
	time.Sleep(time.Second * 3)
}

func TestInsight(t *testing.T) {
	Insight(ctx, InsightLinkCount.SetValue(100))
	Insight(ctx, InsightRoomCount.SetValue(100))
	Insight(ctx, InsightRoomMsgCount.SetValue(100))

	time.Sleep(time.Second * 3)
}
