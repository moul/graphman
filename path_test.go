package graphman

import "fmt"

func ExamplePath() {
	graph := New()

	eab := graph.AddEdge("A", "B")
	ebc := graph.AddEdge("B", "C")
	ecd := graph.AddEdge("C", "D", Attrs{"hello": "world"})

	path := Path{eab, ebc, ecd}
	fmt.Println(path)

	// Output: (A,B,C,D)
}
