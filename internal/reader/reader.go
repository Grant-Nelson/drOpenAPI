package reader

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/grant-nelson/DrOpenAPI/internal/definition/headers"
	"github.com/grant-nelson/DrOpenAPI/internal/definition/source/factory"
)

type unmarshalHandle func(in []byte, out interface{}) error

func ReadYaml(path string) headers.OpenAPI {
	return read(path, yaml.Unmarshal)
}

func ReadJson(path string) headers.OpenAPI {
	return read(path, json.Unmarshal)
}

func read(path string, unmarshal unmarshalHandle) headers.OpenAPI {
	return factory.New().OpenAPI(readRaw(path, unmarshal))
}

func readRaw(path string, unmarshal unmarshalHandle) headers.Raw {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	data := headers.Raw{}
	err = unmarshal(file, &data)
	if err != nil {
		panic(err)
	}

	return data
}
