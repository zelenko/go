package main

// Created on 2017-11-06 to test GET from url
import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	//"log"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	http.HandleFunc("/", index)
	//http.HandleFunc("/process", process)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		fmt.Println("method was post")
	}

	if r.Method == http.MethodGet {
		fmt.Println("method was get")
	}

	next, err := strconv.Atoi(r.FormValue("next"))
	if err != nil {
		fmt.Println(err)
		//log.Println("LOGGED line 38", err)
	}

	senddata := struct {
		Prev int
		Now  int
		Next int
	}{
		next - 1,
		next,
		next + 1,
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", senddata)

	if err != nil {
		fmt.Println(err)
	}
}
