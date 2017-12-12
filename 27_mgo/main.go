package main

import (
	"dblogin"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Content Example
type Content struct {
	Name     string
	Download int
	Date     time.Time
}

var (
	// Session for DB
	Session, _ = mgo.Dial(dblogin.Bookstore) // mongodb://username:yourpasscode@serverip:27017/database?authSource=admin
	// Database for connection
	Database   = "bookstore"
	// Collection for database
	Collection = "test"
	// Coll puts all together
	Coll       = Session.DB(Database).C(Collection)

	content = &Content{
		Name:     "this-is-good-content",
		Download: 1,
		Date:     time.Date(2016, 4, 7, 0, 0, 0, 0, time.UTC),
		//Date: time.Now(),
	}
)

//Drop Database
func dropDatabase() {
	fmt.Println("Drop Database")

	//db.dropDatabase()
	err := Session.DB(Database).DropDatabase()
	if err != nil {
		panic(err)
	}
}

//Insert
func testInsert() {
	fmt.Println("Test Insert into MongoDB")
	c := bson.M{
		"name":     "this-is-good-content",
		"download": 1,
		"date":     time.Date(2016, 4, 7, 0, 0, 0, 0, time.UTC),
	}

	/*
		db.content.insert({
			Name: "this-is-good-content",
			Download:    1,
			Date:        new Date("2016-04-07"),
			}
		)
	*/
	Coll.Insert(c)

}

//Multiple Insert
func testMultipleInsert() {
	fmt.Println("Test Multiple Insert into MongoDB")
	var contentArray []interface{}

	/*
		db.content.insert([
			{
				Name: "this-is-good-content",
				Download:    1,
				Date:        new Date("2016-04-07"),
			},
			{
				Name: "this-is-good-content",
				Download:    2,
				Date:        new Date("2016-04-07"),
			},
			{
				Name: "this-is-good-content",
				Download:    3,
				Date:        new Date("2016-04-07"),
			},
			{
				Name: "this-is-good-content",
				Download:    4,
				Date:        new Date(),
			},
			]
		)
	*/
	//contentArray = append(contentArray, &Content{
	//	Name:     "this-is-good-content",
	//	Download: 1,
	//	Date:     time.Date(2016, 4, 7, 0, 0, 0, 0, time.UTC),
	//})
	contentArray = append(contentArray, bson.M{
		"name":     "this-is-good-content",
		"download": 1,
		"date":     time.Date(2016, 4, 7, 0, 0, 0, 0, time.UTC),
	})

	contentArray = append(contentArray, &Content{
		Name:     "this-is-good-content",
		Download: 2,
		Date:     time.Date(2016, 4, 8, 0, 0, 0, 0, time.UTC),
	})

	//same date
	contentArray = append(contentArray, &Content{
		Name:     "this-is-good-content",
		Download: 3,
		Date:     time.Date(2016, 4, 8, 0, 0, 0, 0, time.UTC),
	})

	contentArray = append(contentArray, &Content{
		Name:     "this-is-good-content",
		Download: 3,
		Date:     time.Date(2016, 4, 9, 0, 0, 0, 0, time.UTC),
	})

	contentArray = append(contentArray, &Content{
		Name:     "this-is-good-content2",
		Download: 4,
		Date:     time.Now(),
	})

	Coll.Insert(contentArray...)
}

//Bulk Insert
func testBulkInsert() {
	fmt.Println("Test Bulk Insert into MongoDB")
	bulk := Coll.Bulk()

	var contentArray []interface{}
	contentArray = append(contentArray, &Content{
		Name:     "this-is-good-content",
		Download: 1,
		Date:     time.Date(2016, 4, 7, 0, 0, 0, 0, time.UTC),
	})

	contentArray = append(contentArray, &Content{
		Name:     "this-is-good-content",
		Download: 2,
		Date:     time.Now(),
	})

	bulk.Insert(contentArray...)
	_, err := bulk.Run()
	if err != nil {
		panic(err)
	}
}

//Update
//db.collection.update(
//   <query>,
//   <update>,
//   {
//     upsert: <boolean>,
//     multi: <boolean>,
//     writeConcern: <document>
//   }
//)
func testUpdate() {
	fmt.Println("Test Update in MongoDB")

	//db.content.update({name: "this-is-good-content"})
	selector := bson.M{"name": "this-is-good-content"}

	//Update One and Replace Doc
	//update := bson.M{"download": 3}
	//err := Coll.Update(selector, update)

	//Update One
	//update := bson.M{"$set": bson.M{"download": 3}}
	//err := Coll.Update(selector, update)

	//Update All
	update := bson.M{"$set": bson.M{"download": 3}}
	_, err := Coll.UpdateAll(selector, update)
	if err != nil {
		panic(err)
	}
}

//Upsert
//db.collection.update(
//   <query>,
//   <update>,
//   {
//     upsert: <boolean>,
//     multi: <boolean>,
//     writeConcern: <document>
//   }
//)
func testUpsert() {
	fmt.Println("Test Upsert in MongoDB")

	//Upsert
	update := bson.M{"$inc": bson.M{"download": 3}}
	selector := bson.M{"name": "this-is-good-content3"}

	_, err := Coll.Upsert(selector, update)
	if err != nil {
		panic(err)
	}
}

//Bulk Upsert
func testBulkUpsert() {
	fmt.Println("Test Bulk Upsert in MongoDB")
	bulk := Coll.Bulk()

	//Upsert
	update := bson.M{"$inc": bson.M{"download": 3}}
	selector := bson.M{"name": "this-is-good-content3"}

	bulk.Upsert(selector, update)
	bulk.Upsert(selector, update)
	bulk.Upsert(selector, update)
	_, err := bulk.Run()
	if err != nil {
		panic(err)
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

//Aggregate
func testAggregate() {
	pipeline := []bson.M{
		{"$match": bson.M{"name": "this-is-good-content"}},
		{"$group": bson.M{"_id": "$date",
			//bson.M{"_id": "$name",
			"download": bson.M{"$sum": "$download"},
		},
		},
		{"$sort": bson.M{"download": 1}},//1: Ascending, -1: Descending

	}
	pipe := Coll.Pipe(pipeline)

	result := []bson.M{}
	//err := pipe.AllowDiskUse().All(&result) //allow disk use
	err := pipe.All(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", result)
}

func main() {
	//dropDatabase()
	//testInsert()
	//testMultipleInsert()
	//testBulkInsert()
	//testUpdate()
	//testUpsert()
	testBulkUpsert()
	//testSelect()
	//testAggregate()

}
