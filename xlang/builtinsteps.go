package xlang

type Echo struct {
	BaseStep
	message string
}

func (step *Echo) init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) {
	step.BaseStep.init(parent, tag, attributes, text, steps...)

	if message, ok := attributes["message"]; ok {
		step.message = message
	}
}

func (step *Echo) execute(scope Scope) (bool, error) {
	// fmt.Printf("Parent: %p\n, Tag: %s\n, Attributes: %#v\n, Text: %s\n, Steps: %T\n", step.parent, step.tag, step.attributes, step.text, step.steps)
	// fmt.Printf("%v", step)
	return step.BaseStep.execute(scope)
}
