package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct {
	Name, Price string
}

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromServer(url)

	tours := toursFromJson(content)

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}
}

// check if error present
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// contentFromServer returns a string
func contentFromServer(url string) string {

	resp, err := http.Get(url)
	check(err)

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	check(err)

	return string(bytes)
}

// toursFromJson parses json data into struct
func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	check(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		check(err)
		tours = append(tours, tour)
	}

	return tours
}
