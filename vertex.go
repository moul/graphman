package graphman

import (
	"fmt"
	"sort"
	"strings"
)

//
// Vertex
//

type Vertex struct {
	id           string
	successors   Edges
	predecessors Edges
	Attrs
	dijkstra struct {
		dist    int64
		prev    *Edge
		visited bool
	}
}

func newVertex(id string, attrs ...Attrs) *Vertex {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	} else {
		a = make(map[string]interface{})
	}
	return &Vertex{
		id:           id,
		Attrs:        a,
		successors:   make(Edges, 0),
		predecessors: make(Edges, 0),
	}
}

func (v Vertex) IsSource() bool          { return v.InDegree() == 0 }
func (v Vertex) IsSink() bool            { return v.OutDegree() == 0 }
func (v Vertex) InDegree() int           { return len(v.predecessors) }
func (v Vertex) OutDegree() int          { return len(v.successors) }
func (v Vertex) Degree() int             { return v.InDegree() + v.OutDegree() }
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

type VerticesWalkFunc func(current, previous *Vertex, depth int) error

func (v *Vertex) WalkSuccessorVertices(fn VerticesWalkFunc) error {
	visited := map[string]bool{}
	return v.walkSuccessorVerticesRec(fn, nil, 0, visited)
}

func (v *Vertex) walkSuccessorVerticesRec(fn VerticesWalkFunc, previous *Vertex, depth int, visited map[string]bool) error {
	if visited[v.id] {
		return nil
	}
	visited[v.id] = true
	if err := fn(v, previous, depth); err != nil {
		return err
	}
	for _, successor := range v.successors {
		if err := successor.dst.walkSuccessorVerticesRec(fn, v, depth+1, visited); err != nil {
			return err
		}
	}
	return nil
}

func (v *Vertex) WalkPredecessorVertices(fn VerticesWalkFunc) error {
	visited := map[string]bool{}
	return v.walkPredecessorVerticesRec(fn, nil, 0, visited)
}

func (v *Vertex) walkPredecessorVerticesRec(fn VerticesWalkFunc, previous *Vertex, depth int, visited map[string]bool) error {
	if visited[v.id] {
		return nil
	}
	visited[v.id] = true
	if err := fn(v, previous, depth); err != nil {
		return err
	}
	for _, predecessor := range v.predecessors {
		if err := predecessor.dst.walkPredecessorVerticesRec(fn, v, depth+1, visited); err != nil {
			return err
		}
	}
	return nil
}

func (v *Vertex) WalkAdjacentVertices(fn VerticesWalkFunc) error {
	visited := map[string]bool{}
	return v.walkAdjacentVerticesRec(fn, nil, 0, visited)
}

func (v *Vertex) walkAdjacentVerticesRec(fn VerticesWalkFunc, previous *Vertex, depth int, visited map[string]bool) error {
	if visited[v.id] {
		return nil
	}
	visited[v.id] = true
	if err := fn(v, previous, depth); err != nil {
		return err
	}
	for _, successor := range v.successors {
		if err := successor.dst.walkAdjacentVerticesRec(fn, v, depth+1, visited); err != nil {
			return err
		}
	}
	for _, predecessor := range v.predecessors {
		if err := predecessor.src.walkAdjacentVerticesRec(fn, v, depth+1, visited); err != nil {
			return err
		}
	}
	return nil
}

//
// Vertices
//

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
