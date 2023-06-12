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
	code int
}

func (exit *Exit) InitalizeAndCheck() error {

	if errorCodeStr, ok := exit.attributes["code"]; ok {
		var err error
		exit.code, err = strconv.Atoi(errorCodeStr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (exit *Exit) Execute(scope *Scope, basedir string) (any, error) {

	if err := exit.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	os.Exit(exit.code)
	return exit.code, nil
}

func createExitStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	exit := &Exit{}
	exit.tag = tag
	exit.attributes = attributes
	exit.text = text
	return exit, nil
}
