package graphman

import (
	"fmt"
	"strings"
)

type Vertex struct {
	ID    string
	Edges Edges
	Attrs
}

func (v *Vertex) String() string {
	ret := v.ID
	if !v.Attrs.IsEmpty() {
		ret += fmt.Sprintf("[%s]", v.Attrs)
	}
	return ret
}

type Vertices []*Vertex

func (v Vertices) String() string {
	ids := []string{}
	for _, vertex := range v {
		ids = append(ids, vertex.ID)
	}
	return fmt.Sprintf("{%s}", strings.Join(ids, ","))
}

func (v Vertices) Len() int           { return len(v) }
func (v Vertices) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v Vertices) Less(i, j int) bool { return v[i].ID < v[j].ID }
