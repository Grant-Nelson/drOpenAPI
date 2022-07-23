package operationType

type Type string

const (
	Delete  Type = `delete`
	Head    Type = `head`
	Get     Type = `get`
	Options Type = `options`
	Patch   Type = `patch`
	Post    Type = `post`
	Put     Type = `put`
	Trace   Type = `trace`
)

func (op Type) Valid() bool {
	switch op {
	case Delete,
		Head,
		Get,
		Options,
		Patch,
		Post,
		Put,
		Trace:
		return true
	}
	return false
}

func All() []Type {
	return []Type{
		Delete,
		Head,
		Get,
		Options,
		Patch,
		Post,
		Put,
		Trace,
	}
}
