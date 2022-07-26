package enumSchema

import (
	"sort"

	"github.com/Grant-Nelson/DrOpenAPI/internal/api"
)

// enumImp is the implementation of the EnumSchema interface.
type enumImp struct {
	api.Schema
	values []string
}

// New creates a new EnumSchema instance.
func New(base api.Schema, data api.Raw) api.EnumSchema {
	imp := &enumImp{Schema: base}
	imp.setValues(data)
	return imp
}

// setValues reads all the enumerator values from the given data,
// then sets them to this EnumSchema implementation.
func (imp *enumImp) setValues(data api.Raw) {
	imp.values = []string{}
	if values, has := data[`enums`]; has {
		imp.values = values.([]string)
	}
	sort.Strings(imp.values)
}

func (imp *enumImp) Values() []string { return imp.values }
