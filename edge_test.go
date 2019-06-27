package graphman

import "fmt"

func ExampleEdge() {
	va := &Vertex{ID: "aaa"}
	vb := &Vertex{ID: "bbb"}
	vc := &Vertex{ID: "ccc", Attrs: Attrs{"ddd": "eee"}}

	eab := &Edge{Src: va, Dst: vb}
	eac := &Edge{Src: va, Dst: vc}

	fmt.Println(eab)
	fmt.Println(eac)

	// Output:
	// (aaa,bbb)
	// (aaa,ccc)
}
