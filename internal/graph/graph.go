package graph

import (
	"OWLm/internal/makefile"
	"OWLm/internal/utils"
	"fmt"
)

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}
func (g *Graph) AddNode(n *Node) {
	if g.Nodes == nil {
		panic("Graph is not initialized")
	}
	if _, ok := g.Nodes[n.Name]; ok {
		fmt.Printf("graph node %q already exists\n", n.Name)
	}

	g.Nodes[n.Name] = n
}
func (g *Graph) RemoveNode(n *Node) {
	if g.Nodes == nil {
		panic("Graph is not initialized")
	}
	if _, ok := g.Nodes[n.Name]; ok {
		delete(g.Nodes, n.Name)
	}
}

func BuildGraph(list []makefile.RawManifest) (*Graph, error) {
	g := NewGraph()
	if len(list) == 0 {
		return nil, fmt.Errorf("no Makefiles to build")
	}
	for _, raw := range list {
		m := utils.ConvertRawManifest(&raw)
		g.AddNode(NewNode(*m))
	}

	return g, nil
}

func BuildDepends(graph *Graph, list []makefile.RawManifest) error {
	for _, raw := range list {
		current := graph.Nodes[raw.Name]

		addDeps(&current.Manifest, raw.Depends, makefile.RuntimeDependency, graph)
		addDeps(&current.Manifest, raw.BuildDepends, makefile.BuildDependency, graph)
		addDeps(&current.Manifest, raw.HostBuildDepends, makefile.HostBuildDependency, graph)
	}

	return nil
}
func addDeps(current *makefile.Manifest, deps []makefile.RawDependency, kind makefile.DependencyType, graph *Graph) {
	for _, dep := range deps {
		target, ok := graph.Nodes[dep.Name]
		if !ok {
			continue
		}

		current.Dependencies = append(current.Dependencies, &makefile.Dependency{
			Type:      kind,
			Condition: dep.Condition,
			Target:    &target.Manifest,
		})
	}
}
