package xlang

import (
	"fmt"
	"strings"
)

func init() {
	stepDefMap["step"] = createCustomStepDefinitionStep
}

type CustomStepDefinition struct {
	BaseStep
	stepName       string
	attributeNames []string
	textAttribute  string
}

func (customStepDefinition *CustomStepDefinition) InitalizeAndCheck() error {

	if stepName, ok := customStepDefinition.parameters["name"]; ok {
		customStepDefinition.stepName = stepName
	} else {
		return fmt.Errorf("name attribute missing")
	}

	if attributeNames, ok := customStepDefinition.parameters["attributeNames"]; ok {
		customStepDefinition.attributeNames = strings.Split(attributeNames, ",")
	}

	if textAttribute, ok := customStepDefinition.parameters["textAttribute"]; ok {
		customStepDefinition.textAttribute = textAttribute
	}

	return nil
}

func (customStepDefinition *CustomStepDefinition) Execute(scope *Scope) (any, error) {

	if err := customStepDefinition.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	// stepTemplateSteps := make([]Step, len(customStepDefinition.nestedSteps))
	// copy(stepTemplateSteps, customStepDefinition.nestedSteps)

	if _, ok := stepDefMap[customStepDefinition.stepName]; !ok {

		stepDefMap[customStepDefinition.stepName] =
			func(name string, parameters map[string]string, body string) (Step, error) {
				// fmt.Printf("Entering the closure.. Steps to Copy %#v", customStepDefinition.nestedSteps)
				customStep := &BaseStep{}
				customStep.name = name
				customStep.parameters = parameters
				customStep.body = body
				customStep.steps = make([]Step, len(customStepDefinition.steps)) //customStepDefinition.nestedSteps
				copy(customStep.steps, customStepDefinition.steps)
				for index := range customStep.steps {
					customStep.steps[index].Parent(customStep)
				}
				return customStep, nil
			}
	} else {
		return nil, fmt.Errorf("step definition already present for name %s", customStepDefinition.stepName)
	}
	return nil, nil
}

func createCustomStepDefinitionStep(name string, parameters map[string]string, body string) (Step, error) {
	customStepDefinition := &CustomStepDefinition{}
	customStepDefinition.name = name
	customStepDefinition.parameters = parameters
	customStepDefinition.body = body
	return customStepDefinition, nil
}
