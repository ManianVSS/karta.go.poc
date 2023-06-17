package xlang

import (
	"fmt"
)

func init() {
	stepDefMap["func"] = createFunctionDefinitionStep
}

type FunctionDefinition struct {
	BaseStep
	name string
}

func (functionDefinition *FunctionDefinition) InitalizeAndCheck() error {

	if name, ok := functionDefinition.attributes["name"]; ok {
		functionDefinition.name = name
	} else {
		return fmt.Errorf("name attribute missing")
	}

	return nil
}

func (functionDefinition *FunctionDefinition) Execute(scope *Scope, basedir string) (any, error) {

	if err := functionDefinition.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	if _, ok := scope.functions[functionDefinition.name]; !ok {
		scope.functions[functionDefinition.name] = functionDefinition.nestedSteps
	} else {
		return nil, fmt.Errorf("function definition already present for name %s", functionDefinition.name)
	}
	return nil, nil
}

func createFunctionDefinitionStep(tag string, attributes map[string]string, text string) (Step, error) {
	functionDefinition := &FunctionDefinition{}
	functionDefinition.tag = tag
	functionDefinition.attributes = attributes
	functionDefinition.text = text
	return functionDefinition, nil
}
