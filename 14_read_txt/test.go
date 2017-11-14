// simplistic example of reading a text file

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileHandle, _ := os.Open("foo.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	i := 1
	for fileScanner.Scan() {
		fmt.Println(i, fileScanner.Text())
		i++
	}
}