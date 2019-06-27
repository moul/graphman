package graphman

import "fmt"

type Edge interface {
	fmt.Stringer

	HasVertex(id string) bool
	IDs() []string
	OtherEnd(id string) string
}

type edge struct {
	a, b string
}

func (e *edge) String() string {
	return fmt.Sprintf("%s->%s", e.a, e.b)
}

func (e *edge) HasVertex(id string) bool {
	return e.a == id || e.b == id
}

func (e *edge) IDs() []string {
	return []string{e.a, e.b}
}

func (e *edge) OtherEnd(id string) string {
	if e.a == id {
		return e.b
	}
	if e.b == id {
		return e.a
	}
	return ""
}

func NewEdge(a, b string) Edge {
	return &edge{a: a, b: b}
}
