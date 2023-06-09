package xlang

import (
	"fmt"

	"github.com/subchen/go-xmldom"
)

var StepDefMap map[string]InitStepFunction = map[string]InitStepFunction{}

func InitStepDefinitions() {
	StepDefMap["echo"] = createEchoStep
	StepDefMap["return"] = createReturnStep
	StepDefMap["func"] = createFunctionDefinitionStep
	StepDefMap["call"] = createFunctionCallStep
	StepDefMap["genericStep"] = createBaseStep
}

func GetStepFromNode(parent *Step, node *xmldom.Node) (Step, error) {
	createStepFunction, ok := StepDefMap[node.Name]
	if !ok {
		createStepFunction = createBaseStep
	}

	if step, err := createStepFunction(node.Name, xmlAttrToAttributes(node.Attributes), node.Text); err == nil {
		step.Parent(parent)
		// steps = append(steps, step)

		for _, child := range node.Children {
			if childsNestedStep, err := GetStepFromNode(&step, child); err == nil {
				step.AddNestedSteps(childsNestedStep)
			}
		}
		return step, err
	} else {
		return step, fmt.Errorf("could not parse step %s using step parsing handler function, %s", node.Name, err.Error())
	}
}

func ExecuteFile(fileName string) error {
	doc := xmldom.Must(xmldom.ParseFile(fileName))
	root := doc.Root

	scope := Scope{}
	scope.variables = map[string]any{}
	scope.functions = map[string]*FunctionDefinition{}

	rootStep, err := GetStepFromNode(nil, root)

	if err != nil {
		return err
	}

	if rootStep == nil {
		return fmt.Errorf("unexpected program parsing error; Nil root step")
	}

	return rootStep.Execute(scope)
}

func Main() {
	InitStepDefinitions()
	if err := ExecuteFile("sampleapp.xml"); err != nil {
		fmt.Printf("Program execution error %s\n", err.Error())
	}
}
