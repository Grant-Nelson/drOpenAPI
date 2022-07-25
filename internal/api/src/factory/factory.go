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

type factoryImp struct {
	titleNum int
}

func New() api.Factory {
	return &factoryImp{
		titleNum: 0,
	}
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
	if len(title) == 0 {
		f.titleNum++
		title = fmt.Sprintf(`unnamed_%d`, f.titleNum)
	}
	return schema.New(f, title, data)
}
