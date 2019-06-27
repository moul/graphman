package graphman

import "fmt"

func ExampleNewVertex_simple() {
	v := &Vertex{ID: "A"}
	fmt.Println(v)
	fmt.Println(v.ID)
	// Output:
	// A
	// A
}

func ExampleNewVertex_withAttrs() {
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
