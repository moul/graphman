package viz // import "moul.io/graphman/viz"

import (
	"fmt"

	graphviz "github.com/awalterschulze/gographviz"
	"moul.io/graphman"
)

type Opts struct {
	CommentsInLabel bool
}

func ToGraphviz(g *graphman.Graph, opts *Opts) (string, error) {
	if opts == nil {
		opts = &Opts{}
	}
	gv := graphviz.NewGraph()
	if err := gv.SetName("G"); err != nil {
		return "", err
	}
	if err := gv.SetDir(true); err != nil {
		return "", err
	}
	for k, v := range attrsFromGraph(g, opts) {
		if err := gv.AddAttr("G", k, v); err != nil {
			return "", err
		}
	}
	for _, vertex := range g.Vertices() {
		if err := gv.AddNode(
			"G",
			escape(vertex.ID()),
			attrsFromVertex(vertex, opts),
		); err != nil {
			return "", err
		}
	}
	for _, edge := range g.Edges() {
		if err := gv.AddEdge(
			escape(edge.Src().ID()),
			escape(edge.Dst().ID()),
			true,
			attrsFromEdge(edge, opts),
		); err != nil {
			return "", err
		}
	}
	return gv.String(), nil
}

func attrsFromGraph(graph *graphman.Graph, opts *Opts) map[string]string {
	attrs := map[string]string{}
	attrs[string(graphviz.RankDir)] = "LR"
	attrsGeneric(graph.Attrs, attrs, opts)
	return attrs
}

func attrsFromVertex(vertex *graphman.Vertex, opts *Opts) map[string]string {
	attrs := map[string]string{}
	attrs[string(graphviz.Shape)] = "box"
	attrs[string(graphviz.Style)] = "rounded"
	attrs[string(graphviz.Label)] = vertex.ID()
	if pert := vertex.Attrs.GetPert(); pert != nil {
		if pert.IsUntitledState {
			attrs[string(graphviz.Label)] = " "
			attrs[string(graphviz.Shape)] = "circle"
			// attrs[string(graphviz.Style)] = "dashed"
		}
	}
	attrsGeneric(vertex.Attrs, attrs, opts)
	return attrs
}

func attrsFromEdge(edge *graphman.Edge, opts *Opts) map[string]string {
	attrs := map[string]string{}
	attrs[string(graphviz.Label)] = ""
	if pert := edge.Attrs.GetPert(); pert != nil {
		if pert.IsZeroTimeActivity {
			attrs[string(graphviz.Style)] = "dashed"
		}
	}
	attrsGeneric(edge.Attrs, attrs, opts)
	return attrs
}

func attrsGeneric(a graphman.Attrs, attrs map[string]string, opts *Opts) {
	ac := a.Clone()
	if color := a.GetColor(); color != "" {
		attrs[string(graphviz.Color)] = color
		ac.Del("color")
	}
	if title := a.GetTitle(); title != "" {
		attrs[string(graphviz.Label)] = title
		ac.Del("title")
	}
	if comment := a.GetComment(); comment != "" {
		attrs[string(graphviz.Comment)] = comment
		ac.Del("comment")
	}
	if len(ac) > 0 {
		attrs[string(graphviz.Comment)] = ""
		for k, v := range ac {
			switch k {
			case "rankdir", "shape", "style":
				attrs[k] = v.(string)
			default:
				if vStr := fmt.Sprintf("%v", v); vStr != "" {
					line := fmt.Sprintf("\n%s: %s", k, vStr)
					attrs[string(graphviz.Comment)] += line
				}
			}
		}
	}
	if opts.CommentsInLabel {
		attrs[string(graphviz.Label)] += attrs[string(graphviz.Comment)]
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
