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
	responseCodes []string
	responses     map[string]api.Response
	tags          []string
}

// New creates a new Operation instance.
func New(factory api.Factory, opType operationType.Type, data api.Raw) api.Operation {
	imp := &operationImp{opType: opType}
	imp.setInfo(factory, data)
	imp.setResponses(factory, data)
	imp.setTags(data)
	return imp
}

// setInfo reads all the basic information from the given data,
// then sets them to this Operation implementation.
func (imp *operationImp) setInfo(factory api.Factory, data api.Raw) {
	if operationId, has := data[`operationId`]; has {
		imp.operationId = fmt.Sprint(operationId)
	}

	if len(imp.operationId) == 0 {
		imp.operationId = factory.UniqueName()
	}

	if summary, has := data[`summary`]; has {
		imp.summary = fmt.Sprint(summary)
	}

	if description, has := data[`description`]; has {
		imp.description = fmt.Sprint(description)
	}
}

// setResponses reads all the response object from the given data,
// then sets them to this Operation implementation.
func (imp *operationImp) setResponses(factory api.Factory, data api.Raw) {
	imp.responseCodes = []string{}
	imp.responses = map[string]api.Response{}
	if responses, has := data[`responses`]; has {
		if resp, ok := responses.(api.Raw); ok {
			for code, response := range resp {
				imp.responses[code] = factory.Response(code, response.(api.Raw))
				imp.responseCodes = append(imp.responseCodes, code)
			}
		} else if resp, ok := responses.(map[interface{}]interface{}); ok {
			for code, response := range resp {
				codeStr := fmt.Sprint(code)
				imp.responses[codeStr] = factory.Response(codeStr, response.(api.Raw))
				imp.responseCodes = append(imp.responseCodes, codeStr)
			}
		}
	}
	sort.Strings(imp.responseCodes)
}

// setTags reads all the tags from the given data,
// then sets them to this OpenAPI implementation.
func (imp *operationImp) setTags(data api.Raw) {
	imp.tags = []string{}
	if tags, has := data[`tags`]; has {
		for _, tag := range tags.([]interface{}) {
			tagStr := fmt.Sprint(tag)
			imp.tags = append(imp.tags, tagStr)
		}
	}
	sort.Strings(imp.tags)
}

func (imp *operationImp) OpType() operationType.Type        { return imp.opType }
func (imp *operationImp) Summary() string                   { return imp.summary }
func (imp *operationImp) Description() string               { return imp.description }
func (imp *operationImp) OperationId() string               { return imp.operationId }
func (imp *operationImp) ResponseCodes() []string           { return imp.responseCodes }
func (imp *operationImp) Response(code string) api.Response { return imp.responses[code] }
func (imp *operationImp) Tags() []string                    { return imp.tags }
