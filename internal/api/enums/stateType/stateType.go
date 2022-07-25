package stateType

type Type string

const (
	Deprecated Type = `deprecated`
	Nullable   Type = `nullable`
	ReadOnly   Type = `readOnly`
	WriteOnly  Type = `writeOnly`
)

func (op Type) Valid() bool {
	switch op {
	case Deprecated,
		Nullable,
		ReadOnly,
		WriteOnly:
		return true
	}
	return false
}

func All() []Type {
	return []Type{
		Deprecated,
		Nullable,
		ReadOnly,
		WriteOnly,
	}
}
