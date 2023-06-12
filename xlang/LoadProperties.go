package xlang

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
	stepDefMap["loadProperties"] = createLoadPropertiesStep
}

type LoadPropertiesStep struct {
	BaseStep
	fileName string
}

func (loadPropertiesStep *LoadPropertiesStep) Initalize() error {

	if file, ok := loadPropertiesStep.attributes["file"]; ok {
		loadPropertiesStep.fileName = file
	} else {
		return fmt.Errorf("file attribute missing")
	}
	return nil
}

func (loadPropertiesStep *LoadPropertiesStep) Execute(scope *Scope, basedir string) (any, error) {

	jsonFile, err := os.Open(loadPropertiesStep.fileName)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

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

func createLoadPropertiesStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	loadPropertiesStep := &LoadPropertiesStep{}
	loadPropertiesStep.tag = tag
	loadPropertiesStep.attributes = attributes
	loadPropertiesStep.text = text
	return loadPropertiesStep, nil
}