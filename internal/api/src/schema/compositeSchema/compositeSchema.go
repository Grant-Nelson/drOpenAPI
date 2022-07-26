package compositeSchema

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/compositeType"
)

// compositeImp is the implementation of the CompositeSchema interface.
// This also implements the Resolvable interface.
type compositeImp struct {
	api.Schema
	compositeType compositeType.Type
	composites    []api.Schema
}

var _ api.Resolvable = (*compositeImp)(nil)

// New creates a new CompositeSchema instance.
func New(base api.Schema, factory api.Factory, data api.Raw) api.CompositeSchema {
	imp := &compositeImp{Schema: base}
	imp.setComponents(factory, data)
	return imp
}

// setComponents reads all the component schemas from the given data,
// then sets them to this CompositeSchema implementation.
func (imp *compositeImp) setComponents(factory api.Factory, data api.Raw) {
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
}

// Resolve runs resolve on all the components in the schema.
// This will return this instance.
func (imp *compositeImp) Resolve(openAPI api.OpenAPI) api.Schema {
	for i, comp := range imp.composites {
		if res, ok := comp.(api.Resolvable); ok {
			imp.composites[i] = res.Resolve(openAPI)
		}
	}
	sort.Slice(imp.composites, func(i, j int) bool {
		return imp.composites[i].Title() < imp.composites[j].Title()
	})
	return imp
}

func (imp *compositeImp) CompositeType() compositeType.Type { return imp.compositeType }
func (imp *compositeImp) Components() []api.Schema          { return imp.composites }
