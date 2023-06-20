package xlang

import (
	"fmt"
)

type Step interface {
	//Set parent if passed non nil and get the set parent
	Parent(Step) Step
	//Set tag for the step if passed non empty and get the set tag
	Name() string
	//Set attributes for the step if passed non nil and get the set attributes
	Parameters() map[string]string
	//Set text for the step if passed non empty and get the set text
	Body() string
	//Adds nested steps for the step if passed non empty and get the nested steps
	Steps(...Step) []Step

	//Execute the step in the scope provided and return if any error
	Execute(*Scope) (any, error)
}

type InitStepFunction func(string, map[string]string, string) (Step, error)

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

func RunSteps(scope *Scope, steps ...Step) (any, error) {

	results := make([]any, len(steps))
	for i, step := range steps {
		if step != nil {
			result, err := step.Execute(scope)
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
	parent     Step
	name       string
	parameters map[string]string
	body       string
	steps      []Step
}

func (baseStep *BaseStep) Parent(parent Step) Step {
	if parent != nil {
		baseStep.parent = parent
	}
	return baseStep.parent
}

func (baseStep *BaseStep) Name() string {
	return baseStep.name
}

func (baseStep *BaseStep) Parameters() map[string]string {
	return baseStep.parameters
}

func (baseStep *BaseStep) Body() string {
	return baseStep.body
}

func (baseStep *BaseStep) Steps(steps ...Step) []Step {
	if baseStep.steps == nil {
		baseStep.steps = []Step{}
	}
	baseStep.steps = append(baseStep.steps, steps...)
	return baseStep.steps
}

func (baseStep *BaseStep) InitalizeAndCheck() error {
	return nil
}

func (baseStep *BaseStep) Execute(scope *Scope) (any, error) {

	if err := baseStep.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	return RunSteps(scope, baseStep.steps...)
}

func createBaseStep(name string, parameters map[string]string, body string) (Step, error) {
	baseStep := &BaseStep{}
	baseStep.name = name
	baseStep.parameters = parameters
	baseStep.body = body
	return baseStep, nil
}
