package vendor

import (
	"gopkg.in/yaml.v2"
)

func (t *Template) Decode(in []byte) {
	yaml.Unmarshal(in, t)
}

type Device struct {
	name        string
	environment string
	types       []string
	tags        []string
}

func (t *Template) GetDevice() *Device {
	return &Device{}
}
