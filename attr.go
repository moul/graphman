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

func (a *Attrs) Del(key string) {
	delete(*a, key)
}

func (a Attrs) IsEmpty() bool { return len(a) == 0 }

func (a Attrs) String() string {
	if a == nil {
		return "[INVALID]"
	}
	if len(a) == 0 {
		return "[]"
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

func (a Attrs) GetTitle() string {
	if attr, found := a["title"]; found {
		return attr.(string)
	}
	return ""
}

func (a Attrs) SetPertEstimates(opt, real, pess float64) Attrs {
	a["pert"] = &PertAttrs{
		Optimistic:  opt,
		Realistic:   real,
		Pessimistic: pess,
	}
	return a
}

func (a Attrs) SetPertUntitledState() Attrs {
	a["pert"] = &PertAttrs{
		IsUntitledState: true,
	}
	return a
}

func (a Attrs) SetPertZeroTimeActivity() Attrs {
	a["pert"] = &PertAttrs{
		IsZeroTimeActivity: true,
	}
	return a
}

func (a Attrs) GetPert() *PertAttrs {
	if attr, found := a["pert"]; found {
		return attr.(*PertAttrs)
	}
	return nil
}

func (a Attrs) Clone() Attrs {
	newAttrs := Attrs{}
	for k, v := range a {
		newAttrs[k] = v
	}
	return newAttrs
}

func (a Attrs) SetColor(color string) Attrs {
	a["color"] = color
	return a
}

func (a Attrs) GetColor() string {
	if attr, found := a["color"]; found {
		return attr.(string)
	}
	return ""
}
