package graphman

import (
	"fmt"
	"strings"
)

type Vertex struct {
	id           string
	successors   Edges
	predecessors Edges
	Attrs
}

func newVertex(id string, attrs ...Attrs) *Vertex {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return &Vertex{
		id:           id,
		Attrs:        a,
		successors:   make(Edges, 0),
		predecessors: make(Edges, 0),
	}
}

func (v *Vertex) Edges() Edges {
	return append(v.predecessors, v.successors...)
}

func (v *Vertex) ID() string { return v.id }

func (v *Vertex) String() string {
	ret := v.id
	if !v.Attrs.IsEmpty() {
		ret += fmt.Sprintf("[%s]", v.Attrs)
	}
	return ret
}

type Vertices []*Vertex

func (v Vertices) String() string {
	ids := []string{}
	for _, vertex := range v {
		ids = append(ids, vertex.id)
	}
	return fmt.Sprintf("{%s}", strings.Join(ids, ","))
}

func (v Vertices) Len() int           { return len(v) }
func (v Vertices) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v Vertices) Less(i, j int) bool { return v[i].id < v[j].id }
