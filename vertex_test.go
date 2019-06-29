package graphman

import "fmt"

func ExampleVertex_simple() {
	graph := New()
	v := graph.AddVertex("A")
	fmt.Println(v)
	fmt.Println(v.ID())
	// Output:
	// A
	// A
}

func ExampleVertex_withAttrs() {
	graph := New()
	v := graph.AddVertex("A", Attrs{
		"ccc": "ddd",
		"eee": []string{"fff", "ggg"},
		"hhh": 42,
	})
	fmt.Println(v)
	fmt.Println(v.ID())

	// Output:
	// A[[ccc:ddd,eee:[fff ggg],hhh:42]]
	// A
}

func ExampleVertices() {
	graph := New()
	a := graph.AddVertex("A")
	b := graph.AddVertex("B")
	c := graph.AddVertex("C")
	d := graph.AddVertex("D")
	vertices := &Vertices{a, b, c, d}
	fmt.Println(vertices)
	// Output: {A,B,C,D}
}
