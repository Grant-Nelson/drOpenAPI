package writer

import (
	"errors"
	"os"
	"strings"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/schemaType"
	"github.com/grant-nelson/DrOpenAPI/internal/markdown"
	"github.com/grant-nelson/DrOpenAPI/internal/markdown/src/factory"
)

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

func addOperation(md markdown.Markdown, path string, op api.Operation) {
	if res := op.Response(`200`); res != nil {
		if schema := res.Content(`application/json`); schema != nil {

			md.Section(op.OperationId())

			if len(op.Summary()) > 0 {
				md.Par().Bold(`Summary:`).Write(` %s`, op.Summary())
			}

			if len(op.Description()) > 0 {
				md.Par().Bold(`Description:`).Write(` %s`, op.Description())
			}

			typeName, base := schemaTypeNameAndBase(schema)

			md.Par().
				Bold(`Path:`).Write(` `).Code(path).LineBreak().
				Bold(`Operation:`).Write(` `).Code(strings.ToUpper(string(op.OpType()))).LineBreak().
				Bold(`Tags:`).Write(` `).Code(strings.Join(op.Tags(), `, `)).LineBreak().
				Bold(`Returns:`).Write(` `).Code(typeName)

			if base != nil {
				diagramOp(md, base)
			}
		}
	}
}

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

func addEnum(dia markdown.Mermaid, schema api.EnumSchema) {
	dia.Enum(schema.Title(), schema.Values()...)
}

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

func addObject(dia markdown.Mermaid, schema api.ObjectSchema) []api.Schema {
	newClasses := []api.Schema{}

	c := dia.Class(schema.Title())
	obj := schema.(api.ObjectSchema)
	for _, name := range obj.PropertyNames() {
		prop := obj.Property(name)
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

func schemaTypeNameAndBase(schema api.Schema) (string, api.Schema) {
	switch schema.Type() {
	case schemaType.Enum,
		schemaType.Object:
		return schema.Title(), schema

	case schemaType.Composite:
		return schema.Title(), schema

	case schemaType.Array:
		typeName, base := schemaTypeNameAndBase(schema.(api.ArraySchema).ItemType())
		return typeName + `[]`, base

	default:
		if len(schema.Format()) > 0 {
			return schema.Format(), nil
		}
		return string(schema.Type()), nil
	}
}
