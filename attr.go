package graphman

type Attr interface {
	Key() interface{}
	Value() interface{}
	SetValue(interface{})
}

// attr uses a standard Go slice for internal storage.
type attr [2]interface{}

func (a *attr) Key() interface{}              { return a[0] }
func (a *attr) Value() interface{}            { return a[1] }
func (a *attr) SetValue(newValue interface{}) { a[1] = newValue }

func NewAttr(key, value interface{}) Attr {
	return &attr{key, value}
}
