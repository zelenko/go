package main

import (
	"fmt"
	"os"
)

// Season used for season of the year
type Season int

// Seasons of the year
const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)

func (s Season) String() string {
	lang := os.Getenv("LANG")
	switch lang {
	case "sp":
		return [...]string{"vesna", "leto", "osen'", "zima"}[s]
	case "en":
		return [...]string{"spring", "summer", "autumn", "winner"}[s]
	default:
		return [...]string{"spring", "summer", "autumn", "winner"}[s]
	}
}

func main() {
	os.Setenv("LANG", "en")

	fmt.Println(Spring)
	fmt.Println(Summer)
	fmt.Println(Autumn)
	fmt.Println(Winter)
}

// https://programming.guide/go/iota.html
