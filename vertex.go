package graphman

import (
	"fmt"
	"sort"
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

func (v Vertex) PredecessorEdges() Edges { return v.predecessors }
func (v Vertex) SuccessorEdges() Edges   { return v.successors }

func (v Vertex) PredecessorVertices() Vertices {
	vertices := Vertices{}
	for _, edge := range v.predecessors {
		vertices = append(vertices, edge.src)
	}
	return vertices.Unique()
}

func (v Vertex) SuccessorVertices() Vertices {
	vertices := Vertices{}
	for _, edge := range v.successors {
		vertices = append(vertices, edge.dst)
	}
	return vertices.Unique()
}

func (v Vertex) IsIsolated() bool {
	return len(v.predecessors) == 0 && len(v.successors) == 0
}

func (v Vertex) Edges() Edges {
	return append(v.predecessors, v.successors...)
}

func (v Vertex) Neighbors() Vertices {
	neighbors := Vertices{}
	for _, edge := range v.predecessors {
		neighbors = append(neighbors, edge.src)
	}
	for _, edge := range v.successors {
		neighbors = append(neighbors, edge.dst)
	}
	return neighbors.Unique()
}

func (v Vertex) ID() string { return v.id }

func (v Vertex) String() string {
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

func (v Vertices) Unique() Vertices {
	m := map[string]*Vertex{}
	for _, v := range v {
		m[v.id] = v
	}
	uniques := Vertices{}
	for _, v := range m {
		uniques = append(uniques, v)
	}
	sort.Sort(uniques)
	return uniques
}
