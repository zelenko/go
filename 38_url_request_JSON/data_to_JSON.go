// How to encode map data into a JSON string?

package main

import (
	"encoding/json" // Encoding and Decoding Package
	"fmt"
)

func main() {
	// Create a map of key/value pairs and parses the data into JSON
	emp := make(map[string]interface{})
	emp["name"] = "John Doe"
	emp["age"] = "30"
	emp["phone"] = map[string]interface{}{
		"mobil":   "555-555-5555",
		"fax": "564-987-654",
	}
	emp["cars"] = map[string]interface{}{
	"car1":"Ford",
	"car2":"BMW",
	"car3":"Fiat",
	}
	emp["email"] = "noreply@gmail.com"

	// Marshal the map into a JSON string.
	empData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	jsonStr := string(empData)
	fmt.Println("The JSON data is:")
	fmt.Println(jsonStr)

}
