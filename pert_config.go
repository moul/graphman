package graphman

import (
	"fmt"
	"log"
	"strings"
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
				graph.pertGetDependencyEnd(dependency),
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
				graph.pertGetDependencyEnd(dependency),
				pertPostID(action.ID),
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
					graph.pertGetDependencyEnd(dependency),
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
		if !pertIsUntitledState(vertex) && vertex.Attrs.GetTitle() == "" {
			vertex.Attrs.SetTitle(vertex.id)
		}
	}
	if !config.Opts.NoSimplify { // simplify the graph
		for {
			pertRemoveDummySteps(graph)
			pertMergeDummyActionGroups(graph)
			if removed := graph.gc(); removed == 0 {
				break
			}
		}
	}

	return graph
}

func pertRemoveDummySteps(graph *Graph) {
	// remove dummy states with only one dummy successor
	for _, vertex := range graph.Vertices() {
		if vertex.deleted || !pertIsUntitledState(vertex) || vertex.OutDegree() != 1 {
			continue
		}
		successor := vertex.SuccessorEdges()[0]
		if pertIsZeroTimeActivity(successor) {
			for _, predecessor := range vertex.PredecessorEdges() {
				predecessor.dst = successor.dst
			}
			graph.RemoveEdge(vertex.id, successor.dst.id)
			vertex.deleted = true
		}
	}
}

func pertMergeDummyActionGroups(graph *Graph) {
	// merge dummy action groups
	for _, vertex := range graph.Vertices() {
		if vertex.deleted || vertex.OutDegree() < 2 {
			continue
		}
		for _, combination := range vertex.SuccessorEdges().AllCombinations().LongestToShortest() {
			if len(combination) < 2 {
				continue
			}
			onlyActiveDummies := true
			for _, edge := range combination {
				if edge.deleted || edge.src.deleted || edge.dst.deleted || !pertIsZeroTimeActivity(edge) {
					onlyActiveDummies = false
					break
				}
			}
			if !onlyActiveDummies {
				continue
			}
			predecessors := combination[0].dst.PredecessorVertices()
			same := true
			for i := 1; i < len(combination); i++ {
				if !predecessors.Equals(combination[i].dst.PredecessorVertices()) {
					same = false
					break
				}
			}
			if !same {
				continue
			}
			successors := Vertices{}
			for _, edge := range combination {
				successors = append(successors, edge.dst)
			}
			predecessors = predecessors.Unique()
			successors = successors.Unique()

			ids := []string{}
			titles := []string{}
			for _, successor := range successors {
				ids = append(ids, successor.id)
				if title := successor.Attrs.GetTitle(); title != "" {
					titles = append(titles, title)
				}
			}
			metaID := strings.Join(ids, ",")
			attrs := Attrs{}
			if len(titles) > 0 {
				attrs.SetTitle(strings.Join(titles, " + "))
			} else {
				attrs.SetPertUntitledState()
			}
			metaVertex := graph.AddVertex(metaID, attrs)
			for _, predecessor := range predecessors {
				depID := graph.pertGetDependencyEnd(predecessor.id)
				graph.AddEdge(depID, metaID, Attrs{}.SetPertZeroTimeActivity())
			}
			for _, successor := range successors {
				for _, successorSuccessor := range successor.successors {
					successorSuccessor.src = metaVertex
				}
				for _, predecessor := range predecessors {
					graph.RemoveEdge(predecessor.id, successor.id)
				}
				successor.deleted = true
			}
			break
		}
	}
}

func pertPreID(id string) string  { return fmt.Sprintf("pre_%s", id) }
func pertPostID(id string) string { return fmt.Sprintf("post_%s", id) }

func (g *Graph) pertGetDependencyEnd(dependency string) string {
	// if dependency is a vertex, the end is the vertex itself
	if vertex := g.GetVertex(dependency); vertex != nil {
		return vertex.id
	}
	// else, we need to take the post_{edge}
	return pertPostID(dependency)
}

func pertIsZeroTimeActivity(edge *Edge) bool {
	pert := edge.GetPert()
	return pert != nil && pert.IsZeroTimeActivity
}

func pertIsUntitledState(vertex *Vertex) bool {
	pert := vertex.GetPert()
	return pert != nil && pert.IsUntitledState
}
