package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	z, _ := t.Zone()
	fmt.Println("ZONE : ", z, " Time : ", t) // local time

	location, err := time.LoadLocation("EST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ZONE : ", location, " Time : ", t.In(location)) // EST

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	fmt.Println("ZONE : ", loc, " Time : ", now) // UTC

	loc, _ = time.LoadLocation("MST")
	now = time.Now().In(loc)
	fmt.Println("ZONE : ", loc, " Time : ", now) // MST

	year, week := t.ISOWeek()
	fmt.Println("week: ", week, ", Year: ", year)
}
