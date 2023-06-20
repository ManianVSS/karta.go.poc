package xlang

import (
	"fmt"
)

func init() {
	stepDefMap["var"] = createVariableDefinitionStep
}

type VariableDefinition struct {
	BaseStep
	variableName string
	value        string
	varType      string
}

func (variableDefinition *VariableDefinition) InitalizeAndCheck() error {

	if name, ok := variableDefinition.parameters["name"]; ok {
		variableDefinition.variableName = name
	} else {
		return fmt.Errorf("name attribute missing")
	}

	if value, ok := variableDefinition.parameters["value"]; ok {
		variableDefinition.value = value
	} else {
		return fmt.Errorf("value attribute missing")
	}

	if varType, ok := variableDefinition.parameters["type"]; ok {
		variableDefinition.varType = varType
	} else {
		return fmt.Errorf("type attribute missing")
	}

	return nil
}

func (variableDefinition *VariableDefinition) Execute(scope *Scope) (any, error) {

	if err := variableDefinition.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	var parentAttributes map[string]string
	if variableDefinition.parent != nil {
		parentAttributes = variableDefinition.parent.Parameters()
	}

	strReplacedWithVariables := replaceVarsInString(variableDefinition.value, scope, parentAttributes)
	if strToVarFunction, ok := variableParserFunctionMap[variableDefinition.varType]; ok {
		if parsedValue, err := strToVarFunction(strReplacedWithVariables); err == nil {
			scope.variables[variableDefinition.variableName] = parsedValue
		} else {
			return parsedValue, err
		}
	} else {
		return variableDefinition.value, fmt.Errorf("undefined type for value conversion %s", variableDefinition.varType)
	}
	return variableDefinition.value, nil
}

func createVariableDefinitionStep(name string, parameters map[string]string, body string) (Step, error) {
	variableDefinition := &VariableDefinition{}
	variableDefinition.name = name
	variableDefinition.parameters = parameters
	variableDefinition.body = body
	return variableDefinition, nil
}
