package xlang

import "fmt"

type Echo struct {
	BaseStep
	message string
}

func (echo *Echo) init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {

	if baseStep, err := echo.BaseStep.init(parent, tag, attributes, text, steps...); err != nil {
		return baseStep, err
	}

	if message, ok := attributes["message"]; ok {
		echo.message = message
	} else if text != "" {
		echo.message = text
	} else {
		return nil, fmt.Errorf("message attribute missing")
	}
	return echo, nil
}

func (echo *Echo) execute(scope Scope) (bool, error) {
	fmt.Println(echo.message)
	return echo.BaseStep.execute(scope)
}

func createEchoStep(tag string, attributes map[string]string, text string) (Step, error) {
	echo := Echo{}
	return echo.init(nil, tag, attributes, text)
}

type ReturnStep struct {
	BaseStep
	// returnValue any
}

func (returnStep *ReturnStep) init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {
	// fmt.Println("Hit return step init")
	return returnStep.BaseStep.init(parent, tag, attributes, text, steps...)
}

func (returnStep *ReturnStep) execute(scope Scope) (bool, error) {
	fmt.Println("Hit return step execute")
	return false, &MethodReturnError{}
}

func createReturnStep(tag string, attributes map[string]string, text string) (Step, error) {
	// fmt.Println("Hit return step create")
	returnStep := ReturnStep{}
	return returnStep.init(nil, tag, attributes, text)
}
