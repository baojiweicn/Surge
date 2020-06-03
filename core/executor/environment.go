package executor

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
	manager SourceManager
}

// SourceManager : is the interface of manager of language required source.
type SourceManager interface {
	Path() string
	Check(string) error   // Check if required source is all installed.
	Install(string) error // Install all required source.
	Update(string) error  // Update all required source.
}
