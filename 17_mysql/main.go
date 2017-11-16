package main

import (
	"database/sql"
	"dblogin"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

var db *sql.DB
var err error

type personaFields struct {
	Id       int
	Username string
}

func signupPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "signup.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var user string

	err := db.QueryRow("SELECT username FROM go_users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		_, err = db.Exec("INSERT INTO go_users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(res, "Server error, unable to create your account.", 500)
			return
		}

		res.Write([]byte("User created!"))
		return
	case err != nil:
		http.Error(res, "Server error, unable to create your account.", 500)
		return
	default:
		http.Redirect(res, req, "/", 301)
	}
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.ServeFile(res, req, "login.html")
		return
	}

	username := req.FormValue("username")
	password := req.FormValue("password")

	var databaseUsername string
	var databasePassword string

	err := db.QueryRow("SELECT username, password FROM go_users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)

	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(res, req, "/login", 301)
		return
	}

	res.Write([]byte("Hello " + databaseUsername))

}

func allPage(res http.ResponseWriter, req *http.Request) {

	var (
		person  personaFields
		persons []personaFields
	)

	rows, err := db.Query("select id, username from go_users;")

	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&person.Id, &person.Username)
		persons = append(persons, person)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()

	out := " "

	for _, v := range persons {
		id := strconv.Itoa(v.Id)
		out += v.Username + "(" + id + "), "
	}

	res.Write([]byte("Hello " + out))

}

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func main() {
	db, err = sql.Open("mysql", dblogin.Catalog) // user_name:password@tcp(192.168.0.2:3306)/database_name
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", allPage)
	http.ListenAndServe(":8080", nil)
}
