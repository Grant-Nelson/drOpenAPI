package enumSchema

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
)

type enumImp struct {
	headers.Schema
	values []string
}

func New(base headers.Schema, data headers.Raw) headers.EnumSchema {
	imp := &enumImp{Schema: base}
	imp.setValues(data)
	return imp
}

func (imp *enumImp) setValues(data headers.Raw) {
	imp.values = []string{}
	if values, has := data[`enums`]; has {
		imp.values = values.([]string)
	}
	sort.Strings(imp.values)
}

func (imp *enumImp) Values() []string { return imp.values }
