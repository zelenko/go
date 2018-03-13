package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"John", "Jeffeson", "jjeff"},
		{"Daniel", "Smith", "dsmith"},
		{"Andrew", "Johnson", "ajohnson"},
	}

	filename := `test3.csv`
	createCSV(filename, records)
}

func createCSV(filename string, records [][]string) {
	csvFile, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create new file", err)
	}
	defer csvFile.Close()

	// w := csv.NewWriter(os.Stdout)
	w := csv.NewWriter(csvFile)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer.
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
