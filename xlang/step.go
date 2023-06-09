package xlang

import (
	"fmt"
	"strconv"
)

type Step interface {
	Parent(Step) Step
	Attributes() map[string]string
	Init(string, map[string]string, string) error
	AddNestedSteps(...Step)
	Execute(*Scope) error
}

type InitStepFunction func(Step, string, map[string]string, string) (Step, error)

var stepDefMap map[string]InitStepFunction = map[string]InitStepFunction{}

func RegisterStepDefinition(tag string, initStepFunction InitStepFunction) error {
	if _, ok := stepDefMap[tag]; !ok {
		stepDefMap[tag] = initStepFunction
	} else {
		return fmt.Errorf("step definition for tag %s already registered", tag)
	}
	return nil
}

func InitStepDefinitions() {
	stepDefMap["echo"] = createEchoStep
	stepDefMap["var"] = createVariableDefinitionStep
	stepDefMap["func"] = createFunctionDefinitionStep
	stepDefMap["call"] = createFunctionCallStep
	stepDefMap["return"] = createReturnStep

	stepDefMap["step"] = createCustomStepDefinitionStep
}

type StrToVarFunction func(string) (any, error)

var variableParserFunctionMap map[string]StrToVarFunction = map[string]StrToVarFunction{}

func RegisterVariableTypeDefinition(varType string, strToVarFunction StrToVarFunction) error {
	if _, ok := stepDefMap[varType]; !ok {
		variableParserFunctionMap[varType] = strToVarFunction
	} else {
		return fmt.Errorf("variable parse function definition for type %s already registered", varType)
	}
	return nil
}

func InitVariableTypeDefinitions() {
	variableParserFunctionMap["int"] = func(strValue string) (any, error) { return strconv.Atoi(strValue) }
	variableParserFunctionMap["float"] = func(strValue string) (any, error) { return strconv.ParseFloat(strValue, 64) }
	variableParserFunctionMap["bool"] = func(strValue string) (any, error) { return strconv.ParseBool(strValue) }
	variableParserFunctionMap["string"] = func(strValue string) (any, error) { return strValue, nil }
}

func RunSteps(scope *Scope, steps ...Step) error {
	for _, step := range steps {
		if step != nil {
			err := step.Execute(scope)
			if err != nil {
				if _, ok := err.(*MethodReturnError); ok {
					// fmt.Printf("Encountered return %t, %s\n", err, err.Error())
					break
				} else {
					fmt.Printf("Program execution error while executing step %t with error %s\n", step, err.Error())
				}
				return err
			}
		} else {
			return fmt.Errorf("encountered an unexpected condition where step is nil")
		}
	}
	return nil
}

type BaseStep struct {
	parent      Step
	tag         string
	attributes  map[string]string
	text        string
	nestedSteps []Step
}

func (baseStep *BaseStep) Parent(parent Step) Step {
	if parent != nil {
		baseStep.parent = parent
	}
	return baseStep.parent
}

func (baseStep *BaseStep) Attributes() map[string]string {
	return baseStep.attributes
}

func (baseStep *BaseStep) Init(tag string, attributes map[string]string, text string) error {
	baseStep.tag = tag
	baseStep.attributes = attributes
	baseStep.text = text
	return nil
}

func (baseStep *BaseStep) AddNestedSteps(steps ...Step) {
	if baseStep.nestedSteps == nil {
		baseStep.nestedSteps = []Step{}
	}
	baseStep.nestedSteps = append(baseStep.nestedSteps, steps...)
}

func (baseStep *BaseStep) Execute(scope *Scope) error {
	return RunSteps(scope, baseStep.nestedSteps...)
}

func createBaseStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	baseStep := &BaseStep{}
	return baseStep, baseStep.Init(tag, attributes, text)
}
