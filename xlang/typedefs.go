package xlang

import "strconv"

func init() {
	variableParserFunctionMap["int"] = func(strValue string) (any, error) { return strconv.Atoi(strValue) }
	variableParserFunctionMap["float"] = func(strValue string) (any, error) { return strconv.ParseFloat(strValue, 64) }
	variableParserFunctionMap["bool"] = func(strValue string) (any, error) { return strconv.ParseBool(strValue) }
	variableParserFunctionMap["string"] = func(strValue string) (any, error) { return strValue, nil }
}
