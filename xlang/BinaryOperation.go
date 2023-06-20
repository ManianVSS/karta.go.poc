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

	binaryFunctionMap["+"] = AddFunction
	binaryFunctionMap["-"] = SubtractFunction
	binaryFunctionMap["*"] = MultiplyFunction
	binaryFunctionMap["/"] = DivideFunction

	stepDefMap["compare"] = createBinaryOperationStep
	stepDefMap["binaryoperation"] = createBinaryOperationStep
}

type BinaryOperation struct {
	BaseStep
	lhs       string
	operator  string
	rhs       string
	varType   string
	resultVar string
}

func (compare *BinaryOperation) InitalizeAndCheck() error {

	if lhs, ok := compare.parameters["lhs"]; ok {
		compare.lhs = lhs
	} else {
		return fmt.Errorf("lhs attribute missing")
	}

	operator, ok := compare.parameters["operator"]
	if !ok {
		operator = ""
	}

	if _, ok := binaryFunctionMap[operator]; ok {
		compare.operator = operator
	} else {
		return fmt.Errorf("binary comparision function definition not registered %s", operator)
	}

	if rhs, ok := compare.parameters["rhs"]; ok {
		compare.rhs = rhs
	} else {
		return fmt.Errorf("rhs attribute missing")
	}

	if varType, ok := compare.parameters["type"]; ok {
		compare.varType = varType
	} else {
		compare.varType = "string"
	}

	if resultVar, ok := compare.parameters["resultVar"]; ok {
		compare.resultVar = resultVar
	} else {
		compare.resultVar = ""
	}

	return nil
}

func (compare *BinaryOperation) Execute(scope *Scope) (any, error) {

	if err := compare.InitalizeAndCheck(); err != nil {
		return nil, err
	}

	var parentAttributes map[string]string
	if compare.parent != nil {
		parentAttributes = compare.parent.Parameters()
	}

	if operatorFunction, ok := binaryFunctionMap[compare.operator]; ok {
		if strToVarFunction, ok := variableParserFunctionMap[compare.varType]; ok {
			if lhsvalue, err := strToVarFunction(replaceVarsInString(compare.lhs, scope, parentAttributes)); err == nil {
				if rhsvalue, err := strToVarFunction(replaceVarsInString(compare.rhs, scope, parentAttributes)); err == nil {

					operationResult, err := operatorFunction(lhsvalue, rhsvalue)
					if (err == nil) && (compare.resultVar != "") {
						scope.variables[compare.resultVar] = operationResult
					}
					return operationResult, err
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

func createBinaryOperationStep(name string, parameters map[string]string, body string) (Step, error) {
	binaryOperation := &BinaryOperation{}
	binaryOperation.name = name
	binaryOperation.parameters = parameters
	binaryOperation.body = body
	return binaryOperation, nil
}
