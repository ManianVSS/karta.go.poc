package xlang

import (
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

func replaceVariablesInString(str string, variables map[string]any) string {
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
		if !replacement {
			break
		}
	}

	return processedString
}
