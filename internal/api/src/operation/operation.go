package operation

import (
	"sort"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/operationType"
)

type operationImp struct {
	opType        operationType.Type
	operationId   string
	summary       string
	description   string
	responseCodes []string
	responses     map[string]api.Response
	tags          []string
}

func New(factory api.Factory, opType operationType.Type, data api.Raw) api.Operation {
	imp := &operationImp{opType: opType}
	imp.setInfo(data)
	imp.setResponses(factory, data)
	imp.setTags(data)
	return imp
}

func (imp *operationImp) setInfo(data api.Raw) {
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

func (imp *operationImp) setResponses(factory api.Factory, data api.Raw) {
	imp.responseCodes = []string{}
	imp.responses = map[string]api.Response{}
	if responses, has := data[`responses`]; has {
		for code, response := range responses.(api.Raw) {
			imp.responseCodes = append(imp.responseCodes, code)
			imp.responses[code] = factory.Response(code, response.(api.Raw))
		}
	}
	sort.Strings(imp.responseCodes)
}

func (imp *operationImp) setTags(data api.Raw) {
	imp.tags = []string{}
	if tags, has := data[`tags`]; has {
		for _, tag := range tags.([]interface{}) {
			imp.tags = append(imp.tags, tag.(string))
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
