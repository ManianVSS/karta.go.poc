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
	name           string
	attributeNames []string
	textAttribute  string
}

func (customStepDefinition *CustomStepDefinition) InitalizeAndCheck() error {

	if name, ok := customStepDefinition.attributes["name"]; ok {
		customStepDefinition.name = name
	} else {
		return fmt.Errorf("name attribute missing")
	}

	if attributeNames, ok := customStepDefinition.attributes["attributeNames"]; ok {
		customStepDefinition.attributeNames = strings.Split(attributeNames, ",")
	}

	if textAttribute, ok := customStepDefinition.attributes["textAttribute"]; ok {
		customStepDefinition.textAttribute = textAttribute
	}

	return nil
}

func (customStepDefinition *CustomStepDefinition) Execute(scope *Scope, basedir string) (any, error) {

	if err := customStepDefinition.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	// stepTemplateSteps := make([]Step, len(customStepDefinition.nestedSteps))
	// copy(stepTemplateSteps, customStepDefinition.nestedSteps)

	if _, ok := stepDefMap[customStepDefinition.name]; !ok {

		stepDefMap[customStepDefinition.name] =
			func(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
				// fmt.Printf("Entering the closure.. Steps to Copy %#v", customStepDefinition.nestedSteps)
				customStep := &BaseStep{}
				customStep.parent = parent
				customStep.tag = tag
				customStep.attributes = attributes
				customStep.text = text
				customStep.nestedSteps = make([]Step, len(customStepDefinition.nestedSteps)) //customStepDefinition.nestedSteps
				copy(customStep.nestedSteps, customStepDefinition.nestedSteps)
				for index := range customStep.nestedSteps {
					customStep.nestedSteps[index].Parent(customStep)
				}
				return customStep, nil
			}
	} else {
		return nil, fmt.Errorf("step definition already present for name %s", customStepDefinition.name)
	}
	return nil, nil
}

func createCustomStepDefinitionStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	customStepDefinition := &CustomStepDefinition{}
	customStepDefinition.tag = tag
	customStepDefinition.attributes = attributes
	customStepDefinition.text = text
	return customStepDefinition, nil
}
