package xlang

import (
	"fmt"
)

type Echo struct {
	BaseStep
	message string
}

func (echo *Echo) Init(tag string, attributes map[string]string, text string) error {

	if err := echo.BaseStep.Init(tag, attributes, text); err != nil {
		return err
	}

	if message, ok := echo.attributes["message"]; ok {
		echo.message = message
	} else if echo.text != "" {
		echo.message = echo.text
	} else {
		return fmt.Errorf("message attribute missing")
	}
	return nil
}

func (echo *Echo) Execute(scope *Scope) error {
	var parentAttributes map[string]string
	if echo.parent != nil {
		parentAttributes = echo.parent.Attributes()
	}
	_, err := fmt.Println(replaceVarsInString(echo.message, scope.variables, parentAttributes))
	return err
}

func createEchoStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	echo := &Echo{}
	return echo, echo.Init(tag, attributes, text)
}
