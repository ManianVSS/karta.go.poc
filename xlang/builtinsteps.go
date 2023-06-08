package xlang

import (
	"fmt"
	"strings"
)

type Echo struct {
	BaseStep
	message string
}

func (echo *Echo) Init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {

	if _, err := echo.BaseStep.Init(parent, tag, attributes, text, steps...); err != nil {
		return echo, err
	}

	if message, ok := attributes["message"]; ok {
		echo.message = message
	} else if text != "" {
		echo.message = text
	} else {
		return nil, fmt.Errorf("message attribute missing")
	}
	return echo, nil
}

func (echo *Echo) Execute(scope Scope) (bool, error) {
	fmt.Println(echo.message)
	return true, nil
}

func createEchoStep(tag string, attributes map[string]string, text string) (Step, error) {
	echo := Echo{}
	return echo.Init(nil, tag, attributes, text)
}

type ReturnStep struct {
	BaseStep
	// returnValue any
}

func (returnStep *ReturnStep) Init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {
	// fmt.Println("Hit return step init")
	_, err := returnStep.BaseStep.Init(parent, tag, attributes, text, steps...)
	return returnStep, err
}

func (returnStep *ReturnStep) Execute(scope Scope) (bool, error) {
	// fmt.Println("Hit return step execute")
	return false, &MethodReturnError{}
}

func createReturnStep(tag string, attributes map[string]string, text string) (Step, error) {
	// fmt.Println("Hit return step create")
	returnStep := ReturnStep{}
	return returnStep.Init(nil, tag, attributes, text)
}

type FunctionDefinition struct {
	BaseStep
	name             string
	inputParameters  []string
	outputParameters []string
}

func (functionDefinition *FunctionDefinition) Init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {

	if _, err := functionDefinition.BaseStep.Init(parent, tag, attributes, text, steps...); err != nil {
		return functionDefinition, err
	}

	if name, ok := attributes["name"]; ok {
		functionDefinition.name = name
	} else {
		return nil, fmt.Errorf("name attribute missing")
	}

	if inputParameters, ok := attributes["inputParameters"]; ok {
		functionDefinition.inputParameters = strings.Split(inputParameters, ",")
	} else {
		functionDefinition.inputParameters = []string{}
	}

	if outputParameters, ok := attributes["outputParameters"]; ok {
		functionDefinition.outputParameters = strings.Split(outputParameters, ",")
	} else {
		functionDefinition.outputParameters = []string{}
	}
	return functionDefinition, nil
}

func (functionDefinition *FunctionDefinition) Execute(scope Scope) (bool, error) {
	fmt.Println("declaring func ", functionDefinition.name, "(", functionDefinition.inputParameters, ")", functionDefinition.outputParameters)
	if _, ok := scope.functions[functionDefinition.name]; !ok {
		scope.functions[functionDefinition.name] = functionDefinition
	} else {
		return false, fmt.Errorf("function definitions already present for name %s", functionDefinition.name)
	}
	return true, nil
}

func createFunctionDefinitionStep(tag string, attributes map[string]string, text string) (Step, error) {
	functionDefinition := FunctionDefinition{}
	return functionDefinition.Init(nil, tag, attributes, text)
}

type FunctionCall struct {
	BaseStep
	name string
}

func (functionCall *FunctionCall) Init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {

	if _, err := functionCall.BaseStep.Init(parent, tag, attributes, text, steps...); err != nil {
		return functionCall, err
	}

	if name, ok := attributes["name"]; ok {
		functionCall.name = name
	} else {
		return nil, fmt.Errorf("name attribute missing")
	}

	return functionCall, nil
}

func (functionCall *FunctionCall) Execute(scope Scope) (bool, error) {

	if functionDefinition, ok := scope.get_function(functionCall.name); ok {
		functionScope := Scope{}
		functionScope.variables = map[string]any{}
		functionScope.functions = map[string]*FunctionDefinition{}

		for _, inputParameter := range functionDefinition.inputParameters {
			if inputParameterValue, result := scope.get_variable(inputParameter); result {
				functionScope.variables[inputParameter] = inputParameterValue
			} else {
				return false, fmt.Errorf("input variable not found in known scope %s", inputParameter)
			}
		}
		fmt.Println("func call", functionCall.name, "(", functionScope.variables, ")")
		if result, err := RunSteps(functionScope, functionDefinition.NestedSteps...); err != nil {

			for _, outputParameter := range functionDefinition.outputParameters {

				if outputParameterValue, result := functionScope.get_variable(outputParameter); result {
					scope.variables[outputParameter] = outputParameterValue
				} else {
					return false, fmt.Errorf("output variable not found in function scope %s", outputParameter)
				}
			}
			return result, err
		} else {
			return result, err
		}

	} else {
		return false, fmt.Errorf("function definitions not present in known scope %s", functionCall.name)
	}
}

func createFunctionCallStep(tag string, attributes map[string]string, text string) (Step, error) {
	functionCall := FunctionCall{}
	return functionCall.Init(nil, tag, attributes, text)
}
