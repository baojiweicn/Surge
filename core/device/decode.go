package device

import (
	"gopkg.in/yaml.v2"
)

func (t *Template) Decode(in []byte) {
	yaml.Unmarshal(in, t)
}

func (t *Template) Encode() []byte {
	return yaml.Marshal(t)
}
