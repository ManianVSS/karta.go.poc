package xlang

import (
	"os"
	"strconv"
)

func init() {
	stepDefMap["exit"] = createExitStep
}

type Exit struct {
	BaseStep
	errorCode int
}

func (exit *Exit) Initalize() error {

	if errorCodeStr, ok := exit.attributes["exitCode"]; ok {
		var err error
		exit.errorCode, err = strconv.Atoi(errorCodeStr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (exit *Exit) Execute(scope *Scope) (any, error) {
	os.Exit(exit.errorCode)
	return exit.errorCode, nil
}

func createExitStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	exit := &Echo{}
	exit.tag = tag
	exit.attributes = attributes
	exit.text = text
	return exit, exit.Initalize()
}
