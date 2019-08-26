package graphman

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	Day  = 24 * time.Hour
	Week = 7 * Day
)

type PertAttrs struct {
	Pessimistic, Realistic, Optimistic float64
	IsZeroTimeActivity                 bool // a.k.a DummyActivity
	IsUntitled                         bool
	IsAction, IsState                  bool
	IsNonStandardGraph                 bool
	IsStart, IsFinish                  bool
}

func (pa PertAttrs) WeightedEstimate() float64 {
	return (pa.Pessimistic + pa.Optimistic + 4*pa.Realistic) / 6
}

func (pa PertAttrs) StandardDeviation() float64 {
	return (pa.Pessimistic - pa.Optimistic) / 6
}

func (pa PertAttrs) Variance() float64 {
	sd := pa.StandardDeviation()
	return sd * sd
}

func (pa PertAttrs) AverageEstimate() float64 {
	return (pa.Pessimistic + pa.Optimistic + pa.Realistic) / 3
}

func (pa PertAttrs) String() string {
	if pa.Optimistic == 0 || pa.Realistic == 0 || pa.Pessimistic == 0 {
		return ""
	}
	return fmt.Sprintf(
		"To=%s,Tm=%s,Tp=%s,Te=%s,σe=%s,Ve=%s",
		prettyFloat(pa.Optimistic),
		prettyFloat(pa.Realistic),
		prettyFloat(pa.Pessimistic),
		prettyFloat(pa.WeightedEstimate()),
		prettyFloat(pa.StandardDeviation()),
		prettyFloat(pa.Variance()),
	)
}

func prettyFloat(f float64) string {
	out := strconv.FormatFloat(f, 'f', 2, 64)
	out = strings.TrimRight(out, "0")
	out = strings.TrimRight(out, ".")
	return out
}

type PertResult struct {
	CriticalPathVariance          float64
	CriticalPathStandardDeviation float64
	CriticalPath                  Path
}

func (pr PertResult) String() string {
	return fmt.Sprintf(
		"Tσe=%s,TVe=%s",
		prettyFloat(pr.CriticalPathStandardDeviation),
		prettyFloat(pr.CriticalPathVariance),
	)
}

func ComputePert(g *Graph) PertResult {
	result := PertResult{}

	// apply pert defaults before computing
	for _, edge := range g.edges {
		pa := edge.GetPert()
		if pa == nil {
			if edge.Attrs == nil {
				edge.Attrs = make(map[string]interface{})
			}
			edge.SetPertEstimates(1, 1, 1)
			pa = edge.GetPert()
		}
		if pa.Realistic == 0 {
			pa.Realistic = 1
		}
		if pa.Pessimistic == 0 && pa.Optimistic == 0 {
			pa.Pessimistic = pa.Realistic
			pa.Optimistic = pa.Realistic
		}
		result.CriticalPathVariance += pa.Variance()
	}
	result.CriticalPathStandardDeviation = math.Sqrt(result.CriticalPathVariance)
	return result
}
