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
