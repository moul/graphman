package graphman

import (
	"encoding/json"
	"fmt"
)

type Vertex interface {
	fmt.Stringer

	ID() string
	Data() interface{}
}

type vertex struct {
	id   string
	data interface{}
}

func (n *vertex) ID() string { return n.id }

func (n *vertex) Data() interface{} { return n.data }

func (n *vertex) String() string {
	out, _ := json.Marshal(n.data)
	return fmt.Sprintf("%s(%s)", n.id, string(out))
}

func NewVertex(id string, data interface{}) Vertex {
	return &vertex{id: id, data: data}
}
