package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser    = ""
	dbPass    = ""
	dbName    = ""
	dbCharset = "utf8"
)

func main() {
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@/"+dbName+"?charset="+dbCharset)
	if err != nil {
		log.Fatal("Cannot open DB connection", err)
	}

	stmt, err := db.Prepare("INSERT data SET content=?")
	if err != nil {
		log.Fatal("Cannot prepare DB statement", err)
	}

	res, err := stmt.Exec("value")
	if err != nil {
		log.Fatal("Cannot run insert statement", err)
	}

	id, _ := res.LastInsertId()

	fmt.Printf("Inserted row: %d", id)
}
