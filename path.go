package graphman

import (
	"fmt"
	"strings"
)

type Path interface {
	fmt.Stringer

	Length() int
	Append(Edge)
	Edges() []Edge
	FirstVertex() Vertex
	LastVertex() Vertex
	FirstEdge() Edge
	LastEdge() Edge
}

type path struct {
	edges []Edge
}

func (p *path) Length() int {
	return len(p.edges)
}

func (p *path) String() string {
	ids := []string{p.FirstVertex().ID()}
	for _, edge := range p.edges {
		ids = append(ids, edge.Dst().ID())
	}
	return fmt.Sprintf("(%s)", strings.Join(ids, ","))
}

func (p *path) Edges() []Edge {
	return p.edges
}

func NewPath(firstEdge Edge) Path {
	return &path{
		edges: []Edge{
			firstEdge,
		},
	}
}

func (p *path) Append(edge Edge) {
	p.edges = append(p.edges, edge)
}

func (p *path) FirstEdge() Edge     { return p.edges[0] }
func (p *path) LastEdge() Edge      { return p.edges[len(p.edges)-1] }
func (p *path) FirstVertex() Vertex { return p.FirstEdge().Src() }
func (p *path) LastVertex() Vertex  { return p.LastEdge().Dst() }
