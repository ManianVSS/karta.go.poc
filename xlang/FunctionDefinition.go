package xlang

import (
	"fmt"
)

type FunctionDefinition struct {
	BaseStep
	name string
}

func (functionDefinition *FunctionDefinition) Init(tag string, attributes map[string]string, text string) error {

	if err := functionDefinition.BaseStep.Init(tag, attributes, text); err != nil {
		return err
	}

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
		return fmt.Errorf("function definitions already present for name %s", functionDefinition.name)
	}
	return nil
}

func createFunctionDefinitionStep(tag string, attributes map[string]string, text string) (Step, error) {
	functionDefinition := &FunctionDefinition{}
	return functionDefinition, functionDefinition.Init(tag, attributes, text)
}
