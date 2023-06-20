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

func createReturnStep(name string, parameters map[string]string, body string) (Step, error) {
	returnStep := &ReturnStep{}
	returnStep.name = name
	returnStep.parameters = parameters
	returnStep.body = body
	return returnStep, nil
}
