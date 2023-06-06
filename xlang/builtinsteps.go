package xlang

import "fmt"

type Echo struct {
	BaseStep
	message string
}

func (echo *Echo) init(parent *Step, tag string, attributes map[string]string, text string, steps ...Step) (Step, error) {
	echo.BaseStep.init(parent, tag, attributes, text, steps...)

	if message, ok := attributes["message"]; ok {
		echo.message = message
	} else {
		return nil, fmt.Errorf("message attribute missing")
	}
	return echo, nil
}

func (echo *Echo) execute(scope Scope) (bool, error) {
	fmt.Println(echo.message)
	return echo.BaseStep.execute(scope)
}
