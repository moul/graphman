package graphman

import (
	"fmt"
	"strings"
)

type Path interface {
	fmt.Stringer

	Clone() Path
	AddTail(Edge)
	Edges() []Edge
}

type path struct {
	edges []Edge
	start string
}

func (p *path) Length() int {
	return len(p.edges)
}

func (p *path) String() string {
	chain := []string{p.start}
	last := p.start
	for _, edge := range p.edges {
		end := edge.OtherEnd(last)
		chain = append(chain, end)
		last = end
	}
	return strings.Join(chain, "->")
}

func (p *path) Edges() []Edge {
	return p.edges
}

func newPath(start string) Path {
	return &path{
		start: start,
		edges: make([]Edge, 0),
	}
}

func (p *path) Clone() Path {
	clone := &path{
		start: p.start,
		edges: make([]Edge, 0),
	}
	for _, edge := range p.Edges() {
		clone.edges = append(clone.edges, edge)
	}
	return clone
}

func (p *path) AddTail(edge Edge) {
	p.edges = append(p.edges, edge)
}
