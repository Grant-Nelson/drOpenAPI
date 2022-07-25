package api

import (
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/compositeType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/operationType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/schemaType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/stateType"
)

// See https://swagger.io/specification
type (
	Raw map[string]interface{}

	Factory interface {
		OpenAPI(data Raw) OpenAPI
		PathItem(path string, data Raw) PathItem
		Operation(opType operationType.Type, data Raw) Operation
		Response(code string, data Raw) Response
		Schema(title string, data Raw) Schema
		UniqueName() string
	}

	Resolvable interface {
		Resolve(openAPI OpenAPI) Schema
	}

	OpenAPI interface {
		Paths() []string
		PathItem(path string) PathItem
		SchemaNames() []string
		Schema(name string) Schema
		Tags() []string
	}

	PathItem interface {
		Path() string
		OperationTypes() []operationType.Type
		Operation(opType operationType.Type) Operation
	}

	Operation interface {
		OpType() operationType.Type
		OperationId() string
		Summary() string
		Description() string
		ResponseCodes() []string
		Response(code string) Response
		Tags() []string
	}

	Response interface {
		Code() string
		Description() string
		MediaTypes() []string
		Content(mediaType string) Schema
	}

	Schema interface {
		Title() string
		Description() string
		Type() schemaType.Type
		Format() string
		Default() string
		State(state stateType.Type) bool
	}

	EnumSchema interface {
		Schema
		Values() []string
	}

	ArraySchema interface {
		Schema
		ItemType() Schema
	}

	ObjectSchema interface {
		Schema
		PropertyNames() []string
		Property(name string) Schema
	}

	CompositeSchema interface {
		ObjectSchema
		CompositeType() compositeType.Type
		Components() []Schema
	}
)
