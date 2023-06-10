package xlang

import (
	"fmt"
	"strings"
)

func init() {
	stepDefMap["call"] = createFunctionCallStep
}

type FunctionCall struct {
	BaseStep
	name             string
	inputParameters  []string
	outputParameters []string
}

func (functionCall *FunctionCall) Initalize() error {

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

func (functionCall *FunctionCall) Execute(scope *Scope) (any, error) {

	if functionSteps, ok := scope.get_function(functionCall.name); ok {
		functionScope := Scope{}
		functionScope.variables = map[string]any{}
		functionScope.functions = map[string][]Step{}

		for _, inputParameter := range functionCall.inputParameters {
			if inputParameterValue, result := scope.get_variable(inputParameter); result {
				functionScope.variables[inputParameter] = inputParameterValue
			} else {
				return nil, fmt.Errorf("input variable not found in known scope %s", inputParameter)
			}
		}

		if results, err := RunSteps(&functionScope, functionSteps...); err == nil {

			for _, outputParameter := range functionCall.outputParameters {

				if outputParameterValue, result := functionScope.get_variable(outputParameter); result {
					scope.variables[outputParameter] = outputParameterValue
				} else {
					return results, fmt.Errorf("output variable not found in function scope %s", outputParameter)
				}
			}
			return results, err
		} else {
			return results, err
		}

	} else {
		return nil, fmt.Errorf("function definitions not present in known scope %s", functionCall.name)
	}
}

func createFunctionCallStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	functionCall := &FunctionCall{}
	functionCall.tag = tag
	functionCall.attributes = attributes
	functionCall.text = text
	return functionCall, functionCall.Initalize()
}
