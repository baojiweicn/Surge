package source

import (
	"encoding/json"

	"github.com/baojiweicn/Surge/util/errors"
	"github.com/baojiweicn/Surge/util/parser"
)

var (
	// CheckInstalledCommand is to get package installed and get the version
	CheckInstalledCommand = parser.NewTemplate("{{pip}} list --format json")
)

// PythonManager : is the python package source manager -> pip
type PythonManager struct {
	path string
}

func NewPythonManager(path string) *PythonManager {
	return &PythonManager{
		path: path,
	}
}

func (m *PythonManager) Path() string {
	return m.path
}

func (m *PythonManager) Check(pack *Package) error {
	packs := make([]*Package, 0)
	err := json.Unmarshal([]byte(CheckInstalledCommand.Render([]parser.Field{
		parser.F("pip", m.Path()),
	})), packs)
	if err != nil {
		return PackageNotInstalledError.Raise(
			[]errors.Field{
				errors.F("package", pack.Name),
			},
		)
	} else {
		for _, p := range packs {
			if p.Name == pack.Name {
				if p.Version == pack.Version {
					return nil
				} else {
					return PackageVersionNotMatchError.Raise(
						[]errors.Field{
							errors.F("package", pack.Name),
							errors.F("current", p.Version),
							errors.F("want", pack.Version),
						},
					)
				}
			}
		}
	}
	return PackageNotInstalledError.Raise(
		[]errors.Field{
			errors.F("package", pack.Name),
		},
	)
}

func (m *PythonManager) Install(pack *Package) error {
	return nil
}

func (m *PythonManager) Update(pack *Package) error {
	return nil
}

func (m *PythonManager) Packages() []string {
	return make([]string, 0)
}
