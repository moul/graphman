package graphman

import (
	"fmt"
)

func Example_pert() {
	graph := New()
	graph.AddEdge("1", "2", Attrs{}.SetTitle("a").SetPert(3, 6, 15))
	graph.AddEdge("1", "3", Attrs{}.SetTitle("b").SetPert(2, 5, 14))
	graph.AddEdge("1", "4", Attrs{}.SetTitle("c").SetPert(6, 12, 30))
	graph.AddEdge("2", "5", Attrs{}.SetTitle("d").SetPert(2, 5, 8))
	graph.AddEdge("2", "6", Attrs{}.SetTitle("e").SetPert(5, 11, 17))
	graph.AddEdge("3", "6", Attrs{}.SetTitle("f").SetPert(3, 6, 15))
	graph.AddEdge("4", "7", Attrs{}.SetTitle("g").SetPert(3, 9, 27))
	graph.AddEdge("5", "7", Attrs{}.SetTitle("h").SetPert(1, 4, 7))
	graph.AddEdge("6", "7", Attrs{}.SetTitle("i").SetPert(4, 19, 28))
	for _, e := range graph.Edges() {
		fmt.Println("*", e)
	}
	// Output:
	// * (1,2)[[pert:To=3,Tm=6,Tp=15,Te=7,σe=2,Ve=4,title:a]]
	// * (1,3)[[pert:To=2,Tm=5,Tp=14,Te=6,σe=2,Ve=4,title:b]]
	// * (1,4)[[pert:To=6,Tm=12,Tp=30,Te=14,σe=4,Ve=16,title:c]]
	// * (2,5)[[pert:To=2,Tm=5,Tp=8,Te=5,σe=1,Ve=1,title:d]]
	// * (2,6)[[pert:To=5,Tm=11,Tp=17,Te=11,σe=2,Ve=4,title:e]]
	// * (3,6)[[pert:To=3,Tm=6,Tp=15,Te=7,σe=2,Ve=4,title:f]]
	// * (4,7)[[pert:To=3,Tm=9,Tp=27,Te=11,σe=4,Ve=16,title:g]]
	// * (5,7)[[pert:To=1,Tm=4,Tp=7,Te=4,σe=1,Ve=1,title:h]]
	// * (6,7)[[pert:To=4,Tm=19,Tp=28,Te=18,σe=4,Ve=16,title:i]]
}
