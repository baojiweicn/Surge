// Copyright 2020 Surge Project. rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Author : baojiwei@live.com.

package bus

import (
	"context"
	"errors"
	"sync"
	"time"
)

// DefaultTopic : default topic
var DefaultTopic = "surge_default_topic"

// Message
type Message struct {
	topic     string
	ctx       context.Context
	timestamp time.Time
}

// NewMessage : create new message.
func NewMessgae(topic string, ctx context.Context) *Message {
	return &Message{
		topic:     topic,
		ctx:       ctx,
		timestamp: time.Now(),
	}
}

// Topic : get topic name of the message
func (msg *Message) Topic() string {
	// if not set topic name
	if msg.topic == "" {
		return DefaultTopic
	}
	return msg.topic
}

// Watcher : any component or any element can be defined as a watcher
type Watcher interface {
	HandleMessage(*Message)
	Id() string
}

// Topic : message topic
type Topic struct {
	lock        *sync.RWMutex
	name        string
	msgs        chan *Message
	subscribers []Watcher
}

// NewTopic : create new topic
func NewTopic(name string) *Topic {
	return &Topic{
		lock:        &sync.RWMutex{},
		name:        name,
		msgs:        make(chan *Message, 1000),
		subscribers: make([]Watcher, 0),
	}
}

// addMsg : add message to topic
func (t *Topic) addMsg(msg *Message) {
	t.msgs <- msg
}

// fire : fire topic to watcher
func (t *Topic) fire() {
	for len(t.msgs) > 0 {
		select {
		case msg := <-t.msgs:
			for _, watcher := range t.GetWatchers() {
				go watcher.HandleMessage(msg)
			}
		}
	}
}

// GetWatchers : get watchers of the topic (with lock)
func (t *Topic) GetWatchers() []Watcher {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.subscribers
}

// Add : add a watcher to the topic
func (t *Topic) Add(watcher Watcher) {
	t.lock.Lock()
	defer t.lock.Unlock()
	for _, i := range t.subscribers {
		if i == watcher {
			return
		}
	}
	t.subscribers = append(t.subscribers, watcher)
}

// Delete : delete a watcher fromt the topic
func (t *Topic) Delete(watcher Watcher) {
	t.lock.Lock()
	defer t.lock.Unlock()
	for i, w := range t.subscribers {
		if w.Id() == watcher.Id() {
			t.subscribers = append(t.subscribers[:i], t.subscribers[i+1:]...)
		}
	}
}

// EventBus : the main event bus center of messages
// Select releated message from topic and send them to wathers.
type EventBus struct {
	lock   *sync.RWMutex
	ticker time.Ticker
	topics map[string]*Topic
}

// GetOrCreateTopic : get or add topic to event bus
func (e *EventBus) GetOrCreateTopic(name string) *Topic {
	if topic, err := e.GetTopic(name); err != nil {
		return topic
	} else {
		topic := NewTopic(name)

		e.lock.Lock()
		defer e.lock.Unlock()
		e.topics[name] = topic

		return topic
	}
}

// GetTopic : get topic from event bus.
func (e *EventBus) GetTopic(name string) (*Topic, error) {
	e.lock.RLock()
	defer e.lock.RUnlock()

	if t, ok := e.topics[name]; ok {
		return t, nil
	}
	return nil, errors.New("topic not exists")
}

// GetTopics : get topics from event bus.
func (e *EventBus) GetTopics() []*Topic {
	e.lock.RLock()
	defer e.lock.RUnlock()

	topics := make([]*Topic, 0)
	for _, topic := range e.topics {
		topics = append(topics, topic)
	}

	return topics
}

// Send : send message to event bus
func (e *EventBus) Send(msg *Message) {
	topic := e.GetOrCreateTopic(msg.topic)
	topic.addMsg(msg)
}

// Start : start event bus
func (e *EventBus) Start() {
	for range e.ticker.C {
		e.lock.RLock()
		for _, topic := range e.topics {
			go topic.fire()
		}
		e.lock.RUnlock()
	}
}
