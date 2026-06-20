package openapidocs

import _ "embed"

//go:embed openapi.yaml
var OpenAPIYAML []byte

//go:embed openapi.json
var OpenAPIJSON []byte
