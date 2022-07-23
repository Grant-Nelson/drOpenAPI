package schemaType

type Type string

const (
	Array     Type = `array`
	Boolean   Type = `boolean`
	Composite Type = `composite`
	Enum      Type = `enum`
	Integer   Type = `integer`
	Number    Type = `number`
	Object    Type = `object`
	String    Type = `string`
)

func (op Type) Valid() bool {
	switch op {
	case Array,
		Boolean,
		Composite,
		Enum,
		Integer,
		Number,
		Object,
		String:
		return true
	}
	return false
}

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
