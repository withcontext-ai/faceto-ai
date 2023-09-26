package liveGPT

import (
	"strings"
	"sync"
	"time"
)

type QueueItem struct {
	Text      string
	Timestamp uint64
}

type QueueCallback struct {
	TimeTicker func(text string)
}

type TimeQueue struct {
	lock       sync.Mutex
	timeTicker *time.Ticker

	timeInterval time.Duration
	lastTime     time.Time
	queue        []*QueueItem
	callback     *QueueCallback
}

func NewTimeQueue(timeInterval time.Duration, callback *QueueCallback) *TimeQueue {
	t := &TimeQueue{
		timeInterval: timeInterval,
		queue:        make([]*QueueItem, 0, 16),
		timeTicker:   time.NewTicker(time.Millisecond * 100),
		callback:     callback,
	}
	t.init()
	return t
}

func (tq *TimeQueue) init() {
	tq.lastTime = time.Time{}
	tq.queue = make([]*QueueItem, 0, 16)
	go tq.setTimeTicker()
}

func (tq *TimeQueue) setTimeTicker() {
	for {
		select {
		case <-tq.timeTicker.C:
			if !tq.canAppend() {
				tq.callback.TimeTicker(tq.flush())
			}
		}
	}
}

func (tq *TimeQueue) SetLastTime() {
	tq.lastTime = time.Now()
}

func (tq *TimeQueue) Append(text string, cover bool) {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	if !tq.canAppend() {
		return
	}

	tq.SetLastTime()
	if cover {
		tq.queue = make([]*QueueItem, 0, 16)
	}

	tq.queue = append(tq.queue, &QueueItem{
		Text:      text,
		Timestamp: uint64(tq.lastTime.Unix()),
	})
}

func (tq *TimeQueue) canAppend() bool {
	if tq.lastTime.IsZero() {
		return true
	}
	return time.Since(tq.lastTime).Seconds() <= tq.timeInterval.Seconds()
}

func (tq *TimeQueue) flush() string {
	tq.lock.Lock()
	defer tq.lock.Unlock()

	ut := make([]string, 0)
	for _, q := range tq.queue {
		ut = append(ut, q.Text)
	}
	tq.init()
	return strings.Join(ut, " ")
}

func (tq *TimeQueue) Close() {
	tq.timeTicker.Stop()
}
