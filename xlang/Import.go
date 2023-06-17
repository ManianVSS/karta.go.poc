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

	if file, ok := importStep.attributes["file"]; ok {
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

func createImportStep(tag string, attributes map[string]string, text string) (Step, error) {
	importStep := &ImportStep{}
	importStep.tag = tag
	importStep.attributes = attributes
	importStep.text = text
	return importStep, nil
}
