package main

import (
	"fmt"
	"time"
)

// defining Day and Week
const (
	Day  = 24 * time.Hour
	Week = 7 * Day
)

func main() {
	now := time.Now()

	// print weeks, from newest to oldest
	for i := 1; i < 50; i++ {
		fmt.Println(Time(now.Add(Week * (-time.Duration(i))))) // subtract week for each iteration
	}
}

// Time return as end of day as a string (example 11:59pm)
func Time(then time.Time) string {
	//then = Bod(then)
	then = Eod(then)
	return then.Format("2006-01-02_03-04-05pm")
	//return then.Format("January 2, 2006")
}

// Bod is Beginning of day
func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// Eod End of day
func Eod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Add(((24 * 60 * 60) - 1) * time.Second)
}
