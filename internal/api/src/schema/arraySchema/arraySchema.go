package arraySchema

import "github.com/grant-nelson/DrOpenAPI/internal/api"

type arrayImp struct {
	api.Schema
	schema api.Schema
}

var _ api.Resolvable = (*arrayImp)(nil)

func New(base api.Schema, factory api.Factory, data api.Raw) api.ArraySchema {
	imp := &arrayImp{Schema: base}
	imp.setItemType(factory, data)
	return imp
}

func (imp *arrayImp) setItemType(factory api.Factory, data api.Raw) {
	if items, has := data[`items`]; has {
		imp.schema = factory.Schema(``, items.(api.Raw))
	}
}

func (imp *arrayImp) Resolve(openAPI api.OpenAPI) api.Schema {
	if res, ok := imp.schema.(api.Resolvable); ok {
		imp.schema = res.Resolve(openAPI)
	}
	return imp
}

func (imp *arrayImp) ItemType() api.Schema { return imp.schema }
