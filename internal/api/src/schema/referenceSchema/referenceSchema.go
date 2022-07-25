package referenceSchema

import (
	"fmt"
	"strings"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/schemaType"
	"github.com/grant-nelson/DrOpenAPI/internal/api/enums/stateType"
)

const refPrefix = `#/components/schemas/`

func throwError(msg string, args ...interface{}) {
	panic(fmt.Errorf(msg, args...))
}

type referenceImp struct {
	title      string
	ref        string
	resolving  bool
	resolution api.Schema
}

var _ api.Resolvable = (*referenceImp)(nil)

func New(title string, data api.Raw) api.Schema {
	imp := &referenceImp{title: title}
	if ref, has := data[`$ref`]; has {
		imp.ref = ref.(string)
	}
	return imp
}

func (imp *referenceImp) Resolve(openAPI api.OpenAPI) api.Schema {
	if imp.resolution != nil {
		return imp.resolution
	}

	if imp.resolving {
		throwError(`loop detected while resolving schema references at reference for %q`, imp.ref)
	}
	imp.resolving = true

	if !strings.HasPrefix(imp.ref, refPrefix) {
		throwError(`expected reference to have the prefix %q but it was %q`, refPrefix, imp.ref)
	}

	title := imp.ref[len(refPrefix):]
	target := openAPI.Schema(title)
	if target == nil {
		throwError(`failed to find schema called %q in components`, title)
	}

	if res, ok := target.(api.Resolvable); ok {
		target = res.Resolve(openAPI)
	}

	imp.resolution = target
	imp.resolving = false
	return target
}

func (imp *referenceImp) badTouch() {
	throwError(`must resolve reference, %q, before trying to read values from schema`, imp.ref)
}

func (imp *referenceImp) Title() string                   { return imp.title }
func (imp *referenceImp) Description() string             { imp.badTouch(); return `` }
func (imp *referenceImp) Type() schemaType.Type           { return schemaType.Type(`ref`) }
func (imp *referenceImp) Format() string                  { imp.badTouch(); return `` }
func (imp *referenceImp) Default() string                 { imp.badTouch(); return `` }
func (imp *referenceImp) State(state stateType.Type) bool { imp.badTouch(); return false }
