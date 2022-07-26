package schemaType

// Type is the variable type for an enumerator for indicating the type of a schema.
type Type string

const (
	// Array indicates the schema is an array type.
	Array Type = `array`

	// Boolean indicates the schema is a boolean.
	Boolean Type = `boolean`

	// Composite indicates the schema is an object made out
	// of oneOf, anyOf, or allOf given component schemas.
	Composite Type = `composite`

	// Enum indicates the schema is an object with enumerated types.
	Enum Type = `enum`

	// Integer indicates the schema is an integer.
	Integer Type = `integer`

	// Number indicates the schema is a floating point number.
	Number Type = `number`

	// Object indicates the schema is an object with properties.
	Object Type = `object`

	// String indicates the schema is a string.
	String Type = `string`
)

// All gets the list of all enumerator values in this enumerator.
func All() []Type {
	return []Type{
		Array,
		Boolean,
		Composite,
		Enum,
		Integer,
		Number,
		Object,
		String,
	}
}
