package response

import "github.com/grant-nelson/DrOpenAPI/internal/api"

type responseImp struct {
	code        string
	description string
	mediaTypes  []string
	content     map[string]api.Schema
}

var _ api.Resolvable = (*responseImp)(nil)

func New(factory api.Factory, code string, data api.Raw) api.Response {
	imp := &responseImp{code: code}
	imp.setInfo(data)
	imp.setContent(factory, data)
	return imp
}

func (imp *responseImp) setInfo(data api.Raw) {
	if description, has := data[`description`]; has {
		imp.description = description.(string)
	}
}

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
