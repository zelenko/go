package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"fmt"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}


func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "page-02.html", nil)
	checkErr(err)
}

func serveTemplate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// Declare template
	//tpl, err := template.ParseFiles(fp)  // can have multiple, separated by comas

	output := params.ByName("tmpl") + " | " + r.URL.Path 

	// Run template
	if err := tpl.ExecuteTemplate(w, "page.html", output); err != nil {
		checkErr(err)
		// Return a generic "Internal Server Error" message
		// http.Error(w, http.StatusText(500), 500)
	}
}


func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}

func main() {
	fmt.Println("Listeting on port 8080")
	r := httprouter.New()
	r.ServeFiles("/static/*filepath", http.Dir("static"))
	r.GET("/", index)
	r.GET("/test/:tmpl", serveTemplate)
	http.ListenAndServe("localhost:8080", r)
}