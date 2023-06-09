package xlang

import (
	"fmt"
	"strings"
)

type FunctionCall struct {
	BaseStep
	name             string
	inputParameters  []string
	outputParameters []string
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

	if inputParameters, ok := functionCall.attributes["inputParameters"]; ok {
		functionCall.inputParameters = strings.Split(inputParameters, ",")
	} else {
		functionCall.inputParameters = []string{}
	}

	if outputParameters, ok := functionCall.attributes["outputParameters"]; ok {
		functionCall.outputParameters = strings.Split(outputParameters, ",")
	} else {
		functionCall.outputParameters = []string{}
	}

	return nil
}

func (functionCall *FunctionCall) Execute(scope *Scope) error {

	if functionSteps, ok := scope.get_function(functionCall.name); ok {
		functionScope := Scope{}
		functionScope.variables = map[string]any{}
		functionScope.functions = map[string][]Step{}

		for _, inputParameter := range functionCall.inputParameters {
			if inputParameterValue, result := scope.get_variable(inputParameter); result {
				functionScope.variables[inputParameter] = inputParameterValue
			} else {
				return fmt.Errorf("input variable not found in known scope %s", inputParameter)
			}
		}

		if err := RunSteps(&functionScope, functionSteps...); err == nil {

			for _, outputParameter := range functionCall.outputParameters {

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

func createFunctionCallStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	functionCall := &FunctionCall{}
	return functionCall, functionCall.Init(tag, attributes, text)
}
