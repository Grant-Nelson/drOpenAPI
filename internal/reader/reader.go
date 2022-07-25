package reader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/grant-nelson/DrOpenAPI/internal/api"
	"github.com/grant-nelson/DrOpenAPI/internal/api/src/factory"
)

type unmarshalHandle func(in []byte, out interface{}) error

func Read(path, fileType string) api.OpenAPI {
	if len(path) == 0 {
		panic(errors.New(`must provide a non-empty input file path`))
	}

	unmarshal := getUnmarshalHandler(path, fileType)
	raw := readRaw(path, unmarshal)
	return factory.New().OpenAPI(raw)
}

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
