package xlang

import (
	"bytes"
	"encoding/gob"
	"fmt"
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

func replaceVariableInString(stringToProcess string, variableName string, variableValue any) string {
	return strings.Replace(stringToProcess, "${"+variableName+"}", fmt.Sprintf("%v", variableValue), -1)
}

func replaceVarsInString(str string, variables map[string]any, xmlattributes map[string]string) string {
	processedString := str

	for {
		var replacement bool = false
		for key, value := range variables {
			newString := replaceVariableInString(processedString, key, value)

			if processedString != newString {
				processedString = newString
				replacement = true
			}
		}

		for key, value := range xmlattributes {
			newString := strings.Replace(processedString, "@{"+key+"}", value, -1)

			if processedString != newString {
				processedString = newString
				replacement = true
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
