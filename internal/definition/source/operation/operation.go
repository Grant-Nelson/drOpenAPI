package operation

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers/operationType"
)

type operationImp struct {
	opType        operationType.Type
	operationId   string
	summary       string
	description   string
	responseCodes []string
	responses     map[string]headers.Response
	tags          []string
}

func New(factory headers.Factory, opType operationType.Type, data headers.Raw) headers.Operation {
	imp := &operationImp{opType: opType}
	imp.setInfo(data)
	imp.setResponses(factory, data)
	imp.setTags(data)
	return imp
}

func (imp *operationImp) setInfo(data headers.Raw) {
	if operationId, has := data[`operationId`]; has {
		imp.operationId = operationId.(string)
	}
	if summary, has := data[`summary`]; has {
		imp.summary = summary.(string)
	}
	if description, has := data[`description`]; has {
		imp.description = description.(string)
	}
}

func (imp *operationImp) setResponses(factory headers.Factory, data headers.Raw) {
	imp.responseCodes = []string{}
	imp.responses = map[string]headers.Response{}
	if responses, has := data[`responses`]; has {
		for code, response := range responses.(headers.Raw) {
			imp.responseCodes = append(imp.responseCodes, code)
			imp.responses[code] = factory.Response(code, response.(headers.Raw))
		}
	}
	sort.Strings(imp.responseCodes)
}

func (imp *operationImp) setTags(data headers.Raw) {
	imp.tags = []string{}
	if tags, has := data[`tags`]; has {
		for _, tag := range tags.([]interface{}) {
			imp.tags = append(imp.tags, tag.(string))
		}
	}
	sort.Strings(imp.tags)
}

func (imp *operationImp) OpType() operationType.Type            { return imp.opType }
func (imp *operationImp) Summary() string                       { return imp.summary }
func (imp *operationImp) Description() string                   { return imp.description }
func (imp *operationImp) OperationId() string                   { return imp.operationId }
func (imp *operationImp) ResponseCodes() []string               { return imp.responseCodes }
func (imp *operationImp) Response(code string) headers.Response { return imp.responses[code] }
func (imp *operationImp) Tags() []string                        { return imp.tags }
