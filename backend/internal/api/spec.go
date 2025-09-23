package api

import (
	"path"

	"github.com/Mroxny/slamIt/internal/config"
	"github.com/getkin/kin-openapi/openapi3"
)

var SpecPath = path.Join(config.RoothPath, "./openapi.yaml")
var SpecUrl = "/swagger/openapi.yaml"

func LoadSpec() (*openapi3.T, error) {

	loader := openapi3.NewLoader()
	spec, err := loader.LoadFromFile(SpecPath)

	if err != nil {
		return nil, err
	}
	return spec, nil
}
