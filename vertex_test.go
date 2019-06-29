package graphman

import "fmt"

func ExampleVertex_simple() {
	v := &Vertex{ID: "A"}
	fmt.Println(v)
	fmt.Println(v.ID)
	// Output:
	// A
	// A
}

func ExampleVertex_withAttrs() {
	v := &Vertex{
		ID: "A",
		Attrs: Attrs{
			"ccc": "ddd",
			"eee": []string{"fff", "ggg"},
			"hhh": 42,
		},
	}
	fmt.Println(v)
	fmt.Println(v.ID)

	// Output:
	// A[[ccc:ddd,eee:[fff ggg],hhh:42]]
	// A
}

func ExampleVertices() {
	vertices := &Vertices{
		&Vertex{ID: "A"},
		&Vertex{ID: "B"},
		&Vertex{ID: "C"},
		&Vertex{ID: "D"},
	}
	fmt.Println(vertices)
	// Output: {A,B,C,D}
}
