package makefile

type Manifest struct {
	Name string
	Path string

	Dependencies []*Dependency
}
type DependencyType int

const (
	RuntimeDependency DependencyType = iota
	BuildDependency
	HostBuildDependency
)

type Dependency struct {
	Target    *Manifest
	Type      DependencyType
	Condition string
}

func NewDependency(condition string, target *Manifest) *Dependency {
	return &Dependency{Condition: condition, Target: target}
}

type RawManifest struct {
	Name string
	Path string

	Depends          []RawDependency
	BuildDepends     []RawDependency
	HostBuildDepends []RawDependency
}

type RawDependency struct {
	Name      string
	Condition string
}
