package graphman

import "fmt"

func Example() {
	graph := New()
	graph.AddNode(NewNode("A", 42))
	graph.AddNode(NewNode("B", "hello world"))
	graph.AddNode(NewNode("C", []string{"a", "b", "c"}))
	graph.AddNode(NewNode("D", nil))
	graph.AddEdge(NewEdge("A", "B"))
	graph.AddEdge(NewEdge("B", "C"))
	graph.AddEdge(NewEdge("C", "D"))
	graph.AddEdge(NewEdge("D", "A"))
	fmt.Println(graph)
	// Output:
	// A(42)
	// B("hello world")
	// C(["a","b","c"])
	// D(null)
	// A->B
	// B->C
	// C->D
	// D->A
}
