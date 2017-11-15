package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

const navigation_menu string = `<a href="/">01</a> | 
<a href="/02">02 login </a> | 
<a href="/03">03 view</a> | 
<a href="/04">04 logout</a> | 
<a href="/hello/wow">hello</a> | <br>

`

// Main page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	html_string := navigation_menu + ` 01 `
	output, err := r.Cookie("username")
	if err == nil {
		html_string += output.Name + "(" + output.Value + ")"
	}

	fmt.Fprint(w, html_string)
}

// Login
func p02(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// set cookies
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)

	// Display all current cookies
	output := ""
	for _, cookie := range r.Cookies() {
		output += " " + cookie.Name + "(" + cookie.Value + ")"
	}
	html_string := navigation_menu + ` 02` + output
	fmt.Fprint(w, html_string)
}

// View
func p03(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Display all current cookies
	output := ""
	for _, cookie := range r.Cookies() {
		output += " " + cookie.Name + "(" + cookie.Value + ")"
	}
	html_string := navigation_menu + ` 03` + output
	fmt.Fprint(w, html_string)
}

// Logout
func p04(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// remove cookie
	expiration := time.Unix(0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)

	fmt.Fprint(w, navigation_menu+` 04 logout`)
}

// Hello page
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	html_string := navigation_menu + ` 03
	` + ps.ByName("name")
	fmt.Fprintf(w, html_string)
	//fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/02", p02)
	router.GET("/03", p03)
	router.GET("/04", p04)
	router.GET("/hello/:name", Hello)

	http.ListenAndServe(":8080", router)
}
