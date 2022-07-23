package arraySchema

import "github.com/grant-nelson/DrOpenAPI/internal/definition/headers"

type arrayImp struct {
	headers.Schema
	schema headers.Schema
}

var _ headers.Resolvable = (*arrayImp)(nil)

func New(base headers.Schema, factory headers.Factory, data headers.Raw) headers.ArraySchema {
	imp := &arrayImp{Schema: base}
	imp.setItemType(factory, data)
	return imp
}

func (imp *arrayImp) setItemType(factory headers.Factory, data headers.Raw) {
	if items, has := data[`items`]; has {
		imp.schema = factory.Schema(``, items.(headers.Raw))
	}
}

func (imp *arrayImp) Resolve(openAPI headers.OpenAPI) headers.Schema {
	if res, ok := imp.schema.(headers.Resolvable); ok {
		imp.schema = res.Resolve(openAPI)
	}
	return imp
}

func (imp *arrayImp) ItemType() headers.Schema { return imp.schema }
