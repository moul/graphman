package graphman

import "fmt"

func ExamplePath() {
	va := &Vertex{ID: "A"}
	vb := &Vertex{ID: "B"}
	vc := &Vertex{ID: "C"}
	vd := &Vertex{ID: "D"}

	path := Path{
		&Edge{Src: va, Dst: vb},
		&Edge{Src: vb, Dst: vc},
		&Edge{Src: vc, Dst: vd},
	}
	fmt.Println(path)

	// Output: (A,B,C,D)
}
