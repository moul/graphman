package graphman

import "fmt"

func Example_NewStringVertex() {
	a := NewStringVertex("aaa")
	fmt.Println(a)
	fmt.Println(a.ID())

	fmt.Println()

	b := NewStringVertex("bbb")
	b.SetAttr("ccc", "ddd")
	b.SetAttr(42, []string{"eee", "fff"})
	b.SetAttr("ggg", "hhh")
	b.DelAttr("ggg")
	b.DelAttr("iii")
	fmt.Println(b)
	fmt.Println(b.ID())

	// Output:
	// aaa
	// aaa

	// bbb(ccc:ddd,42:[eee fff])
	// bbb
}
