package pathItem

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers/operationType"
)

type pathItemImp struct {
	path       string
	opTypes    []operationType.Type
	operations map[operationType.Type]headers.Operation
}

func New(factory headers.Factory, path string, data headers.Raw) headers.PathItem {
	imp := &pathItemImp{path: path}
	imp.setOperations(factory, data)
	return imp
}

func (imp *pathItemImp) setOperations(factory headers.Factory, data headers.Raw) {
	imp.opTypes = []operationType.Type{}
	imp.operations = map[operationType.Type]headers.Operation{}
	for _, opType := range operationType.All() {
		if op, has := data[string(opType)]; has {
			imp.opTypes = append(imp.opTypes, opType)
			imp.operations[opType] = factory.Operation(opType, op.(headers.Raw))
		}
	}
	sort.Slice(imp.opTypes, func(i, j int) bool { return imp.opTypes[i] < imp.opTypes[j] })
}

func (imp *pathItemImp) Path() string                         { return imp.path }
func (imp *pathItemImp) OperationTypes() []operationType.Type { return imp.opTypes }
func (imp *pathItemImp) Operation(opType operationType.Type) headers.Operation {
	return imp.operations[opType]
}
