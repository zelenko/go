package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"

	resp, err := http.Get(url)
	check(err)

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	check(err)

	fmt.Print(string(bytes))

}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
