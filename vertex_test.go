package graphman

import "fmt"

func ExampleNewVertex() {
	a := NewVertex("aaa")
	fmt.Println(a)
	fmt.Println(a.ID())

	fmt.Println()

	b := NewVertex("bbb")
	b.SetAttr(NewAttr("ccc", "ddd"))
	b.SetAttr(NewAttr(42, []string{"eee", "fff"}))
	b.SetAttr(NewAttr("ggg", "hhh"))
	b.DelAttr("ggg")
	b.DelAttr("iii")
	fmt.Println(b)
	fmt.Println(b.ID())

	// Output:
	// aaa
	// aaa
	//
	// bbb(ccc:ddd,42:[eee fff])
	// bbb
}
