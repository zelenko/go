package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Records is an array
var Records []Item

// MaxID is the latest record
var MaxID int

// Item is the record structure
type Item struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

func main() {

	MaxID = 1
	Records = []Item{{ID: MaxID, Description: "Explore Golang"}}

	r := mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	//r.NotFound = http.FileServer(http.Dir("public"))

	r.Handle("/", redirect).Methods("GET")
	r.Handle("/todos", GetAllHandler).Methods("GET")
	r.Handle("/todos/{id}", GetOneHandler).Methods("GET")
	r.Handle("/todos", CreateOneHandler).Methods("POST")
	r.Handle("/todos/{id}", UpdateOneHandler).Methods("PUT")
	r.Handle("/todos/{id}", DeleteHandler).Methods("DELETE")

	// specify port here
	fmt.Println("HTTP port :3000")
	http.ListenAndServe(":3000", r)

}

// redirect
var redirect = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
})

// GetAllHandler list all records
var GetAllHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Records)
})

// CreateOneHandler creates new record, returns record MaxID
var CreateOneHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MaxID)
})

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
var GetOneHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recordID, _ := strconv.Atoi(params["id"])

	oneRecord := Filter(Records, func(t Item) bool {
		// search criteria
		return t.ID == recordID
	})

	json.NewEncoder(w).Encode(oneRecord)
})

// UpdateOneHandler updates record.  ID required.  IF does not exists, crates new record
var UpdateOneHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recordID, _ := strconv.Atoi(params["id"])

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

	w.WriteHeader(http.StatusNoContent)
})

// DeleteHandler deletes record.  ID required.
var DeleteHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	recordID, _ := strconv.Atoi(params["id"])

	Records = Filter(Records, func(t Item) bool {
		return t.ID != recordID
	})

	w.WriteHeader(http.StatusNoContent)
})
