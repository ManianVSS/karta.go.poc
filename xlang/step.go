package xlang

import "fmt"

type Step interface {
	Parent(*Step)
	Init(string, map[string]string, string) error
	AddNestedSteps(...Step)
	Execute(Scope) error
}

func RunSteps(scope Scope, steps ...Step) error {
	for _, step := range steps {
		if step != nil {
			err := step.Execute(scope)
			if err != nil {
				if _, ok := err.(*MethodReturnError); ok {
					// fmt.Printf("Encountered return %t, %s\n", err, err.Error())
					break
				} else {
					fmt.Printf("Program execution error while executing step %t with error %s\n", step, err.Error())
				}
				return err
			}
		} else {
			return fmt.Errorf("encountered an unexpected condition where step is nil")
		}
	}
	return nil
}

type BaseStep struct {
	parent      *Step
	tag         string
	attributes  map[string]string
	text        string
	nestedSteps []Step
}

func (baseStep *BaseStep) Parent(parent *Step) {
	baseStep.parent = parent
}

func (baseStep *BaseStep) Init(tag string, attributes map[string]string, text string) error {
	baseStep.tag = tag
	baseStep.attributes = attributes
	baseStep.text = text
	return nil
}

func (baseStep *BaseStep) AddNestedSteps(steps ...Step) {
	if baseStep.nestedSteps == nil {
		baseStep.nestedSteps = steps
	} else {
		baseStep.nestedSteps = append(baseStep.nestedSteps, steps...)
	}
}

func (baseStep *BaseStep) Execute(scope Scope) error {
	return RunSteps(scope, baseStep.nestedSteps...)
}

func createBaseStep(tag string, attributes map[string]string, text string) (Step, error) {
	baseStep := &BaseStep{}
	return baseStep, baseStep.Init(tag, attributes, text)
}

type InitStepFunction func(string, map[string]string, string) (Step, error)

type MethodReturnError struct {
}

func (methodReturnError MethodReturnError) Error() string {
	return ""
}
