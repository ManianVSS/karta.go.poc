package xlang

import (
	"fmt"
)

func init() {
	stepDefMap["func"] = createFunctionDefinitionStep
}

type FunctionDefinition struct {
	BaseStep
	functionName string
}

func (functionDefinition *FunctionDefinition) InitalizeAndCheck() error {

	if name, ok := functionDefinition.parameters["name"]; ok {
		functionDefinition.functionName = name
	} else {
		return fmt.Errorf("name attribute missing")
	}

	return nil
}

func (functionDefinition *FunctionDefinition) Execute(scope *Scope) (any, error) {

	if err := functionDefinition.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	if _, ok := scope.functions[functionDefinition.functionName]; !ok {
		scope.functions[functionDefinition.functionName] = functionDefinition.steps
	} else {
		return nil, fmt.Errorf("function definition already present for name %s", functionDefinition.functionName)
	}
	return nil, nil
}

func createFunctionDefinitionStep(name string, parameters map[string]string, body string) (Step, error) {
	functionDefinition := &FunctionDefinition{}
	functionDefinition.name = name
	functionDefinition.parameters = parameters
	functionDefinition.body = body
	return functionDefinition, nil
}
