package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Grant-Nelson/drOpenAPI/internal/reader"
	"github.com/Grant-Nelson/drOpenAPI/internal/writer"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(`Error:`, r)
			os.Exit(1)
		}
	}()

	var input, output, fileType, title string
	flag.StringVar(&input, `i`, ``,
		`Specify the input json or yaml bundle file to read.`)
	flag.StringVar(&output, `o`, ``,
		`Specify the output file to write to. `+
			`If not specified then the md is written to standard out.`)
	flag.StringVar(&fileType, `f`, ``,
		`Specify json or yaml input file type. `+
			`If not specified then the input file's extension is used to determine file type.`)
	flag.StringVar(&title, `t`, `Open API Models`,
		`Specify the title to write to the resulting markdown file.`)
	flag.Parse()

	definition := reader.Read(input, fileType)
	writer.Write(output, title, definition)
	os.Exit(0)
}
