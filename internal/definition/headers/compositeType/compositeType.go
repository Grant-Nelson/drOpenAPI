package compositeType

type Type string

const (
	AnyOf Type = `anyOf`
	OneOf Type = `oneOf`
	AllOf Type = `allOf`
)

func (op Type) Valid() bool {
	switch op {
	case AnyOf,
		OneOf,
		AllOf:
		return true
	}
	return false
}

func All() []Type {
	return []Type{
		AnyOf,
		OneOf,
		AllOf,
	}
}
