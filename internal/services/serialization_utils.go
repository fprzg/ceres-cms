package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

type YamlSer struct{}

func (ser *YamlSer) FromJson(asJson string) (string, error) {
	panic(fmt.Errorf("todo(Farid): Implement AsYaml.FromJson()"))
}

func (ser *YamlSer) FromRequest(r *http.Request) (string, error) {
	panic(fmt.Errorf("todo(Farid): Implement AsYaml.FromRequest()"))
}

type JsonSer struct{}

func (ser *JsonSer) FromYaml(asYaml []byte) ([]byte, error) {
	var content map[string]interface{}

	err := yaml.Unmarshal(asYaml, &content)
	if err != nil {
		panic(err)
	}

	asJson, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	return asJson, nil
}

func (ser *JsonSer) FromRequest(r *http.Request) ([]byte, error) {
	var content map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		panic(err)
	}

	asJson, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		panic(err)
	}

	return asJson, nil
}
