package utils

import (
	"OWLm/internal/makefile"
)

func ConvertRawManifest(r *makefile.RawManifest) *makefile.Manifest {
	return &makefile.Manifest{
		Name:         r.Name,
		Path:         r.Path,
		Dependencies: make([]*makefile.Dependency, 0),
	}
}
