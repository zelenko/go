package main

import (
	"fmt"
	"time"
)

func main() {
	year, month, day := time.Now().Date()
	fmt.Println("Year : ", year)
	fmt.Println("Month : ", month)
	fmt.Println("Day : ", day)
	fmt.Println("DayYear:", time.Now().YearDay())

	_, week := time.Now().ISOWeek()
	fmt.Println("Week : ", week)

	hour, min, sec := time.Now().Clock()
	fmt.Println("Hour : ", hour)
	fmt.Println("Min : ", min)
	fmt.Println("Sec : ", sec)

	if month == time.November && day == 10 {
		fmt.Println("Happy Go day!")
	}
}

// https://www.socketloop.com/tutorials/golang-how-to-get-year-month-and-day/?utm_source=socketloop&utm_medium=search
