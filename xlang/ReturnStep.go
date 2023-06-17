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

func (returnStep *ReturnStep) Execute(scope *Scope) (any, error) {

	if err := returnStep.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	return nil, &MethodReturnError{}
}

func createReturnStep(tag string, attributes map[string]string, text string) (Step, error) {
	returnStep := &ReturnStep{}
	returnStep.tag = tag
	returnStep.attributes = attributes
	returnStep.text = text
	return returnStep, nil
}
