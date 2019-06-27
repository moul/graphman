package graphman

import "fmt"

func ExampleNewVertex() {
	va := NewVertex("aaa")
	fmt.Println(va)
	fmt.Println(va.ID())

	fmt.Println()

	vb := NewVertex("bbb")
	vb.AddAttr(
		NewAttr("ccc", "ddd"),
		NewAttr(42, []string{"eee", "fff"}),
		NewAttr("ggg", "hhh"),
	)
	vb.DelAttr("ggg", "iii")
	fmt.Println(vb)
	fmt.Println(vb.ID())

	// Output:
	// aaa
	// aaa
	//
	// bbb(ccc:ddd,42:[eee fff])
	// bbb
}
