package graphman

import (
	"fmt"
)

func ExamplePertResult_withValue() {
	graph := New()
	graph.AddEdge("1", "2", Attrs{}.SetPertEstimates(3, 6, 15))
	graph.AddEdge("1", "3", Attrs{}.SetPertEstimates(2, 5, 14))
	graph.AddEdge("1", "4", Attrs{}.SetPertEstimates(6, 12, 30))
	graph.AddEdge("2", "5", Attrs{}.SetPertEstimates(2, 5, 8))
	graph.AddEdge("2", "6", Attrs{}.SetPertEstimates(5, 11, 17))
	graph.AddEdge("3", "6", Attrs{}.SetPertEstimates(3, 6, 15))
	graph.AddEdge("4", "7", Attrs{}.SetPertEstimates(3, 9, 27))
	graph.AddEdge("5", "7", Attrs{}.SetPertEstimates(1, 4, 7))
	graph.AddEdge("6", "7", Attrs{}.SetPertEstimates(4, 19, 28))

	fmt.Println("graph before computing:")
	for _, e := range graph.Edges() {
		fmt.Println("*", e)
	}
	fmt.Println()

	result := ComputePert(graph)
	fmt.Println("graph after computing:")
	for _, e := range graph.Edges() {
		fmt.Println("*", e)
	}
	fmt.Println("result:", result)

	// Output:
	// graph before computing:
	// * (1,2)[[pert:To=3,Tm=6,Tp=15,Te=7,σe=2,Ve=4]]
	// * (1,3)[[pert:To=2,Tm=5,Tp=14,Te=6,σe=2,Ve=4]]
	// * (1,4)[[pert:To=6,Tm=12,Tp=30,Te=14,σe=4,Ve=16]]
	// * (2,5)[[pert:To=2,Tm=5,Tp=8,Te=5,σe=1,Ve=1]]
	// * (2,6)[[pert:To=5,Tm=11,Tp=17,Te=11,σe=2,Ve=4]]
	// * (3,6)[[pert:To=3,Tm=6,Tp=15,Te=7,σe=2,Ve=4]]
	// * (4,7)[[pert:To=3,Tm=9,Tp=27,Te=11,σe=4,Ve=16]]
	// * (5,7)[[pert:To=1,Tm=4,Tp=7,Te=4,σe=1,Ve=1]]
	// * (6,7)[[pert:To=4,Tm=19,Tp=28,Te=18,σe=4,Ve=16]]
	//
	// graph after computing:
	// * (1,2)[[pert:To=3,Tm=6,Tp=15,Te=7,σe=2,Ve=4]]
	// * (1,3)[[pert:To=2,Tm=5,Tp=14,Te=6,σe=2,Ve=4]]
	// * (1,4)[[pert:To=6,Tm=12,Tp=30,Te=14,σe=4,Ve=16]]
	// * (2,5)[[pert:To=2,Tm=5,Tp=8,Te=5,σe=1,Ve=1]]
	// * (2,6)[[pert:To=5,Tm=11,Tp=17,Te=11,σe=2,Ve=4]]
	// * (3,6)[[pert:To=3,Tm=6,Tp=15,Te=7,σe=2,Ve=4]]
	// * (4,7)[[pert:To=3,Tm=9,Tp=27,Te=11,σe=4,Ve=16]]
	// * (5,7)[[pert:To=1,Tm=4,Tp=7,Te=4,σe=1,Ve=1]]
	// * (6,7)[[pert:To=4,Tm=19,Tp=28,Te=18,σe=4,Ve=16]]
	// result: Tσe=8.12,TVe=66
}

func ExamplePertResult_withoutValue() {
	graph := New()
	graph.AddEdge("1", "2")
	graph.AddEdge("1", "3")
	graph.AddEdge("1", "4")
	graph.AddEdge("2", "5")
	graph.AddEdge("2", "6")
	graph.AddEdge("3", "6")
	graph.AddEdge("4", "7")
	graph.AddEdge("5", "7")
	graph.AddEdge("6", "7")

	fmt.Println("graph before computing:")
	for _, e := range graph.Edges() {
		fmt.Println("*", e)
	}
	fmt.Println()

	result := ComputePert(graph)
	fmt.Println("graph after computing:")
	for _, e := range graph.Edges() {
		fmt.Println("*", e)
	}
	fmt.Println("result:", result)

	// Output:
	// graph before computing:
	// * (1,2)
	// * (1,3)
	// * (1,4)
	// * (2,5)
	// * (2,6)
	// * (3,6)
	// * (4,7)
	// * (5,7)
	// * (6,7)
	//
	// graph after computing:
	// * (1,2)[[pert:To=1,Tm=1,Tp=1,Te=1,σe=0,Ve=0]]
	// * (1,3)[[pert:To=1,Tm=1,Tp=1,Te=1,σe=0,Ve=0]]
	// * (1,4)[[pert:To=1,Tm=1,Tp=1,Te=1,σe=0,Ve=0]]
	// * (2,5)[[pert:To=1,Tm=1,Tp=1,Te=1,σe=0,Ve=0]]
	// * (2,6)[[pert:To=1,Tm=1,Tp=1,Te=1,σe=0,Ve=0]]
	// * (3,6)[[pert:To=1,Tm=1,Tp=1,Te=1,σe=0,Ve=0]]
	// * (4,7)[[pert:To=1,Tm=1,Tp=1,Te=1,σe=0,Ve=0]]
	// * (5,7)[[pert:To=1,Tm=1,Tp=1,Te=1,σe=0,Ve=0]]
	// * (6,7)[[pert:To=1,Tm=1,Tp=1,Te=1,σe=0,Ve=0]]
	// result: Tσe=0,TVe=0
}
