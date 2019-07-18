package graphman

import (
	"fmt"
	"strings"
)

//
// Path
//

type Path Edges

func (p Path) String() string {
	if p == nil {
		return invalidPlaceholder
	}
	vertices := p.Vertices()
	ids := []string{}
	for _, vertex := range vertices {
		ids = append(ids, vertex.id)
	}
	ret := fmt.Sprintf("(%s)", strings.Join(ids, ","))
	if !p.IsValid() {
		ret += invalidPlaceholder
	}
	return ret
}

func (p Path) IsValid() bool {
	for i := 0; i < len(p)-1; i++ {
		if p[i].dst != p[i+1].src {
			return false
		}
	}
	return true
}

func (p Path) Vertices() Vertices {
	if len(p) < 1 {
		return Vertices{}
	}
	vertices := Vertices{p.FirstVertex()}
	for _, edge := range p {
		vertices = append(vertices, edge.dst)
	}
	return vertices
}

func (p Path) HasVertex(id string) bool {
	if len(p) < 1 {
		return false
	}
	if p[0].src.id == id {
		return true
	}
	for _, e := range p {
		if e.dst.id == id {
			return true
		}
	}
	return false
}

func (p Path) FirstEdge() *Edge     { return p[0] }
func (p Path) LastEdge() *Edge      { return p[len(p)-1] }
func (p Path) FirstVertex() *Vertex { return p.FirstEdge().src }
func (p Path) LastVertex() *Vertex  { return p.LastEdge().dst }

func (p Path) Edges() Edges { return Edges(p) }

//
// Paths
//

type Paths []*Path

func (p Paths) String() string {
	items := []string{}
	for _, path := range p {
		items = append(items, path.String())
	}
	return strings.Join(items, ",")
}

func (p Paths) Len() int           { return len(p) }
func (p Paths) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Paths) Less(i, j int) bool { return p[i].String() < p[j].String() }
