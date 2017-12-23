package main

import (
	"bufio"
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// Content Example
type Content struct {
	Name     string
	Download int
	Date     time.Time
}

var (
	start = time.Now()
	// Session for DB
	Session, _ = mgo.Dial(dblogin.Bookstore) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	// Database for connection
	Database = "bookstore"
	// Collection for database
	Collection = "test"
	// Coll puts all together
	Coll = Session.DB(Database).C(Collection)
)

// ReadLine reads file faster
func ReadLine(filename string) {
	// Open file
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()

	// Bulk for MongoDB
	bulk := Coll.Bulk()
	//bulk.Unordered()

	// Loop through each line, line by line
	var i int
	for err == nil && !isPrefix {
		ln := string(line)

		s := strings.Split(ln, "\t")
		//ip, port := s[0], s[1]
		fmt.Println(i, "output: ", s[0], "==>,", s[1], "==>,", s[2])

		//update := bson.M{"$inc": bson.M{"download": 3}}
		id, err := strconv.Atoi(s[0])
		checkErr(err)

		// update := bson.M{"BR01": s[1], "BR02": s[2], "BR03": s[3], "BR04": s[4]}  // works
		// update := bson.M{"test": bson.M{"BR01": s[1], "BR02": s[2], "BR03": s[3], "BR04": s[4]}} // works as object
		//update := bson.M{"$inc": bson.M{"upserts": 1}} // works

		//update := bson.M{"test": bson.M{"BR01": s[1], "BR02": s[2], "BR03": s[3], "BR04": s[4]}} // works
		selector := bson.M{"_id": id}

		//update := bson.M{"test":nil}
		update := bson.M{"test2": nil}

		//update := bson.M{"$inc": bson.M{"upserts": 1}}
		//selector := bson.M{"name": "this-is-good-content3"}
		//['multi' => false, 'upsert' => true]

		bulk.Upsert(selector, update)
		//bulk.Update(selector, update)
		//bulk.Upsert(selector, update, bson.M{"multi": false}, bson.M{"upsert": true})

		i++
		if i > 10 {
			break
		}

		line, isPrefix, err = r.ReadLine()
	}

	// Process Bulk
	out, err := bulk.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Print("Matched: ", out.Matched, "; Modified: ", out.Modified, "; ")

	if isPrefix {
		fmt.Println("buffer size to small")
		return
	}
	if err != io.EOF && err != nil {
		fmt.Println("error: ", err)
		return
	}
}

//Select
func testSelect() {
	fmt.Println("Test Select in MongoDB")
	var result Content
	var results []Content
	var query bson.M
	var err error

	//query := bson.M{"download": 1}
	//Select One
	err = Coll.Find(nil).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Select One: %+v\n", result)

	//Select Limit
	iter := Coll.Find(nil).Limit(2).Iter()
	err = iter.All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Select Limit: %+v\n", results)

	//Select All
	err = Coll.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Select All: %+v\n", results)

	//Select with query
	query = bson.M{"name": "this-is-good-content"}
	err = Coll.Find(query).All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Select contentname: %+v\n", results)

	//Sort (ascending order)
	query = bson.M{"name": "this-is-good-content"}
	err = Coll.Find(query).Sort("download").All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ascending sorted Result: %+v\n", results)

	//Sort (descending order)
	query = bson.M{"name": "this-is-good-content"}
	err = Coll.Find(query).Sort("-download").All(&results) //add - left side
	if err != nil {
		panic(err)
	}
	fmt.Printf("Descending sorted Result: %+v\n", results)
}

func main() {
	//testBulkUpsert()
	//testSelect()
	ReadLine("product_onhand.txt")
	fmt.Print("Time: ", time.Since(start), ";")
}

func checkErr(err error) {
	if err != nil {
		//fmt.Println("Error: ", err.Error())
		//http.Error(w, "Error: " + err.Error(), 500)
		fmt.Println(err)
	}
}
