package enumSchema

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
)

type enumImp struct {
	api.Schema
	values []string
}

func New(base api.Schema, data api.Raw) api.EnumSchema {
	imp := &enumImp{Schema: base}
	imp.setValues(data)
	return imp
}

func (imp *enumImp) setValues(data api.Raw) {
	imp.values = []string{}
	if values, has := data[`enums`]; has {
		imp.values = values.([]string)
	}
	sort.Strings(imp.values)
}

func (imp *enumImp) Values() []string { return imp.values }
