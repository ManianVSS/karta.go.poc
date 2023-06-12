package xlang

// type TypeHandler struct {
// 	parseFunction               ParseFunction
// 	toStringFunction            ToStringFunction
// 	toBoolFunction              UnaryBooleanFunction
// 	notFunction                 UnaryBooleanFunction
// 	equalsFunction              BinaryBooleanFunction
// 	notEqualsFunction           BinaryBooleanFunction
// 	andFunction                 BinaryBooleanFunction
// 	orFunction                  BinaryBooleanFunction
// 	greaterThanFunction         BinaryBooleanFunction
// 	greaterThanOrEqualsFunction BinaryBooleanFunction
// 	lesserThanFunction          BinaryBooleanFunction
// 	lesserThanOrEqualsFunction  BinaryBooleanFunction
// 	incrementFunction           UnaryFunction
// 	decrementFunction           UnaryFunction
// 	unaryMinusFunction          UnaryFunction
// 	addFunction                 BinaryFunction
// 	subtractFunction            BinaryFunction
// 	multiplyFunction            BinaryFunction
// 	divideFunction              BinaryFunction
// 	modulusFunction             BinaryFunction
// 	exponentFunction            BinaryFunction
// }

// var typeHandlerDefinitionMap map[string]TypeHandler = map[string]TypeHandler{}

// func RegisterTypeHandlerDefinition(varType string, typeHandler TypeHandler) error {
// 	if _, ok := typeHandlerDefinitionMap[varType]; !ok {
// 		typeHandlerDefinitionMap[varType] = typeHandler
// 	} else {
// 		return fmt.Errorf("type handler definition for type %s already registered", varType)
// 	}
// 	return nil
// }

// func init() {
// 	genericToStringFunction := func(value any) (string, error) { return fmt.Sprintf("%v", value), nil }
// 	genericEqualsFunction := func(lhs, rhs any) (bool, error) { return lhs == rhs, nil }
// 	genericNotEqualsFunction := func(lhs, rhs any) (bool, error) { return lhs != rhs, nil }

// 	genericUnspportedUnaryFunction := func(value any) (any, error) { return false, fmt.Errorf("unsupported operation for value %v", value) }

// 	genericUnsupportedBinaryBooleanOperation := func(lhs, rhs any) (bool, error) {
// 		return false, fmt.Errorf("unsupported operation for values %v and %v", lhs, rhs)
// 	}

// 	genericUnsupportedBinaryOperation := func(lhs, rhs any) (any, error) {
// 		return nil, fmt.Errorf("unsupported operation for values %v and %v", lhs, rhs)
// 	}

// 	intTypeHandler := TypeHandler{
// 		parseFunction:               func(strValue string) (any, error) { return strconv.Atoi(strValue) },
// 		toStringFunction:            genericToStringFunction,
// 		toBoolFunction:              func(value any) (bool, error) { return value.(int) != 0, nil },
// 		notFunction:                 func(value any) (bool, error) { return value.(int) == 0, nil },
// 		equalsFunction:              genericEqualsFunction,
// 		notEqualsFunction:           genericNotEqualsFunction,
// 		andFunction:                 func(lhs, rhs any) (bool, error) { return (lhs.(int) != 0) && (rhs.(int) != 0), nil },
// 		orFunction:                  func(lhs, rhs any) (bool, error) { return (lhs.(int) != 0) || (rhs.(int) != 0), nil },
// 		greaterThanFunction:         func(lhs, rhs any) (bool, error) { return lhs.(int) > rhs.(int), nil },
// 		greaterThanOrEqualsFunction: func(lhs, rhs any) (bool, error) { return lhs.(int) >= rhs.(int), nil },
// 		lesserThanFunction:          func(lhs, rhs any) (bool, error) { return lhs.(int) < rhs.(int), nil },
// 		lesserThanOrEqualsFunction:  func(lhs, rhs any) (bool, error) { return lhs.(int) <= rhs.(int), nil },
// 		incrementFunction:           func(value any) (any, error) { return value.(int) + 1, nil },
// 		decrementFunction:           func(value any) (any, error) { return value.(int) - 1, nil },
// 		unaryMinusFunction:          func(value any) (any, error) { return -value.(int), nil },
// 		addFunction:                 func(lhs, rhs any) (any, error) { return lhs.(int) + rhs.(int), nil },
// 		subtractFunction:            func(lhs, rhs any) (any, error) { return lhs.(int) - rhs.(int), nil },
// 		multiplyFunction:            func(lhs, rhs any) (any, error) { return lhs.(int) * rhs.(int), nil },
// 		divideFunction:              func(lhs, rhs any) (any, error) { return lhs.(int) / rhs.(int), nil },
// 		modulusFunction:             func(lhs, rhs any) (any, error) { return lhs.(int) % rhs.(int), nil },
// 		exponentFunction:            func(lhs, rhs any) (any, error) { return int(math.Pow(float64(lhs.(int)), float64(rhs.(int)))), nil },
// 	}

