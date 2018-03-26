package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	txtFile, err := os.Create("test.txt")
	if err != nil {
		log.Fatal("Cannot create new file", err)
	}
	defer txtFile.Close()

	fmt.Fprintf(txtFile, "Line 1\nLine2")
}
