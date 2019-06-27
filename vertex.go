package graphman

import (
	"fmt"
	"strings"
)

type Vertex interface {
	fmt.Stringer

	ID() string

	AddEdge(edge Edge)
	Edges() []Edge
	CleanEdges()

	SetAttr(Attr)
	DelAttr(key interface{})
	GetAttr(key interface{}) interface{}
	HasAttr(key interface{}) bool
	CleanAttrs()
}

type vertex struct {
	Vertex

	id    string
	attrs []Attr
	edges []Edge
}

func NewVertex(id string) Vertex {
	vertex := vertex{id: id}
	return &vertex
}

func (n *vertex) ID() string { return n.id }

func (n *vertex) CleanEdges() {
	n.edges = nil
}

func (n *vertex) CleanAttrs() {
	n.attrs = nil
}

func (n *vertex) AddEdge(edge Edge) {
	if n.edges == nil {
		n.edges = make([]Edge, 0)
	}
	n.edges = append(n.edges, edge)
}

func (n *vertex) HasAttr(key interface{}) bool {
	if n.attrs == nil {
		return false
	}
	for _, attr := range n.attrs {
		if attr.Key() == key {
			return true
		}
	}
	return false
}

func (n *vertex) GetAttr(key interface{}) interface{} {
	if n.attrs == nil {
		return nil
	}
	for _, attr := range n.attrs {
		if attr.Key() == key {
			return attr.Value()
		}
	}
	return nil
}

func (n *vertex) SetAttr(attr Attr) {
	n.DelAttr(attr.Key())
	if n.attrs == nil {
		n.attrs = make([]Attr, 0)
	}
	n.attrs = append(n.attrs, attr)
}

func (n *vertex) DelAttr(key interface{}) {
	if n.attrs == nil {
		return
	}
	for idx, attr := range n.attrs {
		if attr.Key() == key {
			n.attrs = append(n.attrs[:idx], n.attrs[idx+1:]...)
			return
		}
	}
}

func (n *vertex) String() string {
	if n.attrs == nil {
		return string(n.id)
	}
	attrs := []string{}
	for _, attr := range n.attrs {
		attrs = append(attrs, fmt.Sprintf("%v:%v", attr.Key(), attr.Value()))
	}
	return fmt.Sprintf("%s(%s)", n.id, strings.Join(attrs, ","))
}

func (n *vertex) Edges() []Edge {
	return n.edges
}

type VertexSorter []Vertex

func (v VertexSorter) Len() int           { return len(v) }
func (v VertexSorter) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v VertexSorter) Less(i, j int) bool { return v[i].ID() < v[j].ID() }
