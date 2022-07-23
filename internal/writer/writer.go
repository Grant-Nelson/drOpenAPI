package writer

import (
	"os"
	"strings"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers/schemaType"
	"github.com/grant-nelson/DrOpenAPI/internal/writer/markdown"
	"github.com/grant-nelson/DrOpenAPI/internal/writer/mermaid"
)

func Write(outputPath string, openAPI headers.OpenAPI) {
	md := markdown.New()
	md.SetTitle(`Content API Models`)

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

func addOperation(md markdown.Markdown, path string, op headers.Operation) {
	if res := op.Response(`200`); res != nil {
		if schema := res.Content(`application/json`); schema != nil {

			md.AddSection(op.OperationId())
			md.Bold(`Summary:`).Plain(` ` + op.Summary())
			md.NewPar()
			md.Bold(`Path:`).Plain(` ` + path).LineBreak()
			md.Bold(`Operation:`).Plain(` ` + strings.ToUpper(string(op.OpType()))).LineBreak()
			md.Bold(`Tags:`).Plain(` ` + strings.Join(op.Tags(), `, `))
			md.NewPar()
			md.Bold(`Description:`).Plain(` ` + op.Description())

			diagramOp(md, schema)
		}
	}
}

func diagramOp(md markdown.Markdown, schema headers.Schema) {
	dia := mermaid.New()

	classesToAdd := []headers.Schema{schema}
	for len(classesToAdd) > 0 {
		schema, classesToAdd = classesToAdd[0], classesToAdd[1:]
		if !dia.Has(schema.Title()) {
			newClasses := addClass(dia, schema)
			classesToAdd = append(classesToAdd, newClasses...)
		}
	}

	md.Mermaid(dia.String())
}

func addClass(dia mermaid.Mermaid, schema headers.Schema) []headers.Schema {
	switch schema.Type() {
	case schemaType.Enum:
		addEnum(dia, schema.(headers.EnumSchema))
		return []headers.Schema{}
	case schemaType.Composite:
		return addComposite(dia, schema.(headers.CompositeSchema))
	case schemaType.Object:
		return addObject(dia, schema.(headers.ObjectSchema))
	default:
		return []headers.Schema{}
	}
}

func addEnum(dia mermaid.Mermaid, schema headers.EnumSchema) {
	dia.Enum(schema.Title(), schema.Values()...)
}

func addComposite(dia mermaid.Mermaid, schema headers.CompositeSchema) []headers.Schema {
	newClasses := []headers.Schema{}

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

func addObject(dia mermaid.Mermaid, schema headers.ObjectSchema) []headers.Schema {
	newClasses := []headers.Schema{}

	c := dia.Class(schema.Title())
	obj := schema.(headers.ObjectSchema)
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

func schemaTypeNameAndBase(schema headers.Schema) (string, headers.Schema) {
	switch schema.Type() {
	case schemaType.Enum,
		schemaType.Object:
		return schema.Title(), schema

	case schemaType.Composite:
		return schema.Title(), schema

	case schemaType.Array:
		typeName, base := schemaTypeNameAndBase(schema.(headers.ArraySchema).ItemType())
		return typeName + `[]`, base

	default:
		if len(schema.Format()) > 0 {
			return schema.Format(), nil
		}
		return string(schema.Type()), nil
	}
}