// 	// byteTypeHandler := TypeHandler{
// 	// 	parseFunction: func(strValue string) (any, error) {
// 	// 		if parsedInt, err := strconv.ParseInt(strValue, 10, 8); err == nil {
// 	// 			return byte(parsedInt), err
// 	// 		} else {
// 	// 			return parsedInt, err
// 	// 		}
// 	// 	},
// 	// 	toStringFunction:            genericToStringFunction,
// 	// 	toBoolFunction:              func(value any) (bool, error) { return value.(byte) != 0, nil },
// 	// 	notFunction:                 func(value any) (bool, error) { return value.(byte) == 0, nil },
// 	// 	equalsFunction:              genericEqualsFunction,
// 	// 	notEqualsFunction:           genericNotEqualsFunction,
// 	//  andFunction:                 func(lhs, rhs any) (bool, error) { return (lhs.(byte) != 0) && (rhs.(byte) != 0), nil },
// 	//  orFunction:                  func(lhs, rhs any) (bool, error) { return (lhs.(byte) != 0) || (rhs.(byte) != 0), nil },
// 	// 	greaterThanFunction:         func(lhs, rhs any) (bool, error) { return lhs.(byte) > rhs.(byte), nil },
// 	// 	greaterThanOrEqualsFunction: func(lhs, rhs any) (bool, error) { return lhs.(byte) >= rhs.(byte), nil },
// 	// 	lesserThanFunction:          func(lhs, rhs any) (bool, error) { return lhs.(byte) < rhs.(byte), nil },
// 	// 	lesserThanOrEqualsFunction:  func(lhs, rhs any) (bool, error) { return lhs.(byte) <= rhs.(byte), nil },
// 	//  incrementFunction:           func(value any) (any, error) { return value.(byte) + 1, nil },
// 	//  decrementFunction:           func(value any) (any, error) { return value.(byte) - 1, nil },
// 	//  unaryMinusFunction:          func(value any) (any, error) { return -value.(byte), nil },
// 	// 	addFunction:                 func(lhs, rhs any) (any, error) { return lhs.(byte) + rhs.(byte), nil },
// 	// 	subtractFunction:            func(lhs, rhs any) (any, error) { return lhs.(byte) - rhs.(byte), nil },
// 	// 	multiplyFunction:            func(lhs, rhs any) (any, error) { return lhs.(byte) * rhs.(byte), nil },
// 	// 	divideFunction:              func(lhs, rhs any) (any, error) { return lhs.(byte) / rhs.(byte), nil },
// 	// 	modulusFunction:             func(lhs, rhs any) (any, error) { return lhs.(byte) % rhs.(byte), nil },
// 	// 	exponentFunction:            genericUnsupportedBinaryOperation,
// 	// }

// 	float64TypeHandler := TypeHandler{
// 		parseFunction:               func(strValue string) (any, error) { return strconv.Atoi(strValue) },
// 		toStringFunction:            genericToStringFunction,
// 		toBoolFunction:              func(value any) (bool, error) { return value.(float64) != 0, nil },
// 		notFunction:                 func(value any) (bool, error) { return value.(float64) == 0, nil },
// 		equalsFunction:              genericEqualsFunction,
// 		notEqualsFunction:           genericNotEqualsFunction,
// 		andFunction:                 func(lhs, rhs any) (bool, error) { return (lhs.(float64) != 0) && (rhs.(float64) != 0), nil },
// 		orFunction:                  func(lhs, rhs any) (bool, error) { return (lhs.(float64) != 0) || (rhs.(float64) != 0), nil },
// 		greaterThanFunction:         func(lhs, rhs any) (bool, error) { return lhs.(float64) > rhs.(float64), nil },
// 		greaterThanOrEqualsFunction: func(lhs, rhs any) (bool, error) { return lhs.(float64) >= rhs.(float64), nil },
// 		lesserThanFunction:          func(lhs, rhs any) (bool, error) { return lhs.(float64) < rhs.(float64), nil },
// 		lesserThanOrEqualsFunction:  func(lhs, rhs any) (bool, error) { return lhs.(float64) <= rhs.(float64), nil },
// 		incrementFunction:           func(value any) (any, error) { return value.(float64) + 1, nil },
// 		decrementFunction:           func(value any) (any, error) { return value.(float64) - 1, nil },
// 		unaryMinusFunction:          func(value any) (any, error) { return -value.(float64), nil },
// 		addFunction:                 func(lhs, rhs any) (any, error) { return lhs.(float64) + rhs.(float64), nil },
// 		subtractFunction:            func(lhs, rhs any) (any, error) { return lhs.(float64) - rhs.(float64), nil },
// 		multiplyFunction:            func(lhs, rhs any) (any, error) { return lhs.(float64) * rhs.(float64), nil },
// 		divideFunction:              func(lhs, rhs any) (any, error) { return lhs.(float64) / rhs.(float64), nil },
// 		modulusFunction:             genericUnsupportedBinaryOperation,
// 		exponentFunction:            func(lhs, rhs any) (any, error) { return math.Pow(lhs.(float64), rhs.(float64)), nil },
// 	}

