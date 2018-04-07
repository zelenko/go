// Setting up authentication using standard library only.  This is just for practice.
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// is user authorised? Pointer to yes or no.
var authorized = new(bool)

// main is the entry point for the program.
func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/member", member)
	http.HandleFunc("/", index)

	// Run the server
	httpServer := http.Server{Addr: ":8080"}
	log.Println(httpServer.ListenAndServe())
}

// index handles the main page
func index(w http.ResponseWriter, r *http.Request) {
	// get username cookie
	name := ""
	username, err := r.Cookie("username")

	// In this example if the username is provided, then user is authenticated.
	if err == nil && username.Value != "" {
		name = "<strong>" + username.Value + `</strong> <a href="/logout">logout</a>`
		*authorized = true
	} else {
		name = `<a href="/login">login</a>`
	}

	// show link to profile
	htmlBytes := []byte(`<p><a href="/member">profile</a> | ` + name + `</p>`)

	// show cookies
	output := ""
	for _, cookie := range r.Cookies() {
		output += "cookie: " + cookie.Name + " = " + cookie.Value + "<br>\n"
	}
	htmlBytes = append(htmlBytes, output...)

	// display page
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlBytes)
}

// member handler displays member profile
func member(w http.ResponseWriter, r *http.Request) {
	username, _ := r.Cookie("username")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	htmlBytes := []byte(`<a href="/">home</a> | `)

	// if not authorized, then display login link
	if !*authorized {
		htmlBytes = append(htmlBytes, `<a href="/login">login</a><br>`...)
		w.Write(htmlBytes)
		return
	}

	htmlBytes = append(htmlBytes, "<strong>"+username.Value+`</strong> <a href="/logout">logout</a><br>`...)

	// display all cookies
	output := ""
	for _, cookie := range r.Cookies() {
		output += cookie.String() + "; <br>\n"
	}
	htmlBytes = append(htmlBytes, output...)

	// display page
	w.Write(htmlBytes)
}

// login handler to process login form
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
			log.Println(err)
			return
		}

		// set username cookie
		cookie := http.Cookie{
			Name:     "username",
			Value:    r.FormValue("username"),
			Expires:  time.Now().Add(365 * 24 * time.Hour),
			HttpOnly: true,
			// Secure: true,
		}
		http.SetCookie(w, &cookie)

		// set username cookie
		cookiePass := http.Cookie{
			Name:     "password",
			Value:    r.PostFormValue("password"),
			Expires:  time.Now().Add(365 * 24 * time.Hour),
			HttpOnly: true,
			// Secure: true,
		}
		http.SetCookie(w, &cookiePass)

		// once cookies are set go to main page
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	} else if r.Method == "GET" && !*authorized {

		formHTML := `
		<p><a href="/">home</a></p>
		<form method="post" action="/login" enctype="application/x-www-form-urlencoded">
			<fieldset>
				<label for="un">Username</label>
				<input type="text" id="un" name="username"><br>
				<label for="pw">Passkey</label>
				<input type="password" id="pw" name="password">

				<input type="submit" name="submit" value="Login">
			</fieldset>
		</form>`

		// display page
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(formHTML))
		return

	} else {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}

// Logout, expire cookies, go to main page.
func logout(w http.ResponseWriter, r *http.Request) {

	// if not GET, then show error
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	// set EXPIRED username cookie
	cookie := http.Cookie{
		Name:     "username",
		Value:    r.FormValue("username"),
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		// Secure: true,
	}
	http.SetCookie(w, &cookie)

	// set EXPIRED password cookie
	cookiePass := http.Cookie{
		Name:     "password",
		Value:    r.PostFormValue("password"),
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		// Secure: true,
	}
	http.SetCookie(w, &cookiePass)

	*authorized = false

	// go to main page
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// loginFormExample is not being used right now.  but serves as an example.
func loginFormExamaple(w http.ResponseWriter, r *http.Request) {
	log.Println(r.ParseForm())

	// list all values from login form
	out1 := ""
	for key, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			out1 += fmt.Sprint(key + " -> " + value + "; <br>\n")
		}
	}
	out := fmt.Sprint(r.FormValue("username") + " and " + r.PostFormValue("password"))
	w.Write([]byte(out + "; " + out1))
}
