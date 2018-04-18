// Regular expression to extract date(YYYY-MM-DD) from string

package main

import (
	"fmt"
	"regexp"
)

// main is the entry point for the program
func main() {
	str1 := "If I am 20 years 10 months and 14 days old as of August 17,2016 then my DOB would be 1995-10-03"

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern

	fmt.Println(re.MatchString(str1)) // true

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
