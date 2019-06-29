package graphman

import (
	"fmt"
	"sort"
	"strings"
)

type Graph struct {
	vertices Vertices
	edges    Edges
	Attrs
}

func New(attrs ...Attrs) *Graph {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return &Graph{
		vertices: make(Vertices, 0),
		edges:    make(Edges, 0),
		Attrs:    a,
	}
}

func (g Graph) Edges() Edges {
	// FIXME: sort.Sort(g.edges)
	return g.edges
}

func (g Graph) Vertices() Vertices {
	sort.Sort(g.vertices)
	return g.vertices
}

func (g *Graph) AddVertex(id string, attrs ...Attrs) *Vertex {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	}

	v := g.GetVertex(id)
	if v != nil {
		v.Attrs.Merge(a)
	} else {
		v = newVertex(id, a)
		g.vertices = append(g.vertices, v)
	}

	return v
}

func (g *Graph) RemoveVertex(id string) bool {
	for k, v := range g.vertices {
		if v.id == id {
			g.vertices = append(g.vertices[:k], g.vertices[k+1:]...)
			return true
		}
	}
	return false
}

func (g Graph) GetVertex(id string) *Vertex {
	for _, v := range g.vertices {
		if v.id == id {
			return v
		}
	}
	return nil
}

func (g *Graph) AddEdge(srcID, dstID string, attrs ...Attrs) *Edge {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	}

	src := g.AddVertex(srcID)
	dst := g.AddVertex(dstID)
	edge := newEdge(src, dst, a)
	src.successors = append(src.successors, edge)
	dst.predecessors = append(src.predecessors, edge)
	g.edges = append(g.edges, edge)
	return edge
}

func (g Graph) IsolatedVertices() Vertices {
	isolated := Vertices{}
	for _, vertex := range g.vertices {
		if len(vertex.Edges()) == 0 {
			isolated = append(isolated, vertex)
		}
	}
	sort.Sort(isolated)
	return isolated
}

func (g Graph) String() string {
	elems := []string{}
	for _, edge := range g.edges {
		elems = append(elems, edge.String())
	}
	for _, vertex := range g.IsolatedVertices() {
		elems = append(elems, vertex.id)
	}
	return fmt.Sprintf("{%s}", strings.Join(elems, ","))
}
