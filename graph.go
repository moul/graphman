package graphman

import (
	"strings"
)

type Graph interface {
	AddNode(Node) error
	AddEdge(Edge) error
	String() string
}

type graph struct {
	nodes []Node
	edges []Edge
}

func (g *graph) String() string {
	lines := []string{}
	for _, node := range g.nodes {
		lines = append(lines, node.String())
	}
	for _, edge := range g.edges {
		lines = append(lines, edge.String())
	}
	return strings.Join(lines, "\n")
}

func (g *graph) AddNode(n Node) error {
	g.nodes = append(g.nodes, n)
	return nil
}

func (g *graph) AddEdge(e Edge) error {
	g.edges = append(g.edges, e)
	return nil
}

func New() Graph {
	return &graph{}
}
