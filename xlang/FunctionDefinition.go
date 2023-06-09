package xlang

import (
	"fmt"
)

type FunctionDefinition struct {
	BaseStep
	name string
}

func (functionDefinition *FunctionDefinition) Initalize() error {

	if name, ok := functionDefinition.attributes["name"]; ok {
		functionDefinition.name = name
	} else {
		return fmt.Errorf("name attribute missing")
	}

	return nil
}

func (functionDefinition *FunctionDefinition) Execute(scope *Scope) error {

	if _, ok := scope.functions[functionDefinition.name]; !ok {
		scope.functions[functionDefinition.name] = functionDefinition.nestedSteps
	} else {
		return fmt.Errorf("function definition already present for name %s", functionDefinition.name)
	}
	return nil
}

func createFunctionDefinitionStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	functionDefinition := &FunctionDefinition{}
	functionDefinition.tag = tag
	functionDefinition.attributes = attributes
	functionDefinition.text = text
	return functionDefinition, functionDefinition.Initalize()
}
