package objectSchema

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
)

type objectImp struct {
	headers.Schema
	propertyNames []string
	properties    map[string]headers.Schema
}

var _ headers.Resolvable = (*objectImp)(nil)

func New(base headers.Schema, factory headers.Factory, data headers.Raw) headers.ObjectSchema {
	imp := &objectImp{Schema: base}
	imp.setProperties(factory, data)
	return imp
}

func (imp *objectImp) setProperties(factory headers.Factory, data headers.Raw) {
	imp.propertyNames = []string{}
	imp.properties = map[string]headers.Schema{}
	if properties, has := data[`properties`]; has {
		for name, prop := range properties.(headers.Raw) {
			imp.propertyNames = append(imp.propertyNames, name)
			imp.properties[name] = factory.Schema(name, prop.(headers.Raw))
		}
	}
	sort.Strings(imp.propertyNames)
}

func (imp *objectImp) Resolve(openAPI headers.OpenAPI) headers.Schema {
	for name, prop := range imp.properties {
		if res, ok := prop.(headers.Resolvable); ok {
			imp.properties[name] = res.Resolve(openAPI)
		}
	}
	return imp
}

func (imp *objectImp) PropertyNames() []string             { return imp.propertyNames }
func (imp *objectImp) Property(name string) headers.Schema { return imp.properties[name] }
