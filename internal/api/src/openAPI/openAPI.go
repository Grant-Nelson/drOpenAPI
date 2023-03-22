package openAPI

import (
	"fmt"
	"sort"

	"github.com/Grant-Nelson/drOpenAPI/internal/api"
)

// openAPIImp is the implementation of the OpenAPI interface.
type openAPIImp struct {
	schemaNames []string
	schemas     map[string]api.Schema
	paths       []string
	pathItems   map[string]api.PathItem
	tags        []string
}

// New creates a new OpenAPI instance.
func New(factory api.Factory, data api.Raw) api.OpenAPI {
	imp := &openAPIImp{}
	imp.setComponents(factory, data)
	imp.setPaths(factory, data)
	imp.setTags(data)

	for title, schema := range imp.schemas {
		if res, ok := schema.(api.Resolvable); ok {
			imp.schemas[title] = res.Resolve(imp)
		}
	}

	for _, path := range imp.pathItems {
		for _, opType := range path.OperationTypes() {
			if op, ok := path.Operation(opType).(api.Resolvable); ok {
				op.Resolve(imp)
			}
		}
	}

	return imp
}

// setComponents reads all the component schemas from the given data,
// then sets them to this OpenAPI implementation.
func (imp *openAPIImp) setComponents(factory api.Factory, data api.Raw) {
	imp.schemaNames = []string{}
	imp.schemas = map[string]api.Schema{}
	if s, has := api.Get[api.Raw](data, `components`, `schemas`); has {
		for title, value := range s {
			imp.schemas[title] = factory.Schema(title, value.(api.Raw))
			imp.schemaNames = append(imp.schemaNames, title)
		}
	}
	sort.Strings(imp.schemaNames)
}

// setPaths reads all the path items from the given data,
// then sets them to this OpenAPI implementation.
func (imp *openAPIImp) setPaths(factory api.Factory, data api.Raw) {
	imp.paths = []string{}
	imp.pathItems = map[string]api.PathItem{}
	if paths, has := api.Get[api.Raw](data, `paths`); has {
		for path, item := range paths {
			imp.pathItems[path] = factory.PathItem(path, item.(api.Raw))
			imp.paths = append(imp.paths, path)
		}
	}
	sort.Strings(imp.paths)
}

// setTags reads all the tags from the given data,
// then sets them to this OpenAPI implementation.
func (imp *openAPIImp) setTags(data api.Raw) {
	imp.tags = []string{}
	if tags, has := api.Get[[]any](data, `tags`); has {
		for _, tag := range tags {
			if name, has := api.Get[api.Raw](tag, `name`); has {
				tagStr := fmt.Sprint(name)
				imp.tags = append(imp.tags, tagStr)
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
