package writer

import (
	"errors"
	"os"
	"strings"

	"github.com/Grant-Nelson/DrOpenAPI/internal/api"
	"github.com/Grant-Nelson/DrOpenAPI/internal/api/enums/schemaType"
	"github.com/Grant-Nelson/DrOpenAPI/internal/api/enums/stateType"
	"github.com/Grant-Nelson/DrOpenAPI/internal/markdown"
	"github.com/Grant-Nelson/DrOpenAPI/internal/markdown/src/factory"
)

// Write creates or overwrites the given output path with the markdown
// file containing the description of the given OpenAPI object.
// If the given title is non-empty, then it will be written as the title
// shown at the top of the created markdown file.
func Write(outputPath, title string, openAPI api.OpenAPI) {
	if len(outputPath) == 0 {
		panic(errors.New(`must provide a non-empty output file path`))
	}

	md := factory.New().Markdown(title)

	for _, path := range openAPI.Paths() {
		item := openAPI.PathItem(path)
		for _, opType := range item.OperationTypes() {
			op := item.Operation(opType)
			addOperation(md, path, op)
		}
	}

	err := os.WriteFile(outputPath, []byte(md.String()), 0644)
	if err != nil {
		panic(err)
	}
}

// addOperation adds the given operation information into the given markdown being created.
// The given path is the path described in the OpenAPI definition the given operation is under.
func addOperation(md markdown.Markdown, path string, op api.Operation) {
	md.Par().HorizontalLine()
	md.Section(op.OperationId())

	if len(op.Summary()) > 0 {
		md.Par().Bold(`Summary:`).Write(` %s`, op.Summary())
	}

	if len(op.Description()) > 0 {
		md.Par().Bold(`Description:`).Write(` %s`, op.Description())
	}

	md.Par().
		Bold(`Path:`).Write(` `).Code(path).LineBreak().
		Bold(`Operation:`).Write(` `).Code(strings.ToUpper(string(op.OpType()))).LineBreak().
		Bold(`Tags:`).Write(` `).Code(strings.Join(op.Tags(), `, `))

	if op.RequestBody() != nil {
		md.Subsection(op.OperationId() + ` Request`)
		diagramOp(md, op.RequestBody())
	}

	for _, code := range op.ResponseCodes() {
		res := op.Response(code)
		if schema := res.Content(`application/json`); schema != nil {
			typeName, base := schemaTypeNameAndBase(schema)

			md.Subsection(op.OperationId() + ` ` + code + ` Response`)

			md.Par().
				Bold(`Code:`).Write(` `).Code(code).LineBreak().
				Bold(`Returns:`).Write(` `).Code(typeName)

			if base != nil {
				diagramOp(md, base)
			}
		}
	}
}

// diagramOp adds a mermaid diagram to the given markdown as a way to describe the given schema.
func diagramOp(md markdown.Markdown, schema api.Schema) {
	dia := md.Mermaid()
	classesToAdd := addClass(dia, schema)
	for len(classesToAdd) > 0 {
		next := classesToAdd[0]
		classesToAdd = classesToAdd[1:]
		if !dia.Has(next.Title()) {
			newClasses := addClass(dia, next)
			classesToAdd = append(classesToAdd, newClasses...)
		}
	}
}

// addClass adds a class for the given schema into the given mermaid diagram.
// The returned schema are the schema referenced by the given schema.
func addClass(dia markdown.Mermaid, schema api.Schema) []api.Schema {
	switch schema.Type() {
	case schemaType.Enum:
		addEnum(dia, schema.(api.EnumSchema))
		return []api.Schema{}
	case schemaType.Composite:
		return addComposite(dia, schema.(api.CompositeSchema))
	case schemaType.Object:
		return addObject(dia, schema.(api.ObjectSchema))
	default:
		return []api.Schema{}
	}
}

// addEnum adds a class for the given enumerator schema into the given mermaid diagram.
func addEnum(dia markdown.Mermaid, schema api.EnumSchema) {
	dia.Enum(schema.Title(), schema.Values()...)
}

// addComposite adds a class for the given composite (e.g. `oneOf`)
// schema into the given mermaid diagram.
func addComposite(dia markdown.Mermaid, schema api.CompositeSchema) []api.Schema {
	newClasses := []api.Schema{}

	c := dia.Interface(schema.Title())
	c.AddEntry(string(schema.CompositeType()) + `:`)
	for _, comp := range schema.Components() {
		c.AddEntry(`- ` + comp.Title())
		c.ConnectTo(comp.Title())
		if !dia.Has(comp.Title()) {
			newClasses = append(newClasses, comp)
		}
	}

	return newClasses
}

// addObject adds a class for the given object schema into the given mermaid diagram.
func addObject(dia markdown.Mermaid, schema api.ObjectSchema) []api.Schema {
	newClasses := []api.Schema{}

	c := dia.Class(schema.Title())
	for _, name := range schema.PropertyNames() {
		prop := schema.Property(name)
		typeName, base := schemaTypeNameAndBase(prop)

		c.AddMember(name, typeName)

		if base != nil {
			c.ConnectTo(base.Title())
			if !dia.Has(base.Title()) {
				newClasses = append(newClasses, base)
			}
		}
	}

	return newClasses
}

// schemaTypeNameAndBase determines the type for the given schema as a string
// and inner most schema that is referenced by this type.
// If there are no inner schema then the returned schema will be nil.
func schemaTypeNameAndBase(schema api.Schema) (string, api.Schema) {
	head := ``
	if schema.State(stateType.Nullable) {
		head = `*`
	}
	switch schema.Type() {
	case schemaType.Enum,
		schemaType.Object:
		return head + schema.Title(), schema

	case schemaType.Composite:
		return head + schema.Title(), schema

	case schemaType.Array:
		typeName, base := schemaTypeNameAndBase(schema.(api.ArraySchema).ItemType())
		return head + typeName + `[]`, base

	default:
		if len(schema.Format()) > 0 {
			return head + schema.Format(), nil
		}
		return head + string(schema.Type()), nil
	}
}
