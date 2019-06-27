package graphman

import (
	"fmt"
)

type Vertex interface {
	fmt.Stringer

	ID() string
	AddEdge(edge Edge) error
	Edges() ([]Edge, error)
	CleanEdges() error
}

type VertexWithAttributes interface {
	Vertex

	SetAttr(key, value interface{}) error
	DelAttr(key interface{}) error
	GetAttr(key interface{}) (interface{}, error)
	HasAttr(key interface{}) (bool, error)
	CleanAttrs() error
}
