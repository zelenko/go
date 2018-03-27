package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	/*
		// Not working on Windows 10
		dateCmd := exec.Command("ver")

		dateOut, err := dateCmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println("> ver")
		fmt.Println(string(dateOut))
	*/

	// Not working on Windows 10
	grepCmd := exec.Command("ver")
	grepIn, err := grepCmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	grepOut, err := grepCmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> ver")
	fmt.Println(string(grepBytes))

}
