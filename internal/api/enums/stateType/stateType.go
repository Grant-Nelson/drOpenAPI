package stateType

// Type is the variable type for an enumerator for looking up
// the boolean states from a schema.
type Type string

const (
	// This is a schema state indicating the schema has been deprecated.
	Deprecated Type = `deprecated`

	// Nullable is a schema state indicating the schema can be nulled.
	Nullable Type = `nullable`

	// ReadOnly is a schema state which can only be read from.
	ReadOnly Type = `readOnly`

	// WriteOnly is a scheme state which can only be written.
	WriteOnly Type = `writeOnly`
)

// All gets the list of all enumerator values in this enumerator.
func All() []Type {
	return []Type{
		Deprecated,
		Nullable,
		ReadOnly,
		WriteOnly,
	}
}
