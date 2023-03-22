package schema

import (
	"github.com/Grant-Nelson/drOpenAPI/internal/api"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/enums/compositeType"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/enums/schemaType"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/enums/stateType"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/src/schema/arraySchema"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/src/schema/compositeSchema"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/src/schema/enumSchema"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/src/schema/objectSchema"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/src/schema/referenceSchema"
)

// baseImp is the implementation of the Schema interface.
type baseImp struct {
	title        string
	description  string
	schemaType   schemaType.Type
	format       string
	defaultValue string
	states       map[stateType.Type]bool
}

// has determines if the given key is in the given raw data.
func has(data api.Raw, key string) bool {
	_, has := data[key]
	return has
}

// New creates one of the Schema interfaces based on the given data.
func New(factory api.Factory, title string, data api.Raw) api.Schema {
	if has(data, `$ref`) {
		return referenceSchema.New(title, data)
	}

	imp := &baseImp{title: title}
	imp.setInfo(factory, data)
	imp.setStates(data)

	if imp.schemaType == schemaType.Array || has(data, `items`) {
		imp.schemaType = schemaType.Array
		return arraySchema.New(imp, factory, data)
	}

	if imp.schemaType == schemaType.Enum || has(data, `enums`) {
		imp.schemaType = schemaType.Enum
		return enumSchema.New(imp, data)
	}

	if imp.schemaType == schemaType.Object || has(data, `properties`) {
		imp.schemaType = schemaType.Object
		return objectSchema.New(imp, factory, data)
	}

	for _, compType := range compositeType.All() {
		if has(data, string(compType)) {
			imp.schemaType = schemaType.Composite
			return compositeSchema.New(imp, factory, data)
		}
	}

	return imp
}

// setInfo reads all the basic information from the given data,
// then sets them to this Schema implementation.
func (imp *baseImp) setInfo(factory api.Factory, data api.Raw) {
	if title, has := api.Get[string](data, `title`); has {
		imp.title = title
	}

	if len(imp.title) == 0 {
		imp.title = factory.UniqueName()
	}

	if description, has := api.Get[string](data, `description`); has {
		imp.description = description
	}

	if st, has := api.Get[string](data, `type`); has {
		imp.schemaType = schemaType.Type(st)
	}

	if format, has := api.Get[string](data, `format`); has {
		imp.format = format
	}

	if value, has := api.Get[string](data, `default`); has {
		imp.defaultValue = value
	}
}

// setStates reads the states (e.g. Readonly) from the given data,
// then sets them to this Schema implementation.
func (imp *baseImp) setStates(data api.Raw) {
	imp.states = map[stateType.Type]bool{}
	for _, st := range stateType.All() {
		if state, has := api.Get[bool](data, string(st)); has && state {
			imp.states[st] = true
		}
	}
}

func (imp *baseImp) Title() string                   { return imp.title }
func (imp *baseImp) Description() string             { return imp.description }
func (imp *baseImp) Type() schemaType.Type           { return imp.schemaType }
func (imp *baseImp) Format() string                  { return imp.format }
func (imp *baseImp) Default() string                 { return imp.defaultValue }
func (imp *baseImp) State(state stateType.Type) bool { return imp.states[state] }
