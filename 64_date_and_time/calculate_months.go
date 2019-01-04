package main

import (
	"fmt"
	"time"
)

func main() {
	//...
	originalTime := time.Date(2011, 9, 2, 0, 0, 0, 0, time.UTC)
	fmt.Println("months:", countMonthsSince(originalTime))

	originalTime = time.Date(2018, 9, 2, 0, 0, 0, 0, time.UTC)
	fmt.Println("months:", countMonthsSince(originalTime))

	originalTime = time.Date(2017, 9, 2, 0, 0, 0, 0, time.UTC)
	fmt.Println("months:", countMonthsSince(originalTime))
}

// countMonthsSince calculates the months between now
// and the createdAtTime time.Time value passed
func countMonthsSince(createdAtTime time.Time) int {
	now := time.Now()
	months := 0
	month := createdAtTime.Month()
	for createdAtTime.Before(now) { // loop through months until current month
		createdAtTime = createdAtTime.Add(time.Hour * 24) // add one day
		nextMonth := createdAtTime.Month()                // get the month
		if nextMonth != month {                           // if not current month then increment
			months++
		}
		month = nextMonth
	}

	return months
}
