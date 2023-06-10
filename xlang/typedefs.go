package xlang

import (
	"fmt"
	"strconv"
)

func init() {
	variableParserFunctionMap["int"] = func(strValue string) (any, error) { return strconv.Atoi(strValue) }
	variableParserFunctionMap["float"] = func(strValue string) (any, error) { return strconv.ParseFloat(strValue, 64) }
	variableParserFunctionMap["bool"] = func(strValue string) (any, error) { return strconv.ParseBool(strValue) }
	variableParserFunctionMap["string"] = func(strValue string) (any, error) { return strValue, nil }
}

type BinaryFunction func(lhs any, rhs any) (any, error)

var binaryFunctionMap map[string]BinaryFunction = map[string]BinaryFunction{}

func RegisterBinaryFunctionDefinition(operator string, binaryFunction BinaryFunction) error {
	if _, ok := stepDefMap[operator]; !ok {
		binaryFunctionMap[operator] = binaryFunction
	} else {
		return fmt.Errorf("binary comparision function definition for operator %s already registered", operator)
	}
	return nil
}

func EqualsFunction(lhs, rhs any) (any, error) { return lhs == rhs, nil }

func NotEqualsFunction(lhs, rhs any) (any, error) { return lhs != rhs, nil }

func GreaterThanFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case byte:
		return lhs.(byte) < rhs.(byte), nil
	case int:
		return lhs.(int) < rhs.(int), nil
	case float64:
		return lhs.(float64) < rhs.(float64), nil
	case float32:
		return lhs.(float32) < rhs.(float32), nil
	case string:
		return lhs.(string) < rhs.(string), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

func GreaterThanOrEqualsFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case byte:
		return lhs.(byte) <= rhs.(byte), nil
	case int:
		return lhs.(int) <= rhs.(int), nil
	case float64:
		return lhs.(float64) <= rhs.(float64), nil
	case float32:
		return lhs.(float32) <= rhs.(float32), nil
	case string:
		return lhs.(string) <= rhs.(string), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

func LesserThanFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case byte:
		return lhs.(byte) < rhs.(byte), nil
	case int:
		return lhs.(int) < rhs.(int), nil
	case float64:
		return lhs.(float64) < rhs.(float64), nil
	case float32:
		return lhs.(float32) < rhs.(float32), nil
	case string:
		return lhs.(string) < rhs.(string), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

func LesserThanOrEqualsFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case byte:
		return lhs.(byte) <= rhs.(byte), nil
	case int:
		return lhs.(int) <= rhs.(int), nil
	case float64:
		return lhs.(float64) <= rhs.(float64), nil
	case float32:
		return lhs.(float32) <= rhs.(float32), nil
	case string:
		return lhs.(string) <= rhs.(string), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

func LogicalAndFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case bool:
		return lhs.(bool) && rhs.(bool), nil
	case int:
		return (lhs.(int) != 0) && (rhs.(int) != 0), nil
	case float64:
		return (lhs.(float64) != 0) && (rhs.(float64) != 0), nil
	case float32:
		return (lhs.(float32) != 0) && (rhs.(float32) != 0), nil
	case string:
		return (lhs.(string) != "") && (rhs.(string) != ""), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

func LogicalOrFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case bool:
		return lhs.(bool) || rhs.(bool), nil
	case int:
		return (lhs.(int) != 0) || (rhs.(int) != 0), nil
	case float64:
		return (lhs.(float64) != 0) || (rhs.(float64) != 0), nil
	case float32:
		return (lhs.(float32) != 0) || (rhs.(float32) != 0), nil
	case string:
		return (lhs.(string) != "") || (rhs.(string) != ""), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

func AddFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case byte:
		return lhs.(byte) + rhs.(byte), nil
	case int:
		return lhs.(int) + rhs.(int), nil
	case float64:
		return lhs.(float64) + rhs.(float64), nil
	case float32:
		return lhs.(float32) + rhs.(float32), nil
	case string:
		return lhs.(string) + rhs.(string), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

func SubtractFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case byte:
		return lhs.(byte) - rhs.(byte), nil
	case int:
		return lhs.(int) - rhs.(int), nil
	case float64:
		return lhs.(float64) - rhs.(float64), nil
	case float32:
		return lhs.(float32) - rhs.(float32), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

func MultiplyFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case byte:
		return lhs.(byte) * rhs.(byte), nil
	case int:
		return lhs.(int) * rhs.(int), nil
	case float64:
		return lhs.(float64) * rhs.(float64), nil
	case float32:
		return lhs.(float32) * rhs.(float32), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

func DivideFunction(lhs, rhs any) (any, error) {
	switch varType := lhs.(type) {
	case byte:
		return lhs.(byte) / rhs.(byte), nil
	case int:
		return lhs.(int) / rhs.(int), nil
	case float64:
		return lhs.(float64) / rhs.(float64), nil
	case float32:
		return lhs.(float32) / rhs.(float32), nil
	default:
		return false, fmt.Errorf("comparision not implemented for type %v", varType)
	}
}

//TODO can we merge type parsing and comparision registration
