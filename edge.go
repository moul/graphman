package graphman

import "fmt"

type Edge interface {
	fmt.Stringer

	Src() Vertex
	Dst() Vertex
	Vertices() []Vertex
	HasVertex(Vertex) bool
	OtherVertex(Vertex) Vertex
}

type edge [2]Vertex

func (e *edge) Src() Vertex        { return e[0] }
func (e *edge) Dst() Vertex        { return e[1] }
func (e *edge) Vertices() []Vertex { return []Vertex{e.Src(), e.Dst()} }

func (e *edge) String() string {
	return fmt.Sprintf("(%s,%s)", e.Src().ID(), e.Dst().ID())
}

func (e *edge) HasVertex(v Vertex) bool {
	return e.Src() == v || e.Dst() == v
}

func (e *edge) OtherVertex(v Vertex) Vertex {
	if e.Src() == v {
		return e.Dst()
	}
	if e.Dst() == v {
		return e.Src()
	}
	return nil
}

func NewEdge(src, dst Vertex) Edge {
	return &edge{src, dst}
}
