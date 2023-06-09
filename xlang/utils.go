package xlang

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/subchen/go-xmldom"
)

func xmlAttrToAttributes(xmlAttributes []*xmldom.Attribute) map[string]string {
	attributesMap := map[string]string{}
	for _, xmlattr := range xmlAttributes {
		attributesMap[xmlattr.Name] = xmlattr.Value
	}
	return attributesMap
}

func replaceVariableInString(stringToProcess string, variableIdenfyingSymbol string, variableName string, variableValue string, caseInsensitive bool) string {
	toReplace := variableIdenfyingSymbol + "{" + variableName + "}"
	if caseInsensitive {
		re := regexp.MustCompile("(?i)" + toReplace)
		return re.ReplaceAllString(stringToProcess, variableValue)
	} else {
		return strings.Replace(stringToProcess, toReplace, variableValue, -1)
	}
}

func replaceVarsInString(str string, scope *Scope, xmlparentattributes map[string]string) string {
	processedString := str

	for {

		var replacement bool = false

		if scope != nil {
			variableNames := scope.getVariableNames()
			for _, key := range variableNames {
				if scopeVariable, ok := scope.getVariable(key); ok {
					newString := replaceVariableInString(processedString, "$", key, fmt.Sprintf("%v", scopeVariable), false)
					if processedString != newString {
						processedString = newString
						replacement = true
					}
				}
			}
		}

		for key, value := range xmlparentattributes {
			newString := replaceVariableInString(processedString, "@", key, value, false)

			if processedString != newString {
				processedString = newString
				replacement = true
			}
		}

		for _, envString := range os.Environ() {
			kvp := strings.SplitN(envString, "=", 2)
			if len(kvp) == 2 {
				newString := replaceVariableInString(processedString, "#", kvp[0], kvp[1], true)

				if processedString != newString {
					processedString = newString
					replacement = true
				}
			}
		}
		if !replacement {
			break
		}
	}

	return processedString
}

func DeepCopy(src, dest interface{}) (err error) {
	buf := bytes.Buffer{}
	if err = gob.NewEncoder(&buf).Encode(src); err != nil {
		return
	}
	return gob.NewDecoder(&buf).Decode(dest)
}

func ToBool(value any) bool {
	switch varType := value.(type) {
	case bool:
		return value.(bool)
	case int:
		return (value.(int) != 0)
	case float64:
		return (value.(float64) != 0)
	case float32:
		return (value.(float32) != 0)
	case string:
		return (value.(string) != "")
	default:
		_ = varType
		return value != nil
	}
}

func checkAtleastNStep(step BaseStep, expectedStepsCount int) (int, error) {
	stepCount := 0

	if step.steps != nil {
		stepCount = len(step.steps)
	}

	if stepCount < expectedStepsCount {
		return stepCount, fmt.Errorf("%s's block needs atleast %d condition", step.name, expectedStepsCount)
	}

	return stepCount, nil
}

func checkOnlyNStep(step BaseStep, expectedStepsCount int) (int, error) {
	stepCount := 0

	if step.steps != nil {
		stepCount = len(step.steps)
	}

	if stepCount != expectedStepsCount {
		return stepCount, fmt.Errorf("%s's block needs only %d step", step.name, expectedStepsCount)
	}

	return stepCount, nil
}

// func CopyStep(anyValue Step) Step {
// 	v := reflect.ValueOf(anyValue).Elem()
// 	return reflect.New(v.Type())
// }

// func CopySteps(source []Step, destination []Step) {

// 	for i, itemToCopy := range source {
// 		destination[i] = CopyStep(itemToCopy)
// 	}
// }
