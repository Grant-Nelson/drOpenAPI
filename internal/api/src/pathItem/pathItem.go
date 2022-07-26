package pathItem

import (
	"sort"

	"github.com/Grant-Nelson/DrOpenAPI/internal/api"
	"github.com/Grant-Nelson/DrOpenAPI/internal/api/enums/operationType"
)

// pathItemImp is the implementation of the PathItem interface.
type pathItemImp struct {
	path       string
	opTypes    []operationType.Type
	operations map[operationType.Type]api.Operation
}

// New creates a new PathItem instance.
func New(factory api.Factory, path string, data api.Raw) api.PathItem {
	imp := &pathItemImp{path: path}
	imp.setOperations(factory, data)
	return imp
}

// setOperations reads all the operations from the given data,
// then sets them to this PathItem implementation.
func (imp *pathItemImp) setOperations(factory api.Factory, data api.Raw) {
	imp.opTypes = []operationType.Type{}
	imp.operations = map[operationType.Type]api.Operation{}
	for _, opType := range operationType.All() {
		if op, has := data[string(opType)]; has {
			imp.opTypes = append(imp.opTypes, opType)
			imp.operations[opType] = factory.Operation(opType, op.(api.Raw))
		}
	}
	sort.Slice(imp.opTypes, func(i, j int) bool {
		return imp.opTypes[i] < imp.opTypes[j]
	})
}

func (imp *pathItemImp) Path() string                         { return imp.path }
func (imp *pathItemImp) OperationTypes() []operationType.Type { return imp.opTypes }
func (imp *pathItemImp) Operation(opType operationType.Type) api.Operation {
	return imp.operations[opType]
}
