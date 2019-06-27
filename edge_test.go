package graphman

import "fmt"

func ExampleNewEdge() {
	va := NewVertex("aaa")
	vb := NewVertex("bbb")
	vc := NewVertex("ccc")
	vc.AddAttr(NewAttr("ddd", "eee"))

	eab := NewEdge(va, vb)
	eac := NewEdge(va, vc)

	fmt.Println(eab)
	fmt.Println(eac)

	// Output:
	// (aaa,bbb)
	// (aaa,ccc)
}
