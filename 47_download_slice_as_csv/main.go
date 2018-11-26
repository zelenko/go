package main

import (
	"bytes"
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

var tmpl = `<html>
<head>
    <title>{{ . }}</title>
</head>
<body>
    {{ . }}
    <p>
      <a href="/">main</a> |
	  <a href="/csv">csv</a> |
	  <a href="/csv2">csv 2</a> |
	  <a href="/tab-delimited">Tab Delimited</a> |
	  <a href="/csv4">csv 4</a> |
    </p>

</body>
</html>
`

// main is the entry point for the program.
func main() {
	server := http.Server{
		Addr: ":80",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/csv", TestCSV)
	http.HandleFunc("/csv2", csv2)
	http.HandleFunc("/tab-delimited", tabDelimited)
	http.HandleFunc("/csv4", csv4)
	server.ListenAndServe()
}

// index shows the main page
func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("main") //name of the template is main
	t, _ = t.Parse(tmpl)      // parsing of template string
	t.Execute(w, "Main page")
}

// TestCSV function
func TestCSV(w http.ResponseWriter, r *http.Request) {

	record := []string{"test1", "test2", "test3"} // just some test data to use for the wr.Writer() method below.

	buffer := &bytes.Buffer{}            // creates IO Writer
	writerToCSV := csv.NewWriter(buffer) // creates a csv writer that uses the io buffer.
	for i := 0; i < 100; i++ {           // make a loop for 100 rows just for testing purposes
		if err := writerToCSV.Write(record); err != nil { // converts array of string to comma separated values for 1 row.
			log.Fatalln("error writing record to csv:", err)
		}
	}
	writerToCSV.Flush() // writes the csv writer data to the buffered data io writer(b(bytes.buffer))

	filename := time.Now().Format("2006-01-02_03-04-05pm") + ".csv"

	w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)
	w.Write(buffer.Bytes())

}

// csv2 function writes directly to http.ResponseWriter
// no buffer here
func csv2(w http.ResponseWriter, r *http.Request) {

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"John", "Jeffeson", "jjeff"},
		{"Daniel", "Smith", "dsmith"},
		{"Andrew", "Johnson", "ajohnson"},
	}

	filename := time.Now().Format("2006-01-02_03-04-05pm") + ".csv"

	w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)

	writer := csv.NewWriter(w) // writes to http.ResponseWriter

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (http.ResponseWriter).
	writer.Flush()

	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}

}

// tabDelimited txt file
func tabDelimited(w http.ResponseWriter, r *http.Request) {

	// These are the records to write
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"John", "Jeffeson", "jjeff"},
		{"Daniel", "Smith", "dsmith"},
		{"Andrew", "Johnson", "ajohnson"},
	}

	buffer := &bytes.Buffer{} // creates IO Writer

	// write each record (line) to buffer
	for _, record := range records {

		// create new line
		line := strings.Join(record, "\t") + "\n"

		// write the line to buffer
		if _, err := buffer.WriteString(line); err != nil {
			log.Fatalln("Error writing to txt file:", err)
		}
	}

	filename := time.Now().Format("2006-01-02_03-04-05pm") + ".txt"

	// Header
	w.Header().Set("Content-Type", "application/csv-tab-delimited-table")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)

	// return file from buffer
	w.Write(buffer.Bytes()) // respond to request with buffer data

}

// csv4 function, writer writes to buffer, then to file
func csv4(w http.ResponseWriter, r *http.Request) {

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"John", "Jeffeson", "jjeff"},
		{"Daniel", "Smith", "dsmith"},
		{"Andrew", "Johnson", "ajohnson"},
	}

	//buffer := bytes.NewBuffer(make([]byte, 0))
	buffer := &bytes.Buffer{}       // creates IO Writer
	writer := csv.NewWriter(buffer) // creates a csv writer that uses the io buffer.

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	writer.Flush() // writes the csv writer data to the buffered data io writer(b(bytes.buffer))

	filename := time.Now().Format("2006-01-02_03-04-05pm") + ".csv"

	w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)
	w.Write(buffer.Bytes())

}
