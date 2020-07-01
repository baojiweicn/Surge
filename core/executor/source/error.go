package source

import (
	"github.com/baojiweicn/Surge/util/errors"
)

var (
	// SourceError : error happened in source manager
	SourceError = errors.NewError("environment {{language}} package manager error {{error}}")
	// SourceNotExistsError : is the error for package manager not founc
	SourceNotExistsError = errors.NewError("environment {{language}} package manager not exists")
	// PackageNotInstalledError : is the error for the package not installed
	PackageNotInstalledError = errors.NewError("{{package}} not installed")

	// PackageVersionNotMatchError : the package installed but the version not match
	PackageVersionNotMatchError = errors.NewError("{{package}} current version {{current}} but want {{want}}")
	// PackageInstallError : the package install failed
	PackageInstallError = errors.NewError("environment {{language}} install {{package}} : {{version}} error {{error}}")
	// PackageUninstallError : the package install failed
	PackageUninstallError = errors.NewError("environment {{language}} uninstall {{package}} error {{error}}")
)
