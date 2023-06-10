package xlang

import (
	"fmt"
)

func init() {
	stepDefMap["equals"] = createEqualsStep
}

type Equals struct {
	BaseStep
	lhs     string
	rhs     string
	varType string
}

func (equals *Equals) Initalize() error {

	if lhs, ok := equals.attributes["lhs"]; ok {
		equals.lhs = lhs
	} else {
		return fmt.Errorf("lhs attribute missing")
	}

	if rhs, ok := equals.attributes["rhs"]; ok {
		equals.rhs = rhs
	} else {
		return fmt.Errorf("rhs attribute missing")
	}

	if varType, ok := equals.attributes["type"]; ok {
		equals.varType = varType
	} else {
		equals.varType = "string"
	}

	return nil
}

func (equals *Equals) Execute(scope *Scope, basedir string) (any, error) {

	var parentAttributes map[string]string
	if equals.parent != nil {
		parentAttributes = equals.parent.Attributes(nil)
	}
	if strToVarFunction, ok := variableParserFunctionMap[equals.varType]; ok {
		if lhsvalue, err := strToVarFunction(replaceVarsInString(equals.lhs, scope.variables, parentAttributes)); err == nil {
			if rhsvalue, err := strToVarFunction(replaceVarsInString(equals.rhs, scope.variables, parentAttributes)); err == nil {
				return lhsvalue == rhsvalue, err
			} else {
				return false, err
			}
		} else {
			return false, err
		}
	} else {
		return false, fmt.Errorf("undefined type for value conversion %s", equals.varType)
	}
}

func createEqualsStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	equals := &Equals{}
	equals.tag = tag
	equals.attributes = attributes
	equals.text = text
	return equals, equals.Initalize()
}
