// One page MGO aggregation
//
// Copyright (c) 2018 - Zelenko <https://github.com/zelenko>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2/bson"
	"html/template"
	"net/http"
)

// Output is the exported data type
type Output struct {
	Message string
}

// itemTemplate can be overridden in each function locally
var itemTemplate = `
	{{- range listFunction }}
	{{ .Message}}<br />
	{{- end}}`

// htmlTemplate is what presented to the client
func htmlTemplate(itemTemplate string) string {
	output := `<html>
<head>
    <title>{{ . }}</title>
</head>
<body>
    {{ . }}
    <p>
    <a href="/">Main</a> |
	<a href="/list/">List</a> |
	<a href="/products/">Products</a> |
	<a href="/aggregate/">Aggregate</a> |
    <a href="/view/">View</a>
    </p>


` + itemTemplate + `
</body>
</html>
`
	return output
}

// main is the entry point for the program.
func main() {
	defer s.Close()

	fmt.Println("HTTP port :80")

	r := httprouter.New()
	r.GET("/", indexHandler)
	r.GET("/view/", viewHandler)
	r.GET("/list/", listHandler)
	r.GET("/products/", productsHandler)
	r.GET("/aggregate/", aggregateHandler)
	//r.GET("/stat/", statHandler)

	//r.NotFound = http.FileServer(http.Dir("public"))

	http.ListenAndServe(":80", r)

}

// viewHandler uses append to populate slice
func viewHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	listFunction := func() (r []Output) {
		r = append(r, Output{Message: fmt.Sprint("one")})
		r = append(r, Output{Message: fmt.Sprint("two")})
		r = append(r, Output{Message: fmt.Sprint("three")})
		r = append(r, Output{Message: fmt.Sprint("four")})
		return
	}

	functionMap := template.FuncMap{
		"listFunction": listFunction,
		//"listFunction3": listFunction3,
	}

	t := template.New("main") //name of the template is main
	t = t.Funcs(functionMap)
	t, _ = t.Parse(htmlTemplate(itemTemplate)) // parsing of template string
	err := t.Execute(w, "View")

	//t := template.Must(template.New("email.html").Funcs(functionMap).ParseFiles("email.html"))
	//err := t.Execute(os.Stdout, createMockStatement())
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

// indexHandler uses append to populate slice
func indexHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	listFunction := func() (r []Output) {
		r = append(r, Output{Message: fmt.Sprint("one -")})
		r = append(r, Output{Message: fmt.Sprint("two-")})
		r = append(r, Output{Message: fmt.Sprint("three-")})
		r = append(r, Output{Message: fmt.Sprint("four-")})
		return
	}

	t := template.New("main").Funcs(template.FuncMap{"listFunction": listFunction})
	t, _ = t.Parse(htmlTemplate(itemTemplate))
	t.Execute(w, "Index")
}

// listHandler is executing html template
func listHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// branch is the data type
	type branch struct {
		Message string
		Note    string
	}

	// listFunction returns slice of branches.  Branches are declared by a literal.
	listFunction := func() []branch {

		// r is a slice of branch data types
		r := []branch{
			{Message: "column[2]", Note: "This is the note!"},
			{Message: "column[3]", Note: "This is the note!"},
			{Message: "column[1]", Note: "This is the note!"},
			{Message: "column[4]", Note: "This is the note!"},
			{Message: "column[5]", Note: "This is the note!"},
			{Message: "column[6]", Note: "This is the note!"},
			{Message: "column[8]", Note: "This is the note!"},
			{Message: "column[9]", Note: "This is the note!"},
			{Message: "column[7]", Note: "This is the note!"},
		}

		return r
	}

	itemTemplate := `
	{{- range listFunction }}
	{{ .Message}} => {{ .Note}}<br />
	{{- end}}`

	t := template.New("main").Funcs(template.FuncMap{"listFunction": listFunction})
	t, _ = t.Parse(htmlTemplate(itemTemplate))
	t.Execute(w, "List")
}

// Anything below here requires database connection.  See db.go for connection details

