package xlang

import "fmt"

type Step interface {
	AddSteps(...Step)
	Init(*Step, string, map[string]string, string, ...Step) (Step, error)
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
	Parent      *Step
	Tag         string
	NestedSteps []Step
}

func (baseStep *BaseStep) AddSteps(steps ...Step) {
	baseStep.NestedSteps = append(baseStep.NestedSteps, steps...)
}

func (baseStep *BaseStep) Init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {
	baseStep.Parent = parent
	baseStep.Tag = tag
	baseStep.NestedSteps = steps
	return baseStep, nil
}

func (baseStep *BaseStep) Execute(scope Scope) (bool, error) {
	return RunSteps(scope, baseStep.NestedSteps...)
}

type GenericStep struct {
	BaseStep
	Attributes map[string]string
	Text       string
}

// func (genericStep *GenericStep) AddSteps(steps ...Step) {
// 	genericStep.NestedSteps = append(genericStep.NestedSteps, steps...)
// }

func (genericStep *GenericStep) init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {
	if _, err := genericStep.BaseStep.Init(parent, tag, attributes, text, steps...); err != nil {
		return genericStep, err
	}
	genericStep.Attributes = attributes
	genericStep.Text = text
	return genericStep, nil
}

func (genericStep *GenericStep) Execute(scope Scope) (bool, error) {
	fmt.Printf("Parent: %p\n, Tag: %s\n, Attributes: %#v\n, Text: %s\n, Steps: %T\n", genericStep.Parent, genericStep.Tag, genericStep.Attributes, genericStep.Text, genericStep.NestedSteps)
	// fmt.Printf("%v", step)
	return genericStep.BaseStep.Execute(scope)
}

type InitStepFunction func(string, map[string]string, string) (Step, error)

func createGenericStep(tag string, attributes map[string]string, text string) (Step, error) {
	genericStep := GenericStep{}
	return genericStep.init(nil, tag, attributes, text)
}

type MethodReturnError struct {
}

func (methodReturnError MethodReturnError) Error() string {
	return ""
}
