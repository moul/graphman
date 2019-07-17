package main

import (
	"fmt"
	"log"

	"moul.io/graphman"
	"moul.io/graphman/viz"
)

func main() {
	graph := graphman.New()
	graph.AddEdge("1", "2")
	graph.AddEdge("2", "3")
	graph.AddEdge("2", "4")
	graph.AddEdge("3", "5")
	graph.AddEdge("3", "6")
	graph.AddEdge("4", "5")
	graph.AddEdge("4", "9")
	graph.AddEdge("5", "7")
	graph.AddEdge("6", "8")
	graph.AddEdge("6", "12")
	graph.AddEdge("7", "8")
	graph.AddEdge("7", "10")
	graph.AddEdge("8", "11")
	graph.AddEdge("9", "10")
	graph.AddEdge("9", "15")
	graph.AddEdge("10", "13")
	graph.AddEdge("11", "12")
	graph.AddEdge("11", "14")
	graph.AddEdge("12", "17")
	graph.AddEdge("13", "14")
	graph.AddEdge("13", "15")
	graph.AddEdge("14", "16")
	graph.AddEdge("15", "18")
	graph.AddEdge("16", "17")
	graph.AddEdge("17", "19")
	graph.AddEdge("18", "19")
	graph.AddEdge("19", "20")

	log.Println("all paths from 1 to 20:")
	for _, path := range graph.FindAllPaths("1", "20") {
		log.Println("-", path)
	}
	log.Println("shortest path from 1 to 20:")
	path, dist := graph.FindShortestPath("1", "20")
	log.Println("-", path, "dist:", dist)

	for _, edge := range path {
		edge.Dst().SetColor("red")
		edge.SetColor("red")
	}
	path.FirstVertex().SetColor("blue")
	path.LastVertex().SetColor("blue")

	s, err := viz.ToGraphviz(graph, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
