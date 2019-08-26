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
		return invalidPlaceholder
	}
	if len(a) == 0 {
		return "[]"
	}
	elems := []string{}
	for key, val := range a {
		valStr := fmt.Sprintf("%v", val)
		if valStr != "" {
			elems = append(elems, fmt.Sprintf("%s:%s", key, valStr))
		}
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

func (a Attrs) getOrCreatePert() *PertAttrs {
	if _, ok := a["pert"]; !ok {
		a["pert"] = &PertAttrs{}
	}
	return a["pert"].(*PertAttrs)
}

func (a Attrs) SetPertEstimates(opt, real, pess float64) Attrs {
	pert := a.getOrCreatePert()
	pert.Optimistic = opt
	pert.Realistic = real
	pert.Pessimistic = pess
	pert.IsAction = true
	pert.IsState = false
	return a
}

func (a Attrs) SetPertState() Attrs {
	pert := a.getOrCreatePert()
	pert.IsAction = false
	pert.IsState = true
	return a
}

func (a Attrs) SetPertAction() Attrs {
	pert := a.getOrCreatePert()
	pert.IsAction = true
	pert.IsState = false
	return a
}

func (a Attrs) SetPertUntitled() Attrs {
	pert := a.getOrCreatePert()
	pert.IsUntitled = true
	return a
}

func (a Attrs) SetPertZeroTimeActivity() Attrs {
	pert := a.getOrCreatePert()
	pert.IsZeroTimeActivity = true
	return a
}

func (a Attrs) SetPertNonStandardGraph() Attrs {
	pert := a.getOrCreatePert()
	pert.IsNonStandardGraph = true
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

func (a Attrs) GetComment() string {
	if attr, found := a["comment"]; found {
		return attr.(string)
	}
	return ""
}

func (a Attrs) SetComment(comment string) Attrs {
	a["comment"] = comment
	return a
}
