package digest

import (
	"os"
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

// FromFile loads a YAML or JSON object from the file at the given path.
func FromFile(path string) *Object {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return FromData(file)
}

// FromData loads a YAML or JSON object from the given bytes.
func FromData(data []byte) *Object {
	return FromString(string(data))
}

// FromString loads a YAML or JSON object from the given string.
func FromString(s string) *Object {
	data := []byte(strings.TrimSpace(s))
	r := &Object{}
	if err := yaml.Unmarshal(data, r); err != nil {
		panic(err)
	}
	return r
}

// getStep reads the value of the given key, if it exists and
// is the given type, otherwise the default is returned with false.
func getStep[T any](data any, key string) (T, bool) {
	var defaultValue T

	obj, ok := data.(Object)
	if !ok {
		return defaultValue, false
	}

	if value, has := obj[key]; has {
		if cast, ok := value.(T); ok {
			return cast, true
		}
	}

	return defaultValue, false
}

// Get reads the value at the given path of objects, if it exists and
// is the given type, otherwise the default is returned with false.
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
