package response

import (
	"fmt"

	"github.com/Grant-Nelson/DrOpenAPI/internal/api"
)

// responseImp is the implementation of the Response interface.
// This also implements the Resolvable interface.
type responseImp struct {
	code        string
	description string
	mediaTypes  []string
	content     map[string]api.Schema
}

var _ api.Resolvable = (*responseImp)(nil)

// New creates a new Response instance.
func New(factory api.Factory, code string, data api.Raw) api.Response {
	imp := &responseImp{code: code}
	imp.setInfo(data)
	imp.setContent(factory, data)
	return imp
}

// setInfo reads all the basic information from the given data,
// then sets them to this Response implementation.
func (imp *responseImp) setInfo(data api.Raw) {
	if description, has := data[`description`]; has {
		imp.description = fmt.Sprint(description)
	}
}

// setContent reads all the content schema from the given data,
// then sets them to this Response implementation.
func (imp *responseImp) setContent(factory api.Factory, data api.Raw) {
	imp.content = map[string]api.Schema{}
	if mediaTypes, has := data[`content`]; has {
		for mediaType, media := range mediaTypes.(api.Raw) {
			if schema, has := media.(api.Raw)[`schema`]; has {
				imp.content[mediaType] = factory.Schema(``, schema.(api.Raw))
			}
		}
	}
}

// Resolve runs resolve on any of the contained schema which is also resolvable.
// This will return nil since this isn't a schema to replace.
func (imp *responseImp) Resolve(openAPI api.OpenAPI) api.Schema {
	for mediaType, content := range imp.content {
		if res, ok := content.(api.Resolvable); ok {
			imp.content[mediaType] = res.Resolve(openAPI)
		}
	}
	return nil
}

func (imp *responseImp) Code() string                        { return imp.code }
func (imp *responseImp) Description() string                 { return imp.description }
func (imp *responseImp) MediaTypes() []string                { return imp.mediaTypes }
func (imp *responseImp) Content(mediaType string) api.Schema { return imp.content[mediaType] }
