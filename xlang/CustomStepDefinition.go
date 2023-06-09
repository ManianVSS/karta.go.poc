package xlang

import (
	"fmt"
	"strings"
)

type CustomStepDefinition struct {
	BaseStep
	name           string
	attributeNames []string
	textAttribute  string
}

func (customStepDefinition *CustomStepDefinition) Init(tag string, attributes map[string]string, text string) error {

	if err := customStepDefinition.BaseStep.Init(tag, attributes, text); err != nil {
		return err
	}

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

func (customStepDefinition *CustomStepDefinition) Execute(scope *Scope) error {

	// stepTemplateSteps := make([]Step, len(customStepDefinition.nestedSteps))
	// copy(stepTemplateSteps, customStepDefinition.nestedSteps)

	if _, ok := stepDefMap[customStepDefinition.name]; !ok {

		stepDefMap[customStepDefinition.name] =
			func(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
				// fmt.Printf("Entering the closure.. Steps to Copy %#v", customStepDefinition.nestedSteps)
				customStep := &struct{ BaseStep }{}
				customStep.parent = parent
				if err := customStep.Init(tag, attributes, text); err != nil {
					return customStep, err
				}
				customStep.nestedSteps = make([]Step, len(customStepDefinition.nestedSteps)) //customStepDefinition.nestedSteps
				copy(customStep.nestedSteps, customStepDefinition.nestedSteps)
				for index := range customStep.nestedSteps {
					customStep.nestedSteps[index].Parent(customStep)
				}
				return customStep, nil
			}
	} else {
		return fmt.Errorf("step definition already present for name %s", customStepDefinition.name)
	}
	return nil
}

func createCustomStepDefinitionStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	customStepDefinition := &CustomStepDefinition{}
	return customStepDefinition, customStepDefinition.Init(tag, attributes, text)
}
