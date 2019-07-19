package graphman

import (
	"fmt"
	"sort"
	"strings"
)

//
// Edge
//

type Edge struct {
	src     *Vertex
	dst     *Vertex
	deleted bool // temporary variable used by gc()
	Attrs
}

type EdgeCostFN func(e *Edge) int64

func newEdge(src, dst *Vertex, attrs ...Attrs) *Edge {
	var a Attrs
	if len(attrs) > 0 {
		a = attrs[0]
	} else {
		a = make(map[string]interface{})
	}
	return &Edge{
		src:   src,
		dst:   dst,
		Attrs: a,
	}
}

func (e Edge) Dst() *Vertex { return e.dst }
func (e Edge) Src() *Vertex { return e.src }

func (e *Edge) Vertices() Vertices {
	return Vertices{e.src, e.dst}
}

func (e *Edge) String() string {
	ret := fmt.Sprintf("(%s,%s)", e.src.id, e.dst.id)
	if !e.Attrs.IsEmpty() {
		ret += fmt.Sprintf("[%s]", e.Attrs)
	}
	return ret
}

func (e *Edge) HasVertex(id string) bool {
	return e.src.id == id || e.dst.id == id
}

func (e *Edge) OtherVertex(id string) *Vertex {
	if e.src.id == id {
		return e.dst
	}
	if e.dst.id == id {
		return e.src
	}
	return nil
}

//
// Edges
//

type Edges []*Edge

func (e Edges) filtered() Edges {
	filtered := Edges{}
	for _, edge := range e {
		if !edge.deleted {
			filtered = append(filtered, edge)
		}
	}
	return filtered
}

func (e Edges) String() string {
	items := []string{}
	for _, edge := range e {
		items = append(items, edge.String())
	}
	return fmt.Sprintf("{%s}", strings.Join(items, ","))
}

func (e Edges) Equals(other Edges) bool {
	tmp := map[*Edge]int{}
	for _, edge := range e {
		tmp[edge]++
	}
	for _, edge := range other {
		tmp[edge]--
	}
	for _, v := range tmp {
		if v != 0 {
			return false
		}
	}
	return true
}

// AllCombinations returns the different combinations of Edges in a group of Edges
//
// Adapted from https://github.com/mxschmitt/golang-combinations
func (e Edges) AllCombinations() EdgesCombinations {
	combinations := EdgesCombinations{}
	length := uint(len(e))

	for combinationBits := 1; combinationBits < (1 << length); combinationBits++ {
		var combination Edges
		for object := uint(0); object < length; object++ {
			if (combinationBits>>object)&1 == 1 {
				combination = append(combination, e[object])
			}
		}
		combinations = append(combinations, combination)
	}

	return combinations
}

//
// EdgesCombinations
//

type EdgesCombinations []Edges

func (ec EdgesCombinations) LongestToShortest() EdgesCombinations {
	sort.Slice(ec, func(i, j int) bool {
		return len(ec[i]) > len(ec[j])
	})
	return ec
}
