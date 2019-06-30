package graphman

import (
	"fmt"
	"sort"
	"strings"
)

type Attrs map[string]interface{}

func (a Attrs) Has(key string) bool {
	_, found := a[key]
	return found
}

func (a *Attrs) Merge(b Attrs) {
	for k, v := range b {
		(*a)[k] = v
	}
}

func (a Attrs) IsEmpty() bool { return len(a) == 0 }

func (a Attrs) String() string {
	if len(a) == 0 {
		return ""
	}
	elems := []string{}
	for key, val := range a {
		elems = append(elems, fmt.Sprintf("%v:%v", key, val))
	}
	sort.Strings(elems)
	return fmt.Sprintf("[%s]", strings.Join(elems, ","))

}

func (a Attrs) SetTitle(title string) Attrs {
	a["title"] = title
	return a
}

func (a Attrs) SetPert(opt, real, pess float64) Attrs {
	a["pert"] = &PertAttrs{
		Optimistic:  opt,
		Realistic:   real,
		Pessimistic: pess,
	}
	return a
}

func (a Attrs) GetPert() *PertAttrs {
	pa, found := a["pert"]
	if !found {
		return nil
	}
	return pa.(*PertAttrs)
}
