package pathItem

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/operationType"
)

type pathItemImp struct {
	path       string
	opTypes    []operationType.Type
	operations map[operationType.Type]api.Operation
}

func New(factory api.Factory, path string, data api.Raw) api.PathItem {
	imp := &pathItemImp{path: path}
	imp.setOperations(factory, data)
	return imp
}

func (imp *pathItemImp) setOperations(factory api.Factory, data api.Raw) {
	imp.opTypes = []operationType.Type{}
	imp.operations = map[operationType.Type]api.Operation{}
	for _, opType := range operationType.All() {
		if op, has := data[string(opType)]; has {
			imp.opTypes = append(imp.opTypes, opType)
			imp.operations[opType] = factory.Operation(opType, op.(api.Raw))
		}
	}
	sort.Slice(imp.opTypes, func(i, j int) bool { return imp.opTypes[i] < imp.opTypes[j] })
}

func (imp *pathItemImp) Path() string                         { return imp.path }
func (imp *pathItemImp) OperationTypes() []operationType.Type { return imp.opTypes }
func (imp *pathItemImp) Operation(opType operationType.Type) api.Operation {
	return imp.operations[opType]
}
