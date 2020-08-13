package models

type Message struct {
	name  string
	items map[string]MessageItem
}

type MessageItem interface {
	Type() string
	Value() interface{}
}