// 	boolTypeHandler := TypeHandler{
// 		parseFunction:               func(strValue string) (any, error) { return strconv.Atoi(strValue) },
// 		toStringFunction:            genericToStringFunction,
// 		toBoolFunction:              func(value any) (bool, error) { return value.(bool), nil },
// 		notFunction:                 func(value any) (bool, error) { return !value.(bool), nil },
// 		equalsFunction:              genericEqualsFunction,
// 		notEqualsFunction:           genericNotEqualsFunction,
// 		andFunction:                 func(lhs, rhs any) (bool, error) { return lhs.(bool) && rhs.(bool), nil },
// 		orFunction:                  func(lhs, rhs any) (bool, error) { return lhs.(bool) || rhs.(bool), nil },
// 		greaterThanFunction:         genericUnsupportedBinaryBooleanOperation,
// 		greaterThanOrEqualsFunction: genericUnsupportedBinaryBooleanOperation,
// 		lesserThanFunction:          genericUnsupportedBinaryBooleanOperation,
// 		lesserThanOrEqualsFunction:  genericUnsupportedBinaryBooleanOperation,
// 		incrementFunction:           genericUnspportedUnaryFunction,
// 		decrementFunction:           genericUnspportedUnaryFunction,
// 		unaryMinusFunction:          genericUnspportedUnaryFunction,
// 		addFunction:                 func(lhs, rhs any) (any, error) { return lhs.(bool) || rhs.(bool), nil },
// 		subtractFunction:            genericUnsupportedBinaryOperation,
// 		multiplyFunction:            func(lhs, rhs any) (any, error) { return lhs.(bool) && rhs.(bool), nil },
// 		divideFunction:              genericUnsupportedBinaryOperation,
// 		modulusFunction:             genericUnsupportedBinaryOperation,
// 		exponentFunction:            genericUnsupportedBinaryOperation,
// 	}

// 	stringTypeHandler := TypeHandler{
// 		parseFunction:               func(strValue string) (any, error) { return fmt.Sprintf("%v", strValue), nil },
// 		toStringFunction:            genericToStringFunction,
// 		toBoolFunction:              func(value any) (bool, error) { return value.(string) != "", nil },
// 		notFunction:                 func(value any) (bool, error) { return value.(string) == "", nil },
// 		equalsFunction:              genericEqualsFunction,
// 		notEqualsFunction:           genericNotEqualsFunction,
// 		andFunction:                 func(lhs, rhs any) (bool, error) { return (lhs.(string) != "") && (rhs.(string) != ""), nil },
// 		orFunction:                  func(lhs, rhs any) (bool, error) { return (lhs.(string) != "") || (rhs.(string) != ""), nil },
// 		greaterThanFunction:         func(lhs, rhs any) (bool, error) { return lhs.(string) > fmt.Sprintf("%v", rhs), nil },
// 		greaterThanOrEqualsFunction: func(lhs, rhs any) (bool, error) { return lhs.(string) >= fmt.Sprintf("%v", rhs), nil },
// 		lesserThanFunction:          func(lhs, rhs any) (bool, error) { return lhs.(string) < fmt.Sprintf("%v", rhs), nil },
// 		lesserThanOrEqualsFunction:  func(lhs, rhs any) (bool, error) { return lhs.(string) <= fmt.Sprintf("%v", rhs), nil },
// 		incrementFunction:           genericUnspportedUnaryFunction, //TODO: Add pattern increment
// 		decrementFunction:           genericUnspportedUnaryFunction, //TODO: Add pattern decrement
// 		unaryMinusFunction:          genericUnspportedUnaryFunction,
// 		addFunction:                 func(lhs, rhs any) (any, error) { return lhs.(string) + fmt.Sprintf("%v", rhs), nil },
// 		subtractFunction:            genericUnsupportedBinaryOperation,
// 		multiplyFunction:            genericUnsupportedBinaryOperation,
// 		divideFunction:              genericUnsupportedBinaryOperation,
// 		modulusFunction:             genericUnsupportedBinaryOperation,
// 		exponentFunction: func(lhs, rhs any) (any, error) {
// 			switch rhs.(type) {
// 			case int:
// 				return strings.Repeat(lhs.(string), rhs.(int)), nil
// 			default:
// 				return lhs, fmt.Errorf("unsupported operation for values %v and %v", lhs, rhs)
// 			}
// 		},
// 	}

