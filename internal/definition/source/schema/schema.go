package schema

import (
	"fmt"
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers/compositeType"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers/schemaType"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers/stateType"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/schema/arraySchema"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/schema/compositeSchema"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/schema/enumSchema"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/schema/objectSchema"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/schema/referenceSchema"
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

func has(data headers.Raw, key string) bool {
	_, has := data[key]
	return has
}

func New(factory headers.Factory, title string, data headers.Raw) headers.Schema {
	if _, has := data[`$ref`]; has {
		return referenceSchema.New(title, data)
	}

	imp := &baseImp{title: title}
	imp.setInfo(data)
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

func (imp *baseImp) setInfo(data headers.Raw) {
	if title, has := data[`title`]; has {
		imp.title = title.(string)
	}
	if description, has := data[`description`]; has {
		imp.description = description.(string)
	}
	if st, has := data[`type`]; has {
		imp.schemaType = schemaType.Type(st.(string))
	}
	if format, has := data[`format`]; has {
		imp.format = format.(string)
	}
	if value, has := data[`default`]; has {
		imp.defaultValue = fmt.Sprint(value)
	}
}

func (imp *baseImp) setRequired(data headers.Raw) {
	imp.required = []string{}
	if required, has := data[`required`]; has {
		for _, req := range required.([]interface{}) {
			imp.required = append(imp.required, req.(string))
		}
	}
	sort.Strings(imp.required)
}

func (imp *baseImp) setStates(data headers.Raw) {
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
