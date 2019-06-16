package graphman

import (
	"encoding/json"
	"fmt"
)

type Node interface {
	String() string
}

type node struct {
	id   string
	data interface{}
}

func (n *node) String() string {
	out, _ := json.Marshal(n.data)
	return fmt.Sprintf("%s(%s)", n.id, string(out))
}

func NewNode(id string, data interface{}) Node {
	return &node{id: id, data: data}
}
