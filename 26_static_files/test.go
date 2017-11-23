package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// RegularPage displays custom page
func RegularPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Regular page!")
}

func main() {
	RunServer(":80")
}

// MyNotFound test
func MyNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound) // StatusNotFound = 404
	w.Write([]byte("My own Not Found handler."))
	w.Write([]byte(" The page you requested could not be found."))
}

// RunServer test
func RunServer(host string) error {
	r := httprouter.New()
	r.GET("/", RegularPage)
	r.GET("/test/", RegularPage)
	//r.NotFound = http.FileServer(http.Dir("public"))
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	r.NotFound = http.HandlerFunc(MyNotFound)
	return http.ListenAndServe(host, r)
}
