package api

import (
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/compositeType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/operationType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/schemaType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/stateType"
)

type (
	// Raw is the typical data type for an object as returned
	// by an unmarshaler method such as JSON or YAML.
	Raw map[string]interface{}

	// Factory is used for creating new instances of objects.
	Factory interface {

		// OpenAPI creates a new root OpenAPI object, populates it with
		// the given raw data, and returns the new object.
		OpenAPI(data Raw) OpenAPI

		// PathItem creates a new PathItem object for the given path.
		PathItem(path string, data Raw) PathItem

		// Operation creates a new Operation object for the given operation type.
		Operation(opType operationType.Type, data Raw) Operation

		// Response crates a new Response object for the given response code.
		Response(code string, data Raw) Response

		// Schema crates a new Schema type for the given optional initial title.
		// The title may be empty if unknown and may be overwritten by the given data.
		Schema(title string, data Raw) Schema

		// UniqueName creates a new unique name to be used as a title
		// or object name if one isn't provided.
		UniqueName() string
	}

	// Resolvable is added to objects which contain schema to
	// handle resolutions of schema references, e.g. `$ref`.
	// Resolution of references occurs prior to a new OpenAPI being returned.
	Resolvable interface {

		// Resolve uses the root OpenAPI object to lookup and replace a
		// resolvable schema or its children schema which are references
		// with a pointer to the actual schema implementation.
		// The returned schema is the called schema, the replacement,
		// or nil if the object being called is not a schema.
		Resolve(openAPI OpenAPI) Schema
	}

	// OpenAPI represents the root object for an OpenAPI definition.
	// See https://swagger.io/specification/#openapi-object
	OpenAPI interface {

		// Paths is the sorted set of all paths defined.
		Paths() []string

		// PathItem gets the path item object for the given path
		// or null if it none exist with that path.
		PathItem(path string) PathItem

		// SchemaNames is the sorted set of all schemas names.
		SchemaNames() []string

		// Schema gets a Schema for the given name (title)
		// or null if it none exist with that name.
		Schema(name string) Schema

		// Tags is the sorted set of all tags.
		Tags() []string
	}

	// PathItem represents a path item object definition.
	// See https://swagger.io/specification/#path-item-object
	PathItem interface {

		// Path is the path string defined with parameters for this item.
		Path() string

		// OperationTypes is the sorted set of operations that can be run on this path.
		OperationTypes() []operationType.Type

		// Operation gets the operation object for the given operation type
		// or null if it none exist with that operation type.
		Operation(opType operationType.Type) Operation
	}

	// Operation represents an operation object definition.
	// See https://swagger.io/specification/#operation-object
	Operation interface {

		// OpType is the type of operation this object is for.
		OpType() operationType.Type

		// OperationId is the unique identifier for this operation.
		OperationId() string

		// Summary in an optional string defined for this operation.
		Summary() string

		// Description is an optional string defined for this operation.
		Description() string

		// ResponseCodes is the sorted response codes (including "default") from this operation.
		ResponseCodes() []string

		// Response gets the Response object for the given code
		// or null if it none exist with that code.
		Response(code string) Response

		// Tags are the sorted set of tags fr this operation.
		Tags() []string
	}

	// Response represents a response object definition.
	// See https://swagger.io/specification/#response-object
	Response interface {

		// Code is the response code this response is for.
		Code() string

		// Description is an optional string defined for this response.
		Description() string

		// MediaTypes is the sorted media types for this response.
		MediaTypes() []string

		// Content gets the schema for the given media type.
		Content(mediaType string) Schema
	}

	// Schema represents the common schema object definition.
	// This is used by itself for integer, boolean, string, and number schemas.
	// See https://swagger.io/specification/#schema-object
	Schema interface {

		// Title is the required title to show for this schema.
		// If no title is defined a generated one will be given.
		Title() string

		// Description is an optional string defined for this schema.
		Description() string

		// Type indicates the type of this schema.
		// This will be adjusted from the defined type based on the data.
		Type() schemaType.Type

		// Format is an optional more specific type for the schema.
		// See https://swagger.io/specification/#data-types
		Format() string

		// Default is the stringified default value for this schema.
		Default() string

		// State gets the status of a schema state.
		State(state stateType.Type) bool
	}

	// EnumSchema represents a schema definition for an enumerator schema.
	EnumSchema interface {
		Schema

		// Values are the sorted strings for each of the enumerator values.
		Values() []string
	}

	// ArraySchema represents a schema definition for an array schema.
	ArraySchema interface {
		Schema

		// ItemType is the schema for all the elements in the array.
		ItemType() Schema
	}

	// ObjectSchema represents a schema definition for an object schema.
	ObjectSchema interface {
		Schema

		// Required is the sorted list of required properties.
		Required() []string

		// PropertyNames is the sorted property names of this object.
		PropertyNames() []string

		// Property gets the schema for the property for the given name
		// or null if it none exist with that name.
		Property(name string) Schema
	}

	// CompositeSchema represents a schema definition for an object schema
	// which is a composite of other schemas.
	CompositeSchema interface {
		Schema

		// CompositeType is the type of the composite for these components.
		CompositeType() compositeType.Type

		// Components gets the schema which all composite into this schema.
		Components() []Schema
	}
)
