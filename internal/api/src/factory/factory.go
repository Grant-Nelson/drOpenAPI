package factory

import (
	"fmt"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/operationType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/openAPI"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/operation"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/pathItem"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/response"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/schema"
)

// factoryImp is the implementation of the Factory interface.
type factoryImp struct {
	num int
}

// New creates a new API factory.
func New() api.Factory {
	return &factoryImp{num: 0}
}

func (f *factoryImp) OpenAPI(data api.Raw) api.OpenAPI {
	return openAPI.New(f, data)
}

func (f *factoryImp) PathItem(path string, data api.Raw) api.PathItem {
	return pathItem.New(f, path, data)
}

func (f *factoryImp) Operation(opType operationType.Type, data api.Raw) api.Operation {
	return operation.New(f, opType, data)
}

func (f *factoryImp) Response(code string, data api.Raw) api.Response {
	return response.New(f, code, data)
}

func (f *factoryImp) Schema(title string, data api.Raw) api.Schema {
	return schema.New(f, title, data)
}

func (f *factoryImp) UniqueName() string {
	f.num++
	return fmt.Sprintf(`unnamed%d`, f.num)
}
