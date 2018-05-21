package configLoader

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// LoadControllerConfigYaml reads the specified YAML config file for Controller Mode
func LoadControllerConfigYaml(filePath string) (*ControllerConfig, error) {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var controllerConfig ControllerConfig

	err = yaml.Unmarshal(yamlFile, &controllerConfig)
	if err != nil {
		return nil, err
	}

	return &controllerConfig, nil
}
