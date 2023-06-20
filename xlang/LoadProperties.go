package xlang

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func init() {
	stepDefMap["loadProperties"] = createLoadPropertiesStep
}

type LoadPropertiesStep struct {
	BaseStep
	fileName string
}

func (loadPropertiesStep *LoadPropertiesStep) InitalizeAndCheck() error {

	if file, ok := loadPropertiesStep.parameters["file"]; ok {
		loadPropertiesStep.fileName = file
	} else {
		return fmt.Errorf("file attribute missing")
	}
	return nil
}

func (loadPropertiesStep *LoadPropertiesStep) Execute(scope *Scope) (any, error) {

	if err := loadPropertiesStep.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	jsonFile, err := os.Open(loadPropertiesStep.fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	var result map[string]any
	json.Unmarshal([]byte(byteValue), &result)

	for key, value := range result {
		scope.variables[key] = value
	}
	return result, nil
}

func createLoadPropertiesStep(name string, parameters map[string]string, body string) (Step, error) {
	loadPropertiesStep := &LoadPropertiesStep{}
	loadPropertiesStep.name = name
	loadPropertiesStep.parameters = parameters
	loadPropertiesStep.body = body
	return loadPropertiesStep, nil
}
