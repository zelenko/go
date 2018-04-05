package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_USER    = ""
	DB_PASS    = ""
	DB_NAME    = ""
	DB_CHARSET = "utf8"
)

func main() {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@/"+DB_NAME+"?charset="+DB_CHARSET)
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
