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
