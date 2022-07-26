package reader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/Grant-Nelson/DrOpenAPI/internal/api"
	"github.com/Grant-Nelson/DrOpenAPI/internal/api/src/factory"
)

// unmarshalHandle is a function signature for a tool to parse
// the given file's data into the object representation of that data.
type unmarshalHandle func(in []byte, out interface{}) error

// Read will open the given file and convert it into an OpenAPI object.
// The given fileType overrides the file path extension, when not empty,
// to select how the file's data is stored (e.g. `yaml`, `json`).
// This will panic on error.
func Read(path, fileType string) api.OpenAPI {
	if len(path) == 0 {
		panic(errors.New(`must provide a non-empty input file path`))
	}

	unmarshal := getUnmarshalHandler(path, fileType)
	raw := readRaw(path, unmarshal)
	return factory.New().OpenAPI(raw)
}

// getUnmarshalHandler determines which unmarshal method to use to parse
// the files data given the path and fileType override.
func getUnmarshalHandler(path, fileType string) unmarshalHandle {
	if len(fileType) == 0 {
		ext := filepath.Ext(fileType)
		if len(ext) == 0 {
			fileType = `yaml`
		} else {
			fileType = ext[1:]
		}
	}

	fileType = strings.ToLower(fileType)
	switch fileType {
	case `json`:
		return json.Unmarshal
	case `yaml`:
		return yaml.Unmarshal
	default:
		panic(fmt.Errorf(`unexpected file type: %q`, fileType))
	}
}

// readRaw reads the raw object representation from the file at the given path
// using the given unmarshal method to parse the file's data.
func readRaw(path string, unmarshal unmarshalHandle) api.Raw {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	data := api.Raw{}
	err = unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	return data
}
