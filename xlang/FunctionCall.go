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
	functionName     string
	inputParameters  []string
	outputParameters []string
}

func (functionCall *FunctionCall) InitalizeAndCheck() error {

	if name, ok := functionCall.parameters["name"]; ok {
		functionCall.functionName = name
	} else {
		return fmt.Errorf("name attribute missing")
	}

	if inputParameters, ok := functionCall.parameters["inputParameters"]; ok {
		functionCall.inputParameters = strings.Split(inputParameters, ",")
	} else {
		functionCall.inputParameters = []string{}
	}

	if outputParameters, ok := functionCall.parameters["outputParameters"]; ok {
		functionCall.outputParameters = strings.Split(outputParameters, ",")
	} else {
		functionCall.outputParameters = []string{}
	}

	return nil
}

func (functionCall *FunctionCall) Execute(scope *Scope) (any, error) {

	if err := functionCall.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	if functionSteps, ok := scope.getFunction(functionCall.functionName); ok {
		functionScope := Scope{}
		functionScope.variables = map[string]any{}
		functionScope.functions = map[string][]Step{}

		for _, inputParameter := range functionCall.inputParameters {
			if inputParameterValue, result := scope.getVariable(inputParameter); result {
				functionScope.variables[inputParameter] = inputParameterValue
			} else {
				return nil, fmt.Errorf("input variable not found in known scope %s", inputParameter)
			}
		}

		if results, err := RunSteps(&functionScope, functionSteps...); err == nil {

			for _, outputParameter := range functionCall.outputParameters {

				if outputParameterValue, result := functionScope.getVariable(outputParameter); result {
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
		return nil, fmt.Errorf("function definitions not present in known scope %s", functionCall.functionName)
	}
}

func createFunctionCallStep(name string, parameters map[string]string, body string) (Step, error) {
	functionCall := &FunctionCall{}
	functionCall.name = name
	functionCall.parameters = parameters
	functionCall.body = body
	return functionCall, nil
}
