package openAPI

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
)

type openAPIImp struct {
	schemaNames []string
	schemas     map[string]api.Schema
	paths       []string
	pathItems   map[string]api.PathItem
	tags        []string
}

func New(factory api.Factory, data api.Raw) api.OpenAPI {
	imp := &openAPIImp{}
	imp.setComponents(factory, data)
	imp.setPaths(factory, data)
	imp.setTags(data)

	for i, schema := range imp.schemas {
		if res, ok := schema.(api.Resolvable); ok {
			imp.schemas[i] = res.Resolve(imp)
		}
	}

	for _, path := range imp.pathItems {
		for _, opType := range path.OperationTypes() {
			op := path.Operation(opType)
			for _, code := range op.ResponseCodes() {
				resp := op.Response(code)
				if res, ok := resp.(api.Resolvable); ok {
					res.Resolve(imp)
				}
			}
		}
	}

	return imp
}

func (imp *openAPIImp) setComponents(factory api.Factory, data api.Raw) {
	imp.schemaNames = []string{}
	imp.schemas = map[string]api.Schema{}
	if comp, has := data[`components`]; has {
		if s, has := comp.(api.Raw)[`schemas`]; has {
			for title, value := range s.(api.Raw) {
				imp.schemas[title] = factory.Schema(title, value.(api.Raw))
				imp.schemaNames = append(imp.schemaNames, title)
			}
		}
	}
	sort.Strings(imp.schemaNames)
}

func (imp *openAPIImp) setPaths(factory api.Factory, data api.Raw) {
	imp.paths = []string{}
	imp.pathItems = map[string]api.PathItem{}
	if paths, has := data[`paths`]; has {
		for path, item := range paths.(api.Raw) {
			imp.paths = append(imp.paths, path)
			imp.pathItems[path] = factory.PathItem(path, item.(api.Raw))
		}
	}
	sort.Strings(imp.paths)
}

func (imp *openAPIImp) setTags(data api.Raw) {
	imp.tags = []string{}
	if tags, has := data[`tags`]; has {
		for _, tag := range tags.([]interface{}) {
			if name, has := tag.(api.Raw)[`name`]; has {
				imp.tags = append(imp.tags, name.(string))
			}
		}
	}
	sort.Strings(imp.tags)
}

func (imp *openAPIImp) Paths() []string                   { return imp.paths }
func (imp *openAPIImp) PathItem(path string) api.PathItem { return imp.pathItems[path] }
func (imp *openAPIImp) SchemaNames() []string             { return imp.schemaNames }
func (imp *openAPIImp) Schema(name string) api.Schema     { return imp.schemas[name] }
func (imp *openAPIImp) Tags() []string                    { return imp.tags }
