package graphman

import (
	"fmt"
	"strings"
)

type Edge struct {
	src *Vertex
	dst *Vertex
	Attrs
}

func newEdge(src, dst *Vertex, attrs ...Attrs) *Edge {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	}
	return &Edge{
		src:   src,
		dst:   dst,
		Attrs: a,
	}
}

func (e *Edge) Vertices() Vertices {
	return Vertices{e.src, e.dst}
}

func (e *Edge) String() string {
	ret := fmt.Sprintf("(%s,%s)", e.src.id, e.dst.id)
	if !e.Attrs.IsEmpty() {
		ret += fmt.Sprintf("[%s]", e.Attrs)
	}
	return ret
}

func (e *Edge) HasVertex(id string) bool {
	return e.src.id == id || e.dst.id == id
}

func (e *Edge) OtherVertex(id string) *Vertex {
	if e.src.id == id {
		return e.dst
	}
	if e.dst.id == id {
		return e.src
	}
	return nil
}

type Edges []*Edge

func (e Edges) String() string {
	items := []string{}
	for _, edge := range e {
		items = append(items, edge.String())
	}
	return fmt.Sprintf("{%s}", strings.Join(items, ","))
}
