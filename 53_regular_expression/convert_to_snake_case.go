package main

import (
	"fmt"
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func main() {
	fmt.Println(ToSnakeCase("JapanCanadaAustralia"))
	fmt.Println(ToSnakeCase("JapanCanadaAUSTRALIA"))
	fmt.Println(ToSnakeCase("JAPANCanadaAUSTRALIA"))
	fmt.Println(ToSnakeCase("Japan125Canada130Australia150"))
}
