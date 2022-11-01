package operation

import (
	"fmt"
	"sort"

	"github.com/Grant-Nelson/DrOpenAPI/internal/api"
	"github.com/Grant-Nelson/DrOpenAPI/internal/api/enums/operationType"
)

// operationImp is the implementation of the Operation interface.
type operationImp struct {
	opType        operationType.Type
	operationId   string
	summary       string
	description   string
	requestBody   api.Schema
	responseCodes []string
	responses     map[string]api.Response
	tags          []string
}

// New creates a new Operation instance.
func New(factory api.Factory, opType operationType.Type, data api.Raw) api.Operation {
	imp := &operationImp{opType: opType}
	imp.setInfo(factory, data)
	imp.setRequestBody(factory, data)
	imp.setResponses(factory, data)
	imp.setTags(data)
	return imp
}

// setInfo reads all the basic information from the given data,
// then sets them to this Operation implementation.
func (imp *operationImp) setInfo(factory api.Factory, data api.Raw) {
	if operationId, has := api.Get[string](data, `operationId`); has {
		imp.operationId = operationId
	}

	if len(imp.operationId) == 0 {
		imp.operationId = factory.UniqueName()
	}

	if summary, has := api.Get[string](data, `summary`); has {
		imp.summary = summary
	}

	if description, has := api.Get[string](data, `description`); has {
		imp.description = description
	}
}

// setRequestBody reads the request body object from the given data,
// then sets it to this Operation implementation.
func (imp *operationImp) setRequestBody(factory api.Factory, data api.Raw) {
	if schema, has := api.Get[api.Raw](data, `requestBody`, `content`, `application/json`, `schema`); has {
		imp.requestBody = factory.Schema(`RequestBody`, schema)
	}
}

// setResponses reads all the response object from the given data,
// then sets them to this Operation implementation.
func (imp *operationImp) setResponses(factory api.Factory, data api.Raw) {
	imp.responseCodes = []string{}
	imp.responses = map[string]api.Response{}
	if resp, ok := api.Get[api.Raw](data, `responses`); ok {
		for code, response := range resp {
			imp.responses[code] = factory.Response(code, response.(api.Raw))
			imp.responseCodes = append(imp.responseCodes, code)
		}
	} else if resp, ok := api.Get[map[any]any](data, `responses`); ok {
		for code, response := range resp {
			codeStr := fmt.Sprint(code)
			imp.responses[codeStr] = factory.Response(codeStr, response.(api.Raw))
			imp.responseCodes = append(imp.responseCodes, codeStr)
		}
	}
	sort.Strings(imp.responseCodes)
}

// setTags reads all the tags from the given data,
// then sets them to this OpenAPI implementation.
func (imp *operationImp) setTags(data api.Raw) {
	imp.tags = []string{}
	if tags, has := api.Get[[]any](data, `tags`); has {
		for _, tag := range tags {
			tagStr := fmt.Sprint(tag)
			imp.tags = append(imp.tags, tagStr)
		}
	}
	sort.Strings(imp.tags)
}

// Resolve runs resolve on any of the contained schema which is also resolvable.
// This will return nil since this isn't a schema to replace.
func (imp *operationImp) Resolve(openAPI api.OpenAPI) api.Schema {
	if res, ok := imp.requestBody.(api.Resolvable); ok {
		imp.requestBody = res.Resolve(openAPI)
	}

	for _, resp := range imp.responses {
		if res, ok := resp.(api.Resolvable); ok {
			res.Resolve(openAPI)
		}
	}
	return nil
}

func (imp *operationImp) OpType() operationType.Type        { return imp.opType }
func (imp *operationImp) Summary() string                   { return imp.summary }
func (imp *operationImp) Description() string               { return imp.description }
func (imp *operationImp) OperationId() string               { return imp.operationId }
func (imp *operationImp) RequestBody() api.Schema           { return imp.requestBody }
func (imp *operationImp) ResponseCodes() []string           { return imp.responseCodes }
func (imp *operationImp) Response(code string) api.Response { return imp.responses[code] }
func (imp *operationImp) Tags() []string                    { return imp.tags }
