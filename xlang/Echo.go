package xlang

import (
	"fmt"
)

func init() {
	stepDefMap["echo"] = createEchoStep
}

type Echo struct {
	BaseStep
	message string
}

func (echo *Echo) InitalizeAndCheck() error {

	if message, ok := echo.parameters["message"]; ok {
		echo.message = message
	} else if echo.body != "" {
		echo.message = echo.body
	} else {
		return fmt.Errorf("message attribute missing")
	}
	return nil
}

func (echo *Echo) Execute(scope *Scope) (any, error) {

	if err := echo.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	var parentAttributes map[string]string
	if echo.parent != nil {
		parentAttributes = echo.parent.Parameters()
	}
	byteWrittenCount, err := fmt.Println(replaceVarsInString(echo.message, scope, parentAttributes))
	return byteWrittenCount > 0, err
}

func createEchoStep(name string, parameters map[string]string, body string) (Step, error) {
	echo := &Echo{}
	echo.name = name
	echo.parameters = parameters
	echo.body = body
	return echo, nil
}
