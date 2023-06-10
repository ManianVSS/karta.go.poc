package xlang

import (
	"fmt"
	"path/filepath"

	"github.com/subchen/go-xmldom"
)

func init() {
	stepDefMap["app"] = createBaseStep
}

func GetStepChildrenForNode(parent Step, nodeChildren []*xmldom.Node) {
	for _, childNode := range nodeChildren {
		if childStep, err := GetStepForNode(parent, childNode); err == nil {
			parent.AddNestedSteps(childStep)
		}
	}
}

func GetStepBaseForNode(parent Step, node *xmldom.Node) (Step, error) {
	createStepFunction, ok := stepDefMap[node.Name]
	if !ok {
		return nil, fmt.Errorf("undefined step %s", node.Name)
	}
	if step, err := createStepFunction(parent, node.Name, xmlAttrToAttributes(node.Attributes), node.Text); err == nil {
		if step.Parent(nil) == nil {
			step.Parent(parent)
		}
		return step, err
	} else {
		return step, fmt.Errorf("could not parse step %s using step parsing handler function, %s", node.Name, err.Error())
	}
}

func GetStepForNode(parent Step, node *xmldom.Node) (Step, error) {
	step, err := GetStepBaseForNode(parent, node)
	if err == nil {
		for _, child := range node.Children {
			if childsNestedStep, err := GetStepForNode(step, child); err == nil {
				step.AddNestedSteps(childsNestedStep)
			}
		}
		GetStepChildrenForNode(step, node.Children)
	}
	return step, err
}

func ExecuteFile(scope *Scope, fileName string) (any, error) {
	doc := xmldom.Must(xmldom.ParseFile(fileName))
	root := doc.Root

	if scope == nil {
		return nil, fmt.Errorf("scope was not provided to run the file")
	}
	scope.variables = map[string]any{}
	scope.functions = map[string][]Step{}

	basedir := filepath.Dir(fileName)

	//Handle non App step
	if root.Name != "app" {
		if rootStep, err := GetStepForNode(nil, root); err == nil {
			return rootStep.Execute(scope, basedir)
		} else {
			return nil, err
		}
	}

	//Else parse full app : This is necessary to process custom step definition before execution
	rootStep, err := GetStepBaseForNode(nil, root)

	if err != nil {
		return nil, err
	}

	overallResult := make([]any, len(root.Children))
	for i, mainStepNode := range root.Children {
		if mainStep, err := GetStepForNode(rootStep, mainStepNode); err == nil {
			if mainStep == nil {
				return overallResult, fmt.Errorf("unexpected program parsing error; Nil root step")
			}
			result, err := mainStep.Execute(scope, basedir)
			overallResult[i] = result
			if err != nil {
				return overallResult, err
			}
		} else {
			return overallResult, err
		}
	}
	return overallResult, nil
}

func Main(mainFile string) {
	scope := &Scope{}
	scope.variables = map[string]any{}
	scope.functions = map[string][]Step{}

	if result, err := ExecuteFile(scope, mainFile); err == nil {
		fmt.Printf("Program exited with results %#v\n", result)
	} else {
		fmt.Printf("Program execution error %s\n", err.Error())
	}
}
