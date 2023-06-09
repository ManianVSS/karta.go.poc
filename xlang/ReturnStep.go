package xlang

type MethodReturnError struct {
}

func (methodReturnError MethodReturnError) Error() string {
	return ""
}

type ReturnStep struct {
	BaseStep
}

func (returnStep *ReturnStep) Init(tag string, attributes map[string]string, text string) error {
	return returnStep.BaseStep.Init(tag, attributes, text)
}

func (returnStep *ReturnStep) Execute(scope *Scope) error {
	return &MethodReturnError{}
}

func createReturnStep(tag string, attributes map[string]string, text string) (Step, error) {
	returnStep := &ReturnStep{}
	return returnStep, returnStep.Init(tag, attributes, text)
}
