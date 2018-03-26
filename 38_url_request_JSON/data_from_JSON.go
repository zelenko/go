// This sample program demonstrates how to decode a JSON string.

package main

import (
	"encoding/json" // Encoding and Decoding Package
	"fmt"
)

// JSON Contains a sample String to unmarshal.
var JSON = `{
    "name":"Mark Taylor",
    "jobtitle":"Software Developer",
    "phone":{
        "home":"123-466-799",
        "office":"564-987-654"
    },
    "email":"markt@gmail.com"
}`

func main() {
	// Unmarshal the JSON string into info map variable.
	var info map[string]interface{}
	json.Unmarshal([]byte(JSON), &info)

	// Print the output from info map.
	fmt.Println(info["name"])
	fmt.Println(info["jobtitle"])
	fmt.Println(info["email"])
	fmt.Println(info["phone"].(map[string]interface{})["home"])
	fmt.Println(info["phone"].(map[string]interface{})["office"])
}
