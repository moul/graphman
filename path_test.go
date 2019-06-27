package graphman

import "fmt"

func ExampleNewPath() {
	va := NewVertex("A")
	vb := NewVertex("B")
	vc := NewVertex("C")
	vd := NewVertex("D")
	eab := NewEdge(va, vb)
	ebc := NewEdge(vb, vc)
	ecd := NewEdge(vc, vd)
	path := NewPath(eab)
	path.Append(ebc)
	path.Append(ecd)
	fmt.Println(path)

	// Output: (A,B,C,D)
}
