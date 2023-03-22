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
- Clone with `go get github.com/Grant-Nelson/drOpenAPI`

## Run from this folder

- Run with `go run main.go <config>`

## Installing to use anywhere

- Install with `go install`
- Run with `drOpenAPI <config>`
drOpenAPI
## Config filesdrOpenAPI

When running Dr OpenAPI a config YAML or JSON file may be passed in.
If no config file is passed in as an argument,
then `./drOpenAPI.yaml` or `./drOpenAPI.json` is used.

The config file may have the following values:

### Input

The config file may have an `input` file path defined.
The input file can be a YAML or JSON file.
If no input file is given then `./bundle.yaml` or `./bundle.json` is used.

```Yaml
input: ./your_bundle_file.yaml
```

### Output

The config file may have zero or more `outputs`.
If no outputs are given then the resulting file will have all of
paths and be written to the input path with the extension changed to `.md`.

```Yaml
outputs:
  - path: ./created_markdown_file.md
    title: Custom title to put at top of file
    paths:
      - path: ./
        ops:  get
        code: 200
```
