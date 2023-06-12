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

	if message, ok := echo.attributes["message"]; ok {
		echo.message = message
	} else if echo.text != "" {
		echo.message = echo.text
	} else {
		return fmt.Errorf("message attribute missing")
	}
	return nil
}

func (echo *Echo) Execute(scope *Scope, basedir string) (any, error) {

	if err := echo.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	var parentAttributes map[string]string
	if echo.parent != nil {
		parentAttributes = echo.parent.Attributes(nil)
	}
	byteWrittenCount, err := fmt.Println(replaceVarsInString(echo.message, scope, parentAttributes))
	return byteWrittenCount > 0, err
}

func createEchoStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	echo := &Echo{}
	echo.tag = tag
	echo.attributes = attributes
	echo.text = text
	return echo, nil
}
