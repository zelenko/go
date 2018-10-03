package main

import (
	"fmt"
	"net/http"
	"time"
)

var body = `<body style="font: 2rem/1.5 monospace;max-width:40rem;margin:0 auto;padding:4rem;">`

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/view", view)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/login.php", login)
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if userNotAuthenticated(r) {
		fmt.Fprintf(w, body+`<form method="post" action="login.php"><input name="name1" placeholder="User Name" type="text">`+
			"<br>\n"+`<input name="lp" type="password" placeholder="Password"><input type="submit" hidden="true" /></form>`)
		return
	}
	w.Write([]byte(body + `member area <a href="/logout">logout</a> | <a href="/">home</a> | <a href="/view">view</a> | <a href="/hello">hello</a>`))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if userNotAuthenticated(r) {
		//fmt.Fprintf(w, `NOT authenticated <a href="/">home</a>`)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, body+`Hello World! <a href="/">home</a>`)
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.Method != http.MethodPost {
		//w.Write([]byte(`method not post`))
		return
	}
	r.ParseForm() // prepare data for parsing
	password := ""
	if len(r.Form["lp"]) > 0 {
		password = r.Form["lp"][0]
	}
	if password == "d5a3ee400994798a2916d10d492f6c1ab90ad5a3ee400994798a" && userNotAuthenticated(r) {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "username",
			Value:    "2916d10d492f6c1ab90ad5a3ee400994798a2916d10d492f6c1ab90ad5a3ee400994798a",
			Expires:  expiration,
			HttpOnly: true,
			Secure: true,
			SameSite: http.SameSiteStrictMode
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	time.Sleep(100 * time.Second)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func view(w http.ResponseWriter, r *http.Request) {
	if userNotAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(body + `member area <a href="/logout">logout</a> | <a href="/">home</a>`))
}

func logout(w http.ResponseWriter, r *http.Request) {
	// remove cookie
	expiration := time.Unix(0, 0)
	cookie := http.Cookie{Name: "username", Value: "2916d10d492f6c1ab90ad5a3ee400994798a2916d10d492f6c1ab90ad5a3ee400994798a", Expires: expiration, HttpOnly: true} //Secure: true,
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func userNotAuthenticated(r *http.Request) bool {
	username, err := r.Cookie("username")
	if err == nil && username.Value == "2916d10d492f6c1ab90ad5a3ee400994798a2916d10d492f6c1ab90ad5a3ee400994798a" {
		return false
	}
	return true
}
