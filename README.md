# Dr OpenAPI

![YAML & JSON to Markdown with Mermaid](./image.png)

Need an MD?

Dr OpenAPI is a simple markdown generator for an OpenAPI bundle.

## What does Dr OpenAPI do?

Dr OpenAPI reads the JSON or YAML bundle file for an OpenAPI v3 definition.
Then it writes a markdown file listing all the paths in the definition.
For each path the result model is drawn using a mermaid class diagram.
The paths must have a `200` response code with `application/json` media type.

## Setup

- Setup [Go v1.17 or later](https://go.dev/dl/)

## Run

- Generate MD with `go run main.go -i <input> -o <output>`
- For help run `go run main.go -h`
- To change the title on the MD file, add the `-t` argument.
  For example `-t "Title for MD file"`
