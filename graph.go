package graphman

import (
	"fmt"
	"sort"
	"strings"
)

type Graph interface {
	fmt.Stringer

	AddNode(Node) error
	AddEdge(Edge) error
	Nodes() []Node
	Edges() []Edge

	DirectConnectionsFor(id string) []string
	EdgesFor(id string) []Edge
	AllConnectionsFor(id string) []string
	AreConnected(a, b string) bool
	AreDirectlyConnected(a, b string) bool
	AllShortestPaths(id string) map[string]Path
	ShortestPath(a, b string) Path
	// AllPaths(a, b string) []Path
}

type graph struct {
	nodes []Node
	edges []Edge
}

func (g *graph) Nodes() []Node { return g.nodes }

func (g *graph) Edges() []Edge { return g.edges }

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

func (g *graph) EdgesFor(id string) []Edge {
	edges := make([]Edge, 0)
	for _, edge := range g.edges {
		if edge.HasNode(id) {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (g *graph) DirectConnectionsFor(id string) []string {
	set := make(map[string]struct{})
	for _, edge := range g.EdgesFor(id) {
		for _, end := range edge.IDs() {
			if id == end {
				continue
			}
			set[end] = struct{}{}
		}
	}
	ids := []string{}
	for id := range set {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func (g *graph) AreDirectlyConnected(a, b string) bool {
	for _, connection := range g.DirectConnectionsFor(a) {
		if connection == b {
			return true
		}
	}
	return false
}

func (g *graph) AreConnected(a, b string) bool {
	for _, connection := range g.AllConnectionsFor(a) {
		if connection == b {
			return true
		}
	}
	return false
}

func (g *graph) ShortestPath(a, b string) Path {
	for end, path := range g.AllShortestPaths(a) {
		if end == b {
			return path
		}
	}
	return nil
}

func (g *graph) AllShortestPaths(id string) map[string]Path {
	paths := map[string]Path{}
	paths[id] = newPath(id)
	g.allShortestPathsRec(paths, id, 0)
	return paths
}

func (g *graph) allShortestPathsRec(paths map[string]Path, currentID string, currentDepth int) {
	currentPath := paths[currentID]
	for _, edge := range g.EdgesFor(currentID) {
		end := edge.OtherEnd(currentID)
		if _, found := paths[end]; found {
		} else {
			paths[end] = currentPath.Clone()
			paths[end].AddTail(edge)
			g.allShortestPathsRec(paths, end, currentDepth+1)
		}
	}
}

func (g *graph) AllConnectionsFor(id string) []string {
	ids := []string{}
	for end := range g.AllShortestPaths(id) {
		ids = append(ids, end)
	}
	sort.Strings(ids)
	return ids
}

func New() Graph {
	return &graph{}
}
