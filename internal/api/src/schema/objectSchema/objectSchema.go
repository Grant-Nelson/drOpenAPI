package objectSchema

import (
	"fmt"
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
)

// objectImp is the implementation of the ObjectSchema interface.
// This also implements the Resolvable interface.
type objectImp struct {
	api.Schema
	required      []string
	propertyNames []string
	properties    map[string]api.Schema
}

var _ api.Resolvable = (*objectImp)(nil)

// New creates a new ObjectSchema instance.
func New(base api.Schema, factory api.Factory, data api.Raw) api.ObjectSchema {
	imp := &objectImp{Schema: base}
	imp.setProperties(factory, data)
	imp.setRequired(data)
	return imp
}

// setProperties reads all the properties from the given data,
// then sets them to the ObjectSchema implementation.
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

// setRequired reads the required properties from the given data,
// then sets them to the ObjectSchema implementation.
func (imp *objectImp) setRequired(data api.Raw) {
	imp.required = []string{}
	if required, has := data[`required`]; has {
		for _, req := range required.([]interface{}) {
			reqStr := fmt.Sprint(req)
			imp.required = append(imp.required, reqStr)
		}
	}
	sort.Strings(imp.required)
}

// Resolve performs resolve on all the schemas in the properties
// and replaces them as needed. This returns this instance.
func (imp *objectImp) Resolve(openAPI api.OpenAPI) api.Schema {
	for name, prop := range imp.properties {
		if res, ok := prop.(api.Resolvable); ok {
			imp.properties[name] = res.Resolve(openAPI)
		}
	}
	return imp
}

func (imp *objectImp) Required() []string              { return imp.required }
func (imp *objectImp) PropertyNames() []string         { return imp.propertyNames }
func (imp *objectImp) Property(name string) api.Schema { return imp.properties[name] }
