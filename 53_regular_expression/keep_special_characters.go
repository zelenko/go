//Regular expression to extract all Non-Alphanumeric Characters from a String

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "We @@@Love@@@@ #Go!$! ****Programming****Language^^^"

	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Println(re.MatchString(str1))        // true

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
