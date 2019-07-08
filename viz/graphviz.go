package viz // import "moul.io/graphman/viz"

import (
	"fmt"

	graphviz "github.com/awalterschulze/gographviz"
	"moul.io/graphman"
)

func ToGraphviz(g *graphman.Graph) (string, error) {
	gv := graphviz.NewGraph()
	if err := gv.SetName("G"); err != nil {
		return "", err
	}
	if err := gv.SetDir(true); err != nil {
		return "", err
	}
	for k, v := range attrsFromGraph(g) {
		if err := gv.AddAttr("G", k, v); err != nil {
			return "", err
		}
	}
	for _, vertex := range g.Vertices() {
		if err := gv.AddNode(
			"G",
			vertex.ID(),
			attrsFromVertex(vertex),
		); err != nil {
			return "", err
		}
	}
	for _, edge := range g.Edges() {
		if err := gv.AddEdge(
			edge.Src().ID(),
			edge.Dst().ID(),
			true,
			attrsFromEdge(edge),
		); err != nil {
			return "", err
		}
	}
	return gv.String(), nil
}

func attrsFromGraph(graph *graphman.Graph) map[string]string {
	attrs := map[string]string{}
	attrs[string(graphviz.RankDir)] = "LR"
	attrsGeneric(graph.Attrs, attrs)
	return attrs
}

func attrsFromVertex(vertex *graphman.Vertex) map[string]string {
	attrs := map[string]string{}
	attrs[string(graphviz.Shape)] = "box"
	attrs[string(graphviz.Style)] = "rounded"
	attrs[string(graphviz.Label)] = vertex.ID()
	if pert := vertex.Attrs.GetPert(); pert != nil {
		if pert.IsUntitledState {
			attrs[string(graphviz.Label)] = " "
			attrs[string(graphviz.Shape)] = "circle"
		}
	}
	attrsGeneric(vertex.Attrs, attrs)
	return attrs
}

func attrsFromEdge(edge *graphman.Edge) map[string]string {
	attrs := map[string]string{}
	attrs[string(graphviz.Label)] = ""
	if pert := edge.Attrs.GetPert(); pert != nil {
		if pert.IsZeroTimeActivity {
			attrs[string(graphviz.Style)] = "dashed"
		}
	}
	attrsGeneric(edge.Attrs, attrs)
	return attrs
}

func attrsGeneric(a graphman.Attrs, attrs map[string]string) {
	ac := a.Clone()
	if color := a.GetColor(); color != "" {
		attrs[string(graphviz.Color)] = color
		ac.Del("color")
	}
	if title := a.GetTitle(); title != "" {
		attrs[string(graphviz.Label)] = title
		ac.Del("title")
	}
	if len(ac) > 0 {
		attrs[string(graphviz.Comment)] = ""
		for k, v := range ac {
			line := fmt.Sprintf("\n%s: %v", k, v)
			attrs[string(graphviz.Comment)] += line
		}
	}

	// sanitize
	for _, key := range []graphviz.Attr{
		graphviz.Label,
		graphviz.Comment,
	} {
		if val := attrs[string(key)]; val != "" {
			attrs[string(key)] = escape(val)
		} else {
			delete(attrs, string(key))
		}
	}
}

func escape(input string) string {
	return fmt.Sprintf("%q", input)
}
