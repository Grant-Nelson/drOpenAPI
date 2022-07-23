package openAPI

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
)

type openAPIImp struct {
	schemaNames []string
	schemas     map[string]headers.Schema
	paths       []string
	pathItems   map[string]headers.PathItem
	tags        []string
}

func New(factory headers.Factory, data headers.Raw) headers.OpenAPI {
	imp := &openAPIImp{}
	imp.setComponents(factory, data)
	imp.setPaths(factory, data)
	imp.setTags(data)

	for i, schema := range imp.schemas {
		if res, ok := schema.(headers.Resolvable); ok {
			imp.schemas[i] = res.Resolve(imp)
		}
	}

	for _, path := range imp.pathItems {
		for _, opType := range path.OperationTypes() {
			op := path.Operation(opType)
			for _, code := range op.ResponseCodes() {
				resp := op.Response(code)
				if res, ok := resp.(headers.Resolvable); ok {
					res.Resolve(imp)
				}
			}
		}
	}

	return imp
}

func (imp *openAPIImp) setComponents(factory headers.Factory, data headers.Raw) {
	imp.schemaNames = []string{}
	imp.schemas = map[string]headers.Schema{}
	if comp, has := data[`components`]; has {
		if s, has := comp.(headers.Raw)[`schemas`]; has {
			for title, value := range s.(headers.Raw) {
				imp.schemas[title] = factory.Schema(title, value.(headers.Raw))
				imp.schemaNames = append(imp.schemaNames, title)
			}
		}
	}
	sort.Strings(imp.schemaNames)
}

func (imp *openAPIImp) setPaths(factory headers.Factory, data headers.Raw) {
	imp.paths = []string{}
	imp.pathItems = map[string]headers.PathItem{}
	if paths, has := data[`paths`]; has {
		for path, item := range paths.(headers.Raw) {
			imp.paths = append(imp.paths, path)
			imp.pathItems[path] = factory.PathItem(path, item.(headers.Raw))
		}
	}
	sort.Strings(imp.paths)
}

func (imp *openAPIImp) setTags(data headers.Raw) {
	imp.tags = []string{}
	if tags, has := data[`tags`]; has {
		for _, tag := range tags.([]interface{}) {
			if name, has := tag.(headers.Raw)[`name`]; has {
				imp.tags = append(imp.tags, name.(string))
			}
		}
	}
	sort.Strings(imp.tags)
}

func (imp *openAPIImp) Paths() []string                       { return imp.paths }
func (imp *openAPIImp) PathItem(path string) headers.PathItem { return imp.pathItems[path] }
func (imp *openAPIImp) SchemaNames() []string                 { return imp.schemaNames }
func (imp *openAPIImp) Schema(name string) headers.Schema     { return imp.schemas[name] }
func (imp *openAPIImp) Tags() []string                        { return imp.tags }
