package referenceSchema

import (
	"fmt"
	"strings"

	"github.com/Grant-Nelson/drOpenAPI/internal/api"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/enums/schemaType"
	"github.com/Grant-Nelson/drOpenAPI/internal/api/enums/stateType"
)

const refPrefix = `#/components/schemas/`

// referenceImp is the implementation of the Schema and Resolvable interface
// but represents a reference schema that needs to be resolved.
// All of these should have been removed prior to the OpenAPI creation being finished.
type referenceImp struct {
	title      string
	ref        string
	resolving  bool
	resolution api.Schema
}

var _ api.Resolvable = (*referenceImp)(nil)

// New creates a new reference to a Schema.
func New(title string, data api.Raw) api.Schema {
	imp := &referenceImp{title: title}
	if ref, has := api.Get[string](data, `$ref`); has {
		imp.ref = ref
	}
	return imp
}

func throwError(msg string, args ...any) {
	panic(fmt.Errorf(msg, args...))
}

// Resolve performs a lookup using the reference and
// returns the found schema in the given OpenAPI.
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

// badTouch is used to keep this reference from being used like a Schema
// even though it has to implement a Schema interface.
func (imp *referenceImp) badTouch() {
	throwError(`must resolve reference, %q, before trying to read values from schema`, imp.ref)
}

func (imp *referenceImp) Title() string                   { return imp.title }
func (imp *referenceImp) Description() string             { imp.badTouch(); return `` }
func (imp *referenceImp) Type() schemaType.Type           { return schemaType.Type(`ref`) }
func (imp *referenceImp) Format() string                  { imp.badTouch(); return `` }
func (imp *referenceImp) Default() string                 { imp.badTouch(); return `` }
func (imp *referenceImp) State(state stateType.Type) bool { imp.badTouch(); return false }
