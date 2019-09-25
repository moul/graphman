package graphman

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
)

const invalidPlaceholder = "[INVALID]"

type Graph struct {
	vertices Vertices
	edges    Edges
	Attrs
	dijkstra struct {
		missing uint64
	}
}

func New(attrs ...Attrs) *Graph {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	} else {
		a = make(map[string]interface{})
	}
	return &Graph{
		vertices: make(Vertices, 0),
		edges:    make(Edges, 0),
		Attrs:    a,
	}
}

func (g Graph) SinkVertices() Vertices {
	sinks := Vertices{}
	for _, vertex := range g.vertices {
		if vertex.IsSink() {
			sinks = append(sinks, vertex)
		}
	}
	return sinks.filtered()
}

func (g Graph) ConnectedSubgraphs() Graphs {
	graphs := Graphs{}
	visited := map[string]bool{}
	for _, vertex := range g.vertices {
		if visited[vertex.id] {
			continue
		}
		subgraph := g.ConnectedSubgraphFromVertex(vertex)
		for _, v := range subgraph.vertices {
			visited[v.id] = true
		}
		graphs = append(graphs, subgraph)
	}
	return graphs
}

func (g Graph) ConnectedSubgraphFromVertex(start *Vertex) *Graph {
	subgraph := New()
	visitedEdges := map[*Edge]bool{}
	if err := start.WalkAdjacentVertices(func(current, previous *Vertex, depth int) error {
		subgraph.vertices = append(subgraph.vertices, current)
		for _, edge := range current.Edges() {
			if visitedEdges[edge] {
				continue
			}
			visitedEdges[edge] = true
			subgraph.edges = append(subgraph.edges, edge)
		}
		return nil
	}); err != nil {
		log.Printf("error while walking vertices: %v", err)
	}
	return subgraph
}

func (g Graph) SourceVertices() Vertices {
	sources := Vertices{}
	for _, vertex := range g.vertices {
		if !vertex.deleted && vertex.IsSource() {
			sources = append(sources, vertex)
		}
	}
	return sources
}

func (g Graph) FindAllPaths(srcID, dstID string) Paths {
	src := g.GetVertex(srcID)
	if src == nil {
		log.Printf("%q does not exist in the graph", srcID)
		return Paths{}
	}
	dst := g.GetVertex(dstID)
	if dst == nil {
		log.Printf("%q does not exist in the graph", dstID)
		return Paths{}
	}
	paths := g.findAllPathsRec(src, dst, Path{})
	sort.Sort(paths)
	return paths
}

func (g Graph) findAllPathsRec(current, target *Vertex, prefix Path) Paths {
	paths := Paths{}
	for _, edge := range current.successors {
		if prefix.HasVertex(edge.dst.id) {
			continue
		}
		newPath := append(prefix, edge)
		if edge.dst == target {
			paths = append(paths, &newPath)
		} else {
			paths = append(paths, g.findAllPathsRec(edge.dst, target, newPath)...)
		}
	}
	return paths
}

func (g Graph) FindShortestPath(srcID, dstID string) (Path, int64) {
	costFN := func(e *Edge) int64 { return 1 }
	return g.FindShortestPathFN(srcID, dstID, costFN)
}

func (g Graph) FindShortestPathFN(srcID, dstID string, fn EdgeCostFN) (Path, int64) {
	src := g.GetVertex(srcID)
	dst := g.GetVertex(dstID)
	if src == nil || dst == nil {
		return nil, -1
	}

	// reset dijkstra
	g.dijkstra.missing = 0
	for _, v := range g.vertices {
		v.dijkstra.dist = math.MaxInt64
		v.dijkstra.prev = nil
		v.dijkstra.visited = false
		g.dijkstra.missing++
	}
	src.dijkstra.dist = 0

	// run dijkstra
	for g.dijkstra.missing > 0 {
		var u *Vertex
		for _, v := range g.vertices {
			if v.dijkstra.visited {
				continue
			}
			if u == nil || u.dijkstra.dist > v.dijkstra.dist {
				u = v
			}
		}
		for _, e := range u.successors {
			n := e.dst
			if n.dijkstra.visited {
				continue
			}
			dist := u.dijkstra.dist + fn(e)
			if dist < n.dijkstra.dist {
				n.dijkstra.dist = dist
				n.dijkstra.prev = e
			}
		}
		u.dijkstra.visited = true
		g.dijkstra.missing--
	}

	if dst.dijkstra.dist == math.MaxInt64 {
		return nil, -1
	}

	path := Path{}
	for cur := dst; cur.dijkstra.prev != nil; cur = cur.dijkstra.prev.src {
		path = append(Path{cur.dijkstra.prev}, path...)
	}
	return path, dst.dijkstra.dist
}

