package xlang

import (
	"fmt"

	"github.com/subchen/go-xmldom"
)

type InitStepFunction func(string, map[string]string, string) (Step, error)

var StepDefMap map[string]InitStepFunction = map[string]InitStepFunction{}

func createGenericStep(tag string, attributes map[string]string, text string) (Step, error) {
	genericStep := GenericStep{}
	return genericStep.init(nil, tag, attributes, text)
}

func createEchoStep(tag string, attributes map[string]string, text string) (Step, error) {
	echo := Echo{}
	return echo.init(nil, tag, attributes, text)
}

func InitStepDefinitions() {
	// StepDefMap = map[string]InitStepFunction{
	// 	"echo": createEchoStep,
	// }
	// StepDefMap["app"] = createGenericStep
	StepDefMap["echo"] = createEchoStep
}

func xmlAttrToAttributes(xmlAttributes []*xmldom.Attribute) map[string]string {
	attributesMap := map[string]string{}
	for _, xmlattr := range xmlAttributes {
		attributesMap[xmlattr.Name] = xmlattr.Value
	}
	return attributesMap
}

func GetSteps(parent *Step, node *xmldom.Node) ([]Step, error) {

	steps := []Step{}
	for _, child := range node.Children {
		if createStepFunction, ok := StepDefMap[child.Name]; ok {
			if step, err := createStepFunction(child.Name, xmlAttrToAttributes(child.Attributes), child.Text); err == nil {
				steps = append(steps, step)
			} else {
				return steps, fmt.Errorf("could not parse step %s using step parsing handler function", node.Name)
			}
		} else {
			return steps, fmt.Errorf("could not find step definition for %s", node.Name)
			// return createGenericStep(node.Name, xmlAttrToAttributes(node.Attributes), node.Text)
		}
	}
	return steps, nil
}

func ExecuteFile(fileName string) bool {
	doc := xmldom.Must(xmldom.ParseFile(fileName))
	root := doc.Root

	scope := Scope{}

	if rootSteps, err := GetSteps(nil, root); (err == nil) && (rootSteps != nil) {
		if result, err := runSteps(scope, rootSteps...); err == nil {
			fmt.Println("Program execution result is ", result)
			return result
		} else {
			fmt.Printf("Program execution error while executing step %t with error %T\n", result, err)
			return false
		}
	} else {
		fmt.Printf("Program execution error while parsing step %T with error %T\n", rootSteps, err)
		return false
	}
}

func Main() {
	InitStepDefinitions()
	ExecuteFile("sampleapp.xml")
}

// func printNode(node *xmldom.Node) {
// 	fmt.Printf("%s, %T, %s\n", node.Name, node.Attributes, node.Text)

// 	for _, child := range node.Children {
// 		printNode(child)
// 	}
// }
