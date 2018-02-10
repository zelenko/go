// Restful API in Go with httprouter

package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"io/ioutil"
	"os"
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

func main() {

	MaxID = 1
	Records = []Item{{ID: MaxID, Description: "Explore Golang"}}

	fmt.Println("HTTP port :80")
	r := httprouter.New()

	// methods GET, POST, PUT, PATCH and DELETE
	//r.GET("/", redirect)
	r.GET("/todos", GetAllHandler)
	r.GET("/todos/:id", GetOneHandler)
	r.POST("/todos", CreateOneHandler)
	r.POST("/test", CreateTest)
	r.PUT("/todos/:id", UpdateOneHandler)
	r.DELETE("/todos/:id", DeleteHandler)
	r.NotFound = http.FileServer(http.Dir("public"))

	http.ListenAndServe(":80", r)

}

// redirect
func redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

// GetAllHandler list all records
func GetAllHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Records)  // return Records to writer
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

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MaxID) // return MaxID to writer
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

// UpdateOneHandler updates record.  ID required.  IF does not exists, crates new record
func UpdateOneHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	recordID, _ := strconv.Atoi(ps.ByName("id"))

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

	w.WriteHeader(http.StatusNoContent) // returns just status, nothing else
}

// DeleteHandler deletes record.  ID required.
func DeleteHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	recordID, _ := strconv.Atoi(ps.ByName("id"))

	Records = Filter(Records, func(t Item) bool {
		return t.ID != recordID
	})

	w.WriteHeader(http.StatusNoContent)
}

// CreateTest2 creates new record, returns record MaxID
func CreateTest2(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


	defer r.Body.Close()

	htmlData, err := ioutil.ReadAll(r.Body) //<--- here!

	if err != nil {
		fmt.Println(err)
	}

	// print out
	fmt.Println(os.Stdout, string(htmlData)) //<-- here !

}


// CreateTest3 creates new record, returns record MaxID
func CreateTest3(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	name := r.FormValue("description")
	address := r.FormValue("complete")
	fmt.Fprintf(os.Stdout, "Name = %s\n", name)
	fmt.Fprintf(os.Stdout, "Address = %s\n", address)

}


// CreateTest4 working
func CreateTest4(_ http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	defer r.Body.Close()
	fmt.Println("response Body:", &r.Body)
	fmt.Println("response Status:", r.Header)
	fmt.Println("response Headers:", r.Response)
	fmt.Println("response Body:", r.Form)
	fmt.Println("response value:", r.PostFormValue("description" ))
	fmt.Println("response PostForm:", r.PostForm)
	fmt.Println("response Encode:", r.Form.Encode())
	fmt.Println("response len:", len(r.PostForm))

	for key, values := range r.Form {   // range over map
		for _, value := range values {    // range over []string
			fmt.Println("out: " + key, value)
		}
	}

	for key, values := range r.PostForm {   // range over map
		for _, value := range values {    // range over []string
			fmt.Println("out: " + key, value)
		}
	}
}

// CreateTest creates new record, returns record MaxID
func CreateTest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	json.NewEncoder(w).Encode(MaxID) // returns MaxID to writer
}
