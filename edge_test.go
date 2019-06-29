package graphman

import "fmt"

func ExampleEdge() {
	graph := New()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C", Attrs{"D": "E"})
	fmt.Println(graph.AddEdge("A", "B"))
	fmt.Println(graph.AddEdge("A", "C"))
	fmt.Println(graph.Edges())

	// Output:
	// (A,B)
	// (A,C)
	// {(A,B),(A,C)}
}
