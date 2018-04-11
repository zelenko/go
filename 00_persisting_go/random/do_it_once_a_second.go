// Do something so often until condition is met
package main

import (
	"fmt"
	"time"
)

const dotsPerLine = 50
const linesPerPage = 20

// main is the entry point for the program.
func main() {

	i, line := 0, 1
	fmt.Print("\n", line, "\t ") // print the first line

	// do it once in two milliseconds
	for range time.Tick(time.Millisecond * 2) {

		fmt.Print(".")
		i++

		// At end of the line do this:
		if i%dotsPerLine == 0 {

			// end loop if all lines are printed
			if line >= linesPerPage {
				break
			}

			i = 0 // reset dot counter

			// start a new line
			line++
			fmt.Print("\n", line, "\t ")
		}
	}
}
