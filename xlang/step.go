package xlang

import "fmt"

type Step interface {
	init(*Step, string, map[string]string, string, ...Step)
	execute(Scope) (bool, error)
}

func runSteps(scope Scope, steps ...Step) (bool, error) {
	for _, step := range steps {
		if _, e := step.execute(scope); e != nil {
			return false, e
		}
	}
	return true, nil
}

type BaseStep struct {
	parent *Step
	tag    string
	steps  []Step
}

func (step *BaseStep) AddSteps(steps ...Step) {
	step.steps = append((*step).steps, steps...)
}

func (step *BaseStep) init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) {
	step.parent = parent
	step.tag = tag
	step.steps = steps
}

func (step *BaseStep) execute(scope Scope) (bool, error) {
	return runSteps(scope, (*step).steps...)
}

type GenericStep struct {
	BaseStep
	attributes map[string]string
	text       string
}

func (step *GenericStep) AddSteps(steps ...Step) {
	step.steps = append((*step).steps, steps...)
}

func (step *GenericStep) init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) {
	step.BaseStep.init(parent, tag, attributes, text, steps...)
	step.attributes = attributes
	step.text = text
}

func (step *GenericStep) execute(scope Scope) (bool, error) {
	fmt.Printf("Parent: %p\n, Tag: %s\n, Attributes: %#v\n, Text: %s\n, Steps: %T\n", step.parent, step.tag, step.attributes, step.text, step.steps)
	// fmt.Printf("%v", step)
	return step.BaseStep.execute(scope)
}
