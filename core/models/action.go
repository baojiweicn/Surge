package models

type Action struct {
	name     string
	executor string
	params   []Param
	message  []Message
}

type Param struct {
	name      string
	value     interface{}
	valueType string
}
