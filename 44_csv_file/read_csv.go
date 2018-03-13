package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	filename := `test2.csv`
	readCSV(filename)

}

func readCSV(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	//r := csv.NewReader(strings.NewReader(in))
	r := csv.NewReader(f)
	r.Comma = ','
	r.Comment = '#'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
