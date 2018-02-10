package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mileusna/crontab"
)

// main is the entry point for the program.
func main() {

	ctab := crontab.New() // create cron table

	// MustAddJob is like AddJob but panics on wrong syntax or problems with func/args
	// use for easier initialization
	ctab.MustAddJob("* * * * *", myFunc)     // every minute
	ctab.MustAddJob("0 12 * * *", myFunc3)   // noon lauch
	ctab.MustAddJob("/2 * * * *", myFunc2, "Monday and Tuesday midnight", 123) // run every two seconds

	// fn with args
	ctab.MustAddJob("0 0 * * 1,2", myFunc2, "Monday and Tuesday midnight", 123)
	ctab.MustAddJob("*/5 * * * *", myFunc2, "every five min", 0)

	// or use AddJob if you want to test the error
	err := ctab.AddJob("0 12 1 * *", myFunc) // on 1st day of month
	if err != nil {
		log.Println(err)
		return
	}

	// all your other app code as usual, or put sleep timer for demo
	 time.Sleep(10 * time.Minute)
	 // time.Sleep(5 * time.Second)
}

func myFunc() {
	fmt.Println("Helo, world")
}

func myFunc3() {
	fmt.Println("Noon!")
}

func myFunc2(s string, n int) {
	//fmt.Println("We have params here, string", s, "and number", n)
	fmt.Println(time.Now(), "- just ticked")

}
