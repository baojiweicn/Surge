package vendor

import (
	"gopkg.in/yaml.v2"
)

type Template struct {
	Device   DeviceTemplate `yaml:"device"`
	Actions  []yaml.MapItem `yaml:"actions"`
	Messages []yaml.MapItem `yaml:"messages"`
}

type DeviceTemplate struct {
	Name        string         `yaml:"name"`
	Environment string         `yaml:"environment"`
	Types       []string       `yaml:"types,omitempty"`
	Tags        []string       `yaml:"tags,omitempty"`
	Infos       []yaml.MapItem `yaml:"infos,omitempty"`
	Storage     []yaml.MapItem `yaml:"storage,omitempty"`
}

type Action struct {
	name     string
	executor string
	params   map[string]Param
	message  Message
}

type Message struct {
	name   string
	params map[string]Param
}

type ParamType uint16

const (
	UnknownType ParamType = 0
	StringType  ParamType = 1
	IntType     ParamType = 2
	FloatType   ParamType = 3
	EnumType    ParamType = 4
	ArrayType   ParamType = 5
)

type Param interface {
	Type() ParamType
	Value() (interface{}, ParamType)
	Values() ([]interface{}, ParamType)
	Name() string
}

type ParamInt struct {
}

func (p *ParamInt) Type() ParamType {
	return IntType
}

type ParamString struct {
}

func (p *ParamString) Type() ParamType {
	return StringType
}

type ParamFloat struct {
}

func (p *ParamFloat) Type() ParamType {
	return FloatType
}

type ParamEnum struct {
}

func (p *ParamEnum) Type() ParamType {
	return EnumType
}

type ParamArray struct {
}

func (p *ParamArray) Type() ParamType {
	return ArrayType
}
