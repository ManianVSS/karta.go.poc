package xlang

import (
	"fmt"
	"strings"
)

type Echo struct {
	BaseStep
	message string
}

func (echo *Echo) Init(tag string, attributes map[string]string, text string) error {

	if err := echo.BaseStep.Init(tag, attributes, text); err != nil {
		return err
	}

	if message, ok := echo.attributes["message"]; ok {
		echo.message = message
	} else if echo.text != "" {
		echo.message = echo.text
	} else {
		return fmt.Errorf("message attribute missing")
	}
	return nil
}

func (echo *Echo) Execute(scope Scope) error {
	_, err := fmt.Println(echo.message)
	return err
}

func createEchoStep(tag string, attributes map[string]string, text string) (Step, error) {
	echo := &Echo{}
	return echo, echo.Init(tag, attributes, text)
}

type ReturnStep struct {
	BaseStep
}

func (returnStep *ReturnStep) Init(tag string, attributes map[string]string, text string) error {
	return returnStep.BaseStep.Init(tag, attributes, text)
}

func (returnStep *ReturnStep) Execute(scope Scope) error {
	return &MethodReturnError{}
}

func createReturnStep(tag string, attributes map[string]string, text string) (Step, error) {
	returnStep := &ReturnStep{}
	return returnStep, returnStep.Init(tag, attributes, text)
}

type FunctionDefinition struct {
	BaseStep
	name             string
	inputParameters  []string
	outputParameters []string
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

	if inputParameters, ok := functionDefinition.attributes["inputParameters"]; ok {
		functionDefinition.inputParameters = strings.Split(inputParameters, ",")
	} else {
		functionDefinition.inputParameters = []string{}
	}

	if outputParameters, ok := functionDefinition.attributes["outputParameters"]; ok {
		functionDefinition.outputParameters = strings.Split(outputParameters, ",")
	} else {
		functionDefinition.outputParameters = []string{}
	}
	return nil
}

func (functionDefinition *FunctionDefinition) Execute(scope Scope) error {

	if _, ok := scope.functions[functionDefinition.name]; !ok {
		scope.functions[functionDefinition.name] = functionDefinition
	} else {
		return fmt.Errorf("function definitions already present for name %s", functionDefinition.name)
	}
	return nil
}

func createFunctionDefinitionStep(tag string, attributes map[string]string, text string) (Step, error) {
	functionDefinition := &FunctionDefinition{}
	return functionDefinition, functionDefinition.Init(tag, attributes, text)
}

type FunctionCall struct {
	BaseStep
	name string
}

func (functionCall *FunctionCall) Init(tag string, attributes map[string]string, text string) error {

	if err := functionCall.BaseStep.Init(tag, attributes, text); err != nil {
		return err
	}

	if name, ok := functionCall.attributes["name"]; ok {
		functionCall.name = name
	} else {
		return fmt.Errorf("name attribute missing")
	}

	return nil
}

func (functionCall *FunctionCall) Execute(scope Scope) error {

	if functionDefinition, ok := scope.get_function(functionCall.name); ok {
		functionScope := Scope{}
		functionScope.variables = map[string]any{}
		functionScope.functions = map[string]*FunctionDefinition{}

		for _, inputParameter := range functionDefinition.inputParameters {
			if inputParameterValue, result := scope.get_variable(inputParameter); result {
				functionScope.variables[inputParameter] = inputParameterValue
			} else {
				return fmt.Errorf("input variable not found in known scope %s", inputParameter)
			}
		}

		if err := RunSteps(functionScope, functionDefinition.nestedSteps...); err != nil {

			for _, outputParameter := range functionDefinition.outputParameters {

				if outputParameterValue, result := functionScope.get_variable(outputParameter); result {
					scope.variables[outputParameter] = outputParameterValue
				} else {
					return fmt.Errorf("output variable not found in function scope %s", outputParameter)
				}
			}
			return err
		} else {
			return err
		}

	} else {
		return fmt.Errorf("function definitions not present in known scope %s", functionCall.name)
	}
}

func createFunctionCallStep(tag string, attributes map[string]string, text string) (Step, error) {
	functionCall := &FunctionCall{}
	return functionCall, functionCall.Init(tag, attributes, text)
}
