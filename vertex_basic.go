package graphman

import (
	"fmt"
	"strings"
)

// stringVertex uses standard Go slices for internal storage.
type stringVertex struct {
	VertexWithAttributes

	id    string
	attrs []vertexAttr
	edges []Edge
}

// NewVertex constructs an Vertex
func NewStringVertex(id string) VertexWithAttributes {
	vertex := stringVertex{id: id}
	return &vertex
}

func (n *stringVertex) ID() string { return n.id }

func (n *stringVertex) CleanEdges() error {
	n.edges = nil
	return nil
}

func (n *stringVertex) CleanAttrs() error {
	n.attrs = nil
	return nil
}

func (n *stringVertex) AddEdge(edge Edge) error {
	if n.edges == nil {
		n.edges = make([]Edge, 0)
	}
	n.edges = append(n.edges, edge)
	return nil
}

func (n *stringVertex) HasAttr(key interface{}) (bool, error) {
	if n.attrs == nil {
		return false, nil
	}
	for _, attr := range n.attrs {
		if attr[0] == key {
			return true, nil
		}
	}
	return false, nil
}

func (n *stringVertex) GetAttr(key interface{}) (interface{}, error) {
	if n.attrs == nil {
		return nil, nil
	}
	for _, attr := range n.attrs {
		if attr[0] == key {
			return attr[1], nil
		}
	}
	return nil, nil
}

func (n *stringVertex) SetAttr(key, value interface{}) error {
	if n.attrs == nil {
		n.attrs = make([]vertexAttr, 0)
	}
	for _, attr := range n.attrs {
		if attr.Key() == key {
			attr.SetValue(value)
			return nil
		}
	}
	attr := [2]interface{}{key, value}
	n.attrs = append(n.attrs, attr)
	return nil
}

func (n *stringVertex) DelAttr(key interface{}) error {
	for idx, attr := range n.attrs {
		if attr.Key() == key {
			n.attrs = append(n.attrs[:idx], n.attrs[idx+1:]...)
			return nil
		}
	}
	return nil
}

func (n *stringVertex) String() string {
	if n.attrs == nil {
		return string(n.id)
	}
	attrs := []string{}
	for _, attr := range n.attrs {
		attrs = append(attrs, fmt.Sprintf("%v:%v", attr.Key(), attr.Value()))
	}
	return fmt.Sprintf("%s(%s)", n.id, strings.Join(attrs, ","))
}

func (n *stringVertex) Edges() ([]Edge, error) {
	return n.edges, nil
}

// vertexAttr uses a standard Go slice for internal storage.
type vertexAttr [2]interface{}

func (va *vertexAttr) Key() interface{}              { return va[0] }
func (va *vertexAttr) Value() interface{}            { return va[1] }
func (va *vertexAttr) SetValue(newValue interface{}) { va[1] = newValue }
