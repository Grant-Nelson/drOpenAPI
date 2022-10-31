package api

import (
	"errors"
)

// Raw is the typical data type for an object as returned
// by an unmarshaler method such as JSON or YAML.
type Raw map[string]any

// getStep reads the value of the given key, if it exists and
// is the given type, otherwise the default is returned.
func getStep[T any](data any, key string) (T, bool) {
	var defaultValue T

	raw, ok := data.(Raw)
	if !ok {
		return defaultValue, false
	}

	if value, has := raw[key]; has {
		if cast, ok := value.(T); ok {
			return cast, true
		}
	}

	return defaultValue, false
}

// Get reads the value at the given path of objects, if it exists and
// is the given type, otherwise the default is returned.
func Get[T any](data any, keys ...string) (T, bool) {
	count := len(keys)
	if count <= 0 {
		panic(errors.New(`must provide at least one key`))
	}

	prior := data
	var ok bool
	for i := 0; i < count-1; i++ {
		prior, ok = getStep[any](prior, keys[i])
		if !ok {
			var defaultValue T
			return defaultValue, false
		}
	}
	return getStep[T](prior, keys[count-1])
}
