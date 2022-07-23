package main

import (
	"os"

	"github.com/grant-nelson/DrOpenAPI/internal/reader"
	"github.com/grant-nelson/DrOpenAPI/internal/writer"
)

const (
	path   = `./input.yaml`
	output = `./output.md`
)

func main() {
	definition := reader.ReadYaml(path)
	writer.Write(output, definition)
	os.Exit(0)
}
