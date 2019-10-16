package main

import (
	"fmt"
	"time"
)

func main() {
	oneDay := 24 * time.Hour
	start := time.Unix(1485307713, 0)
	end := time.Unix(1516411773, 0)
	rounded := time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.UTC)

	// print all dates from start date to end date
	for rounded.Before(end) {
		fmt.Println(rounded)
		rounded = rounded.Add(oneDay)
	}
}

// Bod is Begining of Day
func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// Eod is End of Day
func Eod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).Add((-1 + (24 * 60 * 60)) * time.Second)
}
