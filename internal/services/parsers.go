package services

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type YamlParser struct{}

func (yp *YamlParser) ToJSON(asYaml []byte) (string, error) {
	var result map[string]interface{}
	err := yaml.Unmarshal(asYaml, &result)
	if err != nil {
		panic(err)
	}

	asJson, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	return string(asJson), nil
}

func (yp *YamlParser) FromJson(asJson string) (string, error) {
	return "", nil
}
