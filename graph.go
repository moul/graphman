package graphman

import (
	"fmt"
	"sort"
	"strings"
)

type Graph struct {
	Vertices Vertices
	Edges    Edges
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
