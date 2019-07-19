package vanilla

import (
	"encoding/json"
	"regexp"
)

// JSON converts any data type to pretty JSON string
func JSON(list interface{}) string {
	result, _ := json.MarshalIndent(list, "", "  ")
	return string(result)
}

// AlphaNumeric removes special characters, keeps letters and numbers∏
func AlphaNumeric(input string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(input, "")
}
