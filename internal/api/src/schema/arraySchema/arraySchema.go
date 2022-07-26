package arraySchema

import "github.com/Grant-Nelson/DrOpenAPI/internal/api"

// arrayImp is the implementation of the ArraySchema interface.
// This also implements the Resolvable interface.
type arrayImp struct {
	api.Schema
	schema api.Schema
}

var _ api.Resolvable = (*arrayImp)(nil)

// New creates a new ArraySchema instance.
func New(base api.Schema, factory api.Factory, data api.Raw) api.ArraySchema {
	imp := &arrayImp{Schema: base}
	imp.setItemType(factory, data)
	return imp
}

// setItemType reads the array element type from the given data,
// then sets them to this ArraySchema implementation.
func (imp *arrayImp) setItemType(factory api.Factory, data api.Raw) {
	if items, has := data[`items`]; has {
		imp.schema = factory.Schema(``, items.(api.Raw))
	}
}

// Resolve checks and may replace the item type schema.
func (imp *arrayImp) Resolve(openAPI api.OpenAPI) api.Schema {
	if res, ok := imp.schema.(api.Resolvable); ok {
		imp.schema = res.Resolve(openAPI)
	}
	return imp
}

func (imp *arrayImp) ItemType() api.Schema { return imp.schema }
