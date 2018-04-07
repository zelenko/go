package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

const navigationMenu string = `
	<a href="/">home</a> |
	<a href="/login">login </a> |
	<a href="/view">view</a> |
	<a href="/logout">logout</a> |
	<a href="/hello/page-01">page 01</a> |
	<a href="/hello/page-02">page 02</a> |
	<a href="/hello/page-03">page 03</a> |
	<br>
`

// main is the entry point for the program.
func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	router := httprouter.New()
	router.GET("/", indexH)
	router.GET("/login", loginH)
	router.GET("/view", view)
	router.GET("/logout", logoutH)
	router.GET("/hello/:name", hello)

	log.Println(http.ListenAndServe(":8080", router))
}

// Index Main page
func indexH(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	htmlString := navigationMenu + ` main page `
	output, err := r.Cookie("username")
	if err == nil {
		htmlString += output.Name + " (" + output.Value + ")"
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(htmlString))
}

// Login
func loginH(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// set cookies
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: "username", Value: "Superuser", Expires: expiration, HttpOnly: true} //Secure: true,
	http.SetCookie(w, &cookie)

	// Display all current cookies
	output := ""
	for _, cookie := range r.Cookies() {
		output += " " + cookie.Name + "(" + cookie.Value + ")"
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(navigationMenu + ` login` + output))
}

// View
func view(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Display all current cookies
	output := ""
	for _, cookie := range r.Cookies() {
		output += " " + cookie.Name + "(" + cookie.Value + ")"
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(navigationMenu + ` view` + output))
}

// Logout
func logoutH(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// remove cookie
	expiration := time.Unix(0, 0)
	cookie := http.Cookie{Name: "username", Value: "Superuser", Expires: expiration, HttpOnly: true} //Secure: true,
	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(navigationMenu + ` logout`))
}

// Hello page
func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	htmlString := navigationMenu + ` >>	` + ps.ByName("name")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(htmlString))
}
