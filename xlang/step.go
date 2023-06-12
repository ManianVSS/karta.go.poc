package xlang

import (
	"fmt"
)

type Step interface {
	//Set parent if passed non nil and get the set parent
	Parent(Step) Step
	//Set tag for the step if passed non empty and get the set tag
	Tag(string) string
	//Set attributes for the step if passed non nil and get the set attributes
	Attributes(map[string]string) map[string]string
	//Set text for the step if passed non empty and get the set text
	Text(string) string
	//Adds nested steps for the step if passed non empty and get the nested steps
	NestedSteps(...Step) []Step

	//Execute the step in the scope provided and return if any error
	Execute(*Scope, string) (any, error)
}

type InitStepFunction func(Step, string, map[string]string, string) (Step, error)

var stepDefMap map[string]InitStepFunction = map[string]InitStepFunction{}

func RegisterStepDefinition(tag string, initStepFunction InitStepFunction) error {
	if _, ok := stepDefMap[tag]; !ok {
		stepDefMap[tag] = initStepFunction
	} else {
		return fmt.Errorf("step definition for tag %s already registered", tag)
	}
	return nil
}

type StrToVarFunction func(string) (any, error)

var variableParserFunctionMap map[string]StrToVarFunction = map[string]StrToVarFunction{}

func RegisterVariableTypeDefinition(varType string, strToVarFunction StrToVarFunction) error {
	if _, ok := stepDefMap[varType]; !ok {
		variableParserFunctionMap[varType] = strToVarFunction
	} else {
		return fmt.Errorf("variable parse function definition for type %s already registered", varType)
	}
	return nil
}

func RunSteps(scope *Scope, basedir string, steps ...Step) (any, error) {

	results := make([]any, len(steps))
	for i, step := range steps {
		if step != nil {
			result, err := step.Execute(scope, basedir)
			if err != nil {
				if _, ok := err.(*MethodReturnError); ok {
					results[i] = result
					// fmt.Printf("Encountered return %t, %s\n", err, err.Error())
					break
				} else {
					fmt.Printf("Program execution error while executing step %t with error %s\n", step, err.Error())
				}
				return results, err
			}
		} else {
			return results, fmt.Errorf("encountered an unexpected condition where step is nil")
		}
	}
	return results, nil
}

type BaseStep struct {
	parent      Step
	tag         string
	attributes  map[string]string
	text        string
	nestedSteps []Step
}

func (baseStep *BaseStep) Parent(parent Step) Step {
	if parent != nil {
		baseStep.parent = parent
	}
	return baseStep.parent
}

func (baseStep *BaseStep) Tag(tag string) string {
	if tag != "" {
		baseStep.tag = tag
	}
	return baseStep.tag
}

func (baseStep *BaseStep) Attributes(attributes map[string]string) map[string]string {
	if attributes != nil {
		baseStep.attributes = attributes
	}
	return baseStep.attributes
}
func (baseStep *BaseStep) Text(text string) string {
	if text != "" {
		baseStep.text = text
	}
	return baseStep.text
}

func (baseStep *BaseStep) NestedSteps(steps ...Step) []Step {
	if baseStep.nestedSteps == nil {
		baseStep.nestedSteps = []Step{}
	}
	baseStep.nestedSteps = append(baseStep.nestedSteps, steps...)
	return baseStep.nestedSteps
}

func (baseStep *BaseStep) InitalizeAndCheck() error {
	return nil
}

func (baseStep *BaseStep) Execute(scope *Scope, basedir string) (any, error) {

	if err := baseStep.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	return RunSteps(scope, basedir, baseStep.nestedSteps...)
}

func createBaseStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	baseStep := &BaseStep{}
	baseStep.tag = tag
	baseStep.attributes = attributes
	baseStep.text = text
	return baseStep, nil
}
