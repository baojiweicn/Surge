package core

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type TestWatcher struct {
	current *Message
}

func (w *TestWatcher) Id() string {
	return fmt.Sprintf("test_watcher")
}

func (w *TestWatcher) Message() *Message {
	return w.current
}

func (w *TestWatcher) HandleMessage(msg *Message) {
	w.current = msg
	fmt.Println(msg.ctx.Value("hello"))
}

func Test_EventBus(t *testing.T) {
	bus := NewEventBus(1 * time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	go bus.Start(ctx)
	watcher := &TestWatcher{}
	bus.GetOrCreateTopic("test").Add(watcher)
	bus.Send(NewMessage("test", context.WithValue(context.Background(), "hello", "world")))
	now := time.Now()
	for watcher.Message() == nil {
		if time.Now().Sub(now) > 10*time.Second {
			break
		}
		if watcher.Message() != nil {
			if watcher.Message().ctx.Value("hello") != "world" {
				t.Errorf("except (%v) => while recv (%v)", "world", watcher.Message().ctx.Value("hello"))

			}
			t.Logf("recv message in %v", time.Now().Sub(now))
		}
	}
	cancel()
}
