package main

import (
	"html/template" // https://golang.org/pkg/html/template/
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

// main is the entry point for the program.
func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", abot)
	http.HandleFunc("/contact", cntct)
	http.HandleFunc("/apply", aply)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./pub"))))
	http.ListenAndServe(":80", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Index Page 01",
	}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)

	if err != nil {
		log.Println("LOGGED", err)
		http.Error(w, "Internal serverrrrrr error", http.StatusInternalServerError)
		return
	}
}

func abot(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "About Page",
	}

	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func cntct(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Contact Page",
	}

	err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func aply(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Apply Page",
	}

	var first string
	if req.Method == http.MethodPost {
		first = req.FormValue("fname")
		pd.FirstName = first
	}

	err := tpl.ExecuteTemplate(w, "apply.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
