package graphman

import "fmt"

func ExampleVertex_simple() {
	graph := New()
	v := graph.AddVertex("A")
	fmt.Println(v)
	fmt.Println(v.ID())
	// Output:
	// A
	// A
}

func ExampleVertex_withAttrs() {
	graph := New()
	v := graph.AddVertex("A", Attrs{
		"ccc": "ddd",
		"eee": []string{"fff", "ggg"},
		"hhh": 42,
	})
	fmt.Println(v)
	fmt.Println(v.ID())

	// Output:
	// A[[ccc:ddd,eee:[fff ggg],hhh:42]]
	// A
}

func ExampleVertices() {
	graph := New()
	a := graph.AddVertex("A")
	b := graph.AddVertex("B")
	c := graph.AddVertex("C")
	d := graph.AddVertex("D")
	vertices := &Vertices{a, b, c, d}
	fmt.Println(vertices)
	// Output: {A,B,C,D}
}

func ExampleVertex_helpers() {
	graph := New()
	graph.AddVertex("bob")
	graph.AddVertex("eve")
	graph.AddVertex("joy")
	graph.AddVertex("sam")
	graph.AddVertex("han")      // solo
	graph.AddEdge("bob", "eve") // bob is the predecessor of eve
	graph.AddEdge("bob", "bob") // bob is linked to itself
	graph.AddEdge("eve", "joy") // eve is the predecessor of joy
	graph.AddEdge("joy", "eve") // and also its successor
	graph.AddEdge("sam", "bob") // sam is the predecessor of bob

	for _, vertex := range graph.Vertices() {
		fmt.Println(vertex.ID())
		fmt.Println("  isolated:             ", vertex.IsIsolated())
		fmt.Println("  neighbors:            ", vertex.Neighbors())
		fmt.Println("  predecessor vertices: ", vertex.PredecessorVertices())
		fmt.Println("  predecessor edges:    ", vertex.PredecessorEdges())
		fmt.Println("  successor vertices:   ", vertex.SuccessorVertices())
		fmt.Println("  successor edges:      ", vertex.SuccessorEdges())
		fmt.Println("  edges:                ", vertex.Edges())
	}

	// Output:
	// bob
	//   isolated:              false
	//   neighbors:             {bob,eve,sam}
	//   predecessor vertices:  {bob,sam}
	//   predecessor edges:     {(bob,bob),(sam,bob)}
	//   successor vertices:    {bob,eve}
	//   successor edges:       {(bob,eve),(bob,bob)}
	//   edges:                 {(bob,bob),(sam,bob),(bob,eve),(bob,bob)}
	// eve
	//   isolated:              false
	//   neighbors:             {bob,joy}
	//   predecessor vertices:  {bob,joy}
	//   predecessor edges:     {(bob,eve),(joy,eve)}
	//   successor vertices:    {joy}
	//   successor edges:       {(eve,joy)}
	//   edges:                 {(bob,eve),(joy,eve),(eve,joy)}
	// han
	//   isolated:              true
	//   neighbors:             {}
	//   predecessor vertices:  {}
	//   predecessor edges:     {}
	//   successor vertices:    {}
	//   successor edges:       {}
	//   edges:                 {}
	// joy
	//   isolated:              false
	//   neighbors:             {eve}
	//   predecessor vertices:  {eve}
	//   predecessor edges:     {(eve,joy)}
	//   successor vertices:    {eve}
	//   successor edges:       {(joy,eve)}
	//   edges:                 {(eve,joy),(joy,eve)}
	// sam
	//   isolated:              false
	//   neighbors:             {bob}
	//   predecessor vertices:  {}
	//   predecessor edges:     {}
	//   successor vertices:    {bob}
	//   successor edges:       {(sam,bob)}
	//   edges:                 {(sam,bob)}
}