// productHandler is executing html template
func productsHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// Prod Record fields.  The TYPE!
	type Prod struct {
		// add ID and tags if you need them
		// ID     bson.ObjectId // `json:"id" bson:"_id"`
		Pline    string  // `json:"isbn" bson:"isbn"`
		Bline    string  // `json:"title" bson:"title"`
		Category string  // `json:"author" bson:"author"`
		Price    float32 // `json:"price" bson:"price"`
	}

	// SomeProducts records from products3 collection.  Get the slice of TYPES!
	listFunction := func() []Prod {
		prods := []Prod{{}}

		// same query in mongodb compass:
		//{$and:[{'sales.BR01':{$gt:30000}},{'sales.BR01':{$lt:60000}}]}

		err := Products3.Find(bson.M{
			"$and": []bson.M{
				{"sales.BR01": bson.M{"$gt": 30000}},
				{"sales.BR01": bson.M{"$lt": 60000}},
			},
		}).All(&prods)

		if err != nil {
			return nil
		}
		return prods
	}

	// Present on HTML page
	itemTemplate := `
	{{- range listFunction }}
	{{ .Pline}} => {{ .Bline}} ==> {{ .Category}} ==> {{ .Price}}<br />
	{{- end}}`
	t := template.New("main").Funcs(template.FuncMap{"listFunction": listFunction})
	t, _ = t.Parse(htmlTemplate(itemTemplate))
	t.Execute(w, "List")
}

// aggregateHandler is executing html template
func aggregateHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	// ItemData Record fields.  The TYPE!
	// This is sample output from mongodb shell:
	// { "_id" : 42973, "value" : 11.36, "sales_all" : 598.23, "branch" : "BR19", "maxSale" : 303.96 }
	type ItemData struct {
		// ID     bson.ObjectId // `json:"id" bson:"_id"`
		ID       uint `bson:"_id"` // tags required
		Value    float32
		SalesAll float32 `bson:"sales_all"`
		Branch   string
		MaxSale  float32 `bson:"maxSale"`
	}
	// This is how it will look on HTML template:
	//{{ .ID}} => {{ .Value}} ==> {{ .SalesAll}} ==> {{ .Branch}} ==> {{ .MaxSale}}<br />

	// SomeProducts records from products3 collection.  Get the slice of TYPES!
	listFunction := func() []ItemData {
		/*	This is how the PIPE looks in mongodb shell:
			db.products3.aggregate([
			  {$match: {'onhand.BR03':{"$gt":0},'sales.BR03':{'$lt':2}, sales_all:{'$gt':0}}},
			  {$unwind: '$branch'},
			  {$project: {"_id":1, "value":"$onhand_value.BR03", "sales_all":1,"sales":"$branch.sales","branch":"$branch.id"}},
			  {$sort:{'value':-1,'sales':-1}},
			  {$group: {_id:'$_id', value: { $first: "$value" }, sales_all: { $first: "$sales_all" }, branch: { $first: "$branch" }, "maxSale": { $max: '$sales'} }},
			  {$match: {'branch':'BR19'}},
			  {$sort:{'value':-1}},
			  {$limit: 300}
			])
		*/

		pipeline := []bson.M{
			{
				"$match": bson.M{
					"onhand.BR03": bson.M{"$gt": 0},
					"sales.BR03":  bson.M{"$lt": 2},
					"sales_all":   bson.M{"$gt": 0},
				},
			},
			{
				"$unwind": "$branch",
			},
			{
				"$project": bson.M{
					"_id":       1,
					"value":     "$onhand_value.BR03",
					"sales_all": 1,
					"sales":     "$branch.sales",
					"branch":    "$branch.id",
				},
			},
			{
				"$sort": bson.M{"value": -1, "sales": -1},
			},
			{
				"$group": bson.M{
					"_id":       "$_id",
					"value":     bson.M{"$first": "$value"},
					"sales_all": bson.M{"$first": "$sales_all"},
					"branch":    bson.M{"$first": "$branch"},
					"maxSale":   bson.M{"$max": "$sales"},
				},
			},
			{
				"$match": bson.M{"branch": "BR19"},
			},
			{
				"$sort": bson.M{"value": -1},
			},
			{
				"$limit": 30,
			},
		}

		pipe := Products3.Pipe(pipeline)

		// Declaring type
		prods := []ItemData{{}}
		err := pipe.All(&prods)
		if err != nil {
			//http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
			fmt.Println(err)
			return nil
		}
		return prods
	}

	// Present on HTML page
	itemTemplate := `
	{{- range listFunction }}
	{{ .ID}} => {{ .Value}} ==> {{ .SalesAll}} ==> {{ .Branch}} ==> {{ .MaxSale}}<br />
	{{- end}}`
	t := template.New("main").Funcs(template.FuncMap{"listFunction": listFunction})
	t, _ = t.Parse(htmlTemplate(itemTemplate))
	t.Execute(w, "List")
}
