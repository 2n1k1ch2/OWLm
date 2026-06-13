package graph

import "OWLm/internal/makefile"

type Node struct {
	makefile.Manifest
}

func NewNode(manifest makefile.Manifest) *Node {
	return &Node{manifest}
}
