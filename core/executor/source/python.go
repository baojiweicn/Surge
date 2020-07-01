package source

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"

	"os/exec"

	"github.com/baojiweicn/Surge/util/errors"
	"github.com/baojiweicn/Surge/util/parser"
)

var (
	// CheckInstalledCommand is to get package installed and get the version
	CheckInstalledCommand = NewCommand("--disable-pip-version-check", "list", "--format", "json")
	// InstallPackageCommand is to install package
	InstallPackageCommand = NewCommand("{{name}}", "==", "{{version}}")
)

// PythonManager : is the python package source manager -> pip
type PythonManager struct {
	path string
}

// NewPythonManager : create new python manager
func NewPythonManager(path string) *PythonManager {
	return &PythonManager{
		path: path,
	}
}

// GetDefaultPythonManager : get default python manager
func GetDefaultPythonManager() *PythonManager {
	cmd := exec.Command("which", "pip")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil
	}

	path := strings.Replace(out.String(), "\n", "", -1)
	return &PythonManager{path: path}
}

// Path : get python manager path
func (m *PythonManager) Path() string {
	return m.path
}

// Get : get the package
func (m *PythonManager) Get(pack *Package) (*Package, error) {
	packs, err := m.GetAll()
	if err != nil {
		return nil, err
	}
	for _, p := range packs {
		if p.Name == pack.Name {
			return p, nil
		}
	}
	return nil, PackageNotInstalledError.Raise(
		[]errors.Field{
			errors.F("package", pack.Name),
		},
	)
}

// GetAll : get all packages
func (m *PythonManager) GetAll() ([]*Package, error) {
	packs := make([]*Package, 0)
	cmd := exec.Command(m.Path(), CheckInstalledCommand.Render([]parser.Field{
		parser.F("pip", m.Path()),
	})...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return packs, SourceError.Raise(
			[]errors.Field{
				errors.F("language", "python"),
				errors.F("error", err.Error()),
			},
		)
	}
	if match, err := regexp.Match(`\[.*\]`, out); err != nil || !match {
		return packs, SourceError.Raise(
			[]errors.Field{
				errors.F("language", "python"),
				errors.F("error", err.Error()),
			},
		)
	}
	reg, _ := regexp.Compile(`\[.*\]`)
	if err := json.Unmarshal(reg.Find(out), &packs); err != nil {
		return packs, SourceError.Raise(
			[]errors.Field{
				errors.F("language", "python"),
				errors.F("error", err.Error()),
			},
		)
	}
	return packs, nil
}

// Check : check the package is exist
func (m *PythonManager) Check(pack *Package) error {
	// get the package
	p, err := m.Get(pack)
	if err != nil {
		return err
	}
	// if version match return without error
	if p.Version == pack.Version {
		return nil
	}
	// if version not match, return error
	return PackageVersionNotMatchError.Raise(
		[]errors.Field{
			errors.F("package", pack.Name),
			errors.F("current", p.Version),
			errors.F("want", pack.Version),
		},
	)
}

// Install : install the package
func (m *PythonManager) Install(pack *Package) error {
	cmd := exec.Command(m.Path(), InstallPackageCommand.Render([]parser.Field{
		parser.F("name", pack.Name),
		parser.F("version", pack.Version),
	})...)

	if _, err := cmd.CombinedOutput(); err != nil {
		return PackageInstallError.Raise(
			[]errors.Field{
				errors.F("language", "python"),
				errors.F("package", pack.Name),
				errors.F("version", pack.Version),
				errors.F("error", err.Error()),
			},
		)
	}

	return nil
}

// Update : update the package
func (m *PythonManager) Update(pack *Package) error {
	return m.Install(pack)
}

func (m *PythonManager) Packages() []string {
	return make([]string, 0)
}
