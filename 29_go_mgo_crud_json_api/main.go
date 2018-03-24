// Restful API in Go with httprouter

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Item is the record
type Item struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

// Records is an array
var Records []Item

// MaxID is the latest record
var MaxID int

// main is the entry point for the program.
func main() {

	MaxID = 1
	Records = []Item{{ID: MaxID, Description: "Explore Golang"}}

	fmt.Println("HTTP port :80")
	r := httprouter.New()

	// methods GET, POST, PUT, PATCH and DELETE
	r.GET("/todos", GetAllHandler)
	r.GET("/todos/:id", GetOneHandler)
	r.POST("/todos", CreateOneHandler)
	r.PUT("/todos/:id", UpdateOneHandler)
	r.DELETE("/todos/:id", DeleteHandler)

	r.GET("/tab", tabDelimited)
	r.NotFound = http.FileServer(http.Dir("public"))

	err := http.ListenAndServe(":80", r)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// GetAllHandler list all records
func GetAllHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Records) // return Records to writer
}

// Filter returns an array of records, filtering depends on the function
func Filter(vs []Item, f func(Item) bool) []Item {
	vsf := make([]Item, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// GetOneHandler return one record
func GetOneHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	recordID, _ := strconv.Atoi(ps.ByName("id"))

	oneRecord := Filter(Records, func(t Item) bool {
		// search criteria
		return t.ID == recordID
	})

	json.NewEncoder(w).Encode(oneRecord) // return oneRecord to writer
}

// DeleteHandler deletes record.  ID required.
func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	recordID, _ := strconv.Atoi(ps.ByName("id"))

	// Recreates the slice with one record omitted.
	Records = Filter(Records, func(t Item) bool {
		return t.ID != recordID
	})

	w.WriteHeader(http.StatusNoContent)
}

// CreateOneHandler creates new record, returns record MaxID
func CreateOneHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	decoder := json.NewDecoder(r.Body)
	var newTodo Item

	err := decoder.Decode(&newTodo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	MaxID++
	newTodo.ID = MaxID

	Records = append(Records, newTodo)

	// sort
	sort.Slice(Records, func(i, j int) bool { return Records[i].ID < Records[j].ID })

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MaxID) // return MaxID to writer
}

// UpdateOneHandler updates (replaces) record.  ID required.  IF does not exists, crates new record
func UpdateOneHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	recordID, _ := strconv.Atoi(ps.ByName("id"))

	// Create a new slice, without one with "recordID".
	Records = Filter(Records, func(t Item) bool {
		// search criteria
		return t.ID != recordID
	})

	decoder := json.NewDecoder(r.Body)
	var newTodo Item

	err := decoder.Decode(&newTodo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	newTodo.ID = recordID

	Records = append(Records, newTodo)

	// sort
	sort.Slice(Records, func(i, j int) bool { return Records[i].ID < Records[j].ID })

	// return results
	w.WriteHeader(http.StatusNoContent) // returns just status, nothing else
}

// tabDelimited txt file
func tabDelimited(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	buffer := &bytes.Buffer{} // creates IO Writer

	// add title as first row of txt file
	buffer.WriteString("id\tdescription\tcomplete\n")

	// write each record (line) to buffer
	for _, record := range Records {

		// create new line
		line := fmt.Sprint(strconv.Itoa(record.ID) + "\t" + record.Description + "\t" +
			strconv.FormatBool(record.Complete) + "\n")

		// write the line to buffer
		if _, err := buffer.WriteString(line); err != nil {
			log.Fatalln("Error writing to txt file:", err)
		}
	}

	filename := time.Now().Format("2006-03-02_03-04-05pm") + ".txt"

	w.Header().Set("Content-Type", "application/csv-tab-delimited-table")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)
	w.Write(buffer.Bytes()) // respond to request with buffer data

}
