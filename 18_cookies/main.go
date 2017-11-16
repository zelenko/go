package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

const navigationMenu string = `<a href="/">01</a> | 
<a href="/02">02 login </a> | 
<a href="/03">03 view</a> | 
<a href="/04">04 logout</a> | 
<a href="/hello/wow">hello</a> | <br>

`

// Main page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	htmlString := navigationMenu + ` 01 `
	output, err := r.Cookie("username")
	if err == nil {
		htmlString += output.Name + "(" + output.Value + ")"
	}

	fmt.Fprint(w, htmlString)
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
	htmlString := navigationMenu + ` 02` + output
	fmt.Fprint(w, htmlString)
}

// View
func p03(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Display all current cookies
	output := ""
	for _, cookie := range r.Cookies() {
		output += " " + cookie.Name + "(" + cookie.Value + ")"
	}
	htmlString := navigationMenu + ` 03` + output
	fmt.Fprint(w, htmlString)
}

// Logout
func p04(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// remove cookie
	expiration := time.Unix(0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)

	fmt.Fprint(w, navigationMenu+` 04 logout`)
}

// Hello page
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	htmlString := navigationMenu + ` 03
	` + ps.ByName("name")
	fmt.Fprintf(w, htmlString)
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
