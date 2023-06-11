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

func (importStep *ImportStep) Initalize() error {

	if file, ok := importStep.attributes["file"]; ok {
		importStep.fileName = file
	} else {
		return fmt.Errorf("file attribute missing")
	}
	return nil
}

func (importStep *ImportStep) Execute(scope *Scope, basedir string) (any, error) {
	return ExecuteFile(scope, basedir+"/"+importStep.fileName)
}

func createImportStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	importStep := &ImportStep{}
	importStep.tag = tag
	importStep.attributes = attributes
	importStep.text = text
	return importStep, nil
}
