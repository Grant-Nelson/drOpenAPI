package compositeSchema

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers/compositeType"
)

type compositeImp struct {
	headers.Schema
	propertyNames []string
	properties    map[string]headers.Schema
	compositeType compositeType.Type
	composites    []headers.Schema
}

var _ headers.Resolvable = (*compositeImp)(nil)

func New(base headers.Schema, factory headers.Factory, data headers.Raw) headers.CompositeSchema {
	imp := &compositeImp{Schema: base}
	for _, compType := range compositeType.All() {
		if comp, has := data[string(compType)]; has {
			imp.compositeType = compType
			for _, schema := range comp.([]interface{}) {
				imp.composites = append(imp.composites,
					factory.Schema(``, schema.(headers.Raw)))
			}
			break
		}
	}
	return imp
}

func (imp *compositeImp) Resolve(openAPI headers.OpenAPI) headers.Schema {
	for i, comp := range imp.composites {
		if res, ok := comp.(headers.Resolvable); ok {
			imp.composites[i] = res.Resolve(openAPI)
		}
	}

	imp.properties = map[string]headers.Schema{}
	for _, comp := range imp.composites {
		if obj, ok := comp.(headers.ObjectSchema); ok {
			for _, name := range obj.PropertyNames() {
				imp.properties[name] = obj.Property(name)
			}
		}
	}

	imp.propertyNames = []string{}
	for name := range imp.properties {
		imp.propertyNames = append(imp.propertyNames, name)
	}
	sort.Strings(imp.propertyNames)

	return imp
}

func (imp *compositeImp) PropertyNames() []string             { return imp.propertyNames }
func (imp *compositeImp) Property(name string) headers.Schema { return imp.properties[name] }
func (imp *compositeImp) CompositeType() compositeType.Type   { return imp.compositeType }
func (imp *compositeImp) Components() []headers.Schema        { return imp.composites }
