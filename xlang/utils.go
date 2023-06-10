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

func replaceVarsInString(str string, variables map[string]any, xmlparentattributes map[string]string) string {
	processedString := str

	for {
		var replacement bool = false
		for key, value := range variables {
			newString := replaceVariableInString(processedString, "$", key, fmt.Sprintf("%v", value), false)

			if processedString != newString {
				processedString = newString
				replacement = true
			}
		}

		for key, value := range xmlparentattributes {
			newString := replaceVariableInString(processedString, "@", key, value, false) //strings.Replace(processedString, "@{"+key+"}", value, -1)

			if processedString != newString {
				processedString = newString
				replacement = true
			}
		}

		for _, envString := range os.Environ() {
			kvp := strings.SplitN(envString, "=", 2)
			if len(kvp) == 2 {
				newString := replaceVariableInString(processedString, "#", kvp[0], kvp[1], true) //strings.Replace(processedString, "#{"+kvp[0]+"}", kvp[1], -1)

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

// func CopyStep(anyValue Step) Step {
// 	v := reflect.ValueOf(anyValue).Elem()
// 	return reflect.New(v.Type())
// }

// func CopySteps(source []Step, destination []Step) {

// 	for i, itemToCopy := range source {
// 		destination[i] = CopyStep(itemToCopy)
// 	}
// }
