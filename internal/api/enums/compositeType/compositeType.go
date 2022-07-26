package compositeType

// Type is the variable type for an enumerator for the different
// ways a composite schema can be defined to use its component schemas.
type Type string

const (
	// AnyOf indicates that the composite object may define one or more of its components.
	AnyOf Type = `anyOf`

	// OneOf indicates that the composite object must define one and only one of its components.
	OneOf Type = `oneOf`

	// AllOf indicates that the composite object must define all of its components.
	AllOf Type = `allOf`
)

// All gets the list of all enumerator values in this enumerator.
func All() []Type {
	return []Type{
		AnyOf,
		OneOf,
		AllOf,
	}
}
