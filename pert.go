package graphman

import (
	"fmt"
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
}

func (pa PertAttrs) Estimate() float64 {
	return (pa.Pessimistic + pa.Optimistic + 4*pa.Realistic) / 6
}

func (pa PertAttrs) StandardDeviation() float64 {
	return (pa.Pessimistic - pa.Optimistic) / 6
}

func (pa PertAttrs) Variance() float64 {
	sd := pa.StandardDeviation()
	return sd * sd
}

func (pa PertAttrs) String() string {
	return fmt.Sprintf(
		"To=%s,Tm=%s,Tp=%s,Te=%s,σe=%s,Ve=%s",
		prettyFloat(pa.Optimistic),
		prettyFloat(pa.Realistic),
		prettyFloat(pa.Pessimistic),
		prettyFloat(pa.Estimate()),
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
