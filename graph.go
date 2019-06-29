package graphman

import (
	"fmt"
	"sort"
	"strings"
)

type Graph struct {
	Vertices Vertices
	Edges    Edges
	Attrs
}

func New(attrs ...Attrs) *Graph {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return &Graph{
		Vertices: make(Vertices, 0),
		Edges:    make(Edges, 0),
		Attrs:    a,
	}
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
		v = &Vertex{ID: id, Attrs: a}
	}

	g.Vertices = append(g.Vertices, v)
	return v
}

func (g *Graph) RemoveVertex(id string) bool {
	for k, v := range g.Vertices {
		if v.ID == id {
			g.Vertices = append(g.Vertices[:k], g.Vertices[k+1:]...)
			return true
		}
	}
	return false
}

func (g Graph) GetVertex(id string) *Vertex {
	for _, v := range g.Vertices {
		if v.ID == id {
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
	edge := &Edge{Src: src, Dst: dst, Attrs: a}
	g.Edges = append(g.Edges, edge)
	return edge
}

func (g Graph) IsolatedVertices() Vertices {
	isolatedVertices := map[string]*Vertex{}
	for _, vertex := range g.Vertices {
		isolatedVertices[vertex.ID] = vertex
	}
	for _, edge := range g.Edges {
		isolatedVertices[edge.Src.ID] = nil
		isolatedVertices[edge.Dst.ID] = nil
	}
	filtered := Vertices{}
	for _, vertex := range isolatedVertices {
		if vertex != nil {
			filtered = append(filtered, vertex)
		}
	}

	sort.Sort(filtered)
	return filtered
}

func (g Graph) String() string {
	elems := []string{}
	for _, edge := range g.Edges {
		elems = append(elems, edge.String())
	}
	for _, vertex := range g.IsolatedVertices() {
		elems = append(elems, vertex.ID)
	}
	return fmt.Sprintf("{%s}", strings.Join(elems, ","))
}
