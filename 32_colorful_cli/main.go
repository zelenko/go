package main

import (
	"fmt"
	"github.com/fatih/color"
)

func main() {
	minion := color.New(color.FgBlue).Add(color.BgYellow).Add(color.Bold)
	minion.Println("Minion says: banana!!!!!!")

	m := minion.PrintlnFunc()
	m("I want another banana!!!!!")

	//slantedRed := color.New(color.FgRed, color.BgWhite, color.Italic).SprintFunc()
	slantedRed := color.New(color.FgRed, color.BgWhite, color.Bold).SprintFunc()
	fmt.Println("I've made a huge", slantedRed("mistake"))

	color.Red("Roses are red")
	color.Blue("Violets are blue")

	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")

	// Use handy standard colors
	color.Set(color.FgYellow)

	fmt.Println("Existing text will now be in yellow")
	fmt.Printf("This one %s\n", "too")

	color.Unset() // Don't forget to unset

	// You can mix up parameters
	color.Set(color.FgMagenta, color.Bold)
	fmt.Println("All text will now be bold magenta.")
	color.Unset() // Use it in your function

	red := color.New(color.FgRed)
	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")

}
