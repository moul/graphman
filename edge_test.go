package graphman

import "fmt"

func ExampleEdge() {
	va := &Vertex{ID: "A"}
	vb := &Vertex{ID: "B"}
	vc := &Vertex{ID: "C", Attrs: Attrs{"D": "E"}}

	eab := &Edge{Src: va, Dst: vb}
	eac := &Edge{Src: va, Dst: vc}

	edges := &Edges{eab, eac}

	fmt.Println(eab)
	fmt.Println(eac)
	fmt.Println(edges)

	// Output:
	// (A,B)
	// (A,C)
	// {(A,B),(A,C)}
}
