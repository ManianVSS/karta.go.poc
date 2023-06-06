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
	StepDefMap["app"] = createGenericStep
	StepDefMap["echo"] = createEchoStep
}

func xmlAttrToAttributes(xmlAttributes []*xmldom.Attribute) map[string]string {
	attributesMap := map[string]string{}
	for _, xmlattr := range xmlAttributes {
		attributesMap[xmlattr.Name] = xmlattr.Value
	}
	return attributesMap
}

func GetStep(parent *Step, node *xmldom.Node) (Step, error) {
	if createStepFunction, ok := StepDefMap[node.Name]; ok {
		return createStepFunction(node.Name, xmlAttrToAttributes(node.Attributes), node.Text)
	} else {
		return nil, fmt.Errorf("could not find step definition for %s", node.Name)
		// return createGenericStep(node.Name, xmlAttrToAttributes(node.Attributes), node.Text)
	}
}

func ExecuteFile(fileName string) bool {
	doc := xmldom.Must(xmldom.ParseFile(fileName))
	root := doc.Root

	printNode(root)
	scope := Scope{}

	if rootStep, err := GetStep(nil, root); (err != nil) && (rootStep != nil) {
		if result, err := rootStep.execute(scope); err != nil {
			fmt.Println("Program execution result is ", result)
			return result
		} else {
			fmt.Printf("Program execution error while executing step %t with error %T\n", result, err)
			return false
		}
	} else {
		fmt.Printf("Program execution error while parsing step %T with error %T\n", rootStep, err)
		return false
	}
}

func Main() {
	InitStepDefinitions()
	ExecuteFile("sampleapp.xml")
}

func printNode(node *xmldom.Node) {
	fmt.Printf("%s, %T, %s\n", node.Name, node.Attributes, node.Text)

	for _, child := range node.Children {
		printNode(child)
	}
}
