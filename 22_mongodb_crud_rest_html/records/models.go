package records

import (
	"../config"
	"errors"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

// Book Record fields
type Book struct {
	// add ID and tags if you need them
	// ID     bson.ObjectId // `json:"id" bson:"_id"`
	Isbn   string  // `json:"isbn" bson:"isbn"`
	Title  string  // `json:"title" bson:"title"`
	Author string  // `json:"author" bson:"author"`
	Price  float32 // `json:"price" bson:"price"`
}

// Prod Record fields
type Prod struct {
	// add ID and tags if you need them
	// ID     bson.ObjectId // `json:"id" bson:"_id"`
	Pline    string  // `json:"isbn" bson:"isbn"`
	Bline    string  // `json:"title" bson:"title"`
	Category string  // `json:"author" bson:"author"`
	Price    float32 // `json:"price" bson:"price"`
}

// SomeProducts records from produts3 collection
func SomeProducts() ([]Prod, error) {
	prods := []Prod{}
	err := config.Products3.Find(bson.M{"pline": "53BQ"}).All(&prods)
	if err != nil {
		return nil, err
	}
	return prods, nil
}

// AllBooks finds all records
func AllBooks() ([]Book, error) {
	bks := []Book{}
	err := config.Books.Find(bson.M{}).All(&bks)
	if err != nil {
		return nil, err
	}
	return bks, nil
}

// OneBook Find one record
func OneBook(r *http.Request) (Book, error) {
	bk := Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. Bad Request")
	}
	err := config.Books.Find(bson.M{"isbn": isbn}).One(&bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

// PutBook - Insert
func PutBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	// validate form values
	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number")
	}
	bk.Price = float32(f64)

	// insert values
	err = config.Books.Insert(bk)
	if err != nil {
		return bk, errors.New("500. Internal Server Error" + err.Error())
	}
	return bk, nil
}

// UpdateBook - Update record
func UpdateBook(r *http.Request) (Book, error) {
	// get form values
	bk := Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty")
	}

	// convert form values
	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Enter number for price")
	}
	bk.Price = float32(f64)

	// update values
	err = config.Books.Update(bson.M{"isbn": bk.Isbn}, &bk)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

// DeleteBook Delete record
func DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	err := config.Books.Remove(bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
