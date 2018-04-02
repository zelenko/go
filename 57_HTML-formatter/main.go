package main

import (
	"fmt"
	"io/ioutil"

	"github.com/yosssi/gohtml"
)

func main() {
	html, err := ioutil.ReadFile("original.gohtml")
	checkErr(err, "cannot read file:")

	err = ioutil.WriteFile("original_updated.gohtml", gohtml.FormatBytes(html), 0644)
	checkErr(err, "cannot save file")
}

func checkErr(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
	}
}
