package sdkgo

// not-generated

import (
	_ "embed"
)

//go:embed openapi.json
var openapi []byte

func OpenAPISpec() []byte {
	return openapi
}
