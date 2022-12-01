package digest

// Object is the typical data type for an object as returned
// by an unmarshaler method such as JSON or YAML.
type Object map[string]any

// Object reads an object at the given path of objects,
// if it exists, otherwise the default is returned with false.
func (d Object) Object(keys ...string) (Object, bool) {
	return Get[Object](d, keys...)
}

func (d Object) Array(keys ...string) (Array, bool) {
	return Get[Array](d, keys...)
}

func (d Object) String(keys ...string) (string, bool) {
	return Get[string](d, keys...)
}

