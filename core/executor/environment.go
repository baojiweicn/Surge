package executor

import (
	"github.com/baojiweicn/Surge/core/executor/source"
)

type EnvType uint16

var (
	EnvUnknownType EnvType = 0x00
	GoModEnvType           = 0x10
	PipEnvType             = 0x20
	NpmEnvType             = 0x30
	YarnEnvType            = 0x31
)

// Environment : is the control of language source manager.
type Environment struct {
	envType EnvType
	manager source.SourceManager
}
