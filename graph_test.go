package graphman

import "fmt"

func Example_graph() {
	graph := New()
	graph.AddVertex(NewVertex("A", 42))
	graph.AddVertex(NewVertex("B", "hello world"))
	graph.AddVertex(NewVertex("C", []string{"a", "b", "c"}))
	graph.AddVertex(NewVertex("D", nil))
	graph.AddEdge(NewEdge("A", "B"))
	graph.AddEdge(NewEdge("B", "C"))
	graph.AddEdge(NewEdge("C", "D"))
	graph.AddEdge(NewEdge("D", "A"))
	fmt.Println(graph)
	// Output:
	// A(42)
	// B("hello world")
	// C(["a","b","c"])
	// D(null)
	// A->B
	// B->C
	// C->D
	// D->A
}

func Example_connections() {
	// init graph
	graph := New()
	graph.AddVertex(NewVertex("A", nil))
	graph.AddVertex(NewVertex("B", nil))
	graph.AddVertex(NewVertex("C", nil))
	graph.AddVertex(NewVertex("D", nil))
	graph.AddVertex(NewVertex("E", nil))
	graph.AddVertex(NewVertex("F", nil))
	graph.AddEdge(NewEdge("A", "B"))
	graph.AddEdge(NewEdge("B", "C"))
	graph.AddEdge(NewEdge("D", "E"))
	for _, a := range graph.Vertices() {
		fmt.Printf("=== %s ===\n", a.ID())
		fmt.Printf("edges: %v\n", graph.EdgesFor(a.ID()))
		fmt.Printf("direct connections: %v\n", graph.DirectConnectionsFor(a.ID()))
		fmt.Printf("all connections: %v\n", graph.AllConnectionsFor(a.ID()))
		fmt.Printf("all shortest paths: %v\n", graph.AllShortestPaths(a.ID()))

		for _, b := range graph.Vertices() {
			if a.ID() != b.ID() {
				fmt.Printf(
					"%s & %s, are connected: %v, directly: %v, with the shortest path: %v\n",
					a.ID(), b.ID(),
					graph.AreConnected(a.ID(), b.ID()),
					graph.AreDirectlyConnected(a.ID(), b.ID()),
					graph.ShortestPath(a.ID(), b.ID()),
				)
			}
		}
		fmt.Println()
	}

	// Output:
	// === A ===
	// edges: [A->B]
	// direct connections: [B]
	// all connections: [A B C]
	// all shortest paths: map[A:A B:A->B C:A->B->C]
	// A & B, are connected: true, directly: true, with the shortest path: A->B
	// A & C, are connected: true, directly: false, with the shortest path: A->B->C
	// A & D, are connected: false, directly: false, with the shortest path: <nil>
	// A & E, are connected: false, directly: false, with the shortest path: <nil>
	// A & F, are connected: false, directly: false, with the shortest path: <nil>
	//
	// === B ===
	// edges: [A->B B->C]
	// direct connections: [A C]
	// all connections: [A B C]
	// all shortest paths: map[A:B->A B:B C:B->C]
	// B & A, are connected: true, directly: true, with the shortest path: B->A
	// B & C, are connected: true, directly: true, with the shortest path: B->C
	// B & D, are connected: false, directly: false, with the shortest path: <nil>
	// B & E, are connected: false, directly: false, with the shortest path: <nil>
	// B & F, are connected: false, directly: false, with the shortest path: <nil>
	//
	// === C ===
	// edges: [B->C]
	// direct connections: [B]
	// all connections: [A B C]
	// all shortest paths: map[A:C->B->A B:C->B C:C]
	// C & A, are connected: true, directly: false, with the shortest path: C->B->A
	// C & B, are connected: true, directly: true, with the shortest path: C->B
	// C & D, are connected: false, directly: false, with the shortest path: <nil>
	// C & E, are connected: false, directly: false, with the shortest path: <nil>
	// C & F, are connected: false, directly: false, with the shortest path: <nil>
	//
	// === D ===
	// edges: [D->E]
	// direct connections: [E]
	// all connections: [D E]
	// all shortest paths: map[D:D E:D->E]
	// D & A, are connected: false, directly: false, with the shortest path: <nil>
	// D & B, are connected: false, directly: false, with the shortest path: <nil>
	// D & C, are connected: false, directly: false, with the shortest path: <nil>
	// D & E, are connected: true, directly: true, with the shortest path: D->E
	// D & F, are connected: false, directly: false, with the shortest path: <nil>
	//
	// === E ===
	// edges: [D->E]
	// direct connections: [D]
	// all connections: [D E]
	// all shortest paths: map[D:E->D E:E]
	// E & A, are connected: false, directly: false, with the shortest path: <nil>
	// E & B, are connected: false, directly: false, with the shortest path: <nil>
	// E & C, are connected: false, directly: false, with the shortest path: <nil>
	// E & D, are connected: true, directly: true, with the shortest path: E->D
	// E & F, are connected: false, directly: false, with the shortest path: <nil>
	//
	// === F ===
	// edges: []
	// direct connections: []
	// all connections: [F]
	// all shortest paths: map[F:F]
	// F & A, are connected: false, directly: false, with the shortest path: <nil>
	// F & B, are connected: false, directly: false, with the shortest path: <nil>
	// F & C, are connected: false, directly: false, with the shortest path: <nil>
	// F & D, are connected: false, directly: false, with the shortest path: <nil>
	// F & E, are connected: false, directly: false, with the shortest path: <nil>
}

func Example_bigGraph() {
	// init graph
	graph := New()
	amount := 100
	startVertex := 42

	for i := 0; i <= amount; i++ {
		graph.AddVertex(NewVertex(fmt.Sprintf("%d", i), nil))
	}
	for i := 0; i <= amount-1; i++ {
		graph.AddEdge(NewEdge(
			fmt.Sprintf("%d", i),
			fmt.Sprintf("%d", i+1),
		))
	}

	fmt.Println("edges", graph.EdgesFor(fmt.Sprintf("%d", startVertex)))
	fmt.Println("direct connections", graph.DirectConnectionsFor(fmt.Sprintf("%d", startVertex)))
	fmt.Println("all connections", graph.AllConnectionsFor(fmt.Sprintf("%d", startVertex)))

	fmt.Println(graph.ShortestPath(fmt.Sprintf("%d", startVertex), fmt.Sprintf("%d", amount)))
	fmt.Println(graph.ShortestPath(fmt.Sprintf("%d", startVertex), "0"))

	// Output:
	// edges [41->42 42->43]
	// direct connections [41 43]
	// all connections [0 1 10 100 11 12 13 14 15 16 17 18 19 2 20 21 22 23 24 25 26 27 28 29 3 30 31 32 33 34 35 36 37 38 39 4 40 41 42 43 44 45 46 47 48 49 5 50 51 52 53 54 55 56 57 58 59 6 60 61 62 63 64 65 66 67 68 69 7 70 71 72 73 74 75 76 77 78 79 8 80 81 82 83 84 85 86 87 88 89 9 90 91 92 93 94 95 96 97 98 99]
	// 42->43->44->45->46->47->48->49->50->51->52->53->54->55->56->57->58->59->60->61->62->63->64->65->66->67->68->69->70->71->72->73->74->75->76->77->78->79->80->81->82->83->84->85->86->87->88->89->90->91->92->93->94->95->96->97->98->99->100
	// 42->41->40->39->38->37->36->35->34->33->32->31->30->29->28->27->26->25->24->23->22->21->20->19->18->17->16->15->14->13->12->11->10->9->8->7->6->5->4->3->2->1->0
}
