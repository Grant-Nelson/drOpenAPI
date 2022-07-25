package objectSchema

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
)

type objectImp struct {
	api.Schema
	propertyNames []string
	properties    map[string]api.Schema
}

var _ api.Resolvable = (*objectImp)(nil)

func New(base api.Schema, factory api.Factory, data api.Raw) api.ObjectSchema {
	imp := &objectImp{Schema: base}
	imp.setProperties(factory, data)
	return imp
}

func (imp *objectImp) setProperties(factory api.Factory, data api.Raw) {
	imp.propertyNames = []string{}
	imp.properties = map[string]api.Schema{}
	if properties, has := data[`properties`]; has {
		for name, prop := range properties.(api.Raw) {
			imp.propertyNames = append(imp.propertyNames, name)
			imp.properties[name] = factory.Schema(name, prop.(api.Raw))
		}
	}
	sort.Strings(imp.propertyNames)
}

func (imp *objectImp) Resolve(openAPI api.OpenAPI) api.Schema {
	for name, prop := range imp.properties {
		if res, ok := prop.(api.Resolvable); ok {
			imp.properties[name] = res.Resolve(openAPI)
		}
	}
	return imp
}

func (imp *objectImp) PropertyNames() []string         { return imp.propertyNames }
func (imp *objectImp) Property(name string) api.Schema { return imp.properties[name] }
