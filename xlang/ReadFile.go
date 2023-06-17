package xlang

import (
	"fmt"
	"os"
)

func init() {
	stepDefMap["readFile"] = createReadFileStep
}

type ReadFileStep struct {
	BaseStep
	fileName  string
	resultVar string
}

func (readFileStep *ReadFileStep) InitalizeAndCheck() error {

	if file, ok := readFileStep.attributes["file"]; ok {
		readFileStep.fileName = file
	} else {
		return fmt.Errorf("file attribute missing")
	}

	if resultVar, ok := readFileStep.attributes["resultVar"]; ok {
		readFileStep.resultVar = resultVar
	} else {
		readFileStep.resultVar = ""
	}
	return nil
}

func (readFileStep *ReadFileStep) Execute(scope *Scope) (any, error) {

	if err := readFileStep.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	byteValue, err := os.ReadFile(readFileStep.fileName)

	if err != nil {
		return nil, err
	}

	stringValue := string(byteValue)
	if readFileStep.resultVar != "" {
		scope.variables[readFileStep.resultVar] = string(stringValue)
	}

	return stringValue, nil
}

func createReadFileStep(tag string, attributes map[string]string, text string) (Step, error) {
	readFileStep := &ReadFileStep{}
	readFileStep.tag = tag
	readFileStep.attributes = attributes
	readFileStep.text = text
	return readFileStep, nil
}
