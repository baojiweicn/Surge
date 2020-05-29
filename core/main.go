package core

import (
	"context"
	"time"
)

var DefaultTickDuration = 10 * time.Second

type Surge struct {
	bus    *EventBus
	ctx    context.Context
	cancel context.CancelFunc
}

func NewSurge() *Surge {
	ctx, cancel := context.WithCancel(context.Background())
	return &Surge{
		bus:    NewEventBus(DefaultTickDuration),
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Surge) Start() {
	if s.bus == nil {
		s.bus = NewEventBus(DefaultTickDuration)
	}
	go s.bus.Start(s.ctx)
}

func (s *Surge) Graceful() {
	s.cancel()
}
