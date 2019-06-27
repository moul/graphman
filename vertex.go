package graphman

import (
	"fmt"
)

type Vertex struct {
	ID    string
	Attrs Attrs
	Edges Edges
}

func (v *Vertex) String() string {
	ret := v.ID
	if !v.Attrs.IsEmpty() {
		ret += fmt.Sprintf("[%s]", v.Attrs)
	}
	return ret
}

type Vertices []*Vertex

func (v Vertices) Len() int           { return len(v) }
func (v Vertices) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v Vertices) Less(i, j int) bool { return v[i].ID < v[j].ID }
