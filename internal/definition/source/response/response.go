package response

import "github.com/grant-nelson/DrOpenAPI/internal/definition/headers"

type responseImp struct {
	code        string
	description string
	mediaTypes  []string
	content     map[string]headers.Schema
}

var _ headers.Resolvable = (*responseImp)(nil)

func New(factory headers.Factory, code string, data headers.Raw) headers.Response {
	imp := &responseImp{code: code}
	imp.setInfo(data)
	imp.setContent(factory, data)
	return imp
}

func (imp *responseImp) setInfo(data headers.Raw) {
	if description, has := data[`description`]; has {
		imp.description = description.(string)
	}
}

func (imp *responseImp) setContent(factory headers.Factory, data headers.Raw) {
	imp.content = map[string]headers.Schema{}
	if mediaTypes, has := data[`content`]; has {
		for mediaType, media := range mediaTypes.(headers.Raw) {
			if schema, has := media.(headers.Raw)[`schema`]; has {
				imp.content[mediaType] = factory.Schema(``, schema.(headers.Raw))
			}
		}
	}
}

func (imp *responseImp) Resolve(openAPI headers.OpenAPI) headers.Schema {
	for mediaType, content := range imp.content {
		if res, ok := content.(headers.Resolvable); ok {
			imp.content[mediaType] = res.Resolve(openAPI)
		}
	}
	return nil
}

func (imp *responseImp) Code() string                            { return imp.code }
func (imp *responseImp) Description() string                     { return imp.description }
func (imp *responseImp) MediaTypes() []string                    { return imp.mediaTypes }
func (imp *responseImp) Content(mediaType string) headers.Schema { return imp.content[mediaType] }
