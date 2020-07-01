package source

import (
	"github.com/baojiweicn/Surge/util/errors"
	"github.com/baojiweicn/Surge/util/util"
)

// Manager : is the interface of manager of language required source.
type Manager interface {
	Path() string                // get the source path
	GetAll() ([]*Package, error) // get all packages
	Check(Package) error         // check if required source is all installed.
	Install(Package) error       // install all required source.
	Uninstall(Package) error     // uninstall package
	Update(Package) error        // update all required source.
}

// Package : is the struct for a new package
type Package struct {
	manager Manager
	Name    string
	Version string
}

// NewPackage : create new Package
func NewPackage(manager Manager, name, version string) Package {
	return Package{
		manager: manager,
		Name:    name,
		Version: version,
	}
}

// Install : install the package
func (p Package) Install() error {
	if !util.IsNilInterface(p.manager) {
		return p.manager.Install(p)
	}
	return SourceNotExistsError.Raise(
		[]errors.Field{
			errors.F("language", p.Name),
		},
	)
}

// Installed : if the package already installed
func (p Package) Installed() bool {
	return false
}
