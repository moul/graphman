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

type PertState struct {
	ID        string   `yaml:"id"`
	Title     string   `yaml:"title"`
	DependsOn []string `yaml:"depends_on"`
}

type PertConfig struct {
	Actions []PertAction `yaml:"actions"`
	States  []PertState  `yaml:"states"`
	Opts    struct {
		NoSimplify bool `yaml:"simplify"`
	} `yaml:"opts"`
}

const (
	pertStartVertex  = "Start"
	pertFinishVertex = "Finish"
)

func FromPertConfig(config PertConfig) *Graph {
	graph := New()

	graph.AddVertex(pertStartVertex)
	graph.AddVertex(pertFinishVertex)

	for _, state := range config.States {
		attrs := Attrs{}
		if state.Title != "" {
			attrs.SetTitle(state.Title)
		}
		graph.AddVertex(state.ID, attrs)
		for _, dependency := range state.DependsOn {
			graph.AddEdge(
				pertPostID(dependency),
				state.ID,
				Attrs{}.SetPertZeroTimeActivity(),
			)
		}
	}

	for _, action := range config.Actions {
		attrs := Attrs{}
		if action.Title != "" {
			attrs.SetTitle(action.Title)
		}

		// pert estimates
		switch len(action.Estimate) {
		case 0:
			break
		case 1:
			attrs.SetPertEstimates(action.Estimate[0], action.Estimate[0], action.Estimate[0])
		case 3:
			attrs.SetPertEstimates(action.Estimate[0], action.Estimate[1], action.Estimate[2])
		default:
			log.Printf("invalid pert estimate: %v", action.Estimate)
		}

		// relationships
		switch len(action.DependsOn) {
		case 0: // no dependency, linking with Start
			graph.AddEdge(
				pertStartVertex,
				pertPostID(action.ID),
				attrs,
			)
		case 1: // only one dependency
			dependency := action.DependsOn[0]
			graph.AddEdge(
				pertPostID(dependency),
				pertPreID(action.ID),
				attrs,
			)
		default:
			graph.AddEdge(
				pertPreID(action.ID),
				pertPostID(action.ID),
				attrs,
			)
			for _, dependency := range action.DependsOn {
				graph.AddEdge(
					pertPostID(dependency),
					pertPreID(action.ID),
					Attrs{}.SetPertZeroTimeActivity(),
				)
			}
		}
	}

	// link ending vertices with finish
	for _, vertex := range graph.SinkVertices() {
		if vertex.id == pertFinishVertex {
			continue
		}
		graph.AddEdge(
			vertex.id,
			pertFinishVertex,
			Attrs{}.SetPertZeroTimeActivity(),
		)
	}

	// nice names
	for _, vertex := range graph.Vertices() {
		if vertex.id == pertStartVertex || vertex.id == pertFinishVertex {
			continue
		}
		if vertex.Attrs.GetTitle() == "" {
			vertex.Attrs.SetPertUntitledState()
		}
	}

	for _, vertex := range graph.Vertices() {
		if vertex.Attrs.GetTitle() == "" {
			vertex.Attrs.SetTitle(vertex.id)
		}
	}
	if !config.Opts.NoSimplify { // simplify the graph
		verticesToDelete := map[string]bool{} // creating a map so we can iterate over vertices while deleting some entries
		for _, vertex := range graph.Vertices() {
			if pertIsUntitledState(vertex) || true {
				if vertex.OutDegree() == 1 { // remove dummy states with only one dummy successor
					successor := vertex.SuccessorEdges()[0]
					if pertIsZeroTimeActivity(successor) {
						for _, predecessor := range vertex.PredecessorEdges() {
							predecessor.dst = successor.dst
						}
						graph.RemoveEdge(vertex.id, successor.dst.id)
						verticesToDelete[vertex.id] = true
					}
				}
			}
		}
		for id := range verticesToDelete {
			graph.RemoveVertex(id)
		}
	}

	return graph
}

func pertPreID(id string) string  { return fmt.Sprintf("pre_%s", id) }
func pertPostID(id string) string { return fmt.Sprintf("post_%s", id) }

func pertIsZeroTimeActivity(edge *Edge) bool {
	pert := edge.GetPert()
	return pert != nil && pert.IsZeroTimeActivity
}

func pertIsUntitledState(vertex *Vertex) bool {
	pert := vertex.GetPert()
	return pert != nil && pert.IsUntitledState
}
