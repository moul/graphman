package graphman

import "fmt"

func ExampleGraph_simple() {
	graph := New()
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")
	graph.AddVertex("E")
	graph.AddVertex("F")
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("E", "F")
	fmt.Println(graph)
	// Output: {(A,B),(B,C),(E,F),D}
}

func ExampleGraph_components() {
	graph := New()
	// component 1: loop
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "A")

	// component 2: standard
	graph.AddEdge("L", "M")
	graph.AddEdge("M", "N")
	graph.AddEdge("M", "O")
	graph.AddEdge("M", "P")
	graph.AddEdge("N", "Q")
	graph.AddEdge("O", "Q")
	graph.AddEdge("P", "Q")
	graph.AddEdge("Q", "R")

	// component 3: self loop
	graph.AddEdge("Z", "Z")

	// component 4: reverse (sorted string)
	graph.AddEdge("Y", "X")
	graph.AddEdge("X", "W")
	graph.AddEdge("W", "V")
	graph.AddEdge("V", "U")

	for _, couple := range [][2]string{
		{"A", "D"}, // using the loop from component 1
		{"B", "A"}, // same
		{"A", "L"}, // two components
		{"L", "R"}, // through all the component 2
		{"Z", "Z"}, // self loop
		{"Y", "U"}, // reverse sorted string
	} {
		fmt.Printf("couple %s-%s:\n", couple[0], couple[1])
		for _, path := range graph.FindAllPaths(couple[0], couple[1]) {
			fmt.Println("-", path)
		}
		path, dist := graph.FindShortestPath(couple[0], couple[1])
		fmt.Println("shortest:", path, dist)
		fmt.Println()
	}

	// Output:
	// couple A-D:
	// - (A,B,C,D)
	// shortest: (A,B,C,D) 3
	//
	// couple B-A:
	// - (B,C,D,A)
	// shortest: (B,C,D,A) 3
	//
	// couple A-L:
	// shortest: [INVALID] -1
	//
	// couple L-R:
	// - (L,M,N,Q,R)
	// - (L,M,O,Q,R)
	// - (L,M,P,Q,R)
	// shortest: (L,M,N,Q,R) 4
	//
	// couple Z-Z:
	// - (Z,Z)
	// shortest: () 0
	//
	// couple Y-U:
	// - (Y,X,W,V,U)
	// shortest: (Y,X,W,V,U) 4
}

func ExampleGraphFindAllPaths() {
	graph := New()
	graph.AddEdge("G", "H")
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "B")
	graph.AddEdge("F", "H")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "E")
	graph.AddEdge("B", "G")
	graph.AddEdge("C", "F")
	graph.AddEdge("D", "H")
	graph.AddEdge("A", "D")
	graph.AddEdge("E", "G")
	graph.AddEdge("E", "H")
	graph.AddEdge("F", "G")

	for _, path := range graph.FindAllPaths("A", "H") {
		fmt.Println(path)
	}

	fmt.Println()
	path, dist := graph.FindShortestPath("A", "H")
	fmt.Printf("shortest (distance=%d): %s\n", dist, path)

	// Output:
	// (A,B,E,G,H)
	// (A,B,E,G,H)
	// (A,B,E,H)
	// (A,B,E,H)
	// (A,B,G,H)
	// (A,B,G,H)
	// (A,C,F,G,H)
	// (A,C,F,H)
	// (A,D,H)
	//
	// shortest (distance=2): (A,D,H)
}

func ExampleGraph_complex() {
	graph := New()
	graph.AddVertex("A", Attrs{"hello": "world"}) // create a vertex with attributes
	graph.AddVertex("B")                          // create an empty vertex
	graph.AddVertex("temp")                       // create a vertex that we will delete later
	graph.AddVertex("isolated")                   // create a vertex that won't have any edge
	graph.AddEdge("A", "B")                       // connect existing vertices A and B
	graph.AddEdge("B", "C")                       // connect existing vertex A with newly created vertex C
	graph.RemoveVertex("temp")
	fmt.Println(graph)
	// Output: {(A,B),(B,C),isolated}
}

func ExampleGraph_big() {
	graph := New()
	amount := 100

	for i := 0; i <= amount; i++ {
		graph.AddVertex(fmt.Sprintf("%d", i))
	}

	for i := 0; i <= amount-1; i++ {
		graph.AddEdge(
			fmt.Sprintf("%d", i),
			fmt.Sprintf("%d", i+1),
		)
	}

	fmt.Println(graph)
	// Output: {(0,1),(1,2),(2,3),(3,4),(4,5),(5,6),(6,7),(7,8),(8,9),(9,10),(10,11),(11,12),(12,13),(13,14),(14,15),(15,16),(16,17),(17,18),(18,19),(19,20),(20,21),(21,22),(22,23),(23,24),(24,25),(25,26),(26,27),(27,28),(28,29),(29,30),(30,31),(31,32),(32,33),(33,34),(34,35),(35,36),(36,37),(37,38),(38,39),(39,40),(40,41),(41,42),(42,43),(43,44),(44,45),(45,46),(46,47),(47,48),(48,49),(49,50),(50,51),(51,52),(52,53),(53,54),(54,55),(55,56),(56,57),(57,58),(58,59),(59,60),(60,61),(61,62),(62,63),(63,64),(64,65),(65,66),(66,67),(67,68),(68,69),(69,70),(70,71),(71,72),(72,73),(73,74),(74,75),(75,76),(76,77),(77,78),(78,79),(79,80),(80,81),(81,82),(82,83),(83,84),(84,85),(85,86),(86,87),(87,88),(88,89),(89,90),(90,91),(91,92),(92,93),(93,94),(94,95),(95,96),(96,97),(97,98),(98,99),(99,100)}
}