// 	typeHandlerDefinitionMap["int"] = intTypeHandler
// 	// typeHandlerDefinitionMap["byte"] = byteTypeHandler
// 	typeHandlerDefinitionMap["float"] = float64TypeHandler
// 	typeHandlerDefinitionMap["bool"] = boolTypeHandler
// 	typeHandlerDefinitionMap["string"] = stringTypeHandler
// }

// func parseTypeValue(varType string, value string) (any, error) {
// 	if typeHandler, ok := typeHandlerDefinitionMap[varType]; ok {
// 		return typeHandler.parseFunction(value)
// 	} else {
// 		return nil, fmt.Errorf("type handler not registered for type %s", varType)
// 	}
// }

// func typeToString(varType string, value any) (string, error) {
// 	if typeHandler, ok := typeHandlerDefinitionMap[varType]; ok {
// 		return typeHandler.toStringFunction(value)
// 	} else {
// 		return "", fmt.Errorf("type handler not registered for type %s", varType)
// 	}
// }

// func performUnaryOperation(varType string, operator string, value any) (any, error) {
// 	if typeHandler, ok := typeHandlerDefinitionMap[varType]; ok {
// 		switch operator {
// 		case "!":
// 			return typeHandler.notFunction(value)
// 		case "++":
// 			return typeHandler.incrementFunction(value)
// 		case "--":
// 			return typeHandler.decrementFunction(value)
// 		case "-":
// 			return typeHandler.unaryMinusFunction(value)
// 		default:
// 			return nil, fmt.Errorf("unary operator not registered %s", operator)
// 		}
// 	} else {
// 		return nil, fmt.Errorf("type handler not registered for type %s", varType)
// 	}
// }

// func performBinaryOperation(varType string, operator string, lhs any, rhs any) (any, error) {

// 	if typeHandler, ok := typeHandlerDefinitionMap[varType]; ok {
// 		switch operator {
// 		//Comparative functions
// 		case "=":
// 			return typeHandler.equalsFunction(lhs, rhs)
// 		case "!=":
// 			return typeHandler.notEqualsFunction(lhs, rhs)
// 		case ">":
// 			return typeHandler.greaterThanFunction(lhs, rhs)
// 		case ">=":
// 			return typeHandler.greaterThanOrEqualsFunction(lhs, rhs)
// 		case "<":
// 			return typeHandler.lesserThanFunction(lhs, rhs)
// 		case "<=":
// 			return typeHandler.lesserThanOrEqualsFunction(lhs, rhs)

// 		//Logical Function
// 		case "and":
// 			return typeHandler.andFunction(lhs, rhs)
// 		case "or":
// 			return typeHandler.orFunction(lhs, rhs)

// 		//Arithematic Functions
// 		case "+":
// 			return typeHandler.addFunction(lhs, rhs)
// 		case "-":
// 			return typeHandler.subtractFunction(lhs, rhs)
// 		case "*":
// 			return typeHandler.multiplyFunction(lhs, rhs)
// 		case "/":
// 			return typeHandler.divideFunction(lhs, rhs)
// 		case "%":
// 			return typeHandler.modulusFunction(lhs, rhs)
// 		case "^":
// 			return typeHandler.exponentFunction(lhs, rhs)
// 		default:
// 			return nil, fmt.Errorf("binary operator not registered %s", operator)
// 		}
// 	} else {
// 		return nil, fmt.Errorf("type handler not registered for type %s", varType)
// 	}
// }
