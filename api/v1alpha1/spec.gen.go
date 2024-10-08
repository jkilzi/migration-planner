// Package v1alpha1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package v1alpha1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AP//86gujDweAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