func (g Graph) Edges() Edges {
	// FIXME: sort.Sort(g.edges)
	return g.edges
}

func (g Graph) Vertices() Vertices {
	sort.Sort(g.vertices)
	return g.vertices.filtered()
}

func (g *Graph) gc() uint {
	removed := uint(0)
	verticesToDelete := Vertices{}
	for _, vertex := range g.vertices {
		if vertex.deleted {
			verticesToDelete = append(verticesToDelete, vertex)
		}
	}
	edgesToDelete := Edges{}
	for _, edge := range g.edges {
		if edge.deleted {
			edgesToDelete = append(edgesToDelete, edge)
		}
	}
	for _, edge := range edgesToDelete {
		removed++
		g.RemoveEdge(edge.src.id, edge.dst.id)
	}
	for _, vertex := range verticesToDelete {
		removed++
		g.RemoveVertex(vertex.id)
	}
	return removed
}

func (g *Graph) AddVertex(id string, attrs ...Attrs) *Vertex {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	} else {
		a = make(Attrs)
	}

	v := g.GetVertex(id)
	if v != nil {
		v.Attrs.Merge(a)
	} else {
		v = newVertex(id, a)
		g.vertices = append(g.vertices, v)
	}

	return v
}

func (g *Graph) RemoveVertex(id string) bool {
	for k, v := range g.vertices {
		if v.id == id {
			v.deleted = true
			g.vertices = append(g.vertices[:k], g.vertices[k+1:]...)
			return true
		}
	}
	return false
}

func (g *Graph) RemoveEdge(src, dst string) bool {
	for k, e := range g.edges {
		if e.src.id == src && e.dst.id == dst {
			e.deleted = true
			g.edges = append(g.edges[:k], g.edges[k+1:]...)
			return true
		}
	}
	return false
}

func (g Graph) GetVertex(id string) *Vertex {
	for _, v := range g.vertices {
		if v.id == id {
			return v
		}
	}
	return nil
}

func (g *Graph) AddEdge(srcID, dstID string, attrs ...Attrs) *Edge {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	} else {
		a = make(Attrs)
	}

	src := g.AddVertex(srcID)
	dst := g.AddVertex(dstID)
	edge := newEdge(src, dst, a)
	src.successors = append(src.successors, edge)
	dst.predecessors = append(dst.predecessors, edge)
	g.edges = append(g.edges, edge)
	return edge
}

func (g Graph) IsolatedVertices() Vertices {
	isolated := Vertices{}
	for _, vertex := range g.vertices {
		if len(vertex.Edges()) == 0 {
			isolated = append(isolated, vertex)
		}
	}
	sort.Sort(isolated)
	return isolated
}

func (g *Graph) String() string {
	if g == nil {
		return invalidPlaceholder
	}
	elems := []string{}
	for _, edge := range g.edges {
		elems = append(elems, edge.String())
	}
	for _, vertex := range g.IsolatedVertices() {
		elems = append(elems, vertex.id)
	}
	return fmt.Sprintf("{%s}", strings.Join(elems, ","))
}

// Human returns a diff-friendly and human-readable representation of the vertices and edges
func (g *Graph) Human() string {
	out := ""
	for _, vertex := range g.Vertices() {
		out += fmt.Sprintf("* %s\n", vertex.ID())
		for _, edge := range vertex.Edges() {
			out += fmt.Sprintf("  * %s\n", edge.String())
		}
	}
	return out
}

//
// Graphs
//

type Graphs []*Graph
