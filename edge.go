package graphman

import (
	"fmt"
	"strings"
)

type Edge struct {
	Src *Vertex
	Dst *Vertex
	Attrs
}

func (e *Edge) Vertices() Vertices {
	return Vertices{e.Src, e.Dst}
}

func (e *Edge) String() string {
	ret := fmt.Sprintf("(%s,%s)", e.Src.ID, e.Dst.ID)
	if !e.Attrs.IsEmpty() {
		ret += fmt.Sprintf("[%s]", e.Attrs)
	}
	return ret
}

func (e *Edge) HasVertex(id string) bool {
	return e.Src.ID == id || e.Dst.ID == id
}

func (e *Edge) OtherVertex(id string) *Vertex {
	if e.Src.ID == id {
		return e.Dst
	}
	if e.Dst.ID == id {
		return e.Src
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
