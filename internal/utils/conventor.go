package utils

import (
	"OWLm/internal/makefile"
)

func ConvertRawManifest(r *makefile.RawManifest) *makefile.Manifest {
	return &makefile.Manifest{
		Name:             r.Name,
		Path:             r.Path,
		Depends:          make([]*makefile.Dependency, 0),
		BuildDepends:     make([]*makefile.Dependency, 0),
		HostBuildDepends: make([]*makefile.Dependency, 0),
	}
}
