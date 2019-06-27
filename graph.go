package graphman

import (
	"fmt"
	"sort"
	"strings"
)

type Graph interface {
	fmt.Stringer

	AddVertex(...Vertex)
	AddEdge(...Edge)
	Vertices() []Vertex
	Edges() []Edge
	IsolatedVertices() []Vertex
}

type graph struct {
	vertices []Vertex
	edges    []Edge
}

func NewGraph() Graph {
	return &graph{
		vertices: make([]Vertex, 0),
		edges:    make([]Edge, 0),
	}
}

func (g *graph) Vertices() []Vertex { return g.vertices }
func (g *graph) Edges() []Edge      { return g.edges }

func (g *graph) IsolatedVertices() []Vertex {
	isolatedVertices := map[string]Vertex{}
	for _, vertex := range g.vertices {
		isolatedVertices[vertex.ID()] = vertex
	}
	for _, edge := range g.edges {
		isolatedVertices[edge.Src().ID()] = nil
		isolatedVertices[edge.Dst().ID()] = nil
	}
	filtered := []Vertex{}
	for _, vertex := range isolatedVertices {
		if vertex != nil {
			filtered = append(filtered, vertex)
		}
	}

	sort.Sort(VertexSorter(filtered))
	return filtered
}

func (g *graph) String() string {
	elems := []string{}
	for _, edge := range g.edges {
		elems = append(elems, edge.String())
	}
	for _, vertex := range g.IsolatedVertices() {
		elems = append(elems, vertex.ID())
	}
	return fmt.Sprintf("{%s}", strings.Join(elems, ","))
}

func (g *graph) AddVertex(vertices ...Vertex) {
	g.vertices = append(g.vertices, vertices...)
}

func (g *graph) AddEdge(edges ...Edge) {
	g.edges = append(g.edges, edges...)
}
