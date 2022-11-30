package digest

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func FromFile(path string) *Digest {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return FromData(file)
}

func FromData(data []byte) *Digest {

	// TODO: trim data
	d := &Digest{}
	if err := yaml.Unmarshal(data, d); err != nil {
		panic(err)
	}
	return d
}

type Digest struct {

}
