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

	if errorCodeStr, ok := exit.parameters["code"]; ok {
		var err error
		exit.code, err = strconv.Atoi(errorCodeStr)
		if err != nil {
			return err
		}
	}

	return nil
}

func (exit *Exit) Execute(scope *Scope) (any, error) {

	if err := exit.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	os.Exit(exit.code)
	return exit.code, nil
}

func createExitStep(name string, parameters map[string]string, body string) (Step, error) {
	exit := &Exit{}
	exit.name = name
	exit.parameters = parameters
	exit.body = body
	return exit, nil
}
