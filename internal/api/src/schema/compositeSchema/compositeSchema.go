package compositeSchema

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/compositeType"
)

type compositeImp struct {
	api.Schema
	propertyNames []string
	properties    map[string]api.Schema
	compositeType compositeType.Type
	composites    []api.Schema
}

var _ api.Resolvable = (*compositeImp)(nil)

func New(base api.Schema, factory api.Factory, data api.Raw) api.CompositeSchema {
	imp := &compositeImp{Schema: base}
	for _, compType := range compositeType.All() {
		if comp, has := data[string(compType)]; has {
			imp.compositeType = compType
			for _, schema := range comp.([]interface{}) {
				imp.composites = append(imp.composites,
					factory.Schema(``, schema.(api.Raw)))
			}
			break
		}
	}
	return imp
}

func (imp *compositeImp) Resolve(openAPI api.OpenAPI) api.Schema {
	for i, comp := range imp.composites {
		if res, ok := comp.(api.Resolvable); ok {
			imp.composites[i] = res.Resolve(openAPI)
		}
	}

	imp.properties = map[string]api.Schema{}
	for _, comp := range imp.composites {
		if obj, ok := comp.(api.ObjectSchema); ok {
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

func (imp *compositeImp) PropertyNames() []string           { return imp.propertyNames }
func (imp *compositeImp) Property(name string) api.Schema   { return imp.properties[name] }
func (imp *compositeImp) CompositeType() compositeType.Type { return imp.compositeType }
func (imp *compositeImp) Components() []api.Schema          { return imp.composites }
