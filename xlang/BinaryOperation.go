package xlang

import (
	"fmt"
)

func init() {
	binaryFunctionMap[""] = EqualsFunction
	binaryFunctionMap["="] = EqualsFunction
	binaryFunctionMap["!="] = NotEqualsFunction
	binaryFunctionMap["<"] = GreaterThanFunction
	binaryFunctionMap["<="] = GreaterThanOrEqualsFunction
	binaryFunctionMap[">"] = LesserThanFunction
	binaryFunctionMap[">="] = LesserThanOrEqualsFunction

	binaryFunctionMap["and"] = LogicalAndFunction
	binaryFunctionMap["or"] = LogicalOrFunction

	stepDefMap["compare"] = createBinaryOperationStep
	stepDefMap["binaryoperation"] = createBinaryOperationStep
}

type BinaryOperation struct {
	BaseStep
	lhs      string
	operator string
	rhs      string
	varType  string
}

func (compare *BinaryOperation) Initalize() error {

	if lhs, ok := compare.attributes["lhs"]; ok {
		compare.lhs = lhs
	} else {
		return fmt.Errorf("lhs attribute missing")
	}

	operator, ok := compare.attributes["operator"]
	if !ok {
		operator = ""
	}

	if _, ok := binaryFunctionMap[operator]; ok {
		compare.operator = operator
	} else {
		return fmt.Errorf("binary comparision function definition not registered %s", operator)
	}

	if rhs, ok := compare.attributes["rhs"]; ok {
		compare.rhs = rhs
	} else {
		return fmt.Errorf("rhs attribute missing")
	}

	if varType, ok := compare.attributes["type"]; ok {
		compare.varType = varType
	} else {
		compare.varType = "string"
	}

	return nil
}

func (compare *BinaryOperation) Execute(scope *Scope, basedir string) (any, error) {

	var parentAttributes map[string]string
	if compare.parent != nil {
		parentAttributes = compare.parent.Attributes(nil)
	}

	if operatorFunction, ok := binaryFunctionMap[compare.operator]; ok {
		if strToVarFunction, ok := variableParserFunctionMap[compare.varType]; ok {
			if lhsvalue, err := strToVarFunction(replaceVarsInString(compare.lhs, scope.variables, parentAttributes)); err == nil {
				if rhsvalue, err := strToVarFunction(replaceVarsInString(compare.rhs, scope.variables, parentAttributes)); err == nil {
					return operatorFunction(lhsvalue, rhsvalue)
				} else {
					return false, err
				}
			} else {
				return false, err
			}
		} else {
			return false, fmt.Errorf("undefined type for value conversion and comparision %s", compare.varType)
		}
	} else {
		return false, fmt.Errorf("binary comparision function definition not registered %s", compare.operator)
	}
}

func createBinaryOperationStep(parent Step, tag string, attributes map[string]string, text string) (Step, error) {
	binaryOperation := &BinaryOperation{}
	binaryOperation.tag = tag
	binaryOperation.attributes = attributes
	binaryOperation.text = text
	return binaryOperation, binaryOperation.Initalize()
}
