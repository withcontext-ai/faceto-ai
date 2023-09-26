package memory

import (
	"context"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"

	"faceto-ai/internal/pkg/event"
)

var (
	_ event.Sender   = (*LocalSender)(nil)
	_ event.Receiver = (*LocalReceiver)(nil)
)

const maxQueueSize = 100

type localCache struct {
	Events map[string]chan event.Event
	mutex  sync.RWMutex
}

func newLocalCache() *localCache {
	return &localCache{
		Events: make(map[string]chan event.Event),
	}
}

func (c *localCache) GetWithCreate(key string) chan event.Event {
	c.mutex.RLock()
	v, ok := c.Events[key]
	if !ok {
		c.mutex.RUnlock()
		c.mutex.Lock()
		v, ok = c.Events[key]
		if !ok {
			v = make(chan event.Event, maxQueueSize)
			c.Events[key] = v
		}
		c.mutex.Unlock()
		return v
	}
	c.mutex.RUnlock()
	return v
}

var LocalCache = newLocalCache()

type LocalSender struct {
	log *log.Helper
}

func NewLocalSender(logger log.Logger) *LocalSender {
	return &LocalSender{
		log: log.NewHelper(logger),
	}
}

func (s *LocalSender) Send(ctx context.Context, msg event.Event) error {
	channel := LocalCache.GetWithCreate(msg.Key())
	t := time.After(5 * time.Second)

	select {
	case channel <- msg:
	case <-t:
		s.log.WithContext(ctx).Error("Send timeout")
	}

	s.log.WithContext(ctx).Debugf("Send: %s %s %d", msg.Key(), string(msg.Value()), len(channel))

	return nil
}

func (s *LocalSender) SendBatch(ctx context.Context, msgs []event.Event) error {
	for _, msg := range msgs {

		if err := s.Send(ctx, msg); err != nil {
			return errors.WithMessage(err, "SendBatch")
		}
	}
	return nil
}

func (s *LocalSender) Close(ctx context.Context) error {
	return nil
}

type LocalReceiver struct {
	log *log.Helper
}

func NewLocalReceiver(logger log.Logger) *LocalReceiver {
	return &LocalReceiver{
		log: log.NewHelper(logger),
	}
}

func (s *LocalReceiver) Receive(ctx context.Context, handler event.Handler) error {
	return nil
}

func (s *LocalReceiver) Close(ctx context.Context) error {
	return nil
}
