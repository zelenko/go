// replace special characters with dash
package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {

	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	//reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	string := "#Golang#Python$Php&Kotlin@@"
	fmt.Println(string)
	newStr := reg.ReplaceAllString(string, "-")
	fmt.Println(newStr)
}
