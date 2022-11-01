# Dr OpenAPI

![YAML & JSON to Markdown with Mermaid](./image.png)

Need a MD?

Dr OpenAPI is a simple markdown generator for an OpenAPI bundle.

## What does Dr OpenAPI do?

Dr OpenAPI reads the JSON or YAML bundle file for an OpenAPI v3 definition.
Then it writes a markdown file listing all the paths in the definition.
For each path the result model is drawn using a mermaid class diagram.
The paths must have a `application/json` media type to be shown.

## Setup

- Setup [Go v1.19 or later](https://go.dev/dl/)
- Clone with `go get github.com/Grant-Nelson/DrOpenAPI`

## Run from this folder

- Generate MD with `go run main.go -i <input> -o <output>`
- For help run `go run main.go -h`

## Installing to use anywhere

- Install with `go install`
- Generate MD with `DrOpenAPI -i <input> -o <output>`
- For help run `DrOpenAPI -h`

## Customize Markdown

- To change the title on the MD file, add a `-t <title>` argument.
  For example `-t "Title for MD file"`
