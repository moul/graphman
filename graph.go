package graphman

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strings"
)

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
	}
	return &Graph{
		vertices: make(Vertices, 0),
		edges:    make(Edges, 0),
		Attrs:    a,
	}
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

var i = 200

func (g Graph) findAllPathsRec(current, target *Vertex, prefix Path) Paths {
	i--
	if i < 0 {
		return Paths{}
	}
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
	//src.dijkstra.visited = true
	//g.dijkstra.missing--

	// run dijkstra
	for g.dijkstra.missing > 0 {
		//fmt.Println("missing", g.dijkstra.missing)
		var u *Vertex
		for _, v := range g.vertices {
			//fmt.Println("  v", v)
			if v.dijkstra.visited {
				continue
			}
			if u == nil || u.dijkstra.dist > v.dijkstra.dist {
				u = v
			}
		}
		//fmt.Println("  u (small)", u, u.dijkstra.dist, "prev", u.dijkstra.prev)
		for _, e := range u.successors {
			n := e.dst
			//fmt.Println("    n", n)
			if n.dijkstra.visited {
				continue
			}
			dist := u.dijkstra.dist + 1 // 1 could be replace by a value
			if dist < n.dijkstra.dist {
				n.dijkstra.dist = dist
				n.dijkstra.prev = e
			}
		}
		u.dijkstra.visited = true
		g.dijkstra.missing--
	}

	//fmt.Println("AAA", "missing", g.dijkstra.missing)
	//for _, v := range g.vertices {
	//fmt.Println("v", v, "dist", v.dijkstra.dist, "prev", v.dijkstra.prev)
	//}

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
	return g.vertices
}

func (g *Graph) AddVertex(id string, attrs ...Attrs) *Vertex {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
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
			g.vertices = append(g.vertices[:k], g.vertices[k+1:]...)
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
	}

	src := g.AddVertex(srcID)
	dst := g.AddVertex(dstID)
	edge := newEdge(src, dst, a)
	src.successors = append(src.successors, edge)
	dst.predecessors = append(src.predecessors, edge)
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
		return "[INVALID]"
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
