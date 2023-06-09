package xlang

import "fmt"

type Step interface {
	Init(*Step, string, map[string]string, string, ...Step) (Step, error)
	AddNestedSteps(...Step)
	Execute(Scope) (bool, error)
}

func RunSteps(scope Scope, steps ...Step) (bool, error) {
	overallResult := true
	for _, step := range steps {
		if step != nil {
			result, err := step.Execute(scope)
			overallResult = overallResult && result

			if err != nil {
				if _, ok := err.(*MethodReturnError); ok {
					// fmt.Printf("Encountered return %t, %s\n", err, err.Error())
					break
				} else {
					fmt.Printf("Program execution error while executing step %t with error %s\n", step, err.Error())
				}
				return overallResult, err
			}
			if !overallResult {
				break
			}
		} else {
			fmt.Printf("Encountered an unexpected condition where step is nil")
		}
	}

	return overallResult, nil
}

type BaseStep struct {
	parent      *Step
	tag         string
	attributes  map[string]string
	text        string
	nestedSteps []Step
}

func (baseStep *BaseStep) AddNestedSteps(steps ...Step) {
	if baseStep.nestedSteps == nil {
		baseStep.nestedSteps = steps
	} else {
		baseStep.nestedSteps = append(baseStep.nestedSteps, steps...)
	}
}

func (baseStep *BaseStep) Init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {
	baseStep.parent = parent
	baseStep.tag = tag
	baseStep.attributes = attributes
	baseStep.text = text
	baseStep.nestedSteps = steps
	return baseStep, nil
}

func (baseStep *BaseStep) Execute(scope Scope) (bool, error) {
	return RunSteps(scope, baseStep.nestedSteps...)
}

func createBaseStep(tag string, attributes map[string]string, text string) (Step, error) {
	genericStep := &BaseStep{}
	return genericStep.Init(nil, tag, attributes, text)
}

type InitStepFunction func(string, map[string]string, string) (Step, error)

type MethodReturnError struct {
}

func (methodReturnError MethodReturnError) Error() string {
	return ""
}
