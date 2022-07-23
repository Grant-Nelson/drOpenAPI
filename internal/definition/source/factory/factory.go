package factory

import (
	"fmt"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers/operationType"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/openAPI"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/operation"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/pathItem"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/response"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/schema"
)

type factoryImp struct {
	titleNum int
}

func New() headers.Factory {
	return &factoryImp{
		titleNum: 0,
	}
}

func (f *factoryImp) OpenAPI(data headers.Raw) headers.OpenAPI {
	return openAPI.New(f, data)
}

func (f *factoryImp) PathItem(path string, data headers.Raw) headers.PathItem {
	return pathItem.New(f, path, data)
}

func (f *factoryImp) Operation(opType operationType.Type, data headers.Raw) headers.Operation {
	return operation.New(f, opType, data)
}

func (f *factoryImp) Response(code string, data headers.Raw) headers.Response {
	return response.New(f, code, data)
}

func (f *factoryImp) Schema(title string, data headers.Raw) headers.Schema {
	if len(title) == 0 {
		f.titleNum++
		title = fmt.Sprintf(`unnamed_%d`, f.titleNum)
	}
	return schema.New(f, title, data)
}
