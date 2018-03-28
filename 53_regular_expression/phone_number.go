// Regular expression to validate phone number

package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "1(234)5678901x1234"
	str2 := "(+351) 282 43 50 50"
	str3 := "90191919908"
	//str4 := "555-8909"
	str4 := "a55-8909"
	str5 := "001 6867684"
	str6 := "001 6867684x1"
	str7 := "1 (234) 567-8901"
	str8 := "1-234-567-8901 ext1234"

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nPhone: %v\t:%v\n", str1, re.MatchString(str1))
	fmt.Printf("Phone: %v\t:%v\n", str2, re.MatchString(str2))
	fmt.Printf("Phone: %v\t\t:%v\n", str3, re.MatchString(str3))
	fmt.Printf("Phone: %v\t\t\t:%v\n", str4, re.MatchString(str4))
	fmt.Printf("Phone: %v\t\t:%v\n", str5, re.MatchString(str5))
	fmt.Printf("Phone: %v\t\t:%v\n", str6, re.MatchString(str6))
	fmt.Printf("Phone: %v\t\t:%v\n", str7, re.MatchString(str7))
	fmt.Printf("Phone: %v\t:%v\n", str8, re.MatchString(str8))
}
