package graphman

import (
	"encoding/json"
	"fmt"
)

type Node interface {
	fmt.Stringer

	ID() string
	Data() interface{}
}

type node struct {
	id   string
	data interface{}
}

func (n *node) ID() string { return n.id }

func (n *node) Data() interface{} { return n.data }

func (n *node) String() string {
	out, _ := json.Marshal(n.data)
	return fmt.Sprintf("%s(%s)", n.id, string(out))
}

func NewNode(id string, data interface{}) Node {
	return &node{id: id, data: data}
}
