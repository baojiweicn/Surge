package device

import (
	"gopkg.in/yaml.v2"
)

type Template struct {
	Device  DeviceTemplate            `yaml:"device"`
	Actions map[string]ActionTemplate `yaml:"actions"`
	Objects map[string]ObjectTemplate `yaml:"objects"`
}

type DeviceTemplate struct {
	Name        string         `yaml:"name"`
	Environment string         `yaml:"environment"`
	Types       []string       `yaml:"types,omitempty"`
	Tags        []string       `yaml:"tags,omitempty"`
	Infos       []yaml.MapItem `yaml:"infos,omitempty"`
	Storage     []yaml.MapItem `yaml:"storage,omitempty"`
}

type ActionTemplate struct {
	Name    string         `yaml:"name"`
	Exector string         `yaml:"executor"`
	Params  []yaml.MapItem `yaml:"params"`
	Message string         `yaml:"message"`
}

type ObjectTemplate map[string]string
