package graphman

import "fmt"

type Edge interface {
	fmt.Stringer
}

type edge struct {
	a, b string
}

func (e *edge) String() string {
	return fmt.Sprintf("%s->%s", e.a, e.b)
}

func NewEdge(a, b string) Edge {
	return &edge{a: a, b: b}
}
