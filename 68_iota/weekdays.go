package main

import (
	"fmt"
)

// Weekday has underlying type iota
type Weekday int

// Days of the week
const (
	Sun Weekday = iota
	Mon
	Tue
	Wed
	Thu
	Fri
	Sat
)

func (day Weekday) String() string {
	return [...]string{"Sunday", "Monday", "Tue'", "Wed", "Thursday", "Friday", "Saturday"}[day]
}

// Weekend is it?
func (day Weekday) Weekend() bool {
	switch day {
	case Sun, Sat:
		return true
	default:
		return false
	}
}

func main() {
	fmt.Println(Sun)
	fmt.Println(Mon)
	fmt.Println(Wed, Wed.Weekend())
	fmt.Println(Sat, Sat.Weekend())
	fmt.Println(Sun, Sun.Weekend())
	fmt.Println(Mon, Mon.Weekend())
	fmt.Println(Fri, Fri.Weekend())
}
