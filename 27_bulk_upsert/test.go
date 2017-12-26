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
func readLine(filename string) {

	// Open file
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()

	bulk := Coll.Bulk()
	//bulk.Unordered()

	// Loop through each line, line by line
	var i int
	for err == nil && !isPrefix {
		ln := string(line)

		s := strings.Split(ln, "\t")
		fmt.Println(i, "output: ", s[0], "==>,", s[1], "==>,", s[2])

		column, err := sliceAtoi(s)
		//id, err := strconv.Atoi(s[0])
		check(err)

		// update := bson.M{"BR01": s[1], "BR02": s[2], "BR03": s[3], "BR04": s[4]}  // works
		// update := bson.M{"test": bson.M{"BR01": s[1], "BR02": s[2], "BR03": s[3], "BR04": s[4]}} // works as object
		// update := bson.M{"$inc": bson.M{"upserts": 1}} // works

		// update := bson.M{"test": bson.M{"BR01": s[1], "BR02": s[2], "BR03": s[3], "BR04": s[4]}} // works
		// selector := bson.M{"_id": id, "user.firstName": "Max"} // works
		selector := bson.M{"_id": column[0]} // works
		updateBranch(column[0])

		// update := bson.M{"test":nil}
		// update := bson.M{"$set": bson.M{"test3": nil}} // works
		// update := bson.M{"$set": bson.M{"test3": nil, "test": bson.M{"BR01": s[1], "BR02": s[2], "BR03": s[3], "BR04": s[4]}}} // works

		// update := bson.M{"$set": bson.M{"test": bson.M{"BR01": s[1], "BR02": s[2], "BR03": s[3], "BR04": s[4]}}} // works
		// update := bson.M{"$inc": bson.M{"upserts": 1}}

		// selector := bson.M{"name": "this-is-good-content3"}
		// update := bson.M{"$inc": bson.M{"download": 3}}

		// update := bson.M{"$push":bson.M{"user.$.sales":bson.M{"$each":addsales}}}

		// Declare and initialize datatype
		branches := []struct {
			ID     string `json:"id,omitempty" bson:"id"`
			Onhand int    `json:"onhand,omitempty" bson:"onhand"`
		}{
			{ID: "BR01", Onhand: column[1]},
			{ID: "BR02", Onhand: column[2]},
			{ID: "BR03", Onhand: column[3]},
			{ID: "BR04", Onhand: column[4]},
			{ID: "BR05", Onhand: column[5]},
			{ID: "BR06", Onhand: column[6]},
			{ID: "BR07", Onhand: column[7]},
			{ID: "BR08", Onhand: column[8]},
			{ID: "BR09", Onhand: column[9]},
			{
				ID:     "BR10",
				Onhand: column[10],
			},
		}
		//update := bson.M{"array1": []string{"obj1", "value"}} // works
		update := bson.M{"$set": bson.M{"branch2": branches}}

		//update := bson.M{"$push":bson.M{"user":bson.M{"$each":addsales}}} // works
		//update := bson.M{"$push":bson.M{"user.$.sales":bson.M{"$each":addsales}}}
		//update := bson.M{"user.$.age":99} // does not work

		bulk.Upsert(selector, update)
		//bulk.Update(selector, update)  // Update will not insert new item, if does not exist

		// Reading data from the first ten lines
		i++
		if i > 10 {
			break
		}

		line, isPrefix, err = r.ReadLine()
		check(err)
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

// update one record
func updateBranch(id int) (err error) {
	query := bson.M{
		"_id":       id,
		"branch.id": "BR01",
	}

	update := bson.M{
		"$set": bson.M{
			"branch.$.sales": id,
		},
	}

	out, err := Coll.Upsert(query, update)
	fmt.Print("Matched: ", out.Matched, "; Modified: ", out.Updated, "; ")
	return err
}

func main() {
	readLine("product_onhand.txt")
	fmt.Print("Time: ", time.Since(start), ";")
}

func check(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		//http.Error(w, "Error: " + err.Error(), 500)
	}
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
