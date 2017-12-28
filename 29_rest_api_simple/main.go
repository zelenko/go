// Restful API in Go with httprouter

package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// Todo is the record
type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
}

// Todos is an array
var Todos []Todo

// MaxID is the latest record
var MaxID int

func main() {

	MaxID = 1
	Todos = []Todo{Todo{ID: MaxID, Description: "Explore Golang"}}

	fmt.Println("HTTP port :3000")
	r := httprouter.New()

	// methods GET, POST, PUT, PATCH and DELETE
	r.GET("/", redirect)
	r.GET("/todos", GetAllHandler)
	r.GET("/todos/:id", GetOneHandler)
	r.POST("/todos", CreateOneHandler)
	r.PUT("/todos/:id", UpdateOneHandler)
	r.DELETE("/todos/:id", DeleteHandler)

	http.ListenAndServe(":3000", r)

}

// redirect
func redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

// GetAllHandler list all records
func GetAllHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todos)
}

// CreateOneHandler creates new record, returns record MaxID
func CreateOneHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var newTodo Todo

	err := decoder.Decode(&newTodo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	MaxID++
	newTodo.ID = MaxID

	Todos = append(Todos, newTodo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MaxID)
}

// Filter returns an array of records, filtering depends on the function
func Filter(vs []Todo, f func(Todo) bool) []Todo {
	vsf := make([]Todo, 0)
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

	oneRecord := Filter(Todos, func(t Todo) bool {
		// search criteria
		return t.ID == recordID
	})

	json.NewEncoder(w).Encode(oneRecord)
}

// UpdateOneHandler updates record.  ID required.  IF does not exists, crates new record
func UpdateOneHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	recordID, _ := strconv.Atoi(ps.ByName("id"))

	Todos = Filter(Todos, func(t Todo) bool {
		// search criteria
		return t.ID != recordID
	})

	decoder := json.NewDecoder(r.Body)
	var newTodo Todo

	err := decoder.Decode(&newTodo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	newTodo.ID = recordID

	Todos = append(Todos, newTodo)

	w.WriteHeader(http.StatusNoContent)
}

// DeleteHandler deletes record.  ID required.
func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	recordID, _ := strconv.Atoi(ps.ByName("id"))

	Todos = Filter(Todos, func(t Todo) bool {
		return t.ID != recordID
	})

	w.WriteHeader(http.StatusNoContent)
}
