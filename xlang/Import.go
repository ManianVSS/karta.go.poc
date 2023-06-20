package xlang

import (
	"fmt"
)

func init() {
	stepDefMap["import"] = createImportStep
}

type ImportStep struct {
	BaseStep
	fileName string
}

func (importStep *ImportStep) InitalizeAndCheck() error {

	if file, ok := importStep.parameters["file"]; ok {
		importStep.fileName = file
	} else {
		return fmt.Errorf("file attribute missing")
	}
	return nil
}

func (importStep *ImportStep) Execute(scope *Scope) (any, error) {

	if err := importStep.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	return ExecuteFile(scope, BaseDir+importStep.fileName)
}

func createImportStep(name string, parameters map[string]string, body string) (Step, error) {
	importStep := &ImportStep{}
	importStep.name = name
	importStep.parameters = parameters
	importStep.body = body
	return importStep, nil
}
