package xlang

import (
	"fmt"

	"github.com/subchen/go-xmldom"
)

func GetStepFromNode(parent Step, node *xmldom.Node) (Step, error) {
	createStepFunction, ok := stepDefMap[node.Name]
	if !ok {
		// fmt.Printf("Unregistered tag %s\n", node.Name)
		createStepFunction = createBaseStep
	}

	if step, err := createStepFunction(parent, node.Name, xmlAttrToAttributes(node.Attributes), node.Text); err == nil {
		for _, child := range node.Children {
			if childsNestedStep, err := GetStepFromNode(step, child); err == nil {
				step.AddNestedSteps(childsNestedStep)
			}
		}

		if step.Parent(nil) == nil {
			step.Parent(parent)
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
	scope.functions = map[string][]Step{}

	var rootStep BaseStep = BaseStep{
		tag:        root.Name,
		attributes: xmlAttrToAttributes(root.Attributes),
		text:       root.Text,
	}

	for _, mainStepNode := range root.Children {
		if mainStep, err := GetStepFromNode(&rootStep, mainStepNode); err == nil {

			if mainStep == nil {
				return fmt.Errorf("unexpected program parsing error; Nil root step")
			}

			if err := mainStep.Execute(&scope); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func Main() {
	InitVariableTypeDefinitions()
	InitStepDefinitions()
	if err := ExecuteFile("sampleapp.xml"); err != nil {
		fmt.Printf("Program execution error %s\n", err.Error())
	}
}
