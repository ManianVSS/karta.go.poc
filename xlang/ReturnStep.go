package xlang

func init() {
	stepDefMap["return"] = createReturnStep
}

type MethodReturnError struct {
}

func (methodReturnError MethodReturnError) Error() string {
	return ""
}

type ReturnStep struct {
	BaseStep
}

func (returnStep *ReturnStep) Execute(scope *Scope, basedir string) (any, error) {
	return nil, &MethodReturnError{}
}

func createReturnStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	returnStep := &ReturnStep{}
	returnStep.tag = tag
	returnStep.attributes = attributes
	returnStep.text = text
	return returnStep, nil
}
