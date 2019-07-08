package graphman

import (
	"fmt"
	"log"
)

type PertAction struct {
	ID        string    `yaml:"id"`
	Title     string    `yaml:"title"`
	Estimate  []float64 `yaml:"estimate"`
	DependsOn []string  `yaml:"depends_on"`
}

type PertConfig struct {
	Actions []PertAction `yaml:"actions"`
}

const (
	pertStartVertex  = "Start"
	pertFinishVertex = "Finish"
)

func FromPertConfig(config PertConfig) *Graph {
	graph := New()

	graph.AddVertex(pertStartVertex)
	graph.AddVertex(pertFinishVertex)

	for _, action := range config.Actions {
		attrs := Attrs{}
		if action.Title != "" {
			attrs.SetTitle(action.Title)
		}

		// pert estimates
		switch len(action.Estimate) {
		case 0:
			break
		case 3:
			attrs.SetPertEstimates(action.Estimate[0], action.Estimate[1], action.Estimate[2])
		default:
			log.Printf("invalid pert estimate: %v", action.Estimate)
		}

		// relationships
		postCurrent := fmt.Sprintf("post_%s", action.ID)
		switch len(action.DependsOn) {
		case 0: // no dependency, linking with Start
			graph.AddEdge(pertStartVertex, postCurrent, attrs)
		case 1: // only one dependency
			dependency := action.DependsOn[0]
			graph.AddEdge(
				fmt.Sprintf("post_%s", dependency),
				postCurrent,
				attrs,
			)
		default:
			graph.AddEdge(
				fmt.Sprintf("pre_%s", action.ID),
				postCurrent,
				attrs,
			)
			for _, dependency := range action.DependsOn {
				graph.AddEdge(
					fmt.Sprintf("post_%s", dependency),
					fmt.Sprintf("pre_%s", action.ID),
					Attrs{}.SetPertZeroTimeActivity(),
				)
			}
		}
	}

	// link ending vertices with finish
	for _, vertex := range graph.SinkVertices() {
		if vertex.ID() == pertFinishVertex {
			continue
		}
		graph.AddEdge(
			vertex.ID(),
			pertFinishVertex,
			Attrs{}.SetPertZeroTimeActivity(),
		)
	}

	// optimize
	// FIXME: TODO

	// nice names
	for _, vertex := range graph.Vertices() {
		if vertex.ID() == pertStartVertex || vertex.ID() == pertFinishVertex {
			continue
		}
		if vertex.Attrs.GetTitle() == "" {
			vertex.Attrs.SetPertUntitledState()
		}
	}

	return graph
}
