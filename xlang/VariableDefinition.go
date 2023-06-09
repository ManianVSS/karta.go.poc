package xlang

import (
	"fmt"
)

type VariableDefinition struct {
	BaseStep
	name    string
	value   string
	varType string
}

func (variableDefinition *VariableDefinition) Init(tag string, attributes map[string]string, text string) error {

	if err := variableDefinition.BaseStep.Init(tag, attributes, text); err != nil {
		return err
	}

	if name, ok := variableDefinition.attributes["name"]; ok {
		variableDefinition.name = name
	} else {
		return fmt.Errorf("name attribute missing")
	}

	if value, ok := variableDefinition.attributes["value"]; ok {
		variableDefinition.value = value
	} else {
		return fmt.Errorf("value attribute missing")
	}

	if varType, ok := variableDefinition.attributes["type"]; ok {
		variableDefinition.varType = varType
	} else {
		return fmt.Errorf("type attribute missing")
	}

	return nil
}

func (variableDefinition *VariableDefinition) Execute(scope *Scope) error {
	strReplacedWithVariables := replaceVariablesInString(variableDefinition.value, scope.variables)
	if strToVarFunction, ok := variableParserFunctionMap[variableDefinition.varType]; ok {
		if parsedValue, err := strToVarFunction(strReplacedWithVariables); err == nil {
			scope.variables[variableDefinition.name] = parsedValue
		} else {
			return err
		}
	} else {
		return fmt.Errorf("undefined type for value conversion %s", variableDefinition.varType)
	}
	return nil
}

func createVariableDefinitionStep(tag string, attributes map[string]string, text string) (Step, error) {
	variableDefinition := &VariableDefinition{}
	return variableDefinition, variableDefinition.Init(tag, attributes, text)
}
