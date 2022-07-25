package schema

import (
	"fmt"
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/compositeType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/schemaType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/stateType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/schema/arraySchema"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/schema/compositeSchema"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/schema/enumSchema"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/schema/objectSchema"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/schema/referenceSchema"
)

type baseImp struct {
	title        string
	description  string
	schemaType   schemaType.Type
	format       string
	defaultValue string
	required     []string
	states       map[stateType.Type]bool
}

func has(data api.Raw, key string) bool {
	_, has := data[key]
	return has
}

func New(factory api.Factory, title string, data api.Raw) api.Schema {
	if _, has := data[`$ref`]; has {
		return referenceSchema.New(title, data)
	}

	imp := &baseImp{title: title}
	imp.setInfo(factory, data)
	imp.setRequired(data)
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
		break
	}

	return imp
}

func (imp *baseImp) setInfo(factory api.Factory, data api.Raw) {
	if title, has := data[`title`]; has {
		imp.title = fmt.Sprint(title)
	}

	if len(imp.title) == 0 {
		imp.title = factory.UniqueName()
	}

	if description, has := data[`description`]; has {
		imp.description = fmt.Sprint(description)
	}

	if st, has := data[`type`]; has {
		imp.schemaType = schemaType.Type(fmt.Sprint(st))
	}

	if format, has := data[`format`]; has {
		imp.format = fmt.Sprint(format)
	}

	if value, has := data[`default`]; has {
		imp.defaultValue = fmt.Sprint(value)
	}
}

func (imp *baseImp) setRequired(data api.Raw) {
	imp.required = []string{}
	if required, has := data[`required`]; has {
		for _, req := range required.([]interface{}) {
			reqStr := fmt.Sprint(req)
			imp.required = append(imp.required, reqStr)
		}
	}
	sort.Strings(imp.required)
}

func (imp *baseImp) setStates(data api.Raw) {
	imp.states = map[stateType.Type]bool{}
	for _, st := range stateType.All() {
		if state, has := data[string(st)]; has {
			if state.(bool) {
				imp.states[st] = true
			}
		}
	}
}

func (imp *baseImp) Title() string                   { return imp.title }
func (imp *baseImp) Description() string             { return imp.description }
func (imp *baseImp) Type() schemaType.Type           { return imp.schemaType }
func (imp *baseImp) Format() string                  { return imp.format }
func (imp *baseImp) Default() string                 { return imp.defaultValue }
func (imp *baseImp) State(state stateType.Type) bool { return imp.states[state] }